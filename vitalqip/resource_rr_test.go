package vitalqip

import (
	"fmt"
	cc "terraform-provider-vitalqip/vitalqip/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccRR(t *testing.T) {
	resourceName := "vitalqip_rr.rr_resource"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRRDestroy,
		Steps: []resource.TestStep{
			//  Step 1 create
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_rr" "rr_resource" {
						org_name= "Demo"
						owner="owner"
						rr_type="A"
						data1="9.9.9.9"
						infra_type="ZONE"
						infra_fqdn="com"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRRExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "owner", "owner"),
					resource.TestCheckResourceAttr(resourceName, "rr_type", "A"),
					resource.TestCheckResourceAttr(resourceName, "data1", "9.9.9.9"),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "ZONE"),
					resource.TestCheckResourceAttr(resourceName, "infra_fqdn", "com"),
				),
			},
			// Step 2 update
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_rr" "rr_resource" {
						org_name= "Demo"
						owner="owner"
						rr_type="A"
						data1="8.8.8.8"
						infra_type="ZONE"
						infra_fqdn="com"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRRExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "owner", "owner"),
					resource.TestCheckResourceAttr(resourceName, "rr_type", "A"),
					resource.TestCheckResourceAttr(resourceName, "data1", "8.8.8.8"),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "ZONE"),
					resource.TestCheckResourceAttr(resourceName, "infra_fqdn", "com"),
				),
			},

			// step 3 Update address

			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_rr" "rr_resource" {
						org_name= "Demo"
						owner="owner"
						rr_type="A"
						data1="9.9.9.9"
						infra_type="ZONE"
						infra_fqdn="test"
					}`,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRRExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(resourceName, "owner", "owner"),
					resource.TestCheckResourceAttr(resourceName, "rr_type", "A"),
					resource.TestCheckResourceAttr(resourceName, "data1", "9.9.9.9"),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "ZONE"),
					resource.TestCheckResourceAttr(resourceName, "infra_fqdn", "test"),
				),
			},
		},
	})
}

func testAccCheckRRExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		connector := testAccProvider.Meta().(*cc.Connector)
		objMgr := cc.NewObjectManager(connector)

		query := map[string]string{
			"orgName": rs.Primary.Attributes["org_name"],
			"name":    rs.Primary.Attributes["infra_fqdn"],
			"type":    rs.Primary.Attributes["infra_type"],
		}

		_, err := objMgr.GetRR(query)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckRRDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vitalqip_rr" {
			continue
		}

		query := map[string]string{
			"orgName": rs.Primary.Attributes["org_name"],
			"rrId":    rs.Primary.ID,
		}

		_, err := objMgr.GetRR(query)
		if err == nil {
			return fmt.Errorf("Resource Record still exists")
		}
	}

	return nil
}
