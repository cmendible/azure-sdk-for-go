// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package repo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing"
)

type SDKRepository interface {
	WorkTree
	CreateReleaseBranch(releaseBranchName string) error
	AddReleaseCommit(rpName, namespaceName, specHash, version string) error
}

func OpenSDKRepository(path string) (SDKRepository, error) {
	wt, err := NewWorkTree(path)
	if err != nil {
		return nil, err
	}

	return &sdkRepository{
		WorkTree: wt,
	}, nil
}

func CloneSDKRepository(repoUrl, commitID string) (SDKRepository, error) {
	repoBasePath := filepath.Join(os.TempDir(), "generator_sdk")
	if _, err := os.Stat(repoBasePath); err == nil {
		os.RemoveAll(repoBasePath)
	}
	if err := os.Mkdir(repoBasePath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create tmp folder for generation: %+v", err)
	}

	wt, err := CloneWorkTree(fmt.Sprintf("%s.git", repoUrl), repoBasePath)
	if err != nil {
		return nil, err
	}

	err = wt.Checkout(&CheckoutOptions{
		Hash: plumbing.NewHash(commitID),
	})
	if err != nil {
		return nil, err
	}

	return &sdkRepository{
		WorkTree: wt,
	}, nil
}

type sdkRepository struct {
	WorkTree
}

func (s *sdkRepository) AddReleaseCommit(rpName, namespaceName, specHash, version string) error {
	log.Printf("Add release package and commit")
	if err := s.Add(fmt.Sprintf("sdk\\%s\\%s", rpName, namespaceName)); err != nil {
		return fmt.Errorf("failed to add 'profiles': %+v", err)
	}

	message := fmt.Sprintf("[Release] sdk/%s/%s/%s generation from spec commit: %s", rpName, namespaceName, version, specHash)
	if err := s.Commit(message); err != nil {
		if IsNothingToCommit(err) {
			log.Printf("There is nothing to commit. Message: %s", message)
			return nil
		}
		return fmt.Errorf("failed to commit changes: %+v", err)
	}

	return nil
}
func (s *sdkRepository) CreateReleaseBranch(releaseBranchName string) error {
	log.Printf("Checking out to %s", plumbing.NewBranchReferenceName(releaseBranchName))
	return s.Checkout(&CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(releaseBranchName),
		Create: true,
	})
}