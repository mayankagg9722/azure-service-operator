// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220615storage

import (
	"encoding/json"
	v20221001s "github.com/Azure/azure-service-operator/v2/api/insights/v1api20221001storage"
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

func Test_ScheduledQueryRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRule, ScheduledQueryRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRule runs a test to see if a specific instance of ScheduledQueryRule round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRule(subject ScheduledQueryRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRule
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

// Generator of ScheduledQueryRule instances for property testing - lazily instantiated by ScheduledQueryRuleGenerator()
var scheduledQueryRuleGenerator gopter.Gen

// ScheduledQueryRuleGenerator returns a generator of ScheduledQueryRule instances for property testing.
func ScheduledQueryRuleGenerator() gopter.Gen {
	if scheduledQueryRuleGenerator != nil {
		return scheduledQueryRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduledQueryRule(generators)
	scheduledQueryRuleGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule{}), generators)

	return scheduledQueryRuleGenerator
}

// AddRelatedPropertyGeneratorsForScheduledQueryRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRule(gens map[string]gopter.Gen) {
	gens["Spec"] = ScheduledQueryRule_SpecGenerator()
	gens["Status"] = ScheduledQueryRule_STATUSGenerator()
}

func Test_ScheduledQueryRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRule_Spec, ScheduledQueryRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRule_Spec runs a test to see if a specific instance of ScheduledQueryRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRule_Spec(subject ScheduledQueryRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRule_Spec
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

// Generator of ScheduledQueryRule_Spec instances for property testing - lazily instantiated by
// ScheduledQueryRule_SpecGenerator()
var scheduledQueryRule_SpecGenerator gopter.Gen

// ScheduledQueryRule_SpecGenerator returns a generator of ScheduledQueryRule_Spec instances for property testing.
// We first initialize scheduledQueryRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRule_SpecGenerator() gopter.Gen {
	if scheduledQueryRule_SpecGenerator != nil {
		return scheduledQueryRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec(generators)
	scheduledQueryRule_SpecGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec(generators)
	scheduledQueryRule_SpecGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_Spec{}), generators)

	return scheduledQueryRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRule_Spec(gens map[string]gopter.Gen) {
	gens["AutoMitigate"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["CheckWorkspaceAlertsStorageConfigured"] = gen.PtrOf(gen.Bool())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["EvaluationFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MuteActionsDuration"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["OverrideQueryTimeRange"] = gen.PtrOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.Int())
	gens["SkipQueryValidation"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TargetResourceTypes"] = gen.SliceOf(gen.AlphaString())
	gens["WindowSize"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRule_Spec(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(ActionsGenerator())
	gens["Criteria"] = gen.PtrOf(ScheduledQueryRuleCriteriaGenerator())
}

func Test_ScheduledQueryRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRule_STATUS, ScheduledQueryRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRule_STATUS runs a test to see if a specific instance of ScheduledQueryRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRule_STATUS(subject ScheduledQueryRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRule_STATUS
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

// Generator of ScheduledQueryRule_STATUS instances for property testing - lazily instantiated by
// ScheduledQueryRule_STATUSGenerator()
var scheduledQueryRule_STATUSGenerator gopter.Gen

// ScheduledQueryRule_STATUSGenerator returns a generator of ScheduledQueryRule_STATUS instances for property testing.
// We first initialize scheduledQueryRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRule_STATUSGenerator() gopter.Gen {
	if scheduledQueryRule_STATUSGenerator != nil {
		return scheduledQueryRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS(generators)
	scheduledQueryRule_STATUSGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRule_STATUS(generators)
	scheduledQueryRule_STATUSGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_STATUS{}), generators)

	return scheduledQueryRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS(gens map[string]gopter.Gen) {
	gens["AutoMitigate"] = gen.PtrOf(gen.Bool())
	gens["CheckWorkspaceAlertsStorageConfigured"] = gen.PtrOf(gen.Bool())
	gens["CreatedWithApiVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["EvaluationFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsLegacyLogAnalyticsRule"] = gen.PtrOf(gen.Bool())
	gens["IsWorkspaceAlertsStorageConfigured"] = gen.PtrOf(gen.Bool())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MuteActionsDuration"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["OverrideQueryTimeRange"] = gen.PtrOf(gen.AlphaString())
	gens["Scopes"] = gen.SliceOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.Int())
	gens["SkipQueryValidation"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TargetResourceTypes"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["WindowSize"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForScheduledQueryRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRule_STATUS(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(Actions_STATUSGenerator())
	gens["Criteria"] = gen.PtrOf(ScheduledQueryRuleCriteria_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Actions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Actions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActions, ActionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActions runs a test to see if a specific instance of Actions round trips to JSON and back losslessly
func RunJSONSerializationTestForActions(subject Actions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Actions
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

// Generator of Actions instances for property testing - lazily instantiated by ActionsGenerator()
var actionsGenerator gopter.Gen

// ActionsGenerator returns a generator of Actions instances for property testing.
func ActionsGenerator() gopter.Gen {
	if actionsGenerator != nil {
		return actionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActions(generators)
	actionsGenerator = gen.Struct(reflect.TypeOf(Actions{}), generators)

	return actionsGenerator
}

// AddIndependentPropertyGeneratorsForActions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActions(gens map[string]gopter.Gen) {
	gens["CustomProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_Actions_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Actions_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActions_STATUS, Actions_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActions_STATUS runs a test to see if a specific instance of Actions_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForActions_STATUS(subject Actions_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Actions_STATUS
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

// Generator of Actions_STATUS instances for property testing - lazily instantiated by Actions_STATUSGenerator()
var actions_STATUSGenerator gopter.Gen

// Actions_STATUSGenerator returns a generator of Actions_STATUS instances for property testing.
func Actions_STATUSGenerator() gopter.Gen {
	if actions_STATUSGenerator != nil {
		return actions_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActions_STATUS(generators)
	actions_STATUSGenerator = gen.Struct(reflect.TypeOf(Actions_STATUS{}), generators)

	return actions_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForActions_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActions_STATUS(gens map[string]gopter.Gen) {
	gens["ActionGroups"] = gen.SliceOf(gen.AlphaString())
	gens["CustomProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_ScheduledQueryRuleCriteria_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleCriteria via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleCriteria, ScheduledQueryRuleCriteriaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleCriteria runs a test to see if a specific instance of ScheduledQueryRuleCriteria round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleCriteria(subject ScheduledQueryRuleCriteria) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleCriteria
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

// Generator of ScheduledQueryRuleCriteria instances for property testing - lazily instantiated by
// ScheduledQueryRuleCriteriaGenerator()
var scheduledQueryRuleCriteriaGenerator gopter.Gen

// ScheduledQueryRuleCriteriaGenerator returns a generator of ScheduledQueryRuleCriteria instances for property testing.
func ScheduledQueryRuleCriteriaGenerator() gopter.Gen {
	if scheduledQueryRuleCriteriaGenerator != nil {
		return scheduledQueryRuleCriteriaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria(generators)
	scheduledQueryRuleCriteriaGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleCriteria{}), generators)

	return scheduledQueryRuleCriteriaGenerator
}

// AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(ConditionGenerator())
}

func Test_ScheduledQueryRuleCriteria_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleCriteria_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleCriteria_STATUS, ScheduledQueryRuleCriteria_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleCriteria_STATUS runs a test to see if a specific instance of ScheduledQueryRuleCriteria_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleCriteria_STATUS(subject ScheduledQueryRuleCriteria_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleCriteria_STATUS
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

// Generator of ScheduledQueryRuleCriteria_STATUS instances for property testing - lazily instantiated by
// ScheduledQueryRuleCriteria_STATUSGenerator()
var scheduledQueryRuleCriteria_STATUSGenerator gopter.Gen

// ScheduledQueryRuleCriteria_STATUSGenerator returns a generator of ScheduledQueryRuleCriteria_STATUS instances for property testing.
func ScheduledQueryRuleCriteria_STATUSGenerator() gopter.Gen {
	if scheduledQueryRuleCriteria_STATUSGenerator != nil {
		return scheduledQueryRuleCriteria_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_STATUS(generators)
	scheduledQueryRuleCriteria_STATUSGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleCriteria_STATUS{}), generators)

	return scheduledQueryRuleCriteria_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_STATUS(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(Condition_STATUSGenerator())
}

func Test_SystemData_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_STATUS to SystemData_STATUS via AssignProperties_To_SystemData_STATUS & AssignProperties_From_SystemData_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemData_STATUS tests if a specific instance of SystemData_STATUS can be assigned to v1api20221001storage and back losslessly
func RunPropertyAssignmentTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001s.SystemData_STATUS
	err := copied.AssignProperties_To_SystemData_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_STATUS
	err = actual.AssignProperties_From_SystemData_STATUS(&other)
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

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}

func Test_Condition_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition, ConditionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition runs a test to see if a specific instance of Condition round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition(subject Condition) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition
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

// Generator of Condition instances for property testing - lazily instantiated by ConditionGenerator()
var conditionGenerator gopter.Gen

// ConditionGenerator returns a generator of Condition instances for property testing.
// We first initialize conditionGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConditionGenerator() gopter.Gen {
	if conditionGenerator != nil {
		return conditionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition(generators)
	conditionGenerator = gen.Struct(reflect.TypeOf(Condition{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition(generators)
	AddRelatedPropertyGeneratorsForCondition(generators)
	conditionGenerator = gen.Struct(reflect.TypeOf(Condition{}), generators)

	return conditionGenerator
}

// AddIndependentPropertyGeneratorsForCondition is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition(gens map[string]gopter.Gen) {
	gens["MetricMeasureColumn"] = gen.PtrOf(gen.AlphaString())
	gens["MetricName"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.AlphaString())
	gens["Query"] = gen.PtrOf(gen.AlphaString())
	gens["Threshold"] = gen.PtrOf(gen.Float64())
	gens["TimeAggregation"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForCondition is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCondition(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(DimensionGenerator())
	gens["FailingPeriods"] = gen.PtrOf(Condition_FailingPeriodsGenerator())
}

func Test_Condition_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_STATUS, Condition_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_STATUS runs a test to see if a specific instance of Condition_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_STATUS(subject Condition_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_STATUS
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

// Generator of Condition_STATUS instances for property testing - lazily instantiated by Condition_STATUSGenerator()
var condition_STATUSGenerator gopter.Gen

// Condition_STATUSGenerator returns a generator of Condition_STATUS instances for property testing.
// We first initialize condition_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Condition_STATUSGenerator() gopter.Gen {
	if condition_STATUSGenerator != nil {
		return condition_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_STATUS(generators)
	condition_STATUSGenerator = gen.Struct(reflect.TypeOf(Condition_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_STATUS(generators)
	AddRelatedPropertyGeneratorsForCondition_STATUS(generators)
	condition_STATUSGenerator = gen.Struct(reflect.TypeOf(Condition_STATUS{}), generators)

	return condition_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCondition_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_STATUS(gens map[string]gopter.Gen) {
	gens["MetricMeasureColumn"] = gen.PtrOf(gen.AlphaString())
	gens["MetricName"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.AlphaString())
	gens["Query"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceIdColumn"] = gen.PtrOf(gen.AlphaString())
	gens["Threshold"] = gen.PtrOf(gen.Float64())
	gens["TimeAggregation"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForCondition_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCondition_STATUS(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(Dimension_STATUSGenerator())
	gens["FailingPeriods"] = gen.PtrOf(Condition_FailingPeriods_STATUSGenerator())
}

func Test_Condition_FailingPeriods_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_FailingPeriods via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_FailingPeriods, Condition_FailingPeriodsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_FailingPeriods runs a test to see if a specific instance of Condition_FailingPeriods round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_FailingPeriods(subject Condition_FailingPeriods) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_FailingPeriods
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

// Generator of Condition_FailingPeriods instances for property testing - lazily instantiated by
// Condition_FailingPeriodsGenerator()
var condition_FailingPeriodsGenerator gopter.Gen

// Condition_FailingPeriodsGenerator returns a generator of Condition_FailingPeriods instances for property testing.
func Condition_FailingPeriodsGenerator() gopter.Gen {
	if condition_FailingPeriodsGenerator != nil {
		return condition_FailingPeriodsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_FailingPeriods(generators)
	condition_FailingPeriodsGenerator = gen.Struct(reflect.TypeOf(Condition_FailingPeriods{}), generators)

	return condition_FailingPeriodsGenerator
}

// AddIndependentPropertyGeneratorsForCondition_FailingPeriods is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_FailingPeriods(gens map[string]gopter.Gen) {
	gens["MinFailingPeriodsToAlert"] = gen.PtrOf(gen.Int())
	gens["NumberOfEvaluationPeriods"] = gen.PtrOf(gen.Int())
}

func Test_Condition_FailingPeriods_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_FailingPeriods_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_FailingPeriods_STATUS, Condition_FailingPeriods_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_FailingPeriods_STATUS runs a test to see if a specific instance of Condition_FailingPeriods_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_FailingPeriods_STATUS(subject Condition_FailingPeriods_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_FailingPeriods_STATUS
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

// Generator of Condition_FailingPeriods_STATUS instances for property testing - lazily instantiated by
// Condition_FailingPeriods_STATUSGenerator()
var condition_FailingPeriods_STATUSGenerator gopter.Gen

// Condition_FailingPeriods_STATUSGenerator returns a generator of Condition_FailingPeriods_STATUS instances for property testing.
func Condition_FailingPeriods_STATUSGenerator() gopter.Gen {
	if condition_FailingPeriods_STATUSGenerator != nil {
		return condition_FailingPeriods_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_FailingPeriods_STATUS(generators)
	condition_FailingPeriods_STATUSGenerator = gen.Struct(reflect.TypeOf(Condition_FailingPeriods_STATUS{}), generators)

	return condition_FailingPeriods_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCondition_FailingPeriods_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_FailingPeriods_STATUS(gens map[string]gopter.Gen) {
	gens["MinFailingPeriodsToAlert"] = gen.PtrOf(gen.Int())
	gens["NumberOfEvaluationPeriods"] = gen.PtrOf(gen.Int())
}

func Test_Dimension_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Dimension via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDimension, DimensionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDimension runs a test to see if a specific instance of Dimension round trips to JSON and back losslessly
func RunJSONSerializationTestForDimension(subject Dimension) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Dimension
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

// Generator of Dimension instances for property testing - lazily instantiated by DimensionGenerator()
var dimensionGenerator gopter.Gen

// DimensionGenerator returns a generator of Dimension instances for property testing.
func DimensionGenerator() gopter.Gen {
	if dimensionGenerator != nil {
		return dimensionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDimension(generators)
	dimensionGenerator = gen.Struct(reflect.TypeOf(Dimension{}), generators)

	return dimensionGenerator
}

// AddIndependentPropertyGeneratorsForDimension is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDimension(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.AlphaString())
	gens["Values"] = gen.SliceOf(gen.AlphaString())
}

func Test_Dimension_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Dimension_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDimension_STATUS, Dimension_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDimension_STATUS runs a test to see if a specific instance of Dimension_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDimension_STATUS(subject Dimension_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Dimension_STATUS
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

// Generator of Dimension_STATUS instances for property testing - lazily instantiated by Dimension_STATUSGenerator()
var dimension_STATUSGenerator gopter.Gen

// Dimension_STATUSGenerator returns a generator of Dimension_STATUS instances for property testing.
func Dimension_STATUSGenerator() gopter.Gen {
	if dimension_STATUSGenerator != nil {
		return dimension_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDimension_STATUS(generators)
	dimension_STATUSGenerator = gen.Struct(reflect.TypeOf(Dimension_STATUS{}), generators)

	return dimension_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDimension_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDimension_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.AlphaString())
	gens["Values"] = gen.SliceOf(gen.AlphaString())
}