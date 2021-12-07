//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// ApplicationGroupsClient contains the methods for the ApplicationGroups group.
// Don't use this type directly, use NewApplicationGroupsClient() instead.
type ApplicationGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewApplicationGroupsClient creates a new instance of ApplicationGroupsClient with the specified values.
func NewApplicationGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ApplicationGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ApplicationGroupsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update an applicationGroup.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup, options *ApplicationGroupsCreateOrUpdateOptions) (ApplicationGroupsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationGroupName, applicationGroup, options)
	if err != nil {
		return ApplicationGroupsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ApplicationGroupsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup, options *ApplicationGroupsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, applicationGroup)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplicationGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ApplicationGroupsCreateOrUpdateResponse, error) {
	result := ApplicationGroupsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ApplicationGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Remove an applicationGroup.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGroupsClient) Delete(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsDeleteOptions) (ApplicationGroupsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ApplicationGroupsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ApplicationGroupsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ApplicationGroupsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get an application group.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsGetOptions) (ApplicationGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationGroupsClient) getHandleResponse(resp *http.Response) (ApplicationGroupsGetResponse, error) {
	result := ApplicationGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ApplicationGroupsClient) getHandleError(resp *http.Response) error {
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

// ListByResourceGroup - List applicationGroups.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGroupsClient) ListByResourceGroup(resourceGroupName string, options *ApplicationGroupsListByResourceGroupOptions) *ApplicationGroupsListByResourceGroupPager {
	return &ApplicationGroupsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ApplicationGroupsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationGroupList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationGroupsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationGroupsListByResourceGroupResponse, error) {
	result := ApplicationGroupsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroupList); err != nil {
		return ApplicationGroupsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ApplicationGroupsClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// ListBySubscription - List applicationGroups in subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGroupsClient) ListBySubscription(options *ApplicationGroupsListBySubscriptionOptions) *ApplicationGroupsListBySubscriptionPager {
	return &ApplicationGroupsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ApplicationGroupsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationGroupList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationGroupsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/applicationGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationGroupsListBySubscriptionResponse, error) {
	result := ApplicationGroupsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroupList); err != nil {
		return ApplicationGroupsListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ApplicationGroupsClient) listBySubscriptionHandleError(resp *http.Response) error {
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

// Update - Update an applicationGroup.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationGroupsClient) Update(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsUpdateOptions) (ApplicationGroupsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGroupsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGroupsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ApplicationGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.ApplicationGroup != nil {
		return req, runtime.MarshalAsJSON(req, *options.ApplicationGroup)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ApplicationGroupsClient) updateHandleResponse(resp *http.Response) (ApplicationGroupsUpdateResponse, error) {
	result := ApplicationGroupsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ApplicationGroupsClient) updateHandleError(resp *http.Response) error {
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