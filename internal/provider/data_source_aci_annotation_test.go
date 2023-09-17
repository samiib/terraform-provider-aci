// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceTagAnnotationWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigTagAnnotationDataSourceDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("data.aci_annotation.test", "value", "test_value"),
				),
			},
		},
	})
}

const testConfigTagAnnotationDataSourceDependencyWithFvAEPg = testConfigTagAnnotationMinDependencyWithFvAEPg + `
data "aci_annotation" "test" {
  parent_dn = aci_application_epg.test.id
  key = "test_key"
  depends_on = [aci_annotation.test]
}
`
