// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package release

import (
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/cmd/issue/link"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/cmd/v2/common"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/config"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/config/validate"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/flags"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/repo"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	releaseBranchNamePattern = "release-%s-%s-%s-%v"
	confirmComment           = "Release [PR](%s) is ready. Please check whether the package works well ([doc for testing with the generated Go SDK](https://github.com/Azure/azure-rest-api-specs/blob/main/documentation/code-gen/configure-go-sdk.md#test-with-the-generated-go-sdk)). If you are not a Golang user, you can mainly check whether the changelog meets your requirements."
)

// Release command
func Command() *cobra.Command {
	releaseCmd := &cobra.Command{
		Use:   "release-v2 <azure-sdk-for-go directory/commitid> <azure-rest-api-specs directory/commitid> <rp-name/release-request-json-config> [namespaceName]",
		Short: "Generate a v2 release of azure-sdk-for-go",
		Long: `This command will generate a new v2 release for azure-sdk-for-go with given rp name and namespace name.

azure-sdk-for-go directory/commitid: the directory path of the azure-sdk-for-go with git control or just a commitid for remote repo
azure-rest-api-specs directory: the directory path of the azure-rest-api-specs with git control or just a commitid for remote repo
rp-name/release-request-json-config: name of resource provider to be released (same for the swagger folder name) or release request json config file path generated by 'generator issue'
namespaceName: name of namespace to be released, default value is arm+rp-name

`,
		Args: cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			rpName := args[2]
			namespaceName := "arm" + rpName
			if len(args) == 4 {
				namespaceName = args[3]
			}

			ctx := commandContext{
				rpName:        rpName,
				namespaceName: namespaceName,
				flags:         ParseFlags(cmd.Flags()),
			}
			return ctx.execute(args[0], args[1])
		},
	}

	BindFlags(releaseCmd.Flags())

	return releaseCmd
}

type Flags struct {
	VersionNumber       string
	SwaggerRepo         string
	PackageTitle        string
	SDKRepo             string
	SpecRPName          string
	ReleaseDate         string
	SkipCreateBranch    bool
	SkipGenerateExample bool
	PackageConfig       string
	GoVersion           string
	Token               string
	UpdateSpecVersion   bool
}

func BindFlags(flagSet *pflag.FlagSet) {
	flagSet.String("version-number", "", "Specify the version number of this release")
	flagSet.String("package-title", "", "Specifies the title of this package")
	flagSet.String("sdk-repo", "https://github.com/Azure/azure-sdk-for-go", "Specifies the sdk repo URL for generation")
	flagSet.String("spec-repo", "https://github.com/Azure/azure-rest-api-specs", "Specifies the swagger repo URL for generation")
	flagSet.String("spec-rp-name", "", "Specifies the swagger spec RP name, default is RP name")
	flagSet.String("release-date", "", "Specifies the release date in changelog")
	flagSet.Bool("skip-create-branch", false, "Skip create release branch after generation")
	flagSet.Bool("skip-generate-example", false, "Skip generate example for SDK in the same time")
	flagSet.String("package-config", "", "Additional config for package")
	flagSet.String("go-version", "1.18", "Go version")
	flagSet.StringP("token", "t", "", "Specify the personal access token of Github")
	flagSet.Bool("update-spec-version", true, "Whether to update the commit id, the default is true")
}

func ParseFlags(flagSet *pflag.FlagSet) Flags {
	return Flags{
		VersionNumber:       flags.GetString(flagSet, "version-number"),
		PackageTitle:        flags.GetString(flagSet, "package-title"),
		SDKRepo:             flags.GetString(flagSet, "sdk-repo"),
		SwaggerRepo:         flags.GetString(flagSet, "spec-repo"),
		SpecRPName:          flags.GetString(flagSet, "spec-rp-name"),
		ReleaseDate:         flags.GetString(flagSet, "release-date"),
		SkipCreateBranch:    flags.GetBool(flagSet, "skip-create-branch"),
		SkipGenerateExample: flags.GetBool(flagSet, "skip-generate-example"),
		PackageConfig:       flags.GetString(flagSet, "package-config"),
		GoVersion:           flags.GetString(flagSet, "go-version"),
		Token:               flags.GetString(flagSet, "token"),
		UpdateSpecVersion:   flags.GetBool(flagSet, "update-spec-version"),
	}
}

type commandContext struct {
	rpName            string
	namespaceName     string
	flags             Flags
	pullRequestLabels string
}

func (c *commandContext) execute(sdkRepoParam, specRepoParam string) error {
	sdkRepo, err := common.GetSDKRepo(sdkRepoParam, c.flags.SDKRepo)
	if err != nil {
		return err
	}

	specCommitHash, err := common.GetSpecCommit(specRepoParam)
	if err != nil {
		return err
	}

	if path.Ext(c.rpName) == ".json" {
		return c.generateFromRequest(sdkRepo, specRepoParam, specCommitHash)
	} else {
		return c.generate(sdkRepo, specCommitHash)
	}
}

func (c *commandContext) generate(sdkRepo repo.SDKRepository, specCommitHash string) error {
	log.Printf("Release generation for rp: %s, namespace: %s", c.rpName, c.namespaceName)
	generateCtx := common.GenerateContext{
		SDKPath:           sdkRepo.Root(),
		SDKRepo:           &sdkRepo,
		SpecCommitHash:    specCommitHash,
		SpecRepoURL:       c.flags.SwaggerRepo,
		UpdateSpecVersion: c.flags.UpdateSpecVersion,
	}

	if c.flags.SpecRPName == "" {
		c.flags.SpecRPName = c.rpName
	}
	result, err := generateCtx.GenerateForSingleRPNamespace(&common.GenerateParam{
		RPName:              c.rpName,
		NamespaceName:       c.namespaceName,
		NamespaceConfig:     c.flags.PackageConfig,
		SpecficPackageTitle: c.flags.PackageTitle,
		SpecficVersion:      c.flags.VersionNumber,
		SpecRPName:          c.flags.SpecRPName,
		ReleaseDate:         c.flags.ReleaseDate,
		SkipGenerateExample: c.flags.SkipGenerateExample,
		GoVersion:           c.flags.GoVersion,
	})
	if err != nil {
		return fmt.Errorf("failed to finish release generation process: %+v", err)
	}
	// print generation result
	log.Printf("Generation result: %v", result)
	c.pullRequestLabels = result.PullRequestLabels

	if !c.flags.SkipCreateBranch {
		log.Printf("Create new branch for release")
		releaseBranchName := fmt.Sprintf(releaseBranchNamePattern, c.rpName, c.namespaceName, result.Version, time.Now().Unix())
		if err := sdkRepo.CreateReleaseBranch(releaseBranchName); err != nil {
			return fmt.Errorf("failed to create release branch: %+v", err)
		}

		log.Printf("Include the packages that is about to release in this release and do release commit...")
		// append a time in long to avoid collision of branch names
		if err := sdkRepo.AddReleaseCommit(c.rpName, c.namespaceName, generateCtx.SpecCommitHash, result.Version); err != nil {
			return fmt.Errorf("failed to add release package or do release commit: %+v", err)
		}
	}

	return nil
}

func (c *commandContext) generateFromRequest(sdkRepo repo.SDKRepository, specRepoParam, specCommitHash string) error {
	var generateErr []error
	var pullRequestUrls = make(map[string]string)
	var pushBranch = make(map[string]struct {
		requestLink      string
		pullRequestLabel string
	})
	forkRemote, err := repo.GetForkRemote(sdkRepo)
	if err != nil {
		return err
	}

	log.Printf("track2 parsing the config...")
	cfg, err := config.ParseConfig(c.rpName)
	if err != nil {
		return fmt.Errorf("parse config err: %v", err)
	}
	log.Printf("Configuration: %s", cfg.String())
	armServices, err := validate.ParseTrack2(cfg, specRepoParam)
	if err != nil {
		return err
	}
	for arm, packageInfos := range armServices {
		for _, info := range packageInfos {
			originalHead, err := sdkRepo.Head()
			if err != nil {
				return err
			}

			// run generator
			c.rpName = arm
			c.namespaceName = info.Name
			c.flags.SpecRPName = info.SpecName
			c.flags.PackageConfig = info.Tag
			if info.ReleaseDate != nil {
				c.flags.ReleaseDate = info.ReleaseDate.Format("2006-01-02")
			}
			err = c.generate(sdkRepo, specCommitHash)
			if err != nil {
				generateErr = append(generateErr, err)
				continue
			}

			// get current branch name
			generateHead, err := sdkRepo.Head()
			if err != nil {
				return err
			}
			pushBranch[generateHead.Name().Short()] = struct {
				requestLink      string
				pullRequestLabel string
			}{requestLink: info.RequestLink, pullRequestLabel: c.pullRequestLabels}

			log.Printf("git checkout %v", originalHead.Name().Short())
			if err := sdkRepo.Checkout(&repo.CheckoutOptions{
				Branch: plumbing.ReferenceName(originalHead.Name().Short()),
				Force:  true,
			}); err != nil {
				return err
			}
		}
	}

	for branch := range pushBranch {
		log.Printf("Fixes: %s\n", branch)
	}

	if c.flags.Token != "" {
		for branchName, issue := range pushBranch {
			log.Printf("git push fork %s\n", branchName)
			msg, err := common.ExecuteGitPush(sdkRepo.Root(), forkRemote.Config().Name, branchName)
			if err != nil {
				return fmt.Errorf("git push fork error:%v,msg:%s", err, msg)
			}

			githubUserName := repo.GetRemoteUserName(forkRemote)
			if githubUserName == "" {
				return errors.New("github user name not exist")
			}

			log.Printf("%s: create pull request...\n", branchName)
			pullRequestUrl, err := common.ExecuteCreatePullRequest(sdkRepo.Root(), link.SpecOwner, link.SDKRepo, githubUserName, branchName, repo.ReleaseTitle(branchName), issue.requestLink, c.flags.Token, issue.pullRequestLabel)
			if err != nil {
				return err
			}
			pullRequestUrls[branchName] = pullRequestUrl

			log.Printf("Leave a comment in %s...\n", issue)
			issueNumber := strings.Split(issue.requestLink, "/")
			err = common.ExecuteAddIssueComment(sdkRepo.Root(), link.SpecOwner, link.ReleaseIssueRepo, issueNumber[len(issueNumber)-1], fmt.Sprintf(confirmComment, pullRequestUrl), c.flags.Token)
			if err != nil {
				return err
			}

			log.Printf("Add Labels...\n")
			err = common.ExecuteAddIssueLabels(sdkRepo.Root(), link.SpecOwner, link.ReleaseIssueRepo, issueNumber[len(issueNumber)-1], c.flags.Token, []string{"PRready", issue.pullRequestLabel})
			if err != nil {
				return err
			}
		}
	}

	if len(pullRequestUrls) != 0 {
		log.Println("Fixes:")
		for branch, url := range pullRequestUrls {
			log.Printf("%s : %s", branch, url)
		}
	}

	if len(generateErr) != 0 {
		fmt.Println("generator error:")
		for _, e := range generateErr {
			fmt.Println(e)
		}
	}

	return nil
}
