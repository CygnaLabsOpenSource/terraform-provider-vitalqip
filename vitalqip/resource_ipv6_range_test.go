package vitalqip

import (
	"fmt"
	cc "terraform-provider-vitalqip/vitalqip/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIPv6Range(t *testing.T) {
	resourceName := "vitalqip_ipv6_range.ipv6_range_resource"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPv6RangeDestroy,
		Steps: []resource.TestStep{
			//  Step 1 create
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv6_range" "ipv6_range_resource" {
						org_name= "Demo"
						name="range"
						start_address="2000::4:0"
						range_prefix_length=112
						range_type="DYNAMIC"
						address_selection="NEXT_AVAILABLE"
						subnet_prefix_length=60
						subnet_address="2000::"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv6RangeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "name", "range"),
					resource.TestCheckResourceAttr(resourceName, "start_address", "2000::4:0"),
					resource.TestCheckResourceAttr(resourceName, "range_prefix_length", "112"),
					resource.TestCheckResourceAttr(resourceName, "range_type", "DYNAMIC"),
					resource.TestCheckResourceAttr(resourceName, "address_selection", "NEXT_AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "subnet_prefix_length", "60"),
					resource.TestCheckResourceAttr(resourceName, "subnet_address", "2000::"),
				),
			},
			// Step 2 update
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv6_range" "ipv6_range_resource" {
						org_name= "Demo"
						name="rangeNew"
						start_address="2000::4:0"
						range_prefix_length=112
						range_type="DYNAMIC"
						address_selection="NEXT_AVAILABLE"
						subnet_prefix_length=60
						subnet_address="2000::"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv6RangeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "name", "rangeNew"),
					resource.TestCheckResourceAttr(resourceName, "start_address", "2000::4:0"),
					resource.TestCheckResourceAttr(resourceName, "range_prefix_length", "112"),
					resource.TestCheckResourceAttr(resourceName, "range_type", "DYNAMIC"),
					resource.TestCheckResourceAttr(resourceName, "address_selection", "NEXT_AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "subnet_prefix_length", "60"),
					resource.TestCheckResourceAttr(resourceName, "subnet_address", "2000::"),
				),
			},

			// step 3 Update address

			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv6_range" "ipv6_range_resource" {
						org_name= "Demo"
						name="rangeNew"
						start_address="2000::6:0"
						range_prefix_length=112
						range_type="DYNAMIC"
						address_selection="NEXT_AVAILABLE"
						subnet_prefix_length=60
						subnet_address="2000::"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPv6RangeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "name", "rangeNew"),
					resource.TestCheckResourceAttr(resourceName, "start_address", "2000::6:0"),
					resource.TestCheckResourceAttr(resourceName, "range_prefix_length", "112"),
					resource.TestCheckResourceAttr(resourceName, "range_type", "DYNAMIC"),
					resource.TestCheckResourceAttr(resourceName, "address_selection", "NEXT_AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "subnet_prefix_length", "60"),
					resource.TestCheckResourceAttr(resourceName, "subnet_address", "2000::"),
				),
			},
		},
	})
}

func testAccCheckIPv6RangeExists(n string) resource.TestCheckFunc {
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
			"orgName": rs.Primary.Attributes["org_name"],
			"address": rs.Primary.Attributes["start_address"],
		}

		_, err := objMgr.GetIPv6Range(query)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIPv6RangeDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vitalqip_ipv6_range" {
			continue
		}

		query := map[string]string{
			"orgName": rs.Primary.Attributes["org_name"],
			"address": rs.Primary.Attributes["start_address"],
		}

		_, err := objMgr.GetIPv6Range(query)
		if err == nil {
			return fmt.Errorf("IPv6 Range still exists")
		}
	}

	return nil
}
