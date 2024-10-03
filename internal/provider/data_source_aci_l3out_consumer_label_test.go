// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceL3extConsLblWithL3extOut(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "2.0(1m)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigL3extConsLblDataSourceDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "owner", "infra"),
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_l3out_consumer_label.test", "tag", "yellow-green"),
				),
			},
			{
				Config:      testConfigL3extConsLblNotExistingL3extOut,
				ExpectError: regexp.MustCompile("Failed to read aci_l3out_consumer_label data source"),
			},
		},
	})
}

const testConfigL3extConsLblDataSourceDependencyWithL3extOut = testConfigL3extConsLblMinDependencyWithL3extOut + `
data "aci_l3out_consumer_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "test_name"
  depends_on = [aci_l3out_consumer_label.test]
}
`

const testConfigL3extConsLblNotExistingL3extOut = testConfigL3extOutMin + `
data "aci_l3out_consumer_label" "test_non_existing" {
  parent_dn = aci_l3_outside.test.id
  name = "non_existing_name"
}
`
