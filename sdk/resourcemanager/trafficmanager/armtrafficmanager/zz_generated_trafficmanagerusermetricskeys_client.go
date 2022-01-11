//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrafficmanager

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TrafficManagerUserMetricsKeysClient contains the methods for the TrafficManagerUserMetricsKeys group.
// Don't use this type directly, use NewTrafficManagerUserMetricsKeysClient() instead.
type TrafficManagerUserMetricsKeysClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTrafficManagerUserMetricsKeysClient creates a new instance of TrafficManagerUserMetricsKeysClient with the specified values.
func NewTrafficManagerUserMetricsKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TrafficManagerUserMetricsKeysClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TrafficManagerUserMetricsKeysClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update a subscription-level key used for Real User Metrics collection.
// If the operation fails it returns the *CloudError error type.
func (client *TrafficManagerUserMetricsKeysClient) CreateOrUpdate(ctx context.Context, options *TrafficManagerUserMetricsKeysCreateOrUpdateOptions) (TrafficManagerUserMetricsKeysCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, options)
	if err != nil {
		return TrafficManagerUserMetricsKeysCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrafficManagerUserMetricsKeysCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return TrafficManagerUserMetricsKeysCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TrafficManagerUserMetricsKeysClient) createOrUpdateCreateRequest(ctx context.Context, options *TrafficManagerUserMetricsKeysCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TrafficManagerUserMetricsKeysClient) createOrUpdateHandleResponse(resp *http.Response) (TrafficManagerUserMetricsKeysCreateOrUpdateResponse, error) {
	result := TrafficManagerUserMetricsKeysCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserMetricsModel); err != nil {
		return TrafficManagerUserMetricsKeysCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *TrafficManagerUserMetricsKeysClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete a subscription-level key used for Real User Metrics collection.
// If the operation fails it returns the *CloudError error type.
func (client *TrafficManagerUserMetricsKeysClient) Delete(ctx context.Context, options *TrafficManagerUserMetricsKeysDeleteOptions) (TrafficManagerUserMetricsKeysDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, options)
	if err != nil {
		return TrafficManagerUserMetricsKeysDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrafficManagerUserMetricsKeysDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TrafficManagerUserMetricsKeysDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *TrafficManagerUserMetricsKeysClient) deleteCreateRequest(ctx context.Context, options *TrafficManagerUserMetricsKeysDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *TrafficManagerUserMetricsKeysClient) deleteHandleResponse(resp *http.Response) (TrafficManagerUserMetricsKeysDeleteResponse, error) {
	result := TrafficManagerUserMetricsKeysDeleteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeleteOperationResult); err != nil {
		return TrafficManagerUserMetricsKeysDeleteResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *TrafficManagerUserMetricsKeysClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the subscription-level key used for Real User Metrics collection.
// If the operation fails it returns the *CloudError error type.
func (client *TrafficManagerUserMetricsKeysClient) Get(ctx context.Context, options *TrafficManagerUserMetricsKeysGetOptions) (TrafficManagerUserMetricsKeysGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return TrafficManagerUserMetricsKeysGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrafficManagerUserMetricsKeysGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TrafficManagerUserMetricsKeysGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TrafficManagerUserMetricsKeysClient) getCreateRequest(ctx context.Context, options *TrafficManagerUserMetricsKeysGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TrafficManagerUserMetricsKeysClient) getHandleResponse(resp *http.Response) (TrafficManagerUserMetricsKeysGetResponse, error) {
	result := TrafficManagerUserMetricsKeysGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserMetricsModel); err != nil {
		return TrafficManagerUserMetricsKeysGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TrafficManagerUserMetricsKeysClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}