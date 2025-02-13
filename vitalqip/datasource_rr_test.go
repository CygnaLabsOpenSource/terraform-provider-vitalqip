package vitalqip

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRR(t *testing.T) {
	dataName := "data.vitalqip_rr.rr_data"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigWithProvider(
					`
					data "vitalqip_rr" "rr_data" {
						org_name= "Demo"
						infra_type="ZONE"
						infra_fqdn="com"
						rr_id=111
					}
					`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(dataName, "owner", "owner"),
					resource.TestCheckResourceAttr(dataName, "rr_type", "A"),
					resource.TestCheckResourceAttr(dataName, "data1", "9.9.9.9"),
					resource.TestCheckResourceAttr(dataName, "infra_type", "ZONE"),
					resource.TestCheckResourceAttr(dataName, "infra_fqdn", "com"),
				),
			},
		},
	})
}
