// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20211201p "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20211201preview"
	v1api20211201ps "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20211201previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServersAdministratorExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServersAdministratorExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20211201p.FlexibleServersAdministrator{},
		&v1api20211201ps.FlexibleServersAdministrator{}}
}