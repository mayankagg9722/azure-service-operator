// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

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

func Test_VirtualMachineScaleSets_Extension_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSets_Extension_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSets_Extension_STATUS_ARM, VirtualMachineScaleSets_Extension_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSets_Extension_STATUS_ARM runs a test to see if a specific instance of VirtualMachineScaleSets_Extension_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSets_Extension_STATUS_ARM(subject VirtualMachineScaleSets_Extension_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSets_Extension_STATUS_ARM
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

// Generator of VirtualMachineScaleSets_Extension_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualMachineScaleSets_Extension_STATUS_ARMGenerator()
var virtualMachineScaleSets_Extension_STATUS_ARMGenerator gopter.Gen

// VirtualMachineScaleSets_Extension_STATUS_ARMGenerator returns a generator of VirtualMachineScaleSets_Extension_STATUS_ARM instances for property testing.
// We first initialize virtualMachineScaleSets_Extension_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualMachineScaleSets_Extension_STATUS_ARMGenerator() gopter.Gen {
	if virtualMachineScaleSets_Extension_STATUS_ARMGenerator != nil {
		return virtualMachineScaleSets_Extension_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_STATUS_ARM(generators)
	virtualMachineScaleSets_Extension_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSets_Extension_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualMachineScaleSets_Extension_STATUS_ARM(generators)
	virtualMachineScaleSets_Extension_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSets_Extension_STATUS_ARM{}), generators)

	return virtualMachineScaleSets_Extension_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSets_Extension_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualMachineScaleSets_Extension_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualMachineScaleSets_Extension_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator())
}

func Test_VirtualMachineScaleSetExtensionProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualMachineScaleSetExtensionProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_STATUS_ARM, VirtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_STATUS_ARM runs a test to see if a specific instance of VirtualMachineScaleSetExtensionProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualMachineScaleSetExtensionProperties_STATUS_ARM(subject VirtualMachineScaleSetExtensionProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualMachineScaleSetExtensionProperties_STATUS_ARM
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

// Generator of VirtualMachineScaleSetExtensionProperties_STATUS_ARM instances for property testing - lazily
// instantiated by VirtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator()
var virtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator gopter.Gen

// VirtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator returns a generator of VirtualMachineScaleSetExtensionProperties_STATUS_ARM instances for property testing.
func VirtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator() gopter.Gen {
	if virtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator != nil {
		return virtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_STATUS_ARM(generators)
	virtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualMachineScaleSetExtensionProperties_STATUS_ARM{}), generators)

	return virtualMachineScaleSetExtensionProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualMachineScaleSetExtensionProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradeMinorVersion"] = gen.PtrOf(gen.Bool())
	gens["EnableAutomaticUpgrade"] = gen.PtrOf(gen.Bool())
	gens["ForceUpdateTag"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisionAfterExtensions"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Publisher"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["TypeHandlerVersion"] = gen.PtrOf(gen.AlphaString())
}
