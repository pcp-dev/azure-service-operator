// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

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

func Test_CompositePath_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompositePath_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompositePath_ARM, CompositePath_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompositePath_ARM runs a test to see if a specific instance of CompositePath_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCompositePath_ARM(subject CompositePath_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompositePath_ARM
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

// Generator of CompositePath_ARM instances for property testing - lazily instantiated by CompositePath_ARMGenerator()
var compositePath_ARMGenerator gopter.Gen

// CompositePath_ARMGenerator returns a generator of CompositePath_ARM instances for property testing.
func CompositePath_ARMGenerator() gopter.Gen {
	if compositePath_ARMGenerator != nil {
		return compositePath_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompositePath_ARM(generators)
	compositePath_ARMGenerator = gen.Struct(reflect.TypeOf(CompositePath_ARM{}), generators)

	return compositePath_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCompositePath_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompositePath_ARM(gens map[string]gopter.Gen) {
	gens["Order"] = gen.PtrOf(gen.OneConstOf(CompositePath_Order_ARM_Ascending, CompositePath_Order_ARM_Descending))
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_ConflictResolutionPolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConflictResolutionPolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConflictResolutionPolicy_ARM, ConflictResolutionPolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConflictResolutionPolicy_ARM runs a test to see if a specific instance of ConflictResolutionPolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConflictResolutionPolicy_ARM(subject ConflictResolutionPolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConflictResolutionPolicy_ARM
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

// Generator of ConflictResolutionPolicy_ARM instances for property testing - lazily instantiated by
// ConflictResolutionPolicy_ARMGenerator()
var conflictResolutionPolicy_ARMGenerator gopter.Gen

// ConflictResolutionPolicy_ARMGenerator returns a generator of ConflictResolutionPolicy_ARM instances for property testing.
func ConflictResolutionPolicy_ARMGenerator() gopter.Gen {
	if conflictResolutionPolicy_ARMGenerator != nil {
		return conflictResolutionPolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConflictResolutionPolicy_ARM(generators)
	conflictResolutionPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(ConflictResolutionPolicy_ARM{}), generators)

	return conflictResolutionPolicy_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConflictResolutionPolicy_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConflictResolutionPolicy_ARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPath"] = gen.PtrOf(gen.AlphaString())
	gens["ConflictResolutionProcedure"] = gen.PtrOf(gen.AlphaString())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(ConflictResolutionPolicy_Mode_ARM_Custom, ConflictResolutionPolicy_Mode_ARM_LastWriterWins))
}

func Test_ContainerPartitionKey_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerPartitionKey_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPartitionKey_ARM, ContainerPartitionKey_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPartitionKey_ARM runs a test to see if a specific instance of ContainerPartitionKey_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPartitionKey_ARM(subject ContainerPartitionKey_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerPartitionKey_ARM
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

// Generator of ContainerPartitionKey_ARM instances for property testing - lazily instantiated by
// ContainerPartitionKey_ARMGenerator()
var containerPartitionKey_ARMGenerator gopter.Gen

// ContainerPartitionKey_ARMGenerator returns a generator of ContainerPartitionKey_ARM instances for property testing.
func ContainerPartitionKey_ARMGenerator() gopter.Gen {
	if containerPartitionKey_ARMGenerator != nil {
		return containerPartitionKey_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPartitionKey_ARM(generators)
	containerPartitionKey_ARMGenerator = gen.Struct(reflect.TypeOf(ContainerPartitionKey_ARM{}), generators)

	return containerPartitionKey_ARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPartitionKey_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPartitionKey_ARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ContainerPartitionKey_Kind_ARM_Hash, ContainerPartitionKey_Kind_ARM_MultiHash, ContainerPartitionKey_Kind_ARM_Range))
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.Int())
}

func Test_ExcludedPath_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExcludedPath_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExcludedPath_ARM, ExcludedPath_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExcludedPath_ARM runs a test to see if a specific instance of ExcludedPath_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExcludedPath_ARM(subject ExcludedPath_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExcludedPath_ARM
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

// Generator of ExcludedPath_ARM instances for property testing - lazily instantiated by ExcludedPath_ARMGenerator()
var excludedPath_ARMGenerator gopter.Gen

// ExcludedPath_ARMGenerator returns a generator of ExcludedPath_ARM instances for property testing.
func ExcludedPath_ARMGenerator() gopter.Gen {
	if excludedPath_ARMGenerator != nil {
		return excludedPath_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExcludedPath_ARM(generators)
	excludedPath_ARMGenerator = gen.Struct(reflect.TypeOf(ExcludedPath_ARM{}), generators)

	return excludedPath_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExcludedPath_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExcludedPath_ARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_IncludedPath_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IncludedPath_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIncludedPath_ARM, IncludedPath_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIncludedPath_ARM runs a test to see if a specific instance of IncludedPath_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIncludedPath_ARM(subject IncludedPath_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IncludedPath_ARM
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

// Generator of IncludedPath_ARM instances for property testing - lazily instantiated by IncludedPath_ARMGenerator()
var includedPath_ARMGenerator gopter.Gen

// IncludedPath_ARMGenerator returns a generator of IncludedPath_ARM instances for property testing.
// We first initialize includedPath_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IncludedPath_ARMGenerator() gopter.Gen {
	if includedPath_ARMGenerator != nil {
		return includedPath_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPath_ARM(generators)
	includedPath_ARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPath_ARM(generators)
	AddRelatedPropertyGeneratorsForIncludedPath_ARM(generators)
	includedPath_ARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_ARM{}), generators)

	return includedPath_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIncludedPath_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIncludedPath_ARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIncludedPath_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIncludedPath_ARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(Indexes_ARMGenerator())
}

func Test_Indexes_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Indexes_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexes_ARM, Indexes_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexes_ARM runs a test to see if a specific instance of Indexes_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexes_ARM(subject Indexes_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Indexes_ARM
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

// Generator of Indexes_ARM instances for property testing - lazily instantiated by Indexes_ARMGenerator()
var indexes_ARMGenerator gopter.Gen

// Indexes_ARMGenerator returns a generator of Indexes_ARM instances for property testing.
func Indexes_ARMGenerator() gopter.Gen {
	if indexes_ARMGenerator != nil {
		return indexes_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexes_ARM(generators)
	indexes_ARMGenerator = gen.Struct(reflect.TypeOf(Indexes_ARM{}), generators)

	return indexes_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexes_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexes_ARM(gens map[string]gopter.Gen) {
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		Indexes_DataType_ARM_LineString,
		Indexes_DataType_ARM_MultiPolygon,
		Indexes_DataType_ARM_Number,
		Indexes_DataType_ARM_Point,
		Indexes_DataType_ARM_Polygon,
		Indexes_DataType_ARM_String))
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(Indexes_Kind_ARM_Hash, Indexes_Kind_ARM_Range, Indexes_Kind_ARM_Spatial))
	gens["Precision"] = gen.PtrOf(gen.Int())
}

func Test_IndexingPolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IndexingPolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexingPolicy_ARM, IndexingPolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexingPolicy_ARM runs a test to see if a specific instance of IndexingPolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexingPolicy_ARM(subject IndexingPolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IndexingPolicy_ARM
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

// Generator of IndexingPolicy_ARM instances for property testing - lazily instantiated by IndexingPolicy_ARMGenerator()
var indexingPolicy_ARMGenerator gopter.Gen

// IndexingPolicy_ARMGenerator returns a generator of IndexingPolicy_ARM instances for property testing.
// We first initialize indexingPolicy_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IndexingPolicy_ARMGenerator() gopter.Gen {
	if indexingPolicy_ARMGenerator != nil {
		return indexingPolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicy_ARM(generators)
	indexingPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicy_ARM(generators)
	AddRelatedPropertyGeneratorsForIndexingPolicy_ARM(generators)
	indexingPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_ARM{}), generators)

	return indexingPolicy_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexingPolicy_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexingPolicy_ARM(gens map[string]gopter.Gen) {
	gens["Automatic"] = gen.PtrOf(gen.Bool())
	gens["IndexingMode"] = gen.PtrOf(gen.OneConstOf(IndexingPolicy_IndexingMode_ARM_Consistent, IndexingPolicy_IndexingMode_ARM_Lazy, IndexingPolicy_IndexingMode_ARM_None))
}

// AddRelatedPropertyGeneratorsForIndexingPolicy_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIndexingPolicy_ARM(gens map[string]gopter.Gen) {
	gens["CompositeIndexes"] = gen.SliceOf(gen.SliceOf(CompositePath_ARMGenerator()))
	gens["ExcludedPaths"] = gen.SliceOf(ExcludedPath_ARMGenerator())
	gens["IncludedPaths"] = gen.SliceOf(IncludedPath_ARMGenerator())
	gens["SpatialIndexes"] = gen.SliceOf(SpatialSpec_ARMGenerator())
}

func Test_SpatialSpec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SpatialSpec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSpatialSpec_ARM, SpatialSpec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSpatialSpec_ARM runs a test to see if a specific instance of SpatialSpec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSpatialSpec_ARM(subject SpatialSpec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SpatialSpec_ARM
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

// Generator of SpatialSpec_ARM instances for property testing - lazily instantiated by SpatialSpec_ARMGenerator()
var spatialSpec_ARMGenerator gopter.Gen

// SpatialSpec_ARMGenerator returns a generator of SpatialSpec_ARM instances for property testing.
func SpatialSpec_ARMGenerator() gopter.Gen {
	if spatialSpec_ARMGenerator != nil {
		return spatialSpec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSpatialSpec_ARM(generators)
	spatialSpec_ARMGenerator = gen.Struct(reflect.TypeOf(SpatialSpec_ARM{}), generators)

	return spatialSpec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSpatialSpec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSpatialSpec_ARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Types"] = gen.SliceOf(gen.OneConstOf(
		SpatialType_ARM_LineString,
		SpatialType_ARM_MultiPolygon,
		SpatialType_ARM_Point,
		SpatialType_ARM_Polygon))
}

func Test_SqlContainerCreateUpdateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerCreateUpdateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerCreateUpdateProperties_ARM, SqlContainerCreateUpdateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerCreateUpdateProperties_ARM runs a test to see if a specific instance of SqlContainerCreateUpdateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerCreateUpdateProperties_ARM(subject SqlContainerCreateUpdateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerCreateUpdateProperties_ARM
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

// Generator of SqlContainerCreateUpdateProperties_ARM instances for property testing - lazily instantiated by
// SqlContainerCreateUpdateProperties_ARMGenerator()
var sqlContainerCreateUpdateProperties_ARMGenerator gopter.Gen

// SqlContainerCreateUpdateProperties_ARMGenerator returns a generator of SqlContainerCreateUpdateProperties_ARM instances for property testing.
func SqlContainerCreateUpdateProperties_ARMGenerator() gopter.Gen {
	if sqlContainerCreateUpdateProperties_ARMGenerator != nil {
		return sqlContainerCreateUpdateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlContainerCreateUpdateProperties_ARM(generators)
	sqlContainerCreateUpdateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerCreateUpdateProperties_ARM{}), generators)

	return sqlContainerCreateUpdateProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlContainerCreateUpdateProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerCreateUpdateProperties_ARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_ARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlContainerResource_ARMGenerator())
}

func Test_SqlContainerResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerResource_ARM, SqlContainerResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerResource_ARM runs a test to see if a specific instance of SqlContainerResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerResource_ARM(subject SqlContainerResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerResource_ARM
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

// Generator of SqlContainerResource_ARM instances for property testing - lazily instantiated by
// SqlContainerResource_ARMGenerator()
var sqlContainerResource_ARMGenerator gopter.Gen

// SqlContainerResource_ARMGenerator returns a generator of SqlContainerResource_ARM instances for property testing.
// We first initialize sqlContainerResource_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainerResource_ARMGenerator() gopter.Gen {
	if sqlContainerResource_ARMGenerator != nil {
		return sqlContainerResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResource_ARM(generators)
	sqlContainerResource_ARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResource_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResource_ARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainerResource_ARM(generators)
	sqlContainerResource_ARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResource_ARM{}), generators)

	return sqlContainerResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainerResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainerResource_ARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["DefaultTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlContainerResource_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerResource_ARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPolicy"] = gen.PtrOf(ConflictResolutionPolicy_ARMGenerator())
	gens["IndexingPolicy"] = gen.PtrOf(IndexingPolicy_ARMGenerator())
	gens["PartitionKey"] = gen.PtrOf(ContainerPartitionKey_ARMGenerator())
	gens["UniqueKeyPolicy"] = gen.PtrOf(UniqueKeyPolicy_ARMGenerator())
}

func Test_SqlDatabaseContainer_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainer_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainer_Spec_ARM, SqlDatabaseContainer_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainer_Spec_ARM runs a test to see if a specific instance of SqlDatabaseContainer_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainer_Spec_ARM(subject SqlDatabaseContainer_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainer_Spec_ARM
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

// Generator of SqlDatabaseContainer_Spec_ARM instances for property testing - lazily instantiated by
// SqlDatabaseContainer_Spec_ARMGenerator()
var sqlDatabaseContainer_Spec_ARMGenerator gopter.Gen

// SqlDatabaseContainer_Spec_ARMGenerator returns a generator of SqlDatabaseContainer_Spec_ARM instances for property testing.
// We first initialize sqlDatabaseContainer_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseContainer_Spec_ARMGenerator() gopter.Gen {
	if sqlDatabaseContainer_Spec_ARMGenerator != nil {
		return sqlDatabaseContainer_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainer_Spec_ARM(generators)
	sqlDatabaseContainer_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainer_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainer_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainer_Spec_ARM(generators)
	sqlDatabaseContainer_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainer_Spec_ARM{}), generators)

	return sqlDatabaseContainer_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseContainer_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseContainer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainer_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlContainerCreateUpdateProperties_ARMGenerator())
}

func Test_UniqueKeyPolicy_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKeyPolicy_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyPolicy_ARM, UniqueKeyPolicy_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyPolicy_ARM runs a test to see if a specific instance of UniqueKeyPolicy_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyPolicy_ARM(subject UniqueKeyPolicy_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKeyPolicy_ARM
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

// Generator of UniqueKeyPolicy_ARM instances for property testing - lazily instantiated by
// UniqueKeyPolicy_ARMGenerator()
var uniqueKeyPolicy_ARMGenerator gopter.Gen

// UniqueKeyPolicy_ARMGenerator returns a generator of UniqueKeyPolicy_ARM instances for property testing.
func UniqueKeyPolicy_ARMGenerator() gopter.Gen {
	if uniqueKeyPolicy_ARMGenerator != nil {
		return uniqueKeyPolicy_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUniqueKeyPolicy_ARM(generators)
	uniqueKeyPolicy_ARMGenerator = gen.Struct(reflect.TypeOf(UniqueKeyPolicy_ARM{}), generators)

	return uniqueKeyPolicy_ARMGenerator
}

// AddRelatedPropertyGeneratorsForUniqueKeyPolicy_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUniqueKeyPolicy_ARM(gens map[string]gopter.Gen) {
	gens["UniqueKeys"] = gen.SliceOf(UniqueKey_ARMGenerator())
}

func Test_UniqueKey_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKey_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKey_ARM, UniqueKey_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKey_ARM runs a test to see if a specific instance of UniqueKey_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKey_ARM(subject UniqueKey_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKey_ARM
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

// Generator of UniqueKey_ARM instances for property testing - lazily instantiated by UniqueKey_ARMGenerator()
var uniqueKey_ARMGenerator gopter.Gen

// UniqueKey_ARMGenerator returns a generator of UniqueKey_ARM instances for property testing.
func UniqueKey_ARMGenerator() gopter.Gen {
	if uniqueKey_ARMGenerator != nil {
		return uniqueKey_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUniqueKey_ARM(generators)
	uniqueKey_ARMGenerator = gen.Struct(reflect.TypeOf(UniqueKey_ARM{}), generators)

	return uniqueKey_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUniqueKey_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUniqueKey_ARM(gens map[string]gopter.Gen) {
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
}