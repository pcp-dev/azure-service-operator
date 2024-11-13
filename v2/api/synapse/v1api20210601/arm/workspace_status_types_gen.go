// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// A workspace
type Workspace_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: Identity of the workspace
	Identity *ManagedIdentity_STATUS `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Workspace resource properties
	Properties *WorkspaceProperties_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The workspace managed identity
type ManagedIdentity_STATUS struct {
	// PrincipalId: The principal ID of the workspace managed identity
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of the workspace managed identity
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of managed identity for the workspace
	Type *ManagedIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The user assigned managed identities.
	UserAssignedIdentities map[string]UserAssignedManagedIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Workspace properties
type WorkspaceProperties_STATUS struct {
	// AdlaResourceId: The ADLA resource ID.
	AdlaResourceId *string `json:"adlaResourceId,omitempty"`

	// AzureADOnlyAuthentication: Enable or Disable AzureADOnlyAuthentication on All Workspace subresource
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// ConnectivityEndpoints: Connectivity endpoints
	ConnectivityEndpoints map[string]string `json:"connectivityEndpoints,omitempty"`

	// CspWorkspaceAdminProperties: Initial workspace AAD admin properties for a CSP subscription
	CspWorkspaceAdminProperties *CspWorkspaceAdminProperties_STATUS `json:"cspWorkspaceAdminProperties,omitempty"`

	// DefaultDataLakeStorage: Workspace default data lake storage account details
	DefaultDataLakeStorage *DataLakeStorageAccountDetails_STATUS `json:"defaultDataLakeStorage,omitempty"`

	// Encryption: The encryption details of the workspace
	Encryption *EncryptionDetails_STATUS `json:"encryption,omitempty"`

	// ExtraProperties: Workspace level configs and feature flags
	ExtraProperties map[string]v1.JSON `json:"extraProperties,omitempty"`

	// ManagedResourceGroupName: Workspace managed resource group. The resource group name uniquely identifies the resource
	// group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be
	// alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty"`

	// ManagedVirtualNetwork: Setting this to 'default' will ensure that all compute for this workspace is in a virtual network
	// managed on behalf of the user.
	ManagedVirtualNetwork *string `json:"managedVirtualNetwork,omitempty"`

	// ManagedVirtualNetworkSettings: Managed Virtual Network Settings
	ManagedVirtualNetworkSettings *ManagedVirtualNetworkSettings_STATUS `json:"managedVirtualNetworkSettings,omitempty"`

	// PrivateEndpointConnections: Private endpoint connections to the workspace
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Resource provisioning state
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Enable or Disable public network access to workspace
	PublicNetworkAccess *WorkspaceProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// PurviewConfiguration: Purview Configuration
	PurviewConfiguration *PurviewConfiguration_STATUS `json:"purviewConfiguration,omitempty"`

	// Settings: Workspace settings
	Settings map[string]v1.JSON `json:"settings,omitempty"`

	// SqlAdministratorLogin: Login for workspace SQL active directory administrator
	SqlAdministratorLogin *string `json:"sqlAdministratorLogin,omitempty"`

	// TrustedServiceBypassEnabled: Is trustedServiceBypassEnabled for the workspace
	TrustedServiceBypassEnabled *bool `json:"trustedServiceBypassEnabled,omitempty"`

	// VirtualNetworkProfile: Virtual Network profile
	VirtualNetworkProfile *VirtualNetworkProfile_STATUS `json:"virtualNetworkProfile,omitempty"`

	// WorkspaceRepositoryConfiguration: Git integration settings
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfiguration_STATUS `json:"workspaceRepositoryConfiguration,omitempty"`

	// WorkspaceUID: The workspace unique identifier
	WorkspaceUID *string `json:"workspaceUID,omitempty"`
}

// Initial workspace AAD admin properties for a CSP subscription
type CspWorkspaceAdminProperties_STATUS struct {
	// InitialWorkspaceAdminObjectId: AAD object ID of initial workspace admin
	InitialWorkspaceAdminObjectId *string `json:"initialWorkspaceAdminObjectId,omitempty"`
}

// Details of the data lake storage account associated with the workspace
type DataLakeStorageAccountDetails_STATUS struct {
	// AccountUrl: Account URL
	AccountUrl *string `json:"accountUrl,omitempty"`

	// CreateManagedPrivateEndpoint: Create managed private endpoint to this storage account or not
	CreateManagedPrivateEndpoint *bool `json:"createManagedPrivateEndpoint,omitempty"`

	// Filesystem: Filesystem name
	Filesystem *string `json:"filesystem,omitempty"`

	// ResourceId: ARM resource Id of this storage account
	ResourceId *string `json:"resourceId,omitempty"`
}

// Details of the encryption associated with the workspace
type EncryptionDetails_STATUS struct {
	// Cmk: Customer Managed Key Details
	Cmk *CustomerManagedKeyDetails_STATUS `json:"cmk,omitempty"`

	// DoubleEncryptionEnabled: Double Encryption enabled
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty"`
}

type ManagedIdentity_Type_STATUS string

const (
	ManagedIdentity_Type_STATUS_None                       = ManagedIdentity_Type_STATUS("None")
	ManagedIdentity_Type_STATUS_SystemAssigned             = ManagedIdentity_Type_STATUS("SystemAssigned")
	ManagedIdentity_Type_STATUS_SystemAssignedUserAssigned = ManagedIdentity_Type_STATUS("SystemAssigned,UserAssigned")
)

// Mapping from string to ManagedIdentity_Type_STATUS
var managedIdentity_Type_STATUS_Values = map[string]ManagedIdentity_Type_STATUS{
	"none":                        ManagedIdentity_Type_STATUS_None,
	"systemassigned":              ManagedIdentity_Type_STATUS_SystemAssigned,
	"systemassigned,userassigned": ManagedIdentity_Type_STATUS_SystemAssignedUserAssigned,
}

// Managed Virtual Network Settings
type ManagedVirtualNetworkSettings_STATUS struct {
	// AllowedAadTenantIdsForLinking: Allowed Aad Tenant Ids For Linking
	AllowedAadTenantIdsForLinking []string `json:"allowedAadTenantIdsForLinking,omitempty"`

	// LinkedAccessCheckOnTargetResource: Linked Access Check On Target Resource
	LinkedAccessCheckOnTargetResource *bool `json:"linkedAccessCheckOnTargetResource,omitempty"`

	// PreventDataExfiltration: Prevent Data Exfiltration
	PreventDataExfiltration *bool `json:"preventDataExfiltration,omitempty"`
}

// A private endpoint connection
type PrivateEndpointConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

// Purview Configuration
type PurviewConfiguration_STATUS struct {
	// PurviewResourceId: Purview Resource ID
	PurviewResourceId *string `json:"purviewResourceId,omitempty"`
}

// User Assigned Managed Identity
type UserAssignedManagedIdentity_STATUS struct {
	// ClientId: The client ID.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal ID.
	PrincipalId *string `json:"principalId,omitempty"`
}

// Virtual Network Profile
type VirtualNetworkProfile_STATUS struct {
	// ComputeSubnetId: Subnet ID used for computes in workspace
	ComputeSubnetId *string `json:"computeSubnetId,omitempty"`
}

type WorkspaceProperties_PublicNetworkAccess_STATUS string

const (
	WorkspaceProperties_PublicNetworkAccess_STATUS_Disabled = WorkspaceProperties_PublicNetworkAccess_STATUS("Disabled")
	WorkspaceProperties_PublicNetworkAccess_STATUS_Enabled  = WorkspaceProperties_PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to WorkspaceProperties_PublicNetworkAccess_STATUS
var workspaceProperties_PublicNetworkAccess_STATUS_Values = map[string]WorkspaceProperties_PublicNetworkAccess_STATUS{
	"disabled": WorkspaceProperties_PublicNetworkAccess_STATUS_Disabled,
	"enabled":  WorkspaceProperties_PublicNetworkAccess_STATUS_Enabled,
}

// Git integration settings
type WorkspaceRepositoryConfiguration_STATUS struct {
	// AccountName: Account name
	AccountName *string `json:"accountName,omitempty"`

	// CollaborationBranch: Collaboration branch
	CollaborationBranch *string `json:"collaborationBranch,omitempty"`

	// HostName: GitHub Enterprise host name. For example: `https://github.mydomain.com`
	HostName *string `json:"hostName,omitempty"`

	// LastCommitId: The last commit ID
	LastCommitId *string `json:"lastCommitId,omitempty"`

	// ProjectName: VSTS project name
	ProjectName *string `json:"projectName,omitempty"`

	// RepositoryName: Repository name
	RepositoryName *string `json:"repositoryName,omitempty"`

	// RootFolder: Root folder to use in the repository
	RootFolder *string `json:"rootFolder,omitempty"`

	// TenantId: The VSTS tenant ID
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of workspace repositoryID configuration. Example WorkspaceVSTSConfiguration, WorkspaceGitHubConfiguration
	Type *string `json:"type,omitempty"`
}

// Details of the customer managed key associated with the workspace
type CustomerManagedKeyDetails_STATUS struct {
	// KekIdentity: Key encryption key
	KekIdentity *KekIdentityProperties_STATUS `json:"kekIdentity,omitempty"`

	// Key: The key object of the workspace
	Key *WorkspaceKeyDetails_STATUS `json:"key,omitempty"`

	// Status: The customer managed key status on the workspace
	Status *string `json:"status,omitempty"`
}

// Key encryption key properties
type KekIdentityProperties_STATUS struct {
	// UseSystemAssignedIdentity: Boolean specifying whether to use system assigned identity or not
	UseSystemAssignedIdentity *v1.JSON `json:"useSystemAssignedIdentity,omitempty"`

	// UserAssignedIdentity: User assigned identity resource Id
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Details of the customer managed key associated with the workspace
type WorkspaceKeyDetails_STATUS struct {
	// KeyVaultUrl: Workspace Key sub-resource key vault url
	KeyVaultUrl *string `json:"keyVaultUrl,omitempty"`

	// Name: Workspace Key sub-resource name
	Name *string `json:"name,omitempty"`
}