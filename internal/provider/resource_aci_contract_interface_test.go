// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsConsIfWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsConsIfMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_contract_interface.test", "contract_interface_name", "test_tn_vz_cp_if_name"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "priority", "unspecified"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsConsIfAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_contract_interface.test", "contract_interface_name", "test_tn_vz_cp_if_name"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "priority", "level1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsConsIfMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_contract_interface.test", "contract_interface_name", "test_tn_vz_cp_if_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsConsIfResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_contract_interface.test", "contract_interface_name", "test_tn_vz_cp_if_name"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "priority", "unspecified"),
				),
			},
			// Update with children
			{
				Config:             testConfigFvRsConsIfChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_contract_interface.test", "contract_interface_name", "test_tn_vz_cp_if_name"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "priority", "unspecified"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_contract_interface.test", "annotations.1.value", "value_2"),
				),
			},
		},
	})
}

const testConfigFvRsConsIfMinDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_contract_interface" "test" {
  parent_dn = aci_application_epg.test.id
  contract_interface_name = "test_tn_vz_cp_if_name"
}
`

const testConfigFvRsConsIfAllDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_contract_interface" "test" {
  parent_dn = aci_application_epg.test.id
  contract_interface_name = "test_tn_vz_cp_if_name"
  annotation = "annotation"
  priority = "level1"
}
`

const testConfigFvRsConsIfResetDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_contract_interface" "test" {
  parent_dn = aci_application_epg.test.id
  contract_interface_name = "test_tn_vz_cp_if_name"
  annotation = "orchestrator:terraform"
  priority = "unspecified"
}
`
const testConfigFvRsConsIfChildrenDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_contract_interface" "test" {
  parent_dn = aci_application_epg.test.id
  contract_interface_name = "test_tn_vz_cp_if_name"
  annotations = [
	{
	  key = "annotations_1"
	  value = "value_1"
	},
	{
	  key = "annotations_2"
	  value = "value_2"
	},
  ]
}
`
