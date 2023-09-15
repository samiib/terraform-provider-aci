// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceL3extRsRedistributePolWithL3extOut(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigL3extRsRedistributePolMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "src", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigL3extRsRedistributePolAllDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "src", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "test_annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigL3extRsRedistributePolMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "src", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigL3extRsRedistributePolResetDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "src", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with children
			{
				Config:             testConfigL3extRsRedistributePolChildrenDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "src", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.key", "test_key_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.value", "test_value_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.key", "test_key_2"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.value", "test_value_2"),
				),
			},
		},
	})
}

const testConfigL3extRsRedistributePolMinDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_profile_name = "test_tn_rtctrl_profile_name"
  src = "direct"
}
`

const testConfigL3extRsRedistributePolAllDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_profile_name = "test_tn_rtctrl_profile_name"
  src = "direct"
  annotation = "test_annotation"
}
`

const testConfigL3extRsRedistributePolResetDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_profile_name = "test_tn_rtctrl_profile_name"
  src = "direct"
  annotation = "orchestrator:terraform"
}
`
const testConfigL3extRsRedistributePolChildrenDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_profile_name = "test_tn_rtctrl_profile_name"
  src = "direct"
  annotations = [
	{
	  key = "test_key_1"
	  value = "test_value_1"
	},
	{
	  key = "test_key_2"
	  value = "test_value_2"
	},
  ]
}
`
