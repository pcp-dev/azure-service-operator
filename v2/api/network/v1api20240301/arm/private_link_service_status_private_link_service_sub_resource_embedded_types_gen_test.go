// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
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

func Test_FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded, FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded runs a test to see if a specific instance of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded(subject FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded
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

// Generator of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing -
// lazily instantiated by FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()
var frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator gopter.Gen

// FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator returns a generator of FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing.
func FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator() gopter.Gen {
	if frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator != nil {
		return frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded(generators)
	frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded{}), generators)

	return frontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded, NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded runs a test to see if a specific instance of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded(subject NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded
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

// Generator of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing - lazily
// instantiated by NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()
var networkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator gopter.Gen

// NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator returns a generator of NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing.
func NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator != nil {
		return networkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded(generators)
	networkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded{}), generators)

	return networkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_PrivateLinkService_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrivateEndpointConnection_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS, PrivateEndpointConnection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS runs a test to see if a specific instance of PrivateEndpointConnection_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS(subject PrivateEndpointConnection_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS
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

// Generator of PrivateEndpointConnection_STATUS instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUSGenerator()
var privateEndpointConnection_STATUSGenerator gopter.Gen

// PrivateEndpointConnection_STATUSGenerator returns a generator of PrivateEndpointConnection_STATUS instances for property testing.
func PrivateEndpointConnection_STATUSGenerator() gopter.Gen {
	if privateEndpointConnection_STATUSGenerator != nil {
		return privateEndpointConnection_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(generators)
	privateEndpointConnection_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS{}), generators)

	return privateEndpointConnection_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrivateLinkServiceIpConfigurationProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceIpConfigurationProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceIpConfigurationProperties_STATUS, PrivateLinkServiceIpConfigurationProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceIpConfigurationProperties_STATUS runs a test to see if a specific instance of PrivateLinkServiceIpConfigurationProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceIpConfigurationProperties_STATUS(subject PrivateLinkServiceIpConfigurationProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceIpConfigurationProperties_STATUS
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

// Generator of PrivateLinkServiceIpConfigurationProperties_STATUS instances for property testing - lazily instantiated
// by PrivateLinkServiceIpConfigurationProperties_STATUSGenerator()
var privateLinkServiceIpConfigurationProperties_STATUSGenerator gopter.Gen

// PrivateLinkServiceIpConfigurationProperties_STATUSGenerator returns a generator of PrivateLinkServiceIpConfigurationProperties_STATUS instances for property testing.
// We first initialize privateLinkServiceIpConfigurationProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceIpConfigurationProperties_STATUSGenerator() gopter.Gen {
	if privateLinkServiceIpConfigurationProperties_STATUSGenerator != nil {
		return privateLinkServiceIpConfigurationProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS(generators)
	privateLinkServiceIpConfigurationProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfigurationProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS(generators)
	privateLinkServiceIpConfigurationProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfigurationProperties_STATUS{}), generators)

	return privateLinkServiceIpConfigurationProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Primary"] = gen.PtrOf(gen.Bool())
	gens["PrivateIPAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(IPVersion_STATUS_IPv4, IPVersion_STATUS_IPv6))
	gens["PrivateIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IPAllocationMethod_STATUS_Dynamic, IPAllocationMethod_STATUS_Static))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfigurationProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(Subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator())
}

func Test_PrivateLinkServiceIpConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceIpConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceIpConfiguration_STATUS, PrivateLinkServiceIpConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceIpConfiguration_STATUS runs a test to see if a specific instance of PrivateLinkServiceIpConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceIpConfiguration_STATUS(subject PrivateLinkServiceIpConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceIpConfiguration_STATUS
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

// Generator of PrivateLinkServiceIpConfiguration_STATUS instances for property testing - lazily instantiated by
// PrivateLinkServiceIpConfiguration_STATUSGenerator()
var privateLinkServiceIpConfiguration_STATUSGenerator gopter.Gen

// PrivateLinkServiceIpConfiguration_STATUSGenerator returns a generator of PrivateLinkServiceIpConfiguration_STATUS instances for property testing.
// We first initialize privateLinkServiceIpConfiguration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceIpConfiguration_STATUSGenerator() gopter.Gen {
	if privateLinkServiceIpConfiguration_STATUSGenerator != nil {
		return privateLinkServiceIpConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS(generators)
	privateLinkServiceIpConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfiguration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS(generators)
	privateLinkServiceIpConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceIpConfiguration_STATUS{}), generators)

	return privateLinkServiceIpConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceIpConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateLinkServiceIpConfigurationProperties_STATUSGenerator())
}

func Test_PrivateLinkServiceProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkServiceProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkServiceProperties_STATUS, PrivateLinkServiceProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkServiceProperties_STATUS runs a test to see if a specific instance of PrivateLinkServiceProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkServiceProperties_STATUS(subject PrivateLinkServiceProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkServiceProperties_STATUS
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

// Generator of PrivateLinkServiceProperties_STATUS instances for property testing - lazily instantiated by
// PrivateLinkServiceProperties_STATUSGenerator()
var privateLinkServiceProperties_STATUSGenerator gopter.Gen

// PrivateLinkServiceProperties_STATUSGenerator returns a generator of PrivateLinkServiceProperties_STATUS instances for property testing.
// We first initialize privateLinkServiceProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkServiceProperties_STATUSGenerator() gopter.Gen {
	if privateLinkServiceProperties_STATUSGenerator != nil {
		return privateLinkServiceProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS(generators)
	privateLinkServiceProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkServiceProperties_STATUS(generators)
	privateLinkServiceProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateLinkServiceProperties_STATUS{}), generators)

	return privateLinkServiceProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkServiceProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Alias"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationIPAddress"] = gen.PtrOf(gen.AlphaString())
	gens["EnableProxyProtocol"] = gen.PtrOf(gen.Bool())
	gens["Fqdns"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForPrivateLinkServiceProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkServiceProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AutoApproval"] = gen.PtrOf(ResourceSet_STATUSGenerator())
	gens["IpConfigurations"] = gen.SliceOf(PrivateLinkServiceIpConfiguration_STATUSGenerator())
	gens["LoadBalancerFrontendIpConfigurations"] = gen.SliceOf(FrontendIPConfiguration_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterface_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUSGenerator())
	gens["Visibility"] = gen.PtrOf(ResourceSet_STATUSGenerator())
}

func Test_PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded, PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded runs a test to see if a specific instance of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded(subject PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded
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

// Generator of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing - lazily
// instantiated by PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()
var privateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator gopter.Gen

// PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator returns a generator of PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing.
// We first initialize privateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator() gopter.Gen {
	if privateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator != nil {
		return privateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded(generators)
	privateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded(generators)
	privateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded{}), generators)

	return privateLinkService_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateLinkService_STATUS_PrivateLinkService_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["Properties"] = gen.PtrOf(PrivateLinkServiceProperties_STATUSGenerator())
}

func Test_ResourceSet_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceSet_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceSet_STATUS, ResourceSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceSet_STATUS runs a test to see if a specific instance of ResourceSet_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceSet_STATUS(subject ResourceSet_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceSet_STATUS
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

// Generator of ResourceSet_STATUS instances for property testing - lazily instantiated by ResourceSet_STATUSGenerator()
var resourceSet_STATUSGenerator gopter.Gen

// ResourceSet_STATUSGenerator returns a generator of ResourceSet_STATUS instances for property testing.
func ResourceSet_STATUSGenerator() gopter.Gen {
	if resourceSet_STATUSGenerator != nil {
		return resourceSet_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceSet_STATUS(generators)
	resourceSet_STATUSGenerator = gen.Struct(reflect.TypeOf(ResourceSet_STATUS{}), generators)

	return resourceSet_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForResourceSet_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceSet_STATUS(gens map[string]gopter.Gen) {
	gens["Subscriptions"] = gen.SliceOf(gen.AlphaString())
}

func Test_Subnet_STATUS_PrivateLinkService_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded, Subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded runs a test to see if a specific instance of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded(subject Subnet_STATUS_PrivateLinkService_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_STATUS_PrivateLinkService_SubResourceEmbedded
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

// Generator of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing - lazily
// instantiated by Subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator()
var subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator gopter.Gen

// Subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator returns a generator of Subnet_STATUS_PrivateLinkService_SubResourceEmbedded instances for property testing.
func Subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator() gopter.Gen {
	if subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator != nil {
		return subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded(generators)
	subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(Subnet_STATUS_PrivateLinkService_SubResourceEmbedded{}), generators)

	return subnet_STATUS_PrivateLinkService_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_STATUS_PrivateLinkService_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}