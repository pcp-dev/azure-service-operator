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

func Test_AutoscaleSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings_STATUS, AutoscaleSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings_STATUS runs a test to see if a specific instance of AutoscaleSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings_STATUS(subject AutoscaleSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_STATUS
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

// Generator of AutoscaleSettings_STATUS instances for property testing - lazily instantiated by
// AutoscaleSettings_STATUSGenerator()
var autoscaleSettings_STATUSGenerator gopter.Gen

// AutoscaleSettings_STATUSGenerator returns a generator of AutoscaleSettings_STATUS instances for property testing.
func AutoscaleSettings_STATUSGenerator() gopter.Gen {
	if autoscaleSettings_STATUSGenerator != nil {
		return autoscaleSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS(generators)
	autoscaleSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_STATUS{}), generators)

	return autoscaleSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_MongoDBCollectionGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetProperties_Resource_STATUS, MongoDBCollectionGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetProperties_Resource_STATUS runs a test to see if a specific instance of MongoDBCollectionGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetProperties_Resource_STATUS(subject MongoDBCollectionGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetProperties_Resource_STATUS
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

// Generator of MongoDBCollectionGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// MongoDBCollectionGetProperties_Resource_STATUSGenerator()
var mongoDBCollectionGetProperties_Resource_STATUSGenerator gopter.Gen

// MongoDBCollectionGetProperties_Resource_STATUSGenerator returns a generator of MongoDBCollectionGetProperties_Resource_STATUS instances for property testing.
// We first initialize mongoDBCollectionGetProperties_Resource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if mongoDBCollectionGetProperties_Resource_STATUSGenerator != nil {
		return mongoDBCollectionGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetProperties_Resource_STATUS(generators)
	mongoDBCollectionGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_Resource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionGetProperties_Resource_STATUS(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetProperties_Resource_STATUS(generators)
	mongoDBCollectionGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_Resource_STATUS{}), generators)

	return mongoDBCollectionGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndex_STATUSGenerator())
}

func Test_MongoDBCollectionGetProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionGetProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionGetProperties_STATUS, MongoDBCollectionGetProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionGetProperties_STATUS runs a test to see if a specific instance of MongoDBCollectionGetProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionGetProperties_STATUS(subject MongoDBCollectionGetProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionGetProperties_STATUS
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

// Generator of MongoDBCollectionGetProperties_STATUS instances for property testing - lazily instantiated by
// MongoDBCollectionGetProperties_STATUSGenerator()
var mongoDBCollectionGetProperties_STATUSGenerator gopter.Gen

// MongoDBCollectionGetProperties_STATUSGenerator returns a generator of MongoDBCollectionGetProperties_STATUS instances for property testing.
func MongoDBCollectionGetProperties_STATUSGenerator() gopter.Gen {
	if mongoDBCollectionGetProperties_STATUSGenerator != nil {
		return mongoDBCollectionGetProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBCollectionGetProperties_STATUS(generators)
	mongoDBCollectionGetProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionGetProperties_STATUS{}), generators)

	return mongoDBCollectionGetProperties_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionGetProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionGetProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUSGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionGetProperties_Resource_STATUSGenerator())
}

func Test_MongoIndexKeys_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeys_STATUS, MongoIndexKeys_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeys_STATUS runs a test to see if a specific instance of MongoIndexKeys_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeys_STATUS(subject MongoIndexKeys_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_STATUS
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

// Generator of MongoIndexKeys_STATUS instances for property testing - lazily instantiated by
// MongoIndexKeys_STATUSGenerator()
var mongoIndexKeys_STATUSGenerator gopter.Gen

// MongoIndexKeys_STATUSGenerator returns a generator of MongoIndexKeys_STATUS instances for property testing.
func MongoIndexKeys_STATUSGenerator() gopter.Gen {
	if mongoIndexKeys_STATUSGenerator != nil {
		return mongoIndexKeys_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeys_STATUS(generators)
	mongoIndexKeys_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_STATUS{}), generators)

	return mongoIndexKeys_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeys_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeys_STATUS(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptions_STATUS, MongoIndexOptions_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptions_STATUS runs a test to see if a specific instance of MongoIndexOptions_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptions_STATUS(subject MongoIndexOptions_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_STATUS
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

// Generator of MongoIndexOptions_STATUS instances for property testing - lazily instantiated by
// MongoIndexOptions_STATUSGenerator()
var mongoIndexOptions_STATUSGenerator gopter.Gen

// MongoIndexOptions_STATUSGenerator returns a generator of MongoIndexOptions_STATUS instances for property testing.
func MongoIndexOptions_STATUSGenerator() gopter.Gen {
	if mongoIndexOptions_STATUSGenerator != nil {
		return mongoIndexOptions_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptions_STATUS(generators)
	mongoIndexOptions_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_STATUS{}), generators)

	return mongoIndexOptions_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptions_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptions_STATUS(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}

func Test_MongoIndex_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndex_STATUS, MongoIndex_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndex_STATUS runs a test to see if a specific instance of MongoIndex_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndex_STATUS(subject MongoIndex_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_STATUS
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

// Generator of MongoIndex_STATUS instances for property testing - lazily instantiated by MongoIndex_STATUSGenerator()
var mongoIndex_STATUSGenerator gopter.Gen

// MongoIndex_STATUSGenerator returns a generator of MongoIndex_STATUS instances for property testing.
func MongoIndex_STATUSGenerator() gopter.Gen {
	if mongoIndex_STATUSGenerator != nil {
		return mongoIndex_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndex_STATUS(generators)
	mongoIndex_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoIndex_STATUS{}), generators)

	return mongoIndex_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndex_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndex_STATUS(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeys_STATUSGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptions_STATUSGenerator())
}

func Test_MongodbDatabaseCollection_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollection_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollection_STATUS, MongodbDatabaseCollection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollection_STATUS runs a test to see if a specific instance of MongodbDatabaseCollection_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollection_STATUS(subject MongodbDatabaseCollection_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollection_STATUS
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

// Generator of MongodbDatabaseCollection_STATUS instances for property testing - lazily instantiated by
// MongodbDatabaseCollection_STATUSGenerator()
var mongodbDatabaseCollection_STATUSGenerator gopter.Gen

// MongodbDatabaseCollection_STATUSGenerator returns a generator of MongodbDatabaseCollection_STATUS instances for property testing.
// We first initialize mongodbDatabaseCollection_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongodbDatabaseCollection_STATUSGenerator() gopter.Gen {
	if mongodbDatabaseCollection_STATUSGenerator != nil {
		return mongodbDatabaseCollection_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollection_STATUS(generators)
	mongodbDatabaseCollection_STATUSGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollection_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollection_STATUS(generators)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollection_STATUS(generators)
	mongodbDatabaseCollection_STATUSGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollection_STATUS{}), generators)

	return mongodbDatabaseCollection_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongodbDatabaseCollection_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongodbDatabaseCollection_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollection_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollection_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBCollectionGetProperties_STATUSGenerator())
}

func Test_OptionsResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OptionsResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOptionsResource_STATUS, OptionsResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOptionsResource_STATUS runs a test to see if a specific instance of OptionsResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForOptionsResource_STATUS(subject OptionsResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OptionsResource_STATUS
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

// Generator of OptionsResource_STATUS instances for property testing - lazily instantiated by
// OptionsResource_STATUSGenerator()
var optionsResource_STATUSGenerator gopter.Gen

// OptionsResource_STATUSGenerator returns a generator of OptionsResource_STATUS instances for property testing.
// We first initialize optionsResource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OptionsResource_STATUSGenerator() gopter.Gen {
	if optionsResource_STATUSGenerator != nil {
		return optionsResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResource_STATUS(generators)
	optionsResource_STATUSGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResource_STATUS(generators)
	AddRelatedPropertyGeneratorsForOptionsResource_STATUS(generators)
	optionsResource_STATUSGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUS{}), generators)

	return optionsResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForOptionsResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOptionsResource_STATUS(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForOptionsResource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOptionsResource_STATUS(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettings_STATUSGenerator())
}