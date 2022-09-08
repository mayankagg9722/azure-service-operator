// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220501

type ConfigurationStore_STATUSARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The managed identity information, if configured.
	Identity *ResourceIdentity_STATUSARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a configuration store.
	Properties *ConfigurationStoreProperties_STATUSARM `json:"properties,omitempty"`

	// Sku: The sku of the configuration store.
	Sku *Sku_STATUSARM `json:"sku,omitempty"`

	// SystemData: Resource system metadata.
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ConfigurationStoreProperties_STATUSARM struct {
	// CreateMode: Indicates whether the configuration store need to be recovered.
	CreateMode *ConfigurationStoreProperties_STATUS_CreateMode `json:"createMode,omitempty"`

	// CreationDate: The creation date of configuration store.
	CreationDate *string `json:"creationDate,omitempty"`

	// DisableLocalAuth: Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// EnablePurgeProtection: Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`

	// Encryption: The encryption settings of the configuration store.
	Encryption *EncryptionProperties_STATUSARM `json:"encryption,omitempty"`

	// Endpoint: The DNS endpoint where the configuration store API will be available.
	Endpoint *string `json:"endpoint,omitempty"`

	// PrivateEndpointConnections: The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []PrivateEndpointConnectionReference_STATUSARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The provisioning state of the configuration store.
	ProvisioningState *ConfigurationStoreProperties_STATUS_ProvisioningState `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Control permission for data plane traffic coming from public networks while private endpoint is
	// enabled.
	PublicNetworkAccess *ConfigurationStoreProperties_STATUS_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SoftDeleteRetentionInDays: The amount of time in days that the configuration store will be retained when it is soft
	// deleted.
	SoftDeleteRetentionInDays *int `json:"softDeleteRetentionInDays,omitempty"`
}

type ResourceIdentity_STATUSARM struct {
	// PrincipalId: The principal id of the identity. This property will only be provided for a system-assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant id associated with the resource's identity. This property will only be provided for a
	// system-assigned identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created
	// identity and a set of user-assigned identities. The type 'None' will remove any identities.
	Type *ResourceIdentity_STATUS_Type `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user-assigned identities associated with the resource. The user-assigned identity
	// dictionary keys will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]UserIdentity_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

type Sku_STATUSARM struct {
	// Name: The SKU name of the configuration store.
	Name *string `json:"name,omitempty"`
}

type SystemData_STATUSARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_STATUS_CreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_STATUS_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type EncryptionProperties_STATUSARM struct {
	// KeyVaultProperties: Key vault properties.
	KeyVaultProperties *KeyVaultProperties_STATUSARM `json:"keyVaultProperties,omitempty"`
}

type PrivateEndpointConnectionReference_STATUSARM struct {
	// Id: The resource ID.
	Id *string `json:"id,omitempty"`
}

type ResourceIdentity_STATUS_Type string

const (
	ResourceIdentity_STATUS_Type_None                       = ResourceIdentity_STATUS_Type("None")
	ResourceIdentity_STATUS_Type_SystemAssigned             = ResourceIdentity_STATUS_Type("SystemAssigned")
	ResourceIdentity_STATUS_Type_SystemAssignedUserAssigned = ResourceIdentity_STATUS_Type("SystemAssigned, UserAssigned")
	ResourceIdentity_STATUS_Type_UserAssigned               = ResourceIdentity_STATUS_Type("UserAssigned")
)

type SystemData_STATUS_CreatedByType string

const (
	SystemData_STATUS_CreatedByType_Application     = SystemData_STATUS_CreatedByType("Application")
	SystemData_STATUS_CreatedByType_Key             = SystemData_STATUS_CreatedByType("Key")
	SystemData_STATUS_CreatedByType_ManagedIdentity = SystemData_STATUS_CreatedByType("ManagedIdentity")
	SystemData_STATUS_CreatedByType_User            = SystemData_STATUS_CreatedByType("User")
)

type SystemData_STATUS_LastModifiedByType string

const (
	SystemData_STATUS_LastModifiedByType_Application     = SystemData_STATUS_LastModifiedByType("Application")
	SystemData_STATUS_LastModifiedByType_Key             = SystemData_STATUS_LastModifiedByType("Key")
	SystemData_STATUS_LastModifiedByType_ManagedIdentity = SystemData_STATUS_LastModifiedByType("ManagedIdentity")
	SystemData_STATUS_LastModifiedByType_User            = SystemData_STATUS_LastModifiedByType("User")
)

type UserIdentity_STATUSARM struct {
	// ClientId: The client ID of the user-assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal ID of the user-assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

type KeyVaultProperties_STATUSARM struct {
	// IdentityClientId: The client id of the identity which will be used to access key vault.
	IdentityClientId *string `json:"identityClientId,omitempty"`

	// KeyIdentifier: The URI of the key vault key used to encrypt data.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
