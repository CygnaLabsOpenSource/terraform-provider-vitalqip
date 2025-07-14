package vitalqip

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceZone(t *testing.T) {
	dataName := "data.vitalqip_zone.zone_data"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_zone" "zone_resource" {
						org_name = "Demo"
						name = "test.com"
						negative_cache_ttl = 600
						expire_time = 604800
						email = "example@gmail.com"
						retry_time = 3600
						default_ttl = 86400
						refresh_time = 21600
					}
					
					data "vitalqip_zone" "zone_data" {
						org_name = "Demo"
						name = "test.com"
						depends_on = [vitalqip_zone.zone_resource]
					}
					`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(dataName, "name", "test.com"),
					resource.TestCheckResourceAttr(dataName, "email", "example@gmail.com"),
					resource.TestCheckResourceAttr(dataName, "negative_cache_ttl", "600"),
					resource.TestCheckResourceAttr(dataName, "expire_time", "604800"),
					resource.TestCheckResourceAttr(dataName, "retry_time", "3600"),
					resource.TestCheckResourceAttr(dataName, "default_ttl", "86400"),
					resource.TestCheckResourceAttr(dataName, "refresh_time", "21600"),
				),
			},
		},
	})
}
