// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import (
	"encoding/json"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101previewstorage"
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

func Test_NamespacesQueue_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueue to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesQueue tests if a specific instance of NamespacesQueue round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesQueue(subject NamespacesQueue) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210101ps.NamespacesQueue
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesQueue
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

func Test_NamespacesQueue_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueue to NamespacesQueue via AssignPropertiesToNamespacesQueue & AssignPropertiesFromNamespacesQueue returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesQueue tests if a specific instance of NamespacesQueue can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueue(subject NamespacesQueue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.NamespacesQueue
	err := copied.AssignPropertiesToNamespacesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesQueue
	err = actual.AssignPropertiesFromNamespacesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesQueue_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue, NamespacesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue runs a test to see if a specific instance of NamespacesQueue round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue(subject NamespacesQueue) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue
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

// Generator of NamespacesQueue instances for property testing - lazily instantiated by NamespacesQueueGenerator()
var namespacesQueueGenerator gopter.Gen

// NamespacesQueueGenerator returns a generator of NamespacesQueue instances for property testing.
func NamespacesQueueGenerator() gopter.Gen {
	if namespacesQueueGenerator != nil {
		return namespacesQueueGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesQueue(generators)
	namespacesQueueGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue{}), generators)

	return namespacesQueueGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesQueue is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesQueuesSpecGenerator()
	gens["Status"] = SBQueueStatusGenerator()
}

func Test_NamespacesQueues_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesQueues_Spec to NamespacesQueues_Spec via AssignPropertiesToNamespacesQueuesSpec & AssignPropertiesFromNamespacesQueuesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesQueuesSpec, NamespacesQueuesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesQueuesSpec tests if a specific instance of NamespacesQueues_Spec can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForNamespacesQueuesSpec(subject NamespacesQueues_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.NamespacesQueues_Spec
	err := copied.AssignPropertiesToNamespacesQueuesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesQueues_Spec
	err = actual.AssignPropertiesFromNamespacesQueuesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesQueues_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueues_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueuesSpec, NamespacesQueuesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueuesSpec runs a test to see if a specific instance of NamespacesQueues_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueuesSpec(subject NamespacesQueues_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueues_Spec
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

// Generator of NamespacesQueues_Spec instances for property testing - lazily instantiated by
// NamespacesQueuesSpecGenerator()
var namespacesQueuesSpecGenerator gopter.Gen

// NamespacesQueuesSpecGenerator returns a generator of NamespacesQueues_Spec instances for property testing.
func NamespacesQueuesSpecGenerator() gopter.Gen {
	if namespacesQueuesSpecGenerator != nil {
		return namespacesQueuesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueuesSpec(generators)
	namespacesQueuesSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesQueues_Spec{}), generators)

	return namespacesQueuesSpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueuesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueuesSpec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_SBQueue_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SBQueue_Status to SBQueue_Status via AssignPropertiesToSBQueueStatus & AssignPropertiesFromSBQueueStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSBQueueStatus, SBQueueStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSBQueueStatus tests if a specific instance of SBQueue_Status can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForSBQueueStatus(subject SBQueue_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.SBQueue_Status
	err := copied.AssignPropertiesToSBQueueStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SBQueue_Status
	err = actual.AssignPropertiesFromSBQueueStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SBQueue_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueue_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueueStatus, SBQueueStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueueStatus runs a test to see if a specific instance of SBQueue_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueueStatus(subject SBQueue_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueue_Status
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

// Generator of SBQueue_Status instances for property testing - lazily instantiated by SBQueueStatusGenerator()
var sbQueueStatusGenerator gopter.Gen

// SBQueueStatusGenerator returns a generator of SBQueue_Status instances for property testing.
// We first initialize sbQueueStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBQueueStatusGenerator() gopter.Gen {
	if sbQueueStatusGenerator != nil {
		return sbQueueStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueStatus(generators)
	sbQueueStatusGenerator = gen.Struct(reflect.TypeOf(SBQueue_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueStatus(generators)
	AddRelatedPropertyGeneratorsForSBQueueStatus(generators)
	sbQueueStatusGenerator = gen.Struct(reflect.TypeOf(SBQueue_Status{}), generators)

	return sbQueueStatusGenerator
}

// AddIndependentPropertyGeneratorsForSBQueueStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueueStatus(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["MessageCount"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_StatusActive,
		EntityStatus_StatusCreating,
		EntityStatus_StatusDeleting,
		EntityStatus_StatusDisabled,
		EntityStatus_StatusReceiveDisabled,
		EntityStatus_StatusRenaming,
		EntityStatus_StatusRestoring,
		EntityStatus_StatusSendDisabled,
		EntityStatus_StatusUnknown))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBQueueStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBQueueStatus(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetailsStatusGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_MessageCountDetails_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from MessageCountDetails_Status to MessageCountDetails_Status via AssignPropertiesToMessageCountDetailsStatus & AssignPropertiesFromMessageCountDetailsStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForMessageCountDetailsStatus, MessageCountDetailsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForMessageCountDetailsStatus tests if a specific instance of MessageCountDetails_Status can be assigned to v1beta20210101previewstorage and back losslessly
func RunPropertyAssignmentTestForMessageCountDetailsStatus(subject MessageCountDetails_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210101ps.MessageCountDetails_Status
	err := copied.AssignPropertiesToMessageCountDetailsStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual MessageCountDetails_Status
	err = actual.AssignPropertiesFromMessageCountDetailsStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_MessageCountDetails_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MessageCountDetails_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMessageCountDetailsStatus, MessageCountDetailsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMessageCountDetailsStatus runs a test to see if a specific instance of MessageCountDetails_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMessageCountDetailsStatus(subject MessageCountDetails_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MessageCountDetails_Status
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

// Generator of MessageCountDetails_Status instances for property testing - lazily instantiated by
// MessageCountDetailsStatusGenerator()
var messageCountDetailsStatusGenerator gopter.Gen

// MessageCountDetailsStatusGenerator returns a generator of MessageCountDetails_Status instances for property testing.
func MessageCountDetailsStatusGenerator() gopter.Gen {
	if messageCountDetailsStatusGenerator != nil {
		return messageCountDetailsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMessageCountDetailsStatus(generators)
	messageCountDetailsStatusGenerator = gen.Struct(reflect.TypeOf(MessageCountDetails_Status{}), generators)

	return messageCountDetailsStatusGenerator
}

// AddIndependentPropertyGeneratorsForMessageCountDetailsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMessageCountDetailsStatus(gens map[string]gopter.Gen) {
	gens["ActiveMessageCount"] = gen.PtrOf(gen.Int())
	gens["DeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["ScheduledMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferDeadLetterMessageCount"] = gen.PtrOf(gen.Int())
	gens["TransferMessageCount"] = gen.PtrOf(gen.Int())
}
