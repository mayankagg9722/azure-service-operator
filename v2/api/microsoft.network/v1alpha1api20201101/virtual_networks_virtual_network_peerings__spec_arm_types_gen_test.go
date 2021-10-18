// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_VirtualNetworksVirtualNetworkPeerings_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeerings_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringsSpecARM, VirtualNetworksVirtualNetworkPeeringsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringsSpecARM runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeerings_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringsSpecARM(subject VirtualNetworksVirtualNetworkPeerings_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeerings_SpecARM
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

// Generator of VirtualNetworksVirtualNetworkPeerings_SpecARM instances for property testing - lazily instantiated by
//VirtualNetworksVirtualNetworkPeeringsSpecARMGenerator()
var virtualNetworksVirtualNetworkPeeringsSpecARMGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeeringsSpecARMGenerator returns a generator of VirtualNetworksVirtualNetworkPeerings_SpecARM instances for property testing.
// We first initialize virtualNetworksVirtualNetworkPeeringsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksVirtualNetworkPeeringsSpecARMGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeeringsSpecARMGenerator != nil {
		return virtualNetworksVirtualNetworkPeeringsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpecARM(generators)
	virtualNetworksVirtualNetworkPeeringsSpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeerings_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpecARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpecARM(generators)
	virtualNetworksVirtualNetworkPeeringsSpecARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeerings_SpecARM{}), generators)

	return virtualNetworksVirtualNetworkPeeringsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpecARM(gens map[string]gopter.Gen) {
	gens["APIVersion"] = gen.OneConstOf(VirtualNetworksVirtualNetworkPeeringsSpecAPIVersion20201101)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.OneConstOf(VirtualNetworksVirtualNetworkPeeringsSpecTypeMicrosoftNetworkVirtualNetworksVirtualNetworkPeerings)
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeeringsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = VirtualNetworkPeeringPropertiesFormatARMGenerator()
}

func Test_VirtualNetworkPeeringPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeeringPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatARM, VirtualNetworkPeeringPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatARM runs a test to see if a specific instance of VirtualNetworkPeeringPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormatARM(subject VirtualNetworkPeeringPropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeeringPropertiesFormatARM
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

// Generator of VirtualNetworkPeeringPropertiesFormatARM instances for property testing - lazily instantiated by
//VirtualNetworkPeeringPropertiesFormatARMGenerator()
var virtualNetworkPeeringPropertiesFormatARMGenerator gopter.Gen

// VirtualNetworkPeeringPropertiesFormatARMGenerator returns a generator of VirtualNetworkPeeringPropertiesFormatARM instances for property testing.
// We first initialize virtualNetworkPeeringPropertiesFormatARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringPropertiesFormatARMGenerator() gopter.Gen {
	if virtualNetworkPeeringPropertiesFormatARMGenerator != nil {
		return virtualNetworkPeeringPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(generators)
	virtualNetworkPeeringPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormatARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(generators)
	virtualNetworkPeeringPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormatARM{}), generators)

	return virtualNetworkPeeringPropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormatPeeringStateConnected, VirtualNetworkPeeringPropertiesFormatPeeringStateDisconnected, VirtualNetworkPeeringPropertiesFormatPeeringStateInitiated))
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceARMGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesARMGenerator())
	gens["RemoteVirtualNetwork"] = SubResourceARMGenerator()
}

func Test_VirtualNetworkBgpCommunitiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunitiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunitiesARM, VirtualNetworkBgpCommunitiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunitiesARM runs a test to see if a specific instance of VirtualNetworkBgpCommunitiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunitiesARM(subject VirtualNetworkBgpCommunitiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunitiesARM
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

// Generator of VirtualNetworkBgpCommunitiesARM instances for property testing - lazily instantiated by
//VirtualNetworkBgpCommunitiesARMGenerator()
var virtualNetworkBgpCommunitiesARMGenerator gopter.Gen

// VirtualNetworkBgpCommunitiesARMGenerator returns a generator of VirtualNetworkBgpCommunitiesARM instances for property testing.
func VirtualNetworkBgpCommunitiesARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunitiesARMGenerator != nil {
		return virtualNetworkBgpCommunitiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesARM(generators)
	virtualNetworkBgpCommunitiesARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunitiesARM{}), generators)

	return virtualNetworkBgpCommunitiesARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesARM(gens map[string]gopter.Gen) {
	gens["VirtualNetworkCommunity"] = gen.AlphaString()
}
