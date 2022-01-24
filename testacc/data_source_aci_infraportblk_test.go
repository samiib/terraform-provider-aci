package testacc

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAciAccessPortBlockDataSource_Basic(t *testing.T) {
	resourceName := "aci_access_port_block.test"
	dataSourceName := "data.aci_access_port_block.test"
	randomParameter := acctest.RandStringFromCharSet(10, "abcdefghijklmnopqrstuvwxyz")
	randomValue := acctest.RandString(10)
	rName := makeTestVariable(acctest.RandString(5))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckAciAccessPortBlockDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccessPortBlockDSWithoutRequired(rName, rName, rName, "access_port_selector_dn"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccessPortBlockDSWithoutRequired(rName, rName, rName, "name"),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccAccessPortBlockConfigDataSource(rName, rName, rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "access_port_selector_dn", resourceName, "access_port_selector_dn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "annotation", resourceName, "annotation"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name_alias", resourceName, "name_alias"),
					resource.TestCheckResourceAttrPair(dataSourceName, "from_card", resourceName, "from_card"),
					resource.TestCheckResourceAttrPair(dataSourceName, "from_port", resourceName, "from_port"),
					resource.TestCheckResourceAttrPair(dataSourceName, "to_card", resourceName, "to_card"),
					resource.TestCheckResourceAttrPair(dataSourceName, "to_port", resourceName, "to_port"),
				),
			},
			{
				Config:      CreateAccAccessPortBlockDataSourceUpdate(rName, rName, rName, randomParameter, randomValue),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},

			{
				Config:      CreateAccAccessPortBlockDSWithInvalidParentDn(rName, rName, rName),
				ExpectError: regexp.MustCompile(`(.)+ Object may not exists`),
			},

			{
				Config: CreateAccAccessPortBlockDataSourceUpdatedResource(rName, rName, rName, "annotation", "orchestrator:terraform-testacc"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "annotation", resourceName, "annotation"),
				),
			},
		},
	})
}

func CreateAccAccessPortBlockConfigDataSource(infraAccPortPName, infraHPortSName, rName string) string {
	fmt.Println("=== STEP  testing access_port_block Data Source with required arguments only")
	resource := fmt.Sprintf(`
	
	resource "aci_leaf_interface_profile" "test" {
		name 		= "%s"
	
	}
	
	resource "aci_access_port_selector" "test" {
		name 		= "%s"
		leaf_interface_profile_dn = aci_leaf_interface_profile.test.id
		access_port_selector_type  = "ALL"
	}
	
	resource "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = "%s"
	}

	data "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = aci_access_port_block.test.name
		depends_on = [ aci_access_port_block.test ]
	}
	`, infraAccPortPName, infraHPortSName, rName)
	return resource
}

func CreateAccessPortBlockDSWithoutRequired(infraAccPortPName, infraHPortSName, rName, attrName string) string {
	fmt.Println("=== STEP  Basic: testing access_port_block Data Source without ", attrName)
	rBlock := `
	
	resource "aci_leaf_interface_profile" "test" {
		name 		= "%s"
	
	}
	
	resource "aci_access_port_selector" "test" {
		name 		= "%s"
		leaf_interface_profile_dn = aci_leaf_interface_profile.test.id
		access_port_selector_type  = "ALL"
	}
	
	resource "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = "%s"
	}
	`
	switch attrName {
	case "access_port_selector_dn":
		rBlock += `
	data "aci_access_port_block" "test" {
	#	access_port_selector_dn  = aci_access_port_selector.test.id
		name  = "%s"
		depends_on = [ aci_access_port_block.test ]
	}
		`
	case "name":
		rBlock += `
	data "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
	#	name  = "%s"
		depends_on = [ aci_access_port_block.test ]
	}
		`
	}
	return fmt.Sprintf(rBlock, infraAccPortPName, infraHPortSName, rName)
}

func CreateAccAccessPortBlockDSWithInvalidParentDn(infraAccPortPName, infraHPortSName, rName string) string {
	fmt.Println("=== STEP  testing access_port_block Data Source with Invalid Parent Dn")
	resource := fmt.Sprintf(`
	
	resource "aci_leaf_interface_profile" "test" {
		name 		= "%s"
	
	}
	
	resource "aci_access_port_selector" "test" {
		name 		= "%s"
		leaf_interface_profile_dn = aci_leaf_interface_profile.test.id
		access_port_selector_type  = "ALL"
	}
	
	resource "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = "%s"
	}

	data "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = "${aci_access_port_block.test.name}_invalid"
		depends_on = [ aci_access_port_block.test ]
	}
	`, infraAccPortPName, infraHPortSName, rName)
	return resource
}

func CreateAccAccessPortBlockDataSourceUpdate(infraAccPortPName, infraHPortSName, rName, key, value string) string {
	fmt.Println("=== STEP  testing access_port_block Data Source with random attribute")
	resource := fmt.Sprintf(`
	
	resource "aci_leaf_interface_profile" "test" {
		name 		= "%s"
	
	}
	
	resource "aci_access_port_selector" "test" {
		name 		= "%s"
		leaf_interface_profile_dn = aci_leaf_interface_profile.test.id
		access_port_selector_type  = "ALL"
	}
	
	resource "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = "%s"
	}

	data "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = aci_access_port_block.test.name
		%s = "%s"
		depends_on = [ aci_access_port_block.test ]
	}
	`, infraAccPortPName, infraHPortSName, rName, key, value)
	return resource
}

func CreateAccAccessPortBlockDataSourceUpdatedResource(infraAccPortPName, infraHPortSName, rName, key, value string) string {
	fmt.Println("=== STEP  testing access_port_block Data Source with updated resource")
	resource := fmt.Sprintf(`
	
	resource "aci_leaf_interface_profile" "test" {
		name 		= "%s"
	
	}
	
	resource "aci_access_port_selector" "test" {
		name 		= "%s"
		leaf_interface_profile_dn = aci_leaf_interface_profile.test.id
		access_port_selector_type  = "ALL"
	}
	
	resource "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = "%s"
		%s = "%s"
	}

	data "aci_access_port_block" "test" {
		access_port_selector_dn  = aci_access_port_selector.test.id
		name  = aci_access_port_block.test.name
		depends_on = [ aci_access_port_block.test ]
	}
	`, infraAccPortPName, infraHPortSName, rName, key, value)
	return resource
}
