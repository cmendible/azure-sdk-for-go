//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

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
	"strconv"
	"strings"
)

// ExperimentsClient contains the methods for the Experiments group.
// Don't use this type directly, use NewExperimentsClient() instead.
type ExperimentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewExperimentsClient creates a new instance of ExperimentsClient with the specified values.
func NewExperimentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExperimentsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ExperimentsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCancel - Cancel a running Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) BeginCancel(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsBeginCancelOptions) (ExperimentsCancelPollerResponse, error) {
	resp, err := client.cancel(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsCancelPollerResponse{}, err
	}
	result := ExperimentsCancelPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExperimentsClient.Cancel", "original-uri", resp, client.pl, client.cancelHandleError)
	if err != nil {
		return ExperimentsCancelPollerResponse{}, err
	}
	result.Poller = &ExperimentsCancelPoller{
		pt: pt,
	}
	return result, nil
}

// Cancel - Cancel a running Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) cancel(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsBeginCancelOptions) (*http.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.cancelHandleError(resp)
	}
	return resp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ExperimentsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *ExperimentsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdate - Create or update a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, experimentName string, experiment Experiment, options *ExperimentsBeginCreateOrUpdateOptions) (ExperimentsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, experimentName, experiment, options)
	if err != nil {
		return ExperimentsCreateOrUpdatePollerResponse{}, err
	}
	result := ExperimentsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExperimentsClient.CreateOrUpdate", "original-uri", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ExperimentsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExperimentsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, experimentName string, experiment Experiment, options *ExperimentsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, experimentName, experiment, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExperimentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, experiment Experiment, options *ExperimentsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, experiment)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExperimentsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) Delete(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsDeleteOptions) (ExperimentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExperimentsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ExperimentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ExperimentsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExperimentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ExperimentsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) Get(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsGetOptions) (ExperimentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExperimentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExperimentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExperimentsClient) getHandleResponse(resp *http.Response) (ExperimentsGetResponse, error) {
	result := ExperimentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Experiment); err != nil {
		return ExperimentsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExperimentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetExecutionDetails - Get an execution detail of a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) GetExecutionDetails(ctx context.Context, resourceGroupName string, experimentName string, executionDetailsID string, options *ExperimentsGetExecutionDetailsOptions) (ExperimentsGetExecutionDetailsResponse, error) {
	req, err := client.getExecutionDetailsCreateRequest(ctx, resourceGroupName, experimentName, executionDetailsID, options)
	if err != nil {
		return ExperimentsGetExecutionDetailsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExperimentsGetExecutionDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsGetExecutionDetailsResponse{}, client.getExecutionDetailsHandleError(resp)
	}
	return client.getExecutionDetailsHandleResponse(resp)
}

// getExecutionDetailsCreateRequest creates the GetExecutionDetails request.
func (client *ExperimentsClient) getExecutionDetailsCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, executionDetailsID string, options *ExperimentsGetExecutionDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/executionDetails/{executionDetailsId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	if executionDetailsID == "" {
		return nil, errors.New("parameter executionDetailsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{executionDetailsId}", url.PathEscape(executionDetailsID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getExecutionDetailsHandleResponse handles the GetExecutionDetails response.
func (client *ExperimentsClient) getExecutionDetailsHandleResponse(resp *http.Response) (ExperimentsGetExecutionDetailsResponse, error) {
	result := ExperimentsGetExecutionDetailsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentExecutionDetails); err != nil {
		return ExperimentsGetExecutionDetailsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getExecutionDetailsHandleError handles the GetExecutionDetails error response.
func (client *ExperimentsClient) getExecutionDetailsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetStatus - Get a status of a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) GetStatus(ctx context.Context, resourceGroupName string, experimentName string, statusID string, options *ExperimentsGetStatusOptions) (ExperimentsGetStatusResponse, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, experimentName, statusID, options)
	if err != nil {
		return ExperimentsGetStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExperimentsGetStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsGetStatusResponse{}, client.getStatusHandleError(resp)
	}
	return client.getStatusHandleResponse(resp)
}

// getStatusCreateRequest creates the GetStatus request.
func (client *ExperimentsClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, statusID string, options *ExperimentsGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/statuses/{statusId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	if statusID == "" {
		return nil, errors.New("parameter statusID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{statusId}", url.PathEscape(statusID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getStatusHandleResponse handles the GetStatus response.
func (client *ExperimentsClient) getStatusHandleResponse(resp *http.Response) (ExperimentsGetStatusResponse, error) {
	result := ExperimentsGetStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentStatus); err != nil {
		return ExperimentsGetStatusResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getStatusHandleError handles the GetStatus error response.
func (client *ExperimentsClient) getStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get a list of Experiment resources in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) List(resourceGroupName string, options *ExperimentsListOptions) *ExperimentsListPager {
	return &ExperimentsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ExperimentsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExperimentListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExperimentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ExperimentsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments"
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
	reqQP.Set("api-version", "2021-09-15-preview")
	if options != nil && options.Running != nil {
		reqQP.Set("running", strconv.FormatBool(*options.Running))
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExperimentsClient) listHandleResponse(resp *http.Response) (ExperimentsListResponse, error) {
	result := ExperimentsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentListResult); err != nil {
		return ExperimentsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExperimentsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAll - Get a list of Experiment resources in a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) ListAll(options *ExperimentsListAllOptions) *ExperimentsListAllPager {
	return &ExperimentsListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExperimentsListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExperimentListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *ExperimentsClient) listAllCreateRequest(ctx context.Context, options *ExperimentsListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/experiments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	if options != nil && options.Running != nil {
		reqQP.Set("running", strconv.FormatBool(*options.Running))
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *ExperimentsClient) listAllHandleResponse(resp *http.Response) (ExperimentsListAllResponse, error) {
	result := ExperimentsListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentListResult); err != nil {
		return ExperimentsListAllResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *ExperimentsClient) listAllHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAllStatuses - Get a list of statuses of a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) ListAllStatuses(resourceGroupName string, experimentName string, options *ExperimentsListAllStatusesOptions) *ExperimentsListAllStatusesPager {
	return &ExperimentsListAllStatusesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllStatusesCreateRequest(ctx, resourceGroupName, experimentName, options)
		},
		advancer: func(ctx context.Context, resp ExperimentsListAllStatusesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExperimentStatusListResult.NextLink)
		},
	}
}

// listAllStatusesCreateRequest creates the ListAllStatuses request.
func (client *ExperimentsClient) listAllStatusesCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsListAllStatusesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/statuses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllStatusesHandleResponse handles the ListAllStatuses response.
func (client *ExperimentsClient) listAllStatusesHandleResponse(resp *http.Response) (ExperimentsListAllStatusesResponse, error) {
	result := ExperimentsListAllStatusesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentStatusListResult); err != nil {
		return ExperimentsListAllStatusesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAllStatusesHandleError handles the ListAllStatuses error response.
func (client *ExperimentsClient) listAllStatusesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListExecutionDetails - Get a list of execution details of a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) ListExecutionDetails(resourceGroupName string, experimentName string, options *ExperimentsListExecutionDetailsOptions) *ExperimentsListExecutionDetailsPager {
	return &ExperimentsListExecutionDetailsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listExecutionDetailsCreateRequest(ctx, resourceGroupName, experimentName, options)
		},
		advancer: func(ctx context.Context, resp ExperimentsListExecutionDetailsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExperimentExecutionDetailsListResult.NextLink)
		},
	}
}

// listExecutionDetailsCreateRequest creates the ListExecutionDetails request.
func (client *ExperimentsClient) listExecutionDetailsCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsListExecutionDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/executionDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listExecutionDetailsHandleResponse handles the ListExecutionDetails response.
func (client *ExperimentsClient) listExecutionDetailsHandleResponse(resp *http.Response) (ExperimentsListExecutionDetailsResponse, error) {
	result := ExperimentsListExecutionDetailsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentExecutionDetailsListResult); err != nil {
		return ExperimentsListExecutionDetailsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listExecutionDetailsHandleError handles the ListExecutionDetails error response.
func (client *ExperimentsClient) listExecutionDetailsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Start - Start a Experiment resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExperimentsClient) Start(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsStartOptions) (ExperimentsStartResponse, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsStartResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExperimentsStartResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ExperimentsStartResponse{}, client.startHandleError(resp)
	}
	return client.startHandleResponse(resp)
}

// startCreateRequest creates the Start request.
func (client *ExperimentsClient) startCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startHandleResponse handles the Start response.
func (client *ExperimentsClient) startHandleResponse(resp *http.Response) (ExperimentsStartResponse, error) {
	result := ExperimentsStartResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentStartOperationResult); err != nil {
		return ExperimentsStartResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// startHandleError handles the Start error response.
func (client *ExperimentsClient) startHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}