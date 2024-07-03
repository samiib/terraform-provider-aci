// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceL3extRsOutToFBRGroupWithL3extOut(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigL3extRsOutToFBRGroupMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-vrf_fallback_route_group_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigL3extRsOutToFBRGroupAllDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-vrf_fallback_route_group_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigL3extRsOutToFBRGroupMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-vrf_fallback_route_group_0"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigL3extRsOutToFBRGroupResetDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-vrf_fallback_route_group_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_vrf_fallback_route_group.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-vrf_fallback_route_group_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with children
			{
				Config:             testConfigL3extRsOutToFBRGroupChildrenDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-vrf_fallback_route_group_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_vrf_fallback_route_group.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-vrf_fallback_route_group_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigL3extRsOutToFBRGroupChildrenRemoveFromConfigDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigL3extRsOutToFBRGroupChildrenRemoveOneDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigL3extRsOutToFBRGroupChildrenRemoveAllDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_vrf_fallback_route_group.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testDependencyConfigL3extRsOutToFBRGroup = `
resource "aci_vrf_fallback_route_group" "test_0" {
  parent_dn = aci_vrf.test.id
  name = "vrf_fallback_route_group_0"
}
`

const testConfigL3extRsOutToFBRGroupMinDependencyWithL3extOut = testDependencyConfigL3extRsOutToFBRGroup + testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_relation_to_vrf_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test_0.id
}
`

const testConfigL3extRsOutToFBRGroupAllDependencyWithL3extOut = testDependencyConfigL3extRsOutToFBRGroup + testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_relation_to_vrf_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test_0.id
  annotation = "annotation"
}
`

const testConfigL3extRsOutToFBRGroupResetDependencyWithL3extOut = testDependencyConfigL3extRsOutToFBRGroup + testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_relation_to_vrf_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test_0.id
  annotation = "orchestrator:terraform"
}
`
const testConfigL3extRsOutToFBRGroupChildrenDependencyWithL3extOut = testDependencyConfigL3extRsOutToFBRGroup + testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_relation_to_vrf_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test_0.id
  annotations = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigL3extRsOutToFBRGroupChildrenRemoveFromConfigDependencyWithL3extOut = testDependencyConfigL3extRsOutToFBRGroup + testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_relation_to_vrf_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test_0.id
}
`

const testConfigL3extRsOutToFBRGroupChildrenRemoveOneDependencyWithL3extOut = testDependencyConfigL3extRsOutToFBRGroup + testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_relation_to_vrf_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test_0.id
  annotations = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigL3extRsOutToFBRGroupChildrenRemoveAllDependencyWithL3extOut = testDependencyConfigL3extRsOutToFBRGroup + testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_relation_to_vrf_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test_0.id
  annotations = []
  tags = []
}
`