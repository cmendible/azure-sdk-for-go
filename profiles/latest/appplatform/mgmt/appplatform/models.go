// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package appplatform

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/appplatform/mgmt/2020-07-01/appplatform"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppResourceProvisioningState = original.AppResourceProvisioningState

const (
	Creating  AppResourceProvisioningState = original.Creating
	Failed    AppResourceProvisioningState = original.Failed
	Succeeded AppResourceProvisioningState = original.Succeeded
	Updating  AppResourceProvisioningState = original.Updating
)

type ConfigServerState = original.ConfigServerState

const (
	ConfigServerStateDeleted      ConfigServerState = original.ConfigServerStateDeleted
	ConfigServerStateFailed       ConfigServerState = original.ConfigServerStateFailed
	ConfigServerStateNotAvailable ConfigServerState = original.ConfigServerStateNotAvailable
	ConfigServerStateSucceeded    ConfigServerState = original.ConfigServerStateSucceeded
	ConfigServerStateUpdating     ConfigServerState = original.ConfigServerStateUpdating
)

type DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningState

const (
	DeploymentResourceProvisioningStateCreating  DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateCreating
	DeploymentResourceProvisioningStateFailed    DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateFailed
	DeploymentResourceProvisioningStateSucceeded DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateSucceeded
	DeploymentResourceProvisioningStateUpdating  DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateUpdating
)

type DeploymentResourceStatus = original.DeploymentResourceStatus

const (
	DeploymentResourceStatusAllocating DeploymentResourceStatus = original.DeploymentResourceStatusAllocating
	DeploymentResourceStatusCompiling  DeploymentResourceStatus = original.DeploymentResourceStatusCompiling
	DeploymentResourceStatusFailed     DeploymentResourceStatus = original.DeploymentResourceStatusFailed
	DeploymentResourceStatusRunning    DeploymentResourceStatus = original.DeploymentResourceStatusRunning
	DeploymentResourceStatusStopped    DeploymentResourceStatus = original.DeploymentResourceStatusStopped
	DeploymentResourceStatusUnknown    DeploymentResourceStatus = original.DeploymentResourceStatusUnknown
	DeploymentResourceStatusUpgrading  DeploymentResourceStatus = original.DeploymentResourceStatusUpgrading
)

type ManagedIdentityType = original.ManagedIdentityType

const (
	None                       ManagedIdentityType = original.None
	SystemAssigned             ManagedIdentityType = original.SystemAssigned
	SystemAssignedUserAssigned ManagedIdentityType = original.SystemAssignedUserAssigned
	UserAssigned               ManagedIdentityType = original.UserAssigned
)

type MonitoringSettingState = original.MonitoringSettingState

const (
	MonitoringSettingStateFailed       MonitoringSettingState = original.MonitoringSettingStateFailed
	MonitoringSettingStateNotAvailable MonitoringSettingState = original.MonitoringSettingStateNotAvailable
	MonitoringSettingStateSucceeded    MonitoringSettingState = original.MonitoringSettingStateSucceeded
	MonitoringSettingStateUpdating     MonitoringSettingState = original.MonitoringSettingStateUpdating
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating   ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted    ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting   ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed     ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoved      ProvisioningState = original.ProvisioningStateMoved
	ProvisioningStateMoveFailed ProvisioningState = original.ProvisioningStateMoveFailed
	ProvisioningStateMoving     ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateSucceeded  ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating   ProvisioningState = original.ProvisioningStateUpdating
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
	Zone     ResourceSkuRestrictionsType = original.Zone
)

type RuntimeVersion = original.RuntimeVersion

const (
	Java11    RuntimeVersion = original.Java11
	Java8     RuntimeVersion = original.Java8
	NetCore31 RuntimeVersion = original.NetCore31
)

type SkuScaleType = original.SkuScaleType

const (
	SkuScaleTypeAutomatic SkuScaleType = original.SkuScaleTypeAutomatic
	SkuScaleTypeManual    SkuScaleType = original.SkuScaleTypeManual
	SkuScaleTypeNone      SkuScaleType = original.SkuScaleTypeNone
)

type SupportedRuntimePlatform = original.SupportedRuntimePlatform

const (
	Java    SupportedRuntimePlatform = original.Java
	NETCore SupportedRuntimePlatform = original.NETCore
)

type SupportedRuntimeValue = original.SupportedRuntimeValue

const (
	SupportedRuntimeValueJava11    SupportedRuntimeValue = original.SupportedRuntimeValueJava11
	SupportedRuntimeValueJava8     SupportedRuntimeValue = original.SupportedRuntimeValueJava8
	SupportedRuntimeValueNetCore31 SupportedRuntimeValue = original.SupportedRuntimeValueNetCore31
)

type TestKeyType = original.TestKeyType

const (
	Primary   TestKeyType = original.Primary
	Secondary TestKeyType = original.Secondary
)

type UserSourceType = original.UserSourceType

const (
	Jar        UserSourceType = original.Jar
	NetCoreZip UserSourceType = original.NetCoreZip
	Source     UserSourceType = original.Source
)

type AppResource = original.AppResource
type AppResourceCollection = original.AppResourceCollection
type AppResourceCollectionIterator = original.AppResourceCollectionIterator
type AppResourceCollectionPage = original.AppResourceCollectionPage
type AppResourceProperties = original.AppResourceProperties
type AppsClient = original.AppsClient
type AppsCreateOrUpdateFuture = original.AppsCreateOrUpdateFuture
type AppsDeleteFuture = original.AppsDeleteFuture
type AppsUpdateFuture = original.AppsUpdateFuture
type AvailableOperations = original.AvailableOperations
type AvailableOperationsIterator = original.AvailableOperationsIterator
type AvailableOperationsPage = original.AvailableOperationsPage
type AvailableRuntimeVersions = original.AvailableRuntimeVersions
type BaseClient = original.BaseClient
type BindingResource = original.BindingResource
type BindingResourceCollection = original.BindingResourceCollection
type BindingResourceCollectionIterator = original.BindingResourceCollectionIterator
type BindingResourceCollectionPage = original.BindingResourceCollectionPage
type BindingResourceProperties = original.BindingResourceProperties
type BindingsClient = original.BindingsClient
type BindingsCreateOrUpdateFuture = original.BindingsCreateOrUpdateFuture
type BindingsDeleteFuture = original.BindingsDeleteFuture
type BindingsUpdateFuture = original.BindingsUpdateFuture
type CertificateProperties = original.CertificateProperties
type CertificateResource = original.CertificateResource
type CertificateResourceCollection = original.CertificateResourceCollection
type CertificateResourceCollectionIterator = original.CertificateResourceCollectionIterator
type CertificateResourceCollectionPage = original.CertificateResourceCollectionPage
type CertificatesClient = original.CertificatesClient
type CertificatesCreateOrUpdateFuture = original.CertificatesCreateOrUpdateFuture
type CertificatesDeleteFuture = original.CertificatesDeleteFuture
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ClusterResourceProperties = original.ClusterResourceProperties
type ConfigServerGitProperty = original.ConfigServerGitProperty
type ConfigServerProperties = original.ConfigServerProperties
type ConfigServerResource = original.ConfigServerResource
type ConfigServerSettings = original.ConfigServerSettings
type ConfigServerSettingsErrorRecord = original.ConfigServerSettingsErrorRecord
type ConfigServerSettingsValidateResult = original.ConfigServerSettingsValidateResult
type ConfigServersClient = original.ConfigServersClient
type ConfigServersUpdatePatchFuture = original.ConfigServersUpdatePatchFuture
type ConfigServersUpdatePutFuture = original.ConfigServersUpdatePutFuture
type ConfigServersValidateFuture = original.ConfigServersValidateFuture
type CustomDomainProperties = original.CustomDomainProperties
type CustomDomainResource = original.CustomDomainResource
type CustomDomainResourceCollection = original.CustomDomainResourceCollection
type CustomDomainResourceCollectionIterator = original.CustomDomainResourceCollectionIterator
type CustomDomainResourceCollectionPage = original.CustomDomainResourceCollectionPage
type CustomDomainValidatePayload = original.CustomDomainValidatePayload
type CustomDomainValidateResult = original.CustomDomainValidateResult
type CustomDomainsClient = original.CustomDomainsClient
type CustomDomainsCreateOrUpdateFuture = original.CustomDomainsCreateOrUpdateFuture
type CustomDomainsDeleteFuture = original.CustomDomainsDeleteFuture
type CustomDomainsUpdateFuture = original.CustomDomainsUpdateFuture
type DeploymentInstance = original.DeploymentInstance
type DeploymentResource = original.DeploymentResource
type DeploymentResourceCollection = original.DeploymentResourceCollection
type DeploymentResourceCollectionIterator = original.DeploymentResourceCollectionIterator
type DeploymentResourceCollectionPage = original.DeploymentResourceCollectionPage
type DeploymentResourceProperties = original.DeploymentResourceProperties
type DeploymentSettings = original.DeploymentSettings
type DeploymentsClient = original.DeploymentsClient
type DeploymentsCreateOrUpdateFuture = original.DeploymentsCreateOrUpdateFuture
type DeploymentsDeleteFuture = original.DeploymentsDeleteFuture
type DeploymentsRestartFuture = original.DeploymentsRestartFuture
type DeploymentsStartFuture = original.DeploymentsStartFuture
type DeploymentsStopFuture = original.DeploymentsStopFuture
type DeploymentsUpdateFuture = original.DeploymentsUpdateFuture
type Error = original.Error
type GitPatternRepository = original.GitPatternRepository
type LogFileURLResponse = original.LogFileURLResponse
type LogSpecification = original.LogSpecification
type ManagedIdentityProperties = original.ManagedIdentityProperties
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type MonitoringSettingProperties = original.MonitoringSettingProperties
type MonitoringSettingResource = original.MonitoringSettingResource
type MonitoringSettingsClient = original.MonitoringSettingsClient
type MonitoringSettingsUpdatePatchFuture = original.MonitoringSettingsUpdatePatchFuture
type MonitoringSettingsUpdatePutFuture = original.MonitoringSettingsUpdatePutFuture
type NameAvailability = original.NameAvailability
type NameAvailabilityParameters = original.NameAvailabilityParameters
type NetworkProfile = original.NetworkProfile
type NetworkProfileOutboundIPs = original.NetworkProfileOutboundIPs
type OperationDetail = original.OperationDetail
type OperationDisplay = original.OperationDisplay
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PersistentDisk = original.PersistentDisk
type ProxyResource = original.ProxyResource
type RegenerateTestKeyRequestPayload = original.RegenerateTestKeyRequestPayload
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuCollection = original.ResourceSkuCollection
type ResourceSkuCollectionIterator = original.ResourceSkuCollectionIterator
type ResourceSkuCollectionPage = original.ResourceSkuCollectionPage
type ResourceSkuLocationInfo = original.ResourceSkuLocationInfo
type ResourceSkuRestrictionInfo = original.ResourceSkuRestrictionInfo
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkuZoneDetails = original.ResourceSkuZoneDetails
type ResourceUploadDefinition = original.ResourceUploadDefinition
type RuntimeVersionsClient = original.RuntimeVersionsClient
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceListIterator = original.ServiceResourceListIterator
type ServiceResourceListPage = original.ServiceResourceListPage
type ServiceSpecification = original.ServiceSpecification
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesUpdateFuture = original.ServicesUpdateFuture
type Sku = original.Sku
type SkuCapacity = original.SkuCapacity
type SkusClient = original.SkusClient
type SupportedRuntimeVersion = original.SupportedRuntimeVersion
type TemporaryDisk = original.TemporaryDisk
type TestKeys = original.TestKeys
type TrackedResource = original.TrackedResource
type UserSourceInfo = original.UserSourceInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAppResourceCollectionIterator(page AppResourceCollectionPage) AppResourceCollectionIterator {
	return original.NewAppResourceCollectionIterator(page)
}
func NewAppResourceCollectionPage(cur AppResourceCollection, getNextPage func(context.Context, AppResourceCollection) (AppResourceCollection, error)) AppResourceCollectionPage {
	return original.NewAppResourceCollectionPage(cur, getNextPage)
}
func NewAppsClient(subscriptionID string) AppsClient {
	return original.NewAppsClient(subscriptionID)
}
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return original.NewAppsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailableOperationsIterator(page AvailableOperationsPage) AvailableOperationsIterator {
	return original.NewAvailableOperationsIterator(page)
}
func NewAvailableOperationsPage(cur AvailableOperations, getNextPage func(context.Context, AvailableOperations) (AvailableOperations, error)) AvailableOperationsPage {
	return original.NewAvailableOperationsPage(cur, getNextPage)
}
func NewBindingResourceCollectionIterator(page BindingResourceCollectionPage) BindingResourceCollectionIterator {
	return original.NewBindingResourceCollectionIterator(page)
}
func NewBindingResourceCollectionPage(cur BindingResourceCollection, getNextPage func(context.Context, BindingResourceCollection) (BindingResourceCollection, error)) BindingResourceCollectionPage {
	return original.NewBindingResourceCollectionPage(cur, getNextPage)
}
func NewBindingsClient(subscriptionID string) BindingsClient {
	return original.NewBindingsClient(subscriptionID)
}
func NewBindingsClientWithBaseURI(baseURI string, subscriptionID string) BindingsClient {
	return original.NewBindingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCertificateResourceCollectionIterator(page CertificateResourceCollectionPage) CertificateResourceCollectionIterator {
	return original.NewCertificateResourceCollectionIterator(page)
}
func NewCertificateResourceCollectionPage(cur CertificateResourceCollection, getNextPage func(context.Context, CertificateResourceCollection) (CertificateResourceCollection, error)) CertificateResourceCollectionPage {
	return original.NewCertificateResourceCollectionPage(cur, getNextPage)
}
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return original.NewCertificatesClient(subscriptionID)
}
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return original.NewCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigServersClient(subscriptionID string) ConfigServersClient {
	return original.NewConfigServersClient(subscriptionID)
}
func NewConfigServersClientWithBaseURI(baseURI string, subscriptionID string) ConfigServersClient {
	return original.NewConfigServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewCustomDomainResourceCollectionIterator(page CustomDomainResourceCollectionPage) CustomDomainResourceCollectionIterator {
	return original.NewCustomDomainResourceCollectionIterator(page)
}
func NewCustomDomainResourceCollectionPage(cur CustomDomainResourceCollection, getNextPage func(context.Context, CustomDomainResourceCollection) (CustomDomainResourceCollection, error)) CustomDomainResourceCollectionPage {
	return original.NewCustomDomainResourceCollectionPage(cur, getNextPage)
}
func NewCustomDomainsClient(subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClient(subscriptionID)
}
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentResourceCollectionIterator(page DeploymentResourceCollectionPage) DeploymentResourceCollectionIterator {
	return original.NewDeploymentResourceCollectionIterator(page)
}
func NewDeploymentResourceCollectionPage(cur DeploymentResourceCollection, getNextPage func(context.Context, DeploymentResourceCollection) (DeploymentResourceCollection, error)) DeploymentResourceCollectionPage {
	return original.NewDeploymentResourceCollectionPage(cur, getNextPage)
}
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitoringSettingsClient(subscriptionID string) MonitoringSettingsClient {
	return original.NewMonitoringSettingsClient(subscriptionID)
}
func NewMonitoringSettingsClientWithBaseURI(baseURI string, subscriptionID string) MonitoringSettingsClient {
	return original.NewMonitoringSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkuCollectionIterator(page ResourceSkuCollectionPage) ResourceSkuCollectionIterator {
	return original.NewResourceSkuCollectionIterator(page)
}
func NewResourceSkuCollectionPage(cur ResourceSkuCollection, getNextPage func(context.Context, ResourceSkuCollection) (ResourceSkuCollection, error)) ResourceSkuCollectionPage {
	return original.NewResourceSkuCollectionPage(cur, getNextPage)
}
func NewRuntimeVersionsClient(subscriptionID string) RuntimeVersionsClient {
	return original.NewRuntimeVersionsClient(subscriptionID)
}
func NewRuntimeVersionsClientWithBaseURI(baseURI string, subscriptionID string) RuntimeVersionsClient {
	return original.NewRuntimeVersionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceResourceListIterator(page ServiceResourceListPage) ServiceResourceListIterator {
	return original.NewServiceResourceListIterator(page)
}
func NewServiceResourceListPage(cur ServiceResourceList, getNextPage func(context.Context, ServiceResourceList) (ServiceResourceList, error)) ServiceResourceListPage {
	return original.NewServiceResourceListPage(cur, getNextPage)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAppResourceProvisioningStateValues() []AppResourceProvisioningState {
	return original.PossibleAppResourceProvisioningStateValues()
}
func PossibleConfigServerStateValues() []ConfigServerState {
	return original.PossibleConfigServerStateValues()
}
func PossibleDeploymentResourceProvisioningStateValues() []DeploymentResourceProvisioningState {
	return original.PossibleDeploymentResourceProvisioningStateValues()
}
func PossibleDeploymentResourceStatusValues() []DeploymentResourceStatus {
	return original.PossibleDeploymentResourceStatusValues()
}
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return original.PossibleManagedIdentityTypeValues()
}
func PossibleMonitoringSettingStateValues() []MonitoringSettingState {
	return original.PossibleMonitoringSettingStateValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return original.PossibleResourceSkuRestrictionsReasonCodeValues()
}
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return original.PossibleResourceSkuRestrictionsTypeValues()
}
func PossibleRuntimeVersionValues() []RuntimeVersion {
	return original.PossibleRuntimeVersionValues()
}
func PossibleSkuScaleTypeValues() []SkuScaleType {
	return original.PossibleSkuScaleTypeValues()
}
func PossibleSupportedRuntimePlatformValues() []SupportedRuntimePlatform {
	return original.PossibleSupportedRuntimePlatformValues()
}
func PossibleSupportedRuntimeValueValues() []SupportedRuntimeValue {
	return original.PossibleSupportedRuntimeValueValues()
}
func PossibleTestKeyTypeValues() []TestKeyType {
	return original.PossibleTestKeyTypeValues()
}
func PossibleUserSourceTypeValues() []UserSourceType {
	return original.PossibleUserSourceTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}