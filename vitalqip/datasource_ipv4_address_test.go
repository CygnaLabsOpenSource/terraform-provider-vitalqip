package vitalqip

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIPv4Address(t *testing.T) {
	dataName := "data.vitalqip_ipv4_address.ipv4_address_data"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
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
					}
					
					data "vitalqip_ipv4_address" "ipv4_address_data" {
						org_name= "Demo"
						object_addr="123.0.0.5"
						depends_on = [vitalqip_ipv4_address.ipv4_address_resource]
					}
					`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(dataName, "object_addr", "123.0.0.5"),
					resource.TestCheckResourceAttr(dataName, "dynamic_config", "Static"),
					resource.TestCheckResourceAttr(dataName, "object_class", "Workstation"),
					resource.TestCheckResourceAttr(dataName, "object_desc", "desc"),
					resource.TestCheckResourceAttr(dataName, "object_name", "obj5"),
				),
			},
		},
	})
}
