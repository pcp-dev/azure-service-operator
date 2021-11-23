// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers_firewallRules
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersFirewallRules_Spec `json:"spec,omitempty"`
	Status            FirewallRule_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersFirewallRule{}

// GetConditions returns the conditions of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetConditions() conditions.Conditions {
	return flexibleServersFirewallRule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) SetConditions(conditions conditions.Conditions) {
	flexibleServersFirewallRule.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersFirewallRule{}

// ConvertFrom populates our FlexibleServersFirewallRule from the provided hub FlexibleServersFirewallRule
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20210501storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected storage:dbformysql/v1alpha1api20210501storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return flexibleServersFirewallRule.AssignPropertiesFromFlexibleServersFirewallRule(source)
}

// ConvertTo populates the provided hub FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210501storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected storage:dbformysql/v1alpha1api20210501storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return flexibleServersFirewallRule.AssignPropertiesToFlexibleServersFirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-dbformysql-azure-com-v1alpha1api20210501-flexibleserversfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210501,name=default.v1alpha1api20210501.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &FlexibleServersFirewallRule{}

// Default applies defaults to the FlexibleServersFirewallRule resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) Default() {
	flexibleServersFirewallRule.defaultImpl()
	var temp interface{} = flexibleServersFirewallRule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) defaultAzureName() {
	if flexibleServersFirewallRule.Spec.AzureName == "" {
		flexibleServersFirewallRule.Spec.AzureName = flexibleServersFirewallRule.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersFirewallRule resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) defaultImpl() {
	flexibleServersFirewallRule.defaultAzureName()
}

var _ genruntime.KubernetesResource = &FlexibleServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) AzureName() string {
	return flexibleServersFirewallRule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (flexibleServersFirewallRule FlexibleServersFirewallRule) GetAPIVersion() string {
	return "2021-05-01"
}

// GetResourceKind returns the kind of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &flexibleServersFirewallRule.Spec
}

// GetStatus returns the status of this resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &flexibleServersFirewallRule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/firewallRules"
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FirewallRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(flexibleServersFirewallRule.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: flexibleServersFirewallRule.Namespace,
		Name:      flexibleServersFirewallRule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FirewallRule_Status); ok {
		flexibleServersFirewallRule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FirewallRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	flexibleServersFirewallRule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformysql-azure-com-v1alpha1api20210501-flexibleserversfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210501,name=validate.v1alpha1api20210501.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &FlexibleServersFirewallRule{}

// ValidateCreate validates the creation of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) ValidateCreate() error {
	validations := flexibleServersFirewallRule.createValidations()
	var temp interface{} = flexibleServersFirewallRule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) ValidateDelete() error {
	validations := flexibleServersFirewallRule.deleteValidations()
	var temp interface{} = flexibleServersFirewallRule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) ValidateUpdate(old runtime.Object) error {
	validations := flexibleServersFirewallRule.updateValidations()
	var temp interface{} = flexibleServersFirewallRule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) createValidations() []func() error {
	return []func() error{flexibleServersFirewallRule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return flexibleServersFirewallRule.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&flexibleServersFirewallRule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromFlexibleServersFirewallRule populates our FlexibleServersFirewallRule from the provided source FlexibleServersFirewallRule
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) AssignPropertiesFromFlexibleServersFirewallRule(source *v1alpha1api20210501storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	flexibleServersFirewallRule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersFirewallRules_Spec
	err := spec.AssignPropertiesFromFlexibleServersFirewallRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromFlexibleServersFirewallRulesSpec()")
	}
	flexibleServersFirewallRule.Spec = spec

	// Status
	var status FirewallRule_Status
	err = status.AssignPropertiesFromFirewallRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromFirewallRuleStatus()")
	}
	flexibleServersFirewallRule.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule populates the provided destination FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) AssignPropertiesToFlexibleServersFirewallRule(destination *v1alpha1api20210501storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *flexibleServersFirewallRule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210501storage.FlexibleServersFirewallRules_Spec
	err := flexibleServersFirewallRule.Spec.AssignPropertiesToFlexibleServersFirewallRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToFlexibleServersFirewallRulesSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210501storage.FirewallRule_Status
	err = flexibleServersFirewallRule.Status.AssignPropertiesToFirewallRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToFirewallRuleStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: flexibleServersFirewallRule.Spec.OriginalVersion(),
		Kind:    "FlexibleServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers_firewallRules
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

type FirewallRule_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//EndIpAddress: The end IP address of the server firewall rule. Must be IPv4
	//format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//StartIpAddress: The start IP address of the server firewall rule. Must be IPv4
	//format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`

	//SystemData: The system metadata relating to this resource.
	SystemData *SystemData_Status `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FirewallRule_Status{}

// ConvertStatusFrom populates our FirewallRule_Status from the provided source
func (firewallRuleStatus *FirewallRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210501storage.FirewallRule_Status)
	if ok {
		// Populate our instance from source
		return firewallRuleStatus.AssignPropertiesFromFirewallRuleStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210501storage.FirewallRule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = firewallRuleStatus.AssignPropertiesFromFirewallRuleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FirewallRule_Status
func (firewallRuleStatus *FirewallRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210501storage.FirewallRule_Status)
	if ok {
		// Populate destination from our instance
		return firewallRuleStatus.AssignPropertiesToFirewallRuleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210501storage.FirewallRule_Status{}
	err := firewallRuleStatus.AssignPropertiesToFirewallRuleStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &FirewallRule_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (firewallRuleStatus *FirewallRule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FirewallRule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (firewallRuleStatus *FirewallRule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FirewallRule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FirewallRule_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		firewallRuleStatus.EndIpAddress = &typedInput.Properties.EndIpAddress
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		firewallRuleStatus.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		firewallRuleStatus.Name = &name
	}

	// Set property ‘StartIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		firewallRuleStatus.StartIpAddress = &typedInput.Properties.StartIpAddress
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_Status
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		firewallRuleStatus.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		firewallRuleStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromFirewallRuleStatus populates our FirewallRule_Status from the provided source FirewallRule_Status
func (firewallRuleStatus *FirewallRule_Status) AssignPropertiesFromFirewallRuleStatus(source *v1alpha1api20210501storage.FirewallRule_Status) error {

	// Conditions
	firewallRuleStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIpAddress
	firewallRuleStatus.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

	// Id
	firewallRuleStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	firewallRuleStatus.Name = genruntime.ClonePointerToString(source.Name)

	// StartIpAddress
	firewallRuleStatus.StartIpAddress = genruntime.ClonePointerToString(source.StartIpAddress)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesFromSystemDataStatus()")
		}
		firewallRuleStatus.SystemData = &systemDatum
	} else {
		firewallRuleStatus.SystemData = nil
	}

	// Type
	firewallRuleStatus.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToFirewallRuleStatus populates the provided destination FirewallRule_Status from our FirewallRule_Status
func (firewallRuleStatus *FirewallRule_Status) AssignPropertiesToFirewallRuleStatus(destination *v1alpha1api20210501storage.FirewallRule_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(firewallRuleStatus.Conditions)

	// EndIpAddress
	destination.EndIpAddress = genruntime.ClonePointerToString(firewallRuleStatus.EndIpAddress)

	// Id
	destination.Id = genruntime.ClonePointerToString(firewallRuleStatus.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(firewallRuleStatus.Name)

	// StartIpAddress
	destination.StartIpAddress = genruntime.ClonePointerToString(firewallRuleStatus.StartIpAddress)

	// SystemData
	if firewallRuleStatus.SystemData != nil {
		var systemDatum v1alpha1api20210501storage.SystemData_Status
		err := (*firewallRuleStatus.SystemData).AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesToSystemDataStatus()")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(firewallRuleStatus.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2021-05-01"}
type FlexibleServersFirewallRulesSpecAPIVersion string

const FlexibleServersFirewallRulesSpecAPIVersion20210501 = FlexibleServersFirewallRulesSpecAPIVersion("2021-05-01")

type FlexibleServersFirewallRules_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	//EndIpAddress: The end IP address of the server firewall rule. Must be IPv4
	//format.
	EndIpAddress string `json:"endIpAddress"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner" kind:"FlexibleServer"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	//StartIpAddress: The start IP address of the server firewall rule. Must be IPv4
	//format.
	StartIpAddress string `json:"startIpAddress"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &FlexibleServersFirewallRules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if flexibleServersFirewallRulesSpec == nil {
		return nil, nil
	}
	var result FlexibleServersFirewallRules_SpecARM

	// Set property ‘Location’:
	if flexibleServersFirewallRulesSpec.Location != nil {
		location := *flexibleServersFirewallRulesSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	result.Properties.EndIpAddress = flexibleServersFirewallRulesSpec.EndIpAddress
	result.Properties.StartIpAddress = flexibleServersFirewallRulesSpec.StartIpAddress

	// Set property ‘Tags’:
	if flexibleServersFirewallRulesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range flexibleServersFirewallRulesSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersFirewallRules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersFirewallRules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersFirewallRules_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	flexibleServersFirewallRulesSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	flexibleServersFirewallRulesSpec.EndIpAddress = typedInput.Properties.EndIpAddress

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		flexibleServersFirewallRulesSpec.Location = &location
	}

	// Set property ‘Owner’:
	flexibleServersFirewallRulesSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘StartIpAddress’:
	// copying flattened property:
	flexibleServersFirewallRulesSpec.StartIpAddress = typedInput.Properties.StartIpAddress

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		flexibleServersFirewallRulesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			flexibleServersFirewallRulesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRules_Spec{}

// ConvertSpecFrom populates our FlexibleServersFirewallRules_Spec from the provided source
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210501storage.FlexibleServersFirewallRules_Spec)
	if ok {
		// Populate our instance from source
		return flexibleServersFirewallRulesSpec.AssignPropertiesFromFlexibleServersFirewallRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210501storage.FlexibleServersFirewallRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = flexibleServersFirewallRulesSpec.AssignPropertiesFromFlexibleServersFirewallRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRules_Spec
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210501storage.FlexibleServersFirewallRules_Spec)
	if ok {
		// Populate destination from our instance
		return flexibleServersFirewallRulesSpec.AssignPropertiesToFlexibleServersFirewallRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210501storage.FlexibleServersFirewallRules_Spec{}
	err := flexibleServersFirewallRulesSpec.AssignPropertiesToFlexibleServersFirewallRulesSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromFlexibleServersFirewallRulesSpec populates our FlexibleServersFirewallRules_Spec from the provided source FlexibleServersFirewallRules_Spec
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) AssignPropertiesFromFlexibleServersFirewallRulesSpec(source *v1alpha1api20210501storage.FlexibleServersFirewallRules_Spec) error {

	// AzureName
	flexibleServersFirewallRulesSpec.AzureName = source.AzureName

	// EndIpAddress
	if source.EndIpAddress != nil {
		flexibleServersFirewallRulesSpec.EndIpAddress = *source.EndIpAddress
	} else {
		flexibleServersFirewallRulesSpec.EndIpAddress = ""
	}

	// Location
	flexibleServersFirewallRulesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	flexibleServersFirewallRulesSpec.Owner = source.Owner.Copy()

	// StartIpAddress
	if source.StartIpAddress != nil {
		flexibleServersFirewallRulesSpec.StartIpAddress = *source.StartIpAddress
	} else {
		flexibleServersFirewallRulesSpec.StartIpAddress = ""
	}

	// Tags
	flexibleServersFirewallRulesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRulesSpec populates the provided destination FlexibleServersFirewallRules_Spec from our FlexibleServersFirewallRules_Spec
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) AssignPropertiesToFlexibleServersFirewallRulesSpec(destination *v1alpha1api20210501storage.FlexibleServersFirewallRules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = flexibleServersFirewallRulesSpec.AzureName

	// EndIpAddress
	endIpAddress := flexibleServersFirewallRulesSpec.EndIpAddress
	destination.EndIpAddress = &endIpAddress

	// Location
	destination.Location = genruntime.ClonePointerToString(flexibleServersFirewallRulesSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = flexibleServersFirewallRulesSpec.OriginalVersion()

	// Owner
	destination.Owner = flexibleServersFirewallRulesSpec.Owner.Copy()

	// StartIpAddress
	startIpAddress := flexibleServersFirewallRulesSpec.StartIpAddress
	destination.StartIpAddress = &startIpAddress

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(flexibleServersFirewallRulesSpec.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) SetAzureName(azureName string) {
	flexibleServersFirewallRulesSpec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
