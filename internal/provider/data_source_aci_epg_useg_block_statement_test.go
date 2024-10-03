// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvCrtrnWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvCrtrnDataSourceDependencyWithFvAEPg + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "match", "any"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "name", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "owner_tag", ""),
					composeAggregateTestCheckFuncWithVersion(t, "4.1(1i)", ">",
						resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "precedence", "0"),
						resource.TestCheckResourceAttr("data.aci_epg_useg_block_statement.test", "scope", "scope-bd")),
				),
			},
			{
				Config:      testConfigFvCrtrnNotExistingFvAEPg + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_block_statement data source"),
			},
		},
	})
}

const testConfigFvCrtrnDataSourceDependencyWithFvAEPg = testConfigFvCrtrnMinDependencyWithFvAEPg + `
data "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
  depends_on = [aci_epg_useg_block_statement.test]
}
`

const testConfigFvCrtrnNotExistingFvAEPg = testConfigFvAEPgMin + `
data "aci_epg_useg_block_statement" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
}
`
