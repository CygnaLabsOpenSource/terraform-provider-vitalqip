package vitalqip

import (
	"fmt"
	cc "terraform-provider-vitalqip/vitalqip/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccZone(t *testing.T) {
	resourceName := "vitalqip_zone.zone_resource"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			//  Step 1 create
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
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "name", "test.com"),
					resource.TestCheckResourceAttr(resourceName, "email", "example@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "negative_cache_ttl", "600"),
					resource.TestCheckResourceAttr(resourceName, "expire_time", "604800"),
					resource.TestCheckResourceAttr(resourceName, "retry_time", "3600"),
					resource.TestCheckResourceAttr(resourceName, "default_ttl", "86400"),
					resource.TestCheckResourceAttr(resourceName, "refresh_time", "21600"),
				),
			},
			// Step 2 update name and email
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_zone" "zone_resource" {
						org_name = "Demo"
						name = "testNew.com"
						negative_cache_ttl = 600
						expire_time = 604800
						email = "update@gmail.com"
						retry_time = 3600
						default_ttl = 86400
						refresh_time = 21600
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "name", "testNew.com"),
					resource.TestCheckResourceAttr(resourceName, "email", "update@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "negative_cache_ttl", "600"),
					resource.TestCheckResourceAttr(resourceName, "expire_time", "604800"),
					resource.TestCheckResourceAttr(resourceName, "retry_time", "3600"),
					resource.TestCheckResourceAttr(resourceName, "default_ttl", "86400"),
					resource.TestCheckResourceAttr(resourceName, "refresh_time", "21600"),
				),
			},
		},
	})
}

func testAccCheckZoneExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		connector := testAccProvider.Meta().(*cc.Connector)
		objMgr := cc.NewObjectManager(connector)

		query := map[string]string{
			"orgName":  rs.Primary.Attributes["org_name"],
			"zoneName": rs.Primary.Attributes["name"],
		}

		_, err := objMgr.GetZone(query)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckZoneDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vitalqip_zone" {
			continue
		}

		query := map[string]string{
			"orgName":  rs.Primary.Attributes["org_name"],
			"zoneName": rs.Primary.Attributes["name"],
		}

		_, err := objMgr.GetZone(query)
		if err == nil {
			return fmt.Errorf("Zone still exists")
		}
	}

	return nil
}
