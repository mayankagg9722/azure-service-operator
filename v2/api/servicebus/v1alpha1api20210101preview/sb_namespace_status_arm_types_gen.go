// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

// Deprecated version of SBNamespace_STATUS. Use v1beta20210101preview.SBNamespace_STATUS instead
type SBNamespace_STATUSARM struct {
	Id         *string                          `json:"id,omitempty"`
	Identity   *Identity_STATUSARM              `json:"identity,omitempty"`
	Location   *string                          `json:"location,omitempty"`
	Name       *string                          `json:"name,omitempty"`
	Properties *SBNamespaceProperties_STATUSARM `json:"properties,omitempty"`
	Sku        *SBSku_STATUSARM                 `json:"sku,omitempty"`
	SystemData *SystemData_STATUSARM            `json:"systemData,omitempty"`
	Tags       map[string]string                `json:"tags,omitempty"`
	Type       *string                          `json:"type,omitempty"`
}

// Deprecated version of Identity_STATUS. Use v1beta20210101preview.Identity_STATUS instead
type Identity_STATUSARM struct {
	PrincipalId            *string                              `json:"principalId,omitempty"`
	TenantId               *string                              `json:"tenantId,omitempty"`
	Type                   *Identity_STATUS_Type                `json:"type,omitempty"`
	UserAssignedIdentities map[string]DictionaryValue_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of SBNamespaceProperties_STATUS. Use v1beta20210101preview.SBNamespaceProperties_STATUS instead
type SBNamespaceProperties_STATUSARM struct {
	CreatedAt                  *string                                                   `json:"createdAt,omitempty"`
	Encryption                 *Encryption_STATUSARM                                     `json:"encryption,omitempty"`
	MetricId                   *string                                                   `json:"metricId,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                                                   `json:"provisioningState,omitempty"`
	ServiceBusEndpoint         *string                                                   `json:"serviceBusEndpoint,omitempty"`
	Status                     *string                                                   `json:"status,omitempty"`
	UpdatedAt                  *string                                                   `json:"updatedAt,omitempty"`
	ZoneRedundant              *bool                                                     `json:"zoneRedundant,omitempty"`
}

// Deprecated version of SBSku_STATUS. Use v1beta20210101preview.SBSku_STATUS instead
type SBSku_STATUSARM struct {
	Capacity *int               `json:"capacity,omitempty"`
	Name     *SBSku_STATUS_Name `json:"name,omitempty"`
	Tier     *SBSku_STATUS_Tier `json:"tier,omitempty"`
}

// Deprecated version of SystemData_STATUS. Use v1beta20210101preview.SystemData_STATUS instead
type SystemData_STATUSARM struct {
	CreatedAt          *string                               `json:"createdAt,omitempty"`
	CreatedBy          *string                               `json:"createdBy,omitempty"`
	CreatedByType      *SystemData_STATUS_CreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                               `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                               `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_STATUS_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of DictionaryValue_STATUS. Use v1beta20210101preview.DictionaryValue_STATUS instead
type DictionaryValue_STATUSARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

// Deprecated version of Encryption_STATUS. Use v1beta20210101preview.Encryption_STATUS instead
type Encryption_STATUSARM struct {
	KeySource                       *Encryption_STATUS_KeySource   `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties_STATUSARM `json:"keyVaultProperties,omitempty"`
	RequireInfrastructureEncryption *bool                          `json:"requireInfrastructureEncryption,omitempty"`
}

// Deprecated version of Identity_STATUS_Type. Use v1beta20210101preview.Identity_STATUS_Type instead
type Identity_STATUS_Type string

const (
	Identity_STATUS_Type_None                       = Identity_STATUS_Type("None")
	Identity_STATUS_Type_SystemAssigned             = Identity_STATUS_Type("SystemAssigned")
	Identity_STATUS_Type_SystemAssignedUserAssigned = Identity_STATUS_Type("SystemAssigned, UserAssigned")
	Identity_STATUS_Type_UserAssigned               = Identity_STATUS_Type("UserAssigned")
)

// Deprecated version of PrivateEndpointConnection_STATUS_SubResourceEmbedded. Use v1beta20210101preview.PrivateEndpointConnection_STATUS_SubResourceEmbedded instead
type PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM struct {
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`
}

// Deprecated version of SBSku_STATUS_Name. Use v1beta20210101preview.SBSku_STATUS_Name instead
type SBSku_STATUS_Name string

const (
	SBSku_STATUS_Name_Basic    = SBSku_STATUS_Name("Basic")
	SBSku_STATUS_Name_Premium  = SBSku_STATUS_Name("Premium")
	SBSku_STATUS_Name_Standard = SBSku_STATUS_Name("Standard")
)

// Deprecated version of SBSku_STATUS_Tier. Use v1beta20210101preview.SBSku_STATUS_Tier instead
type SBSku_STATUS_Tier string

const (
	SBSku_STATUS_Tier_Basic    = SBSku_STATUS_Tier("Basic")
	SBSku_STATUS_Tier_Premium  = SBSku_STATUS_Tier("Premium")
	SBSku_STATUS_Tier_Standard = SBSku_STATUS_Tier("Standard")
)

// Deprecated version of SystemData_STATUS_CreatedByType. Use v1beta20210101preview.SystemData_STATUS_CreatedByType instead
type SystemData_STATUS_CreatedByType string

const (
	SystemData_STATUS_CreatedByType_Application     = SystemData_STATUS_CreatedByType("Application")
	SystemData_STATUS_CreatedByType_Key             = SystemData_STATUS_CreatedByType("Key")
	SystemData_STATUS_CreatedByType_ManagedIdentity = SystemData_STATUS_CreatedByType("ManagedIdentity")
	SystemData_STATUS_CreatedByType_User            = SystemData_STATUS_CreatedByType("User")
)

// Deprecated version of SystemData_STATUS_LastModifiedByType. Use
// v1beta20210101preview.SystemData_STATUS_LastModifiedByType instead
type SystemData_STATUS_LastModifiedByType string

const (
	SystemData_STATUS_LastModifiedByType_Application     = SystemData_STATUS_LastModifiedByType("Application")
	SystemData_STATUS_LastModifiedByType_Key             = SystemData_STATUS_LastModifiedByType("Key")
	SystemData_STATUS_LastModifiedByType_ManagedIdentity = SystemData_STATUS_LastModifiedByType("ManagedIdentity")
	SystemData_STATUS_LastModifiedByType_User            = SystemData_STATUS_LastModifiedByType("User")
)

// Deprecated version of KeyVaultProperties_STATUS. Use v1beta20210101preview.KeyVaultProperties_STATUS instead
type KeyVaultProperties_STATUSARM struct {
	Identity    *UserAssignedIdentityProperties_STATUSARM `json:"identity,omitempty"`
	KeyName     *string                                   `json:"keyName,omitempty"`
	KeyVaultUri *string                                   `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                                   `json:"keyVersion,omitempty"`
}

// Deprecated version of UserAssignedIdentityProperties_STATUS. Use v1beta20210101preview.UserAssignedIdentityProperties_STATUS instead
type UserAssignedIdentityProperties_STATUSARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}