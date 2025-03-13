package vitalqip

import (

	// "regexp"

	"context"
	"fmt"
	"log"
	"strings"

	en "terraform-provider-vitalqip/vitalqip/entities"
	cc "terraform-provider-vitalqip/vitalqip/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPv4Address() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPv4AddressRead,
		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
			},
			"object_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"object_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of IPv4 object.",
			},
			"subnet_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Subnet of IPv4 object.",
			},
			"object_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Class type of IPv4 object.",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Domain of IPv4 object.",
			},
			"object_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Description of IPv4 object.",
			},
			"dynamic_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Dynamic Configuration of IPv4 object.",
			},
			"mac_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "MAC address of IPv4 object.",
			},
			"a_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Time to live value of default A resource record.",
			},
			"ptr_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Time to live value of default PTR resource record.",
			},
			"publish_a": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A resource record option: Always, None, Push Only.",
			},
			"publish_ptr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "PTR resource record option: Always, None, Push Only.",
			},
		},
	}
}

func dataSourceIPv4AddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	var err error
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	objectAddr := strings.TrimSpace(d.Get("object_addr").(string))
	objectName := strings.TrimSpace(d.Get("object_name").(string))

	query := map[string]string{
		"orgName":        orgName,
		"addressVersion": "4",
	}

	if objectAddr != "" {
		query["objectAddr"] = objectAddr
	} else if objectName != "" {
		query["objectName"] = objectName
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv4 Address Failed",
			Detail:   "Missing object_addr and object_name field",
		})
		return diags
	}
	log.Println("[DEBUG] Get IPv4 Address: " + fmt.Sprintf("%v", query))

	response, err := objMgr.GetIPv4Address(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv4 Address Failed",
			Detail:   fmt.Sprintf("Getting QIP IPv4 Address failed : %s", err),
		})
		return diags
	}

	if response == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API returns a nil/empty QIP IPv4 Address",
			Detail:   "API returns a nil/empty Address response. Getting QIP IPv4 Address failed",
		})
		return diags
	}

	flattenIPv4Address(d, response)
	log.Println("[DEBUG] Response Get IPv4 Address: " + fmt.Sprintf("%v", response))
	return diags
}

func flattenIPv4Address(d *schema.ResourceData, ipv4Address *en.IPv4Address) {

	d.SetId(ipv4Address.ObjectAddr)
	d.Set("object_addr", ipv4Address.ObjectAddr)
	d.Set("object_name", ipv4Address.ObjectName)
	d.Set("subnet_addr", ipv4Address.SubnetAddr)
	d.Set("object_class", ipv4Address.ObjectClass)
	d.Set("domain_name", ipv4Address.DomainName)
	d.Set("object_desc", ipv4Address.ObjectDesc)
	d.Set("dynamic_config", ipv4Address.DynamicConfig)
	d.Set("mac_addr", ipv4Address.MacAddr)
	d.Set("a_ttl", ipv4Address.ATTL)
	d.Set("ptr_ttl", ipv4Address.PTRTTL)
	d.Set("publish_a", ipv4Address.PublishA)
	d.Set("publish_ptr", ipv4Address.PublishPTR)
}
