// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Profile_Spec struct {
	// Identity: Managed service identity (system assigned and/or user assigned identities).
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties required to create a profile.
	Properties *ProfileProperties `json:"properties,omitempty"`

	// Sku: The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the
	// profile.
	Sku *Sku `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Profile_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (profile Profile_Spec) GetAPIVersion() string {
	return "2023-05-01"
}

// GetName returns the Name of the resource
func (profile *Profile_Spec) GetName() string {
	return profile.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles"
func (profile *Profile_Spec) GetType() string {
	return "Microsoft.Cdn/profiles"
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// Type: Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type                   *ManagedServiceIdentityType            `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// The JSON object that contains the properties required to create a profile.
type ProfileProperties struct {
	// OriginResponseTimeoutSeconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the
	// request fails and returns.
	OriginResponseTimeoutSeconds *int `json:"originResponseTimeoutSeconds,omitempty"`
}

// Standard_Verizon = The SKU name for a Standard Verizon CDN profile.
// Premium_Verizon = The SKU name for a Premium Verizon
// CDN profile.
// Custom_Verizon = The SKU name for a Custom Verizon CDN profile.
// Standard_Akamai = The SKU name for an
// Akamai CDN profile.
// Standard_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download scenarios using
// GB based billing model.
// Standard_Microsoft = The SKU name for a Standard Microsoft CDN profile.
// Standard_AzureFrontDoor
// =  The SKU name for an Azure Front Door Standard profile.
// Premium_AzureFrontDoor = The SKU name for an Azure Front Door
// Premium profile.
// Standard_955BandWidth_ChinaCdn = The SKU name for a China CDN profile for VOD, Web and download
// scenarios using 95-5 peak bandwidth billing model.
// Standard_AvgBandWidth_ChinaCdn = The SKU name for a China CDN profile
// for VOD, Web and download scenarios using monthly average peak bandwidth billing model.
// StandardPlus_ChinaCdn = The SKU
// name for a China CDN profile for live-streaming using GB based billing model.
// StandardPlus_955BandWidth_ChinaCdn = The
// SKU name for a China CDN live-streaming profile using 95-5 peak bandwidth billing
// model.
// StandardPlus_AvgBandWidth_ChinaCdn = The SKU name for a China CDN live-streaming profile using monthly average
// peak bandwidth billing model.
type Sku struct {
	// Name: Name of the pricing tier.
	Name *Sku_Name `json:"name,omitempty"`
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityType_None                       = ManagedServiceIdentityType("None")
	ManagedServiceIdentityType_SystemAssigned             = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityType_SystemAssignedUserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
	ManagedServiceIdentityType_UserAssigned               = ManagedServiceIdentityType("UserAssigned")
)

// Mapping from string to ManagedServiceIdentityType
var managedServiceIdentityType_Values = map[string]ManagedServiceIdentityType{
	"none":                         ManagedServiceIdentityType_None,
	"systemassigned":               ManagedServiceIdentityType_SystemAssigned,
	"systemassigned, userassigned": ManagedServiceIdentityType_SystemAssignedUserAssigned,
	"userassigned":                 ManagedServiceIdentityType_UserAssigned,
}

// +kubebuilder:validation:Enum={"Custom_Verizon","Premium_AzureFrontDoor","Premium_Verizon","StandardPlus_955BandWidth_ChinaCdn","StandardPlus_AvgBandWidth_ChinaCdn","StandardPlus_ChinaCdn","Standard_955BandWidth_ChinaCdn","Standard_Akamai","Standard_AvgBandWidth_ChinaCdn","Standard_AzureFrontDoor","Standard_ChinaCdn","Standard_Microsoft","Standard_Verizon"}
type Sku_Name string

const (
	Sku_Name_Custom_Verizon                     = Sku_Name("Custom_Verizon")
	Sku_Name_Premium_AzureFrontDoor             = Sku_Name("Premium_AzureFrontDoor")
	Sku_Name_Premium_Verizon                    = Sku_Name("Premium_Verizon")
	Sku_Name_StandardPlus_955BandWidth_ChinaCdn = Sku_Name("StandardPlus_955BandWidth_ChinaCdn")
	Sku_Name_StandardPlus_AvgBandWidth_ChinaCdn = Sku_Name("StandardPlus_AvgBandWidth_ChinaCdn")
	Sku_Name_StandardPlus_ChinaCdn              = Sku_Name("StandardPlus_ChinaCdn")
	Sku_Name_Standard_955BandWidth_ChinaCdn     = Sku_Name("Standard_955BandWidth_ChinaCdn")
	Sku_Name_Standard_Akamai                    = Sku_Name("Standard_Akamai")
	Sku_Name_Standard_AvgBandWidth_ChinaCdn     = Sku_Name("Standard_AvgBandWidth_ChinaCdn")
	Sku_Name_Standard_AzureFrontDoor            = Sku_Name("Standard_AzureFrontDoor")
	Sku_Name_Standard_ChinaCdn                  = Sku_Name("Standard_ChinaCdn")
	Sku_Name_Standard_Microsoft                 = Sku_Name("Standard_Microsoft")
	Sku_Name_Standard_Verizon                   = Sku_Name("Standard_Verizon")
)

// Mapping from string to Sku_Name
var sku_Name_Values = map[string]Sku_Name{
	"custom_verizon":                     Sku_Name_Custom_Verizon,
	"premium_azurefrontdoor":             Sku_Name_Premium_AzureFrontDoor,
	"premium_verizon":                    Sku_Name_Premium_Verizon,
	"standardplus_955bandwidth_chinacdn": Sku_Name_StandardPlus_955BandWidth_ChinaCdn,
	"standardplus_avgbandwidth_chinacdn": Sku_Name_StandardPlus_AvgBandWidth_ChinaCdn,
	"standardplus_chinacdn":              Sku_Name_StandardPlus_ChinaCdn,
	"standard_955bandwidth_chinacdn":     Sku_Name_Standard_955BandWidth_ChinaCdn,
	"standard_akamai":                    Sku_Name_Standard_Akamai,
	"standard_avgbandwidth_chinacdn":     Sku_Name_Standard_AvgBandWidth_ChinaCdn,
	"standard_azurefrontdoor":            Sku_Name_Standard_AzureFrontDoor,
	"standard_chinacdn":                  Sku_Name_Standard_ChinaCdn,
	"standard_microsoft":                 Sku_Name_Standard_Microsoft,
	"standard_verizon":                   Sku_Name_Standard_Verizon,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
}
