// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceMldSnoopPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "4.1(1i)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigMldSnoopPolDataSourceDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "control.#", "0"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "query_interval", "125"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "response_interval", "10"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "start_query_count", "2"),
					resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "start_query_interval", "31"),
					composeAggregateTestCheckFuncWithVersion(t, "5.1(1h)", ">",
						resource.TestCheckResourceAttr("data.aci_mld_snooping_policy.test", "ver", "v2")),
				),
			},
			{
				Config:      testConfigMldSnoopPolNotExistingFvTenant + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Failed to read aci_mld_snooping_policy data source"),
			},
		},
	})
}

const testConfigMldSnoopPolDataSourceDependencyWithFvTenant = testConfigMldSnoopPolMinDependencyWithFvTenant + `
data "aci_mld_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_mld_snooping_policy.test]
}
`

const testConfigMldSnoopPolNotExistingFvTenant = testConfigFvTenantMin + `
data "aci_mld_snooping_policy" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "non_existing_name"
}
`