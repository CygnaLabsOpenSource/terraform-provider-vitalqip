package vitalqip

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIPv6Range(t *testing.T) {
	dataName := "data.vitalqip_ipv6_range.ipv6_range_data"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
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
					}
					
					data "vitalqip_ipv6_range" "ipv6_range_data" {
						org_name= "Demo"
						start_address="2000::4:0"
						depends_on = [vitalqip_ipv6_range.ipv6_range_resource]
					}
					`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(dataName, "name", "range"),
					resource.TestCheckResourceAttr(dataName, "start_address", "2000::4:0"),
					resource.TestCheckResourceAttr(dataName, "range_prefix_length", "112"),
					resource.TestCheckResourceAttr(dataName, "range_type", "DYNAMIC"),
					resource.TestCheckResourceAttr(dataName, "address_selection", "NEXT_AVAILABLE"),
					resource.TestCheckResourceAttr(dataName, "subnet_prefix_length", "60"),
					resource.TestCheckResourceAttr(dataName, "subnet_address", "2000::"),
				),
			},
		},
	})
}
