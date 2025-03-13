package vitalqip

import (
	"fmt"
	cc "terraform-provider-vitalqip/vitalqip/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIPv4Address(t *testing.T) {
	resourceName := "vitalqip_ipv4_address.ipv4_address_resource"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPv4AddressDestroy,
		Steps: []resource.TestStep{
			//  Step 1 create
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv4_address" "ipv4_address_resource" {
						org_name= "Demo"
						object_addr = "123.0.0.5"
						dynamic_config = "Static"
						object_class = "Workstation"
						object_desc = "desc"
						object_name = "obj5"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv4AddressExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "object_addr", "123.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_config", "Static"),
					resource.TestCheckResourceAttr(resourceName, "object_class", "Workstation"),
					resource.TestCheckResourceAttr(resourceName, "object_desc", "desc"),
					resource.TestCheckResourceAttr(resourceName, "object_name", "obj5"),
				),
			},
			// Step 2 update
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv4_address" "ipv4_address_resource" {
						org_name= "Demo"
						object_addr = "123.0.0.5"
						dynamic_config = "Static"
						object_class = "Workstation"
						object_desc = "desc update"
						object_name = "objNew"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv4AddressExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "object_addr", "123.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_config", "Static"),
					resource.TestCheckResourceAttr(resourceName, "object_class", "Workstation"),
					resource.TestCheckResourceAttr(resourceName, "object_desc", "desc update"),
					resource.TestCheckResourceAttr(resourceName, "object_name", "objNew"),
				),
			},

			// step 3 Update object_addr

			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv4_address" "ipv4_address_resource" {
						org_name= "Demo"
						object_addr = "123.0.0.8"
						dynamic_config = "Static"
						object_class = "Workstation"
						object_desc = "desc update"
						object_name = "objNew"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv4AddressExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "object_addr", "123.0.0.8"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_config", "Static"),
					resource.TestCheckResourceAttr(resourceName, "object_class", "Workstation"),
					resource.TestCheckResourceAttr(resourceName, "object_desc", "desc update"),
					resource.TestCheckResourceAttr(resourceName, "object_name", "objNew"),
				),
			},
		},
	})
}

func testAccCheckIPv4AddressExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Address ID is set")
		}

		connector := testAccProvider.Meta().(*cc.Connector)
		objMgr := cc.NewObjectManager(connector)

		query := map[string]string{
			"orgName":        rs.Primary.Attributes["org_name"],
			"objectAddr":     rs.Primary.Attributes["object_addr"],
			"addressVersion": "4",
		}

		_, err := objMgr.GetIPv4Address(query)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIPv4AddressDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vitalqip_ipv4_address" {
			continue
		}

		query := map[string]string{
			"orgName":        rs.Primary.Attributes["org_name"],
			"objectAddr":     rs.Primary.Attributes["object_addr"],
			"addressVersion": "4",
		}

		_, err := objMgr.GetIPv4Address(query)
		if err == nil {
			return fmt.Errorf("Address still exists")
		}
	}

	return nil
}
