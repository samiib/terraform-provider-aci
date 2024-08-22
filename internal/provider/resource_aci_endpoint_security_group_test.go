// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvESgWithFvAp(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvESgMinDependencyWithFvApAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "exception_tag", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "exception_tag", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.allow_test", "pc_tag"),
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
				Config:      testConfigFvESgMinDependencyWithFvApAllowExisting,
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
				Config:             testConfigFvESgMinDependencyWithFvApAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "exception_tag", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "exception_tag", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.allow_test_2", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.allow_test", "pc_tag"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.allow_test_2", "pc_tag"),
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
				Config:             testConfigFvESgMinDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "exception_tag", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvESgAllDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "exception_tag", "exception_tag_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "intra_esg_isolation", "enforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "match_criteria", "All"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvESgMinDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name", "test_name"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvESgResetDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "exception_tag", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_endpoint_security_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvESgChildrenDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "admin_state", "no"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "exception_tag", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "preferred_group_member", "exclude"),
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.contract_name", "contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.1.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.0.target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.1.target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.imported_contract_name", "imported_contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.1.imported_contract_name", "imported_contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.0.contract_name", "contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.1.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.contract_name", "contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.match_criteria", "All"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.vrf_name", "vrf_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_endpoint_security_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvESgChildrenRemoveFromConfigDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.contract_name", "contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.1.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.#", "2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.0.target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.1.target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.#", "2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.imported_contract_name", "imported_contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.1.imported_contract_name", "imported_contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.#", "2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.0.contract_name", "contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.1.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.#", "2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.contract_name", "contract_name_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.match_criteria", "All"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.#", "2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.vrf_name", "vrf_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvESgChildrenRemoveOneDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.0.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.0.target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.imported_contract_name", "imported_contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.0.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.0.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.contract_name", "contract_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.0.priority", "level2"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.vrf_name", "vrf_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvESgChildrenRemoveAllDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("aci_endpoint_security_group.test", "pc_tag"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_consumed_contracts.#", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_contract_masters.#", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_imported_contracts.#", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_intra_epg_contracts.#", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_provided_contracts.#", "0"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.0.vrf_name", "vrf_name_1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "relation_to_vrf.#", "1"),
					resource.TestCheckResourceAttr("aci_endpoint_security_group.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testChildDependencyConfigFvESg = `
resource "aci_endpoint_security_group" "test_endpoint_security_group_0"{
  application_profile_dn = aci_application_profile.test.id
  name = "esg_0"
}
resource "aci_endpoint_security_group" "test_endpoint_security_group_1"{
  application_profile_dn = aci_application_profile.test.id
  name = "esg_1"
}
`

const testConfigFvESgMinDependencyWithFvApAllowExisting = testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "allow_test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
}
resource "aci_endpoint_security_group" "allow_test_2" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
  depends_on = [aci_endpoint_security_group.allow_test]
}
`

const testConfigFvESgMinDependencyWithFvAp = testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
}
`

const testConfigFvESgAllDependencyWithFvAp = testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
  admin_state = "no"
  annotation = "annotation"
  description = "description_1"
  exception_tag = "exception_tag_1"
  intra_esg_isolation = "enforced"
  match_criteria = "All"
  name_alias = "name_alias_1"
  preferred_group_member = "exclude"
}
`

const testConfigFvESgResetDependencyWithFvAp = testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
  admin_state = "no"
  annotation = "orchestrator:terraform"
  description = ""
  exception_tag = ""
  intra_esg_isolation = "unenforced"
  match_criteria = "AtleastOne"
  name_alias = ""
  preferred_group_member = "exclude"
}
`
const testConfigFvESgChildrenDependencyWithFvAp = testChildDependencyConfigFvESg + testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
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
  relation_to_consumed_contracts = [
	{
	  annotation = "annotation_1"
	  contract_name = "contract_name_0"
	  priority = "level1"
	},
	{
	  annotation = "annotation_2"
	  contract_name = "contract_name_1"
	  priority = "level2"
	},
  ]
  relation_to_contract_masters = [
	{
	  annotation = "annotation_1"
	  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
	},
	{
	  annotation = "annotation_2"
	  target_dn = aci_endpoint_security_group.test_endpoint_security_group_1.id
	},
  ]
  relation_to_imported_contracts = [
	{
	  annotation = "annotation_1"
	  imported_contract_name = "imported_contract_name_0"
	  priority = "level1"
	},
	{
	  annotation = "annotation_2"
	  imported_contract_name = "imported_contract_name_1"
	  priority = "level2"
	},
  ]
  relation_to_intra_epg_contracts = [
	{
	  annotation = "annotation_1"
	  contract_name = "contract_name_0"
	},
	{
	  annotation = "annotation_2"
	  contract_name = "contract_name_1"
	},
  ]
  relation_to_provided_contracts = [
	{
	  annotation = "annotation_1"
	  contract_name = "contract_name_0"
	  match_criteria = "All"
	  priority = "level1"
	},
	{
	  annotation = "annotation_2"
	  contract_name = "contract_name_1"
	  match_criteria = "AtleastOne"
	  priority = "level2"
	},
  ]
  relation_to_vrf = [
	{
	  annotation = "annotation_1"
	  vrf_name = "vrf_name_1"
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

const testConfigFvESgChildrenRemoveFromConfigDependencyWithFvAp = testChildDependencyConfigFvESg + testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
}
`

const testConfigFvESgChildrenRemoveOneDependencyWithFvAp = testChildDependencyConfigFvESg + testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  relation_to_consumed_contracts = [ 
	{
	  annotation = "annotation_2"
	  contract_name = "contract_name_1"
	  priority = "level2"
	},
  ]
  relation_to_contract_masters = [ 
	{
	  annotation = "annotation_2"
	  target_dn = aci_endpoint_security_group.test_endpoint_security_group_1.id
	},
  ]
  relation_to_imported_contracts = [ 
	{
	  annotation = "annotation_2"
	  imported_contract_name = "imported_contract_name_1"
	  priority = "level2"
	},
  ]
  relation_to_intra_epg_contracts = [ 
	{
	  annotation = "annotation_2"
	  contract_name = "contract_name_1"
	},
  ]
  relation_to_provided_contracts = [ 
	{
	  annotation = "annotation_2"
	  contract_name = "contract_name_1"
	  match_criteria = "AtleastOne"
	  priority = "level2"
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

const testConfigFvESgChildrenRemoveAllDependencyWithFvAp = testChildDependencyConfigFvESg + testConfigFvApMinDependencyWithFvTenant + `
resource "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
  annotations = []
  relation_to_consumed_contracts = []
  relation_to_contract_masters = []
  relation_to_imported_contracts = []
  relation_to_intra_epg_contracts = []
  relation_to_provided_contracts = []
  tags = []
}
`
