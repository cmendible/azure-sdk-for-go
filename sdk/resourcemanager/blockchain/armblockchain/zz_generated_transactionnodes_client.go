//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TransactionNodesClient contains the methods for the TransactionNodes group.
// Don't use this type directly, use NewTransactionNodesClient() instead.
type TransactionNodesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTransactionNodesClient creates a new instance of TransactionNodesClient with the specified values.
func NewTransactionNodesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TransactionNodesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TransactionNodesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Create or update the transaction node.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) BeginCreate(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesBeginCreateOptions) (TransactionNodesCreatePollerResponse, error) {
	resp, err := client.create(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesCreatePollerResponse{}, err
	}
	result := TransactionNodesCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TransactionNodesClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return TransactionNodesCreatePollerResponse{}, err
	}
	result.Poller = &TransactionNodesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create or update the transaction node.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) create(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *TransactionNodesClient) createCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.TransactionNode != nil {
		return req, runtime.MarshalAsJSON(req, *options.TransactionNode)
	}
	return req, nil
}

// createHandleError handles the Create error response.
func (client *TransactionNodesClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete the transaction node.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) BeginDelete(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesBeginDeleteOptions) (TransactionNodesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesDeletePollerResponse{}, err
	}
	result := TransactionNodesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TransactionNodesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return TransactionNodesDeletePollerResponse{}, err
	}
	result.Poller = &TransactionNodesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete the transaction node.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) deleteOperation(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TransactionNodesClient) deleteCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *TransactionNodesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get the details of the transaction node.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) Get(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesGetOptions) (TransactionNodesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TransactionNodesClient) getCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
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
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TransactionNodesClient) getHandleResponse(resp *http.Response) (TransactionNodesGetResponse, error) {
	result := TransactionNodesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionNode); err != nil {
		return TransactionNodesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TransactionNodesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Lists the transaction nodes for a blockchain member.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) List(blockchainMemberName string, resourceGroupName string, options *TransactionNodesListOptions) *TransactionNodesListPager {
	return &TransactionNodesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp TransactionNodesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TransactionNodeCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *TransactionNodesClient) listCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *TransactionNodesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
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
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TransactionNodesClient) listHandleResponse(resp *http.Response) (TransactionNodesListResponse, error) {
	result := TransactionNodesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionNodeCollection); err != nil {
		return TransactionNodesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *TransactionNodesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListAPIKeys - List the API keys for the transaction node.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) ListAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesListAPIKeysOptions) (TransactionNodesListAPIKeysResponse, error) {
	req, err := client.listAPIKeysCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesListAPIKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesListAPIKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesListAPIKeysResponse{}, client.listAPIKeysHandleError(resp)
	}
	return client.listAPIKeysHandleResponse(resp)
}

// listAPIKeysCreateRequest creates the ListAPIKeys request.
func (client *TransactionNodesClient) listAPIKeysCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesListAPIKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}/listApiKeys"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAPIKeysHandleResponse handles the ListAPIKeys response.
func (client *TransactionNodesClient) listAPIKeysHandleResponse(resp *http.Response) (TransactionNodesListAPIKeysResponse, error) {
	result := TransactionNodesListAPIKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyCollection); err != nil {
		return TransactionNodesListAPIKeysResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAPIKeysHandleError handles the ListAPIKeys error response.
func (client *TransactionNodesClient) listAPIKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListRegenerateAPIKeys - Regenerate the API keys for the blockchain member.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) ListRegenerateAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesListRegenerateAPIKeysOptions) (TransactionNodesListRegenerateAPIKeysResponse, error) {
	req, err := client.listRegenerateAPIKeysCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesListRegenerateAPIKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesListRegenerateAPIKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesListRegenerateAPIKeysResponse{}, client.listRegenerateAPIKeysHandleError(resp)
	}
	return client.listRegenerateAPIKeysHandleResponse(resp)
}

// listRegenerateAPIKeysCreateRequest creates the ListRegenerateAPIKeys request.
func (client *TransactionNodesClient) listRegenerateAPIKeysCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesListRegenerateAPIKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}/regenerateApiKeys"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.APIKey != nil {
		return req, runtime.MarshalAsJSON(req, *options.APIKey)
	}
	return req, nil
}

// listRegenerateAPIKeysHandleResponse handles the ListRegenerateAPIKeys response.
func (client *TransactionNodesClient) listRegenerateAPIKeysHandleResponse(resp *http.Response) (TransactionNodesListRegenerateAPIKeysResponse, error) {
	result := TransactionNodesListRegenerateAPIKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyCollection); err != nil {
		return TransactionNodesListRegenerateAPIKeysResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listRegenerateAPIKeysHandleError handles the ListRegenerateAPIKeys error response.
func (client *TransactionNodesClient) listRegenerateAPIKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Update the transaction node.
// If the operation fails it returns a generic error.
func (client *TransactionNodesClient) Update(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesUpdateOptions) (TransactionNodesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *TransactionNodesClient) updateCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.TransactionNode != nil {
		return req, runtime.MarshalAsJSON(req, *options.TransactionNode)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TransactionNodesClient) updateHandleResponse(resp *http.Response) (TransactionNodesUpdateResponse, error) {
	result := TransactionNodesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionNode); err != nil {
		return TransactionNodesUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *TransactionNodesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}