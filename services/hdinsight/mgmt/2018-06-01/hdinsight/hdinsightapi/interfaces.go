package hdinsightapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/hdinsight/mgmt/2018-06-01/hdinsight"
	"github.com/Azure/go-autorest/autorest"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.ClusterCreateParametersExtended) (result hdinsight.ClustersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ClustersDeleteFuture, err error)
	ExecuteScriptActions(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.ExecuteScriptActionParameters) (result hdinsight.ClustersExecuteScriptActionsFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.Cluster, err error)
	GetAzureAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, operationID string) (result hdinsight.AsyncOperationResult, err error)
	GetGatewaySettings(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.GatewaySettings, err error)
	List(ctx context.Context) (result hdinsight.ClusterListResultPage, err error)
	ListComplete(ctx context.Context) (result hdinsight.ClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result hdinsight.ClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result hdinsight.ClusterListResultIterator, err error)
	Resize(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.ClusterResizeParameters) (result hdinsight.ClustersResizeFuture, err error)
	RotateDiskEncryptionKey(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.ClusterDiskEncryptionParameters) (result hdinsight.ClustersRotateDiskEncryptionKeyFuture, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.ClusterPatchParameters) (result hdinsight.Cluster, err error)
	UpdateAutoScaleConfiguration(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.AutoscaleConfigurationUpdateParameter) (result hdinsight.ClustersUpdateAutoScaleConfigurationFuture, err error)
	UpdateGatewaySettings(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.UpdateGatewaySettingsParameters) (result hdinsight.ClustersUpdateGatewaySettingsFuture, err error)
	UpdateIdentityCertificate(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.UpdateClusterIdentityCertificateParameters) (result hdinsight.ClustersUpdateIdentityCertificateFuture, err error)
}

var _ ClustersClientAPI = (*hdinsight.ClustersClient)(nil)

// ApplicationsClientAPI contains the set of methods on the ApplicationsClient type.
type ApplicationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters hdinsight.Application) (result hdinsight.ApplicationsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result hdinsight.ApplicationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result hdinsight.Application, err error)
	GetAzureAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, operationID string) (result hdinsight.AsyncOperationResult, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ApplicationListResultPage, err error)
	ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ApplicationListResultIterator, err error)
}

var _ ApplicationsClientAPI = (*hdinsight.ApplicationsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, parameters hdinsight.NameAvailabilityCheckRequestParameters) (result hdinsight.NameAvailabilityCheckResult, err error)
	GetAzureAsyncOperationStatus(ctx context.Context, location string, operationID string) (result hdinsight.AsyncOperationResult, err error)
	GetCapabilities(ctx context.Context, location string) (result hdinsight.CapabilitiesResult, err error)
	ListBillingSpecs(ctx context.Context, location string) (result hdinsight.BillingResponseListResult, err error)
	ListUsages(ctx context.Context, location string) (result hdinsight.UsagesListResult, err error)
	ValidateClusterCreateRequest(ctx context.Context, location string, parameters hdinsight.ClusterCreateRequestValidationParameters) (result hdinsight.ClusterCreateValidationResult, err error)
}

var _ LocationsClientAPI = (*hdinsight.LocationsClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, clusterName string, configurationName string) (result hdinsight.SetString, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ClusterConfigurations, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, configurationName string, parameters map[string]*string) (result hdinsight.ConfigurationsUpdateFuture, err error)
}

var _ ConfigurationsClientAPI = (*hdinsight.ConfigurationsClient)(nil)

// ExtensionsClientAPI contains the set of methods on the ExtensionsClient type.
type ExtensionsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, parameters hdinsight.Extension) (result hdinsight.ExtensionsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, extensionName string) (result hdinsight.ExtensionsDeleteFuture, err error)
	DisableMonitoring(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ExtensionsDisableMonitoringFuture, err error)
	EnableMonitoring(ctx context.Context, resourceGroupName string, clusterName string, parameters hdinsight.ClusterMonitoringRequest) (result hdinsight.ExtensionsEnableMonitoringFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, extensionName string) (result hdinsight.ClusterMonitoringResponse, err error)
	GetAzureAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, operationID string) (result hdinsight.AsyncOperationResult, err error)
	GetMonitoringStatus(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ClusterMonitoringResponse, err error)
}

var _ ExtensionsClientAPI = (*hdinsight.ExtensionsClient)(nil)

// ScriptActionsClientAPI contains the set of methods on the ScriptActionsClient type.
type ScriptActionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, clusterName string, scriptName string) (result autorest.Response, err error)
	GetExecutionAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, operationID string) (result hdinsight.AsyncOperationResult, err error)
	GetExecutionDetail(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string) (result hdinsight.RuntimeScriptActionDetail, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ScriptActionsListPage, err error)
	ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ScriptActionsListIterator, err error)
}

var _ ScriptActionsClientAPI = (*hdinsight.ScriptActionsClient)(nil)

// ScriptExecutionHistoryClientAPI contains the set of methods on the ScriptExecutionHistoryClient type.
type ScriptExecutionHistoryClientAPI interface {
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ScriptActionExecutionHistoryListPage, err error)
	ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ScriptActionExecutionHistoryListIterator, err error)
	Promote(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string) (result autorest.Response, err error)
}

var _ ScriptExecutionHistoryClientAPI = (*hdinsight.ScriptExecutionHistoryClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result hdinsight.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result hdinsight.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*hdinsight.OperationsClient)(nil)

// VirtualMachinesClientAPI contains the set of methods on the VirtualMachinesClient type.
type VirtualMachinesClientAPI interface {
	GetAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, operationID string) (result hdinsight.AsyncOperationResult, err error)
	ListHosts(ctx context.Context, resourceGroupName string, clusterName string) (result hdinsight.ListHostInfo, err error)
	RestartHosts(ctx context.Context, resourceGroupName string, clusterName string, hosts []string) (result hdinsight.VirtualMachinesRestartHostsFuture, err error)
}

var _ VirtualMachinesClientAPI = (*hdinsight.VirtualMachinesClient)(nil)