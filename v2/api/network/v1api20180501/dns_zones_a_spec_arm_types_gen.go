// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DnsZones_A_Spec_ARM struct {
	// Etag: The etag of the record set.
	Etag *string `json:"etag,omitempty"`
	Name string  `json:"name,omitempty"`

	// Properties: The properties of the record set.
	Properties *RecordSetProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DnsZones_A_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01"
func (zonesA DnsZones_A_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (zonesA *DnsZones_A_Spec_ARM) GetName() string {
	return zonesA.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsZones/A"
func (zonesA *DnsZones_A_Spec_ARM) GetType() string {
	return "Microsoft.Network/dnsZones/A"
}
