package vitalqip

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIPv6Address(t *testing.T) {
	dataName := "data.vitalqip_ipv6_address.ipv6_address_data"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
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
					}
					
					data "vitalqip_ipv6_address" "ipv6_address_data" {
						org_name= "Demo"
						address="2000::5"
						depends_on = [vitalqip_ipv6_address.ipv6_address_resource]
					}
					`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(dataName, "host_name", "v6obj"),
					resource.TestCheckResourceAttr(dataName, "address", "2000::5"),
					resource.TestCheckResourceAttr(dataName, "domain_name", "com"),
					resource.TestCheckResourceAttr(dataName, "range_address", "2000::/112"),
					resource.TestCheckResourceAttr(dataName, "address_type", "STATIC"),
					resource.TestCheckResourceAttr(dataName, "publish_a", "ALWAYS"),
					resource.TestCheckResourceAttr(dataName, "publish_ptr", "ALWAYS"),
					resource.TestCheckResourceAttr(dataName, "class_type", "Workstation"),
				),
			},
		},
	})
}
