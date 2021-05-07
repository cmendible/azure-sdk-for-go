// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// HTTPPoller provides polling facilities until the operation reaches a terminal state.
type HTTPPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final *http.Response will be returned.
	FinalResponse(ctx context.Context) (*http.Response, error)
}

type httpPoller struct {
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

func (p *httpPoller) Done() bool {
	return p.pt.Done()
}

func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *httpPoller) FinalResponse(ctx context.Context) (*http.Response, error) {
	return p.pt.FinalResponse(ctx, p.pipeline, nil)
}

func (p *httpPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *httpPoller) pollUntilDone(ctx context.Context, freq time.Duration) (*http.Response, error) {
	return p.pt.PollUntilDone(ctx, freq, p.pipeline, nil)
}

// PrivateEndpointConnectionPoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final PrivateEndpointConnectionResponse will be returned.
	FinalResponse(ctx context.Context) (PrivateEndpointConnectionResponse, error)
}

type privateEndpointConnectionPoller struct {
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

func (p *privateEndpointConnectionPoller) Done() bool {
	return p.pt.Done()
}

func (p *privateEndpointConnectionPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *privateEndpointConnectionPoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionResponse, error) {
	respType := PrivateEndpointConnectionResponse{PrivateEndpointConnection: &PrivateEndpointConnection{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *privateEndpointConnectionPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *privateEndpointConnectionPoller) pollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionResponse, error) {
	respType := PrivateEndpointConnectionResponse{PrivateEndpointConnection: &PrivateEndpointConnection{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, p.pipeline, respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// VaultPoller provides polling facilities until the operation reaches a terminal state.
type VaultPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final VaultResponse will be returned.
	FinalResponse(ctx context.Context) (VaultResponse, error)
}

type vaultPoller struct {
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

func (p *vaultPoller) Done() bool {
	return p.pt.Done()
}

func (p *vaultPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *vaultPoller) FinalResponse(ctx context.Context) (VaultResponse, error) {
	respType := VaultResponse{Vault: &Vault{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.Vault)
	if err != nil {
		return VaultResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *vaultPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *vaultPoller) pollUntilDone(ctx context.Context, freq time.Duration) (VaultResponse, error) {
	respType := VaultResponse{Vault: &Vault{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, p.pipeline, respType.Vault)
	if err != nil {
		return VaultResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}