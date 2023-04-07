// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20200601 "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	v1api20200601s "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601storage"
	v20200601 "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ResourceGroupExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ResourceGroupExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20200601.ResourceGroup{},
		&v1api20200601s.ResourceGroup{},
		&v20200601.ResourceGroup{},
		&v20200601s.ResourceGroup{}}
}