// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

// Deprecated version of Namespaces_Topic_STATUS. Use v1beta20210101preview.Namespaces_Topic_STATUS instead
type Namespaces_Topic_STATUS_ARM struct {
	Id         *string                       `json:"id,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *SBTopicProperties_STATUS_ARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUS_ARM        `json:"systemData,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}

// Deprecated version of SBTopicProperties_STATUS. Use v1beta20210101preview.SBTopicProperties_STATUS instead
type SBTopicProperties_STATUS_ARM struct {
	AccessedAt                          *string                         `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                         `json:"autoDeleteOnIdle,omitempty"`
	CountDetails                        *MessageCountDetails_STATUS_ARM `json:"countDetails,omitempty"`
	CreatedAt                           *string                         `json:"createdAt,omitempty"`
	DefaultMessageTimeToLive            *string                         `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                         `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                           `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                           `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                           `json:"enablePartitioning,omitempty"`
	MaxSizeInMegabytes                  *int                            `json:"maxSizeInMegabytes,omitempty"`
	RequiresDuplicateDetection          *bool                           `json:"requiresDuplicateDetection,omitempty"`
	SizeInBytes                         *int                            `json:"sizeInBytes,omitempty"`
	Status                              *EntityStatus_STATUS            `json:"status,omitempty"`
	SubscriptionCount                   *int                            `json:"subscriptionCount,omitempty"`
	SupportOrdering                     *bool                           `json:"supportOrdering,omitempty"`
	UpdatedAt                           *string                         `json:"updatedAt,omitempty"`
}