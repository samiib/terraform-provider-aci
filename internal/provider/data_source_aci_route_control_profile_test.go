// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceRtctrlProfileWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigRtctrlProfileDataSourceDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "route_control_profile_type", "combinable"),
					composeAggregateTestCheckFuncWithVersion(t, "4.2(6d)-4.2(7w),5.1(3e)", ">",
						resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "route_map_continue", "no")),
				),
			},
			{
				Config:      testConfigRtctrlProfileNotExistingFvTenant + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Failed to read aci_route_control_profile data source"),
			},
		},
	})
}
func TestAccDataSourceRtctrlProfileWithL3extOut(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigRtctrlProfileDataSourceDependencyWithL3extOut + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "route_control_profile_type", "combinable"),
					composeAggregateTestCheckFuncWithVersion(t, "4.2(6d)-4.2(7w),5.1(3e)", ">",
						resource.TestCheckResourceAttr("data.aci_route_control_profile.test", "route_map_continue", "no")),
				),
			},
			{
				Config:      testConfigRtctrlProfileNotExistingL3extOut + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Failed to read aci_route_control_profile data source"),
			},
		},
	})
}

const testConfigRtctrlProfileDataSourceDependencyWithFvTenant = testConfigRtctrlProfileMinDependencyWithFvTenant + `
data "aci_route_control_profile" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_route_control_profile.test]
}
`

const testConfigRtctrlProfileNotExistingFvTenant = testConfigFvTenantMin + `
data "aci_route_control_profile" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "non_existing_name"
}
`
const testConfigRtctrlProfileDataSourceDependencyWithL3extOut = testConfigRtctrlProfileMinDependencyWithL3extOut + `
data "aci_route_control_profile" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "test_name"
  depends_on = [aci_route_control_profile.test]
}
`

const testConfigRtctrlProfileNotExistingL3extOut = testConfigL3extOutMin + `
data "aci_route_control_profile" "test_non_existing" {
  parent_dn = aci_l3_outside.test.id
  name = "non_existing_name"
}
`
