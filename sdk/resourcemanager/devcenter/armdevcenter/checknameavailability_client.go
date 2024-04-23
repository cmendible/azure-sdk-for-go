//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CheckNameAvailabilityClient contains the methods for the CheckNameAvailability group.
// Don't use this type directly, use NewCheckNameAvailabilityClient() instead.
type CheckNameAvailabilityClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCheckNameAvailabilityClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CheckNameAvailabilityClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CheckNameAvailabilityClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Execute - Check the availability of name for resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - nameAvailabilityRequest - The required parameters for checking if resource name is available.
//   - options - CheckNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityClient.Execute
//     method.
func (client *CheckNameAvailabilityClient) Execute(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, options *CheckNameAvailabilityClientExecuteOptions) (CheckNameAvailabilityClientExecuteResponse, error) {
	var err error
	const operationName = "CheckNameAvailabilityClient.Execute"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.executeCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return CheckNameAvailabilityClientExecuteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CheckNameAvailabilityClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CheckNameAvailabilityClientExecuteResponse{}, err
	}
	resp, err := client.executeHandleResponse(httpResp)
	return resp, err
}

// executeCreateRequest creates the Execute request.
func (client *CheckNameAvailabilityClient) executeCreateRequest(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, options *CheckNameAvailabilityClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevCenter/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, nameAvailabilityRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *CheckNameAvailabilityClient) executeHandleResponse(resp *http.Response) (CheckNameAvailabilityClientExecuteResponse, error) {
	result := CheckNameAvailabilityClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return CheckNameAvailabilityClientExecuteResponse{}, err
	}
	return result, nil
}
