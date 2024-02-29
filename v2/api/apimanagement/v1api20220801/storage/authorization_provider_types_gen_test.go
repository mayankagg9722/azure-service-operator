// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_AuthorizationProvider_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvider via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvider, AuthorizationProviderGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvider runs a test to see if a specific instance of AuthorizationProvider round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvider(subject AuthorizationProvider) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvider
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

// Generator of AuthorizationProvider instances for property testing - lazily instantiated by
// AuthorizationProviderGenerator()
var authorizationProviderGenerator gopter.Gen

// AuthorizationProviderGenerator returns a generator of AuthorizationProvider instances for property testing.
func AuthorizationProviderGenerator() gopter.Gen {
	if authorizationProviderGenerator != nil {
		return authorizationProviderGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAuthorizationProvider(generators)
	authorizationProviderGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvider{}), generators)

	return authorizationProviderGenerator
}

// AddRelatedPropertyGeneratorsForAuthorizationProvider is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvider(gens map[string]gopter.Gen) {
	gens["Spec"] = Service_AuthorizationProvider_SpecGenerator()
	gens["Status"] = Service_AuthorizationProvider_STATUSGenerator()
}

func Test_Service_AuthorizationProvider_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProvider_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProvider_Spec, Service_AuthorizationProvider_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProvider_Spec runs a test to see if a specific instance of Service_AuthorizationProvider_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProvider_Spec(subject Service_AuthorizationProvider_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProvider_Spec
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

// Generator of Service_AuthorizationProvider_Spec instances for property testing - lazily instantiated by
// Service_AuthorizationProvider_SpecGenerator()
var service_AuthorizationProvider_SpecGenerator gopter.Gen

// Service_AuthorizationProvider_SpecGenerator returns a generator of Service_AuthorizationProvider_Spec instances for property testing.
// We first initialize service_AuthorizationProvider_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_AuthorizationProvider_SpecGenerator() gopter.Gen {
	if service_AuthorizationProvider_SpecGenerator != nil {
		return service_AuthorizationProvider_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProvider_Spec(generators)
	service_AuthorizationProvider_SpecGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProvider_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProvider_Spec(generators)
	AddRelatedPropertyGeneratorsForService_AuthorizationProvider_Spec(generators)
	service_AuthorizationProvider_SpecGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProvider_Spec{}), generators)

	return service_AuthorizationProvider_SpecGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProvider_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProvider_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["IdentityProvider"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForService_AuthorizationProvider_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_AuthorizationProvider_Spec(gens map[string]gopter.Gen) {
	gens["Oauth2"] = gen.PtrOf(AuthorizationProviderOAuth2SettingsGenerator())
}

func Test_Service_AuthorizationProvider_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Service_AuthorizationProvider_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForService_AuthorizationProvider_STATUS, Service_AuthorizationProvider_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForService_AuthorizationProvider_STATUS runs a test to see if a specific instance of Service_AuthorizationProvider_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForService_AuthorizationProvider_STATUS(subject Service_AuthorizationProvider_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Service_AuthorizationProvider_STATUS
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

// Generator of Service_AuthorizationProvider_STATUS instances for property testing - lazily instantiated by
// Service_AuthorizationProvider_STATUSGenerator()
var service_AuthorizationProvider_STATUSGenerator gopter.Gen

// Service_AuthorizationProvider_STATUSGenerator returns a generator of Service_AuthorizationProvider_STATUS instances for property testing.
// We first initialize service_AuthorizationProvider_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Service_AuthorizationProvider_STATUSGenerator() gopter.Gen {
	if service_AuthorizationProvider_STATUSGenerator != nil {
		return service_AuthorizationProvider_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS(generators)
	service_AuthorizationProvider_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProvider_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS(generators)
	AddRelatedPropertyGeneratorsForService_AuthorizationProvider_STATUS(generators)
	service_AuthorizationProvider_STATUSGenerator = gen.Struct(reflect.TypeOf(Service_AuthorizationProvider_STATUS{}), generators)

	return service_AuthorizationProvider_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForService_AuthorizationProvider_STATUS(gens map[string]gopter.Gen) {
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IdentityProvider"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForService_AuthorizationProvider_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForService_AuthorizationProvider_STATUS(gens map[string]gopter.Gen) {
	gens["Oauth2"] = gen.PtrOf(AuthorizationProviderOAuth2Settings_STATUSGenerator())
}

func Test_AuthorizationProviderOAuth2Settings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProviderOAuth2Settings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProviderOAuth2Settings, AuthorizationProviderOAuth2SettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProviderOAuth2Settings runs a test to see if a specific instance of AuthorizationProviderOAuth2Settings round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProviderOAuth2Settings(subject AuthorizationProviderOAuth2Settings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProviderOAuth2Settings
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

// Generator of AuthorizationProviderOAuth2Settings instances for property testing - lazily instantiated by
// AuthorizationProviderOAuth2SettingsGenerator()
var authorizationProviderOAuth2SettingsGenerator gopter.Gen

// AuthorizationProviderOAuth2SettingsGenerator returns a generator of AuthorizationProviderOAuth2Settings instances for property testing.
// We first initialize authorizationProviderOAuth2SettingsGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProviderOAuth2SettingsGenerator() gopter.Gen {
	if authorizationProviderOAuth2SettingsGenerator != nil {
		return authorizationProviderOAuth2SettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings(generators)
	authorizationProviderOAuth2SettingsGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2Settings{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings(generators)
	authorizationProviderOAuth2SettingsGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2Settings{}), generators)

	return authorizationProviderOAuth2SettingsGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings(gens map[string]gopter.Gen) {
	gens["RedirectUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings(gens map[string]gopter.Gen) {
	gens["GrantTypes"] = gen.PtrOf(AuthorizationProviderOAuth2GrantTypesGenerator())
}

func Test_AuthorizationProviderOAuth2Settings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProviderOAuth2Settings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProviderOAuth2Settings_STATUS, AuthorizationProviderOAuth2Settings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProviderOAuth2Settings_STATUS runs a test to see if a specific instance of AuthorizationProviderOAuth2Settings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProviderOAuth2Settings_STATUS(subject AuthorizationProviderOAuth2Settings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProviderOAuth2Settings_STATUS
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

// Generator of AuthorizationProviderOAuth2Settings_STATUS instances for property testing - lazily instantiated by
// AuthorizationProviderOAuth2Settings_STATUSGenerator()
var authorizationProviderOAuth2Settings_STATUSGenerator gopter.Gen

// AuthorizationProviderOAuth2Settings_STATUSGenerator returns a generator of AuthorizationProviderOAuth2Settings_STATUS instances for property testing.
// We first initialize authorizationProviderOAuth2Settings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProviderOAuth2Settings_STATUSGenerator() gopter.Gen {
	if authorizationProviderOAuth2Settings_STATUSGenerator != nil {
		return authorizationProviderOAuth2Settings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS(generators)
	authorizationProviderOAuth2Settings_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2Settings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS(generators)
	authorizationProviderOAuth2Settings_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2Settings_STATUS{}), generators)

	return authorizationProviderOAuth2Settings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS(gens map[string]gopter.Gen) {
	gens["RedirectUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProviderOAuth2Settings_STATUS(gens map[string]gopter.Gen) {
	gens["GrantTypes"] = gen.PtrOf(AuthorizationProviderOAuth2GrantTypes_STATUSGenerator())
}

func Test_AuthorizationProviderOAuth2GrantTypes_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProviderOAuth2GrantTypes via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes, AuthorizationProviderOAuth2GrantTypesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes runs a test to see if a specific instance of AuthorizationProviderOAuth2GrantTypes round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes(subject AuthorizationProviderOAuth2GrantTypes) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProviderOAuth2GrantTypes
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

// Generator of AuthorizationProviderOAuth2GrantTypes instances for property testing - lazily instantiated by
// AuthorizationProviderOAuth2GrantTypesGenerator()
var authorizationProviderOAuth2GrantTypesGenerator gopter.Gen

// AuthorizationProviderOAuth2GrantTypesGenerator returns a generator of AuthorizationProviderOAuth2GrantTypes instances for property testing.
func AuthorizationProviderOAuth2GrantTypesGenerator() gopter.Gen {
	if authorizationProviderOAuth2GrantTypesGenerator != nil {
		return authorizationProviderOAuth2GrantTypesGenerator
	}

	generators := make(map[string]gopter.Gen)
	authorizationProviderOAuth2GrantTypesGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2GrantTypes{}), generators)

	return authorizationProviderOAuth2GrantTypesGenerator
}

func Test_AuthorizationProviderOAuth2GrantTypes_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProviderOAuth2GrantTypes_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes_STATUS, AuthorizationProviderOAuth2GrantTypes_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes_STATUS runs a test to see if a specific instance of AuthorizationProviderOAuth2GrantTypes_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProviderOAuth2GrantTypes_STATUS(subject AuthorizationProviderOAuth2GrantTypes_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProviderOAuth2GrantTypes_STATUS
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

// Generator of AuthorizationProviderOAuth2GrantTypes_STATUS instances for property testing - lazily instantiated by
// AuthorizationProviderOAuth2GrantTypes_STATUSGenerator()
var authorizationProviderOAuth2GrantTypes_STATUSGenerator gopter.Gen

// AuthorizationProviderOAuth2GrantTypes_STATUSGenerator returns a generator of AuthorizationProviderOAuth2GrantTypes_STATUS instances for property testing.
func AuthorizationProviderOAuth2GrantTypes_STATUSGenerator() gopter.Gen {
	if authorizationProviderOAuth2GrantTypes_STATUSGenerator != nil {
		return authorizationProviderOAuth2GrantTypes_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2GrantTypes_STATUS(generators)
	authorizationProviderOAuth2GrantTypes_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationProviderOAuth2GrantTypes_STATUS{}), generators)

	return authorizationProviderOAuth2GrantTypes_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2GrantTypes_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProviderOAuth2GrantTypes_STATUS(gens map[string]gopter.Gen) {
	gens["AuthorizationCode"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["ClientCredentials"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
