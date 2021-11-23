// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_VirtualNetworksVirtualNetworkPeering_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering to hub returns original",
		prop.ForAll(RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering tests if a specific instance of VirtualNetworksVirtualNetworkPeering round trips to the hub storage version and back losslessly
func RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1alpha1api20201101storage.VirtualNetworksVirtualNetworkPeering
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual VirtualNetworksVirtualNetworkPeering
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworksVirtualNetworkPeering_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering to VirtualNetworksVirtualNetworkPeering via AssignPropertiesToVirtualNetworksVirtualNetworkPeering & AssignPropertiesFromVirtualNetworksVirtualNetworkPeering returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering tests if a specific instance of VirtualNetworksVirtualNetworkPeering can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.VirtualNetworksVirtualNetworkPeering
	err := copied.AssignPropertiesToVirtualNetworksVirtualNetworkPeering(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeering
	err = actual.AssignPropertiesFromVirtualNetworksVirtualNetworkPeering(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworksVirtualNetworkPeering_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeering via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeering
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworksVirtualNetworkPeering instances for property testing - lazily instantiated by
//VirtualNetworksVirtualNetworkPeeringGenerator()
var virtualNetworksVirtualNetworkPeeringGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeeringGenerator returns a generator of VirtualNetworksVirtualNetworkPeering instances for property testing.
func VirtualNetworksVirtualNetworkPeeringGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeeringGenerator != nil {
		return virtualNetworksVirtualNetworkPeeringGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering(generators)
	virtualNetworksVirtualNetworkPeeringGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering{}), generators)

	return virtualNetworksVirtualNetworkPeeringGenerator
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualNetworksVirtualNetworkPeeringsSpecGenerator()
	gens["Status"] = VirtualNetworkPeeringStatusGenerator()
}

func Test_VirtualNetworkPeering_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworkPeering_Status to VirtualNetworkPeering_Status via AssignPropertiesToVirtualNetworkPeeringStatus & AssignPropertiesFromVirtualNetworkPeeringStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkPeeringStatus, VirtualNetworkPeeringStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkPeeringStatus tests if a specific instance of VirtualNetworkPeering_Status can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkPeeringStatus(subject VirtualNetworkPeering_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.VirtualNetworkPeering_Status
	err := copied.AssignPropertiesToVirtualNetworkPeeringStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworkPeering_Status
	err = actual.AssignPropertiesFromVirtualNetworkPeeringStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworkPeering_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeering_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringStatus, VirtualNetworkPeeringStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringStatus runs a test to see if a specific instance of VirtualNetworkPeering_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringStatus(subject VirtualNetworkPeering_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeering_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworkPeering_Status instances for property testing - lazily instantiated by
//VirtualNetworkPeeringStatusGenerator()
var virtualNetworkPeeringStatusGenerator gopter.Gen

// VirtualNetworkPeeringStatusGenerator returns a generator of VirtualNetworkPeering_Status instances for property testing.
// We first initialize virtualNetworkPeeringStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringStatusGenerator() gopter.Gen {
	if virtualNetworkPeeringStatusGenerator != nil {
		return virtualNetworkPeeringStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatus(generators)
	virtualNetworkPeeringStatusGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatus(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringStatus(generators)
	virtualNetworkPeeringStatusGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_Status{}), generators)

	return virtualNetworkPeeringStatusGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringStatus(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormatStatusPeeringStateConnected, VirtualNetworkPeeringPropertiesFormatStatusPeeringStateDisconnected, VirtualNetworkPeeringPropertiesFormatStatusPeeringStateInitiated))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringStatus(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceStatusGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesStatusGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceStatusGenerator())
}

func Test_VirtualNetworksVirtualNetworkPeerings_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeerings_Spec to VirtualNetworksVirtualNetworkPeerings_Spec via AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec & AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeeringsSpec, VirtualNetworksVirtualNetworkPeeringsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeeringsSpec tests if a specific instance of VirtualNetworksVirtualNetworkPeerings_Spec can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeeringsSpec(subject VirtualNetworksVirtualNetworkPeerings_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.VirtualNetworksVirtualNetworkPeerings_Spec
	err := copied.AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeerings_Spec
	err = actual.AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworksVirtualNetworkPeerings_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeerings_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringsSpec, VirtualNetworksVirtualNetworkPeeringsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringsSpec runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeerings_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringsSpec(subject VirtualNetworksVirtualNetworkPeerings_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeerings_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworksVirtualNetworkPeerings_Spec instances for property testing - lazily instantiated by
//VirtualNetworksVirtualNetworkPeeringsSpecGenerator()
var virtualNetworksVirtualNetworkPeeringsSpecGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeeringsSpecGenerator returns a generator of VirtualNetworksVirtualNetworkPeerings_Spec instances for property testing.
// We first initialize virtualNetworksVirtualNetworkPeeringsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksVirtualNetworkPeeringsSpecGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeeringsSpecGenerator != nil {
		return virtualNetworksVirtualNetworkPeeringsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpec(generators)
	virtualNetworksVirtualNetworkPeeringsSpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeerings_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpec(generators)
	virtualNetworksVirtualNetworkPeeringsSpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeerings_Spec{}), generators)

	return virtualNetworksVirtualNetworkPeeringsSpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpec(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormatPeeringStateConnected, VirtualNetworkPeeringPropertiesFormatPeeringStateDisconnected, VirtualNetworkPeeringPropertiesFormatPeeringStateInitiated))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpec(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["RemoteVirtualNetwork"] = SubResourceGenerator()
}
