// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourcePimRouteMapPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigPimRouteMapPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_tag", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigPimRouteMapPolAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_tag", "owner_tag"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigPimRouteMapPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name", "test_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigPimRouteMapPolResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_tag", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigPimRouteMapPolChildrenDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_pim_route_map_policy.test", "annotations.1.value", "value_2"),
				),
			},
		},
	})
}

const testConfigPimRouteMapPolMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_pim_route_map_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigPimRouteMapPolAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_pim_route_map_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotation = "annotation"
  description = "description"
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
}
`

const testConfigPimRouteMapPolResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_pim_route_map_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
}
`
const testConfigPimRouteMapPolChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_pim_route_map_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
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
