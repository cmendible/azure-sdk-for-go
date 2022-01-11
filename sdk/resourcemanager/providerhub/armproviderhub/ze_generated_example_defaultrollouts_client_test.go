//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_Get.json
func ExampleDefaultRolloutsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewDefaultRolloutsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<provider-namespace>",
		"<rollout-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DefaultRollout.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_Delete.json
func ExampleDefaultRolloutsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewDefaultRolloutsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<provider-namespace>",
		"<rollout-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_CreateOrUpdate.json
func ExampleDefaultRolloutsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewDefaultRolloutsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<provider-namespace>",
		"<rollout-name>",
		armproviderhub.DefaultRollout{
			Properties: &armproviderhub.DefaultRolloutProperties{
				DefaultRolloutPropertiesAutoGenerated: armproviderhub.DefaultRolloutPropertiesAutoGenerated{
					Specification: &armproviderhub.DefaultRolloutPropertiesSpecification{
						DefaultRolloutSpecification: armproviderhub.DefaultRolloutSpecification{
							Canary: &armproviderhub.DefaultRolloutSpecificationCanary{
								CanaryTrafficRegionRolloutConfiguration: armproviderhub.CanaryTrafficRegionRolloutConfiguration{
									SkipRegions: []*string{
										to.StringPtr("eastus2euap")},
								},
							},
							RestOfTheWorldGroupTwo: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwo{
								TrafficRegionRolloutConfiguration: armproviderhub.TrafficRegionRolloutConfiguration{
									WaitDuration: to.StringPtr("<wait-duration>"),
								},
							},
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DefaultRollout.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_ListByProviderRegistration.json
func ExampleDefaultRolloutsClient_ListByProviderRegistration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewDefaultRolloutsClient("<subscription-id>", cred, nil)
	pager := client.ListByProviderRegistration("<provider-namespace>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DefaultRollout.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_Stop.json
func ExampleDefaultRolloutsClient_Stop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewDefaultRolloutsClient("<subscription-id>", cred, nil)
	_, err = client.Stop(ctx,
		"<provider-namespace>",
		"<rollout-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}