package vitalqip

import (
	"fmt"
	cc "terraform-provider-vitalqip/vitalqip/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIPv6Address(t *testing.T) {
	resourceName := "vitalqip_ipv6_address.ipv6_address_resource"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPv6AddressDestroy,
		Steps: []resource.TestStep{
			//  Step 1 create
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv6_address" "ipv6_address_resource" {
						org_name= "Demo"
						host_name="v6obj"
						address="2000::5"
						domain_name="com"
						range_address="2000::/112"
						address_type="STATIC"
						publish_a="ALWAYS"
						publish_ptr="ALWAYS"
						class_type="Workstation"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv6AddressExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "host_name", "v6obj"),
					resource.TestCheckResourceAttr(resourceName, "address", "2000::5"),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "com"),
					resource.TestCheckResourceAttr(resourceName, "range_address", "2000::/112"),
					resource.TestCheckResourceAttr(resourceName, "address_type", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "publish_a", "ALWAYS"),
					resource.TestCheckResourceAttr(resourceName, "publish_ptr", "ALWAYS"),
					resource.TestCheckResourceAttr(resourceName, "class_type", "Workstation"),
				),
			},
			// Step 2 update
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv6_address" "ipv6_address_resource" {
						org_name= "Demo"
						host_name="v6objNew"
						address="2000::5"
						domain_name="com"
						range_address="2000::/112"
						address_type="STATIC"
						publish_a="ALWAYS"
						publish_ptr="ALWAYS"
						class_type="Workstation"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv6AddressExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "host_name", "v6objNew"),
					resource.TestCheckResourceAttr(resourceName, "address", "2000::5"),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "com"),
					resource.TestCheckResourceAttr(resourceName, "range_address", "2000::/112"),
					resource.TestCheckResourceAttr(resourceName, "address_type", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "publish_a", "ALWAYS"),
					resource.TestCheckResourceAttr(resourceName, "publish_ptr", "ALWAYS"),
					resource.TestCheckResourceAttr(resourceName, "class_type", "Workstation"),
				),
			},

			// step 3 Update address

			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv6_address" "ipv6_address_resource" {
						org_name= "Demo"
						host_name="v6objNew"
						address="2000::8"
						domain_name="com"
						range_address="2000::/112"
						address_type="STATIC"
						publish_a="ALWAYS"
						publish_ptr="ALWAYS"
						class_type="Workstation"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv6AddressExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "host_name", "v6objNew"),
					resource.TestCheckResourceAttr(resourceName, "address", "2000::8"),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "com"),
					resource.TestCheckResourceAttr(resourceName, "range_address", "2000::/112"),
					resource.TestCheckResourceAttr(resourceName, "address_type", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "publish_a", "ALWAYS"),
					resource.TestCheckResourceAttr(resourceName, "publish_ptr", "ALWAYS"),
					resource.TestCheckResourceAttr(resourceName, "class_type", "Workstation"),
				),
			},
		},
	})
}

func testAccCheckIPv6AddressExists(n string) resource.TestCheckFunc {
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
			"objectAddr":     rs.Primary.Attributes["address"],
			"addressVersion": "6",
		}

		_, err := objMgr.GetIPv6Address(query)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIPv6AddressDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vitalqip_ipv6_address" {
			continue
		}

		query := map[string]string{
			"orgName":        rs.Primary.Attributes["org_name"],
			"objectAddr":     rs.Primary.Attributes["address"],
			"addressVersion": "6",
		}

		_, err := objMgr.GetIPv6Address(query)
		if err == nil {
			return fmt.Errorf("Address still exists")
		}
	}

	return nil
}
