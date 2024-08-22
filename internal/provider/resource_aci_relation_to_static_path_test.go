// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsPathAttWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsPathAttMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "primary_encapsulation", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "primary_encapsulation", "unknown"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRsPathAttMinDependencyWithFvAEPgAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsPathAttMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test", "primary_encapsulation", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.allow_test_2", "primary_encapsulation", "unknown"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsPathAttMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "primary_encapsulation", "unknown"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsPathAttAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "deployment_immediacy", "immediate"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "encapsulation", "vlan-202"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "mode", "native"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "primary_encapsulation", "vlan-203"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsPathAttMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsPathAttResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "primary_encapsulation", "unknown"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_static_path.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsPathAttChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "primary_encapsulation", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_static_path.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsPathAttChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsPathAttChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsPathAttChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_static_path.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testDependencyConfigFvRsPathAtt = `
`

const testConfigFvRsPathAttMinDependencyWithFvAEPgAllowExisting = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "allow_test" {
  parent_dn = aci_application_epg.test.id
  encapsulation = "vlan-201"
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
}
resource "aci_relation_to_static_path" "allow_test_2" {
  parent_dn = aci_application_epg.test.id
  encapsulation = "vlan-201"
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  depends_on = [aci_relation_to_static_path.allow_test]
}
`

const testConfigFvRsPathAttMinDependencyWithFvAEPg = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  encapsulation = "vlan-201"
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
}
`

const testConfigFvRsPathAttAllDependencyWithFvAEPg = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotation = "annotation"
  deployment_immediacy = "immediate"
  description = "description_1"
  encapsulation = "vlan-202"
  mode = "native"
  primary_encapsulation = "vlan-203"
}
`

const testConfigFvRsPathAttResetDependencyWithFvAEPg = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotation = "orchestrator:terraform"
  deployment_immediacy = "lazy"
  description = ""
  encapsulation = "vlan-201"
  mode = "regular"
  primary_encapsulation = "unknown"
}
`
const testConfigFvRsPathAttChildrenDependencyWithFvAEPg = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  encapsulation = "vlan-201"
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotations = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigFvRsPathAttChildrenRemoveFromConfigDependencyWithFvAEPg = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  encapsulation = "vlan-201"
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
}
`

const testConfigFvRsPathAttChildrenRemoveOneDependencyWithFvAEPg = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  encapsulation = "vlan-201"
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigFvRsPathAttChildrenRemoveAllDependencyWithFvAEPg = testDependencyConfigFvRsPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  encapsulation = "vlan-201"
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotations = []
  tags = []
}
`
