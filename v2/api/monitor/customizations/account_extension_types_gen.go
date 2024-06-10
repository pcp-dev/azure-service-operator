// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20230403 "github.com/Azure/azure-service-operator/v2/api/monitor/v1api20230403"
	storage "github.com/Azure/azure-service-operator/v2/api/monitor/v1api20230403/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type AccountExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *AccountExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20230403.Account{},
		&storage.Account{}}
}