// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/storage"
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

func Test_AgentPoolSecurityProfile_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AgentPoolSecurityProfile_STATUS to AgentPoolSecurityProfile_STATUS via AssignProperties_To_AgentPoolSecurityProfile_STATUS & AssignProperties_From_AgentPoolSecurityProfile_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAgentPoolSecurityProfile_STATUS, AgentPoolSecurityProfile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAgentPoolSecurityProfile_STATUS tests if a specific instance of AgentPoolSecurityProfile_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAgentPoolSecurityProfile_STATUS(subject AgentPoolSecurityProfile_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AgentPoolSecurityProfile_STATUS
	err := copied.AssignProperties_To_AgentPoolSecurityProfile_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AgentPoolSecurityProfile_STATUS
	err = actual.AssignProperties_From_AgentPoolSecurityProfile_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AgentPoolSecurityProfile_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolSecurityProfile_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolSecurityProfile_STATUS, AgentPoolSecurityProfile_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolSecurityProfile_STATUS runs a test to see if a specific instance of AgentPoolSecurityProfile_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolSecurityProfile_STATUS(subject AgentPoolSecurityProfile_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolSecurityProfile_STATUS
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

// Generator of AgentPoolSecurityProfile_STATUS instances for property testing - lazily instantiated by
// AgentPoolSecurityProfile_STATUSGenerator()
var agentPoolSecurityProfile_STATUSGenerator gopter.Gen

// AgentPoolSecurityProfile_STATUSGenerator returns a generator of AgentPoolSecurityProfile_STATUS instances for property testing.
func AgentPoolSecurityProfile_STATUSGenerator() gopter.Gen {
	if agentPoolSecurityProfile_STATUSGenerator != nil {
		return agentPoolSecurityProfile_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolSecurityProfile_STATUS(generators)
	agentPoolSecurityProfile_STATUSGenerator = gen.Struct(reflect.TypeOf(AgentPoolSecurityProfile_STATUS{}), generators)

	return agentPoolSecurityProfile_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolSecurityProfile_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolSecurityProfile_STATUS(gens map[string]gopter.Gen) {
	gens["EnableSecureBoot"] = gen.PtrOf(gen.Bool())
	gens["EnableVTPM"] = gen.PtrOf(gen.Bool())
	gens["SshAccess"] = gen.PtrOf(gen.AlphaString())
}