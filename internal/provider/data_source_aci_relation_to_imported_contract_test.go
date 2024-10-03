// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvRsConsIfWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsConsIfDataSourceDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_imported_contract.test", "imported_contract_name", "test_tn_vz_cp_if_name"),
					resource.TestCheckResourceAttr("data.aci_relation_to_imported_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_relation_to_imported_contract.test", "priority", "unspecified"),
				),
			},
			{
				Config:      testConfigFvRsConsIfNotExistingFvAEPg,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_imported_contract data source"),
			},
		},
	})
}
func TestAccDataSourceFvRsConsIfWithFvESg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsConsIfDataSourceDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_imported_contract.test", "imported_contract_name", "test_tn_vz_cp_if_name"),
					resource.TestCheckResourceAttr("data.aci_relation_to_imported_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_relation_to_imported_contract.test", "priority", "unspecified"),
				),
			},
			{
				Config:      testConfigFvRsConsIfNotExistingFvESg,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_imported_contract data source"),
			},
		},
	})
}

const testConfigFvRsConsIfDataSourceDependencyWithFvAEPg = testConfigFvRsConsIfMinDependencyWithFvAEPg + `
data "aci_relation_to_imported_contract" "test" {
  parent_dn = aci_application_epg.test.id
  imported_contract_name = "test_tn_vz_cp_if_name"
  depends_on = [aci_relation_to_imported_contract.test]
}
`

const testConfigFvRsConsIfNotExistingFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
data "aci_relation_to_imported_contract" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
  imported_contract_name = "non_existing_tn_vz_cp_if_name"
}
`
const testConfigFvRsConsIfDataSourceDependencyWithFvESg = testConfigFvRsConsIfMinDependencyWithFvESg + `
data "aci_relation_to_imported_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  imported_contract_name = "test_tn_vz_cp_if_name"
  depends_on = [aci_relation_to_imported_contract.test]
}
`

const testConfigFvRsConsIfNotExistingFvESg = testConfigFvESgMinDependencyWithFvAp + `
data "aci_relation_to_imported_contract" "test_non_existing" {
  parent_dn = aci_endpoint_security_group.test.id
  imported_contract_name = "non_existing_tn_vz_cp_if_name"
}
`
