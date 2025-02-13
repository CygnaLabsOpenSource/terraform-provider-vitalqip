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

func dataSourceIPv6Address() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPv6AddressRead,
		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
			},
			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Hostname of address.",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Domain name of address.",
			},
			"range_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Address range.",
			},
			"address_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Type of address.",
			},
			"publish_a": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "AAAA resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL",
			},
			"publish_ptr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "PTR resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL",
			},
			"class_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Class of object.",
			},
			"iaid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Identity Association Identifier (IAID).",
			},
			"range_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of range.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Time to live.",
			},
			"node_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of node.",
			},
			"unique_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Unique ID.",
			},
			"duid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "DUID. Note: IPv6 address using DUID will be dynamic updated to DHCPv6 Server.",
			},
			"use_mac_address": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "True: use MAC address, False: use DUID value instead.",
			},
			"mac_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "MAC address.",
			},
		},
	}
}

func dataSourceIPv6AddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	var err error
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	hostName := strings.TrimSpace(d.Get("host_name").(string))
	address := strings.TrimSpace(d.Get("address").(string))
	query := map[string]string{
		"orgName":        orgName,
		"addressVersion": "6",
	}

	if address != "" {
		query["objectAddr"] = address
	} else if hostName != "" {
		query["objectName"] = hostName
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv6 Address Failed",
			Detail:   "Missing address and host_name field",
		})
		return diags
	}
	log.Println("[DEBUG] Get IPv6 Address: " + fmt.Sprintf("%v", query))

	response, err := objMgr.GetIPv6Address(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv6 Address Failed",
			Detail:   fmt.Sprintf("Getting QIP IPv6 Address failed : %s", err),
		})
		return diags
	}

	if response == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API returns a nil/empty QIP IPv6 Address",
			Detail:   "API returns a nil/empty Address response. Getting QIP IPv6 Address failed",
		})
		return diags
	}

	flattenIPv6Address(d, response)
	log.Println("[DEBUG] Response Get IPv6 Address: " + fmt.Sprintf("%v", response))
	return diags
}

func flattenIPv6Address(d *schema.ResourceData, ipv6Address *en.IPv6Address) {

	d.SetId(ipv6Address.Address)
	d.Set("host_name", ipv6Address.HostName)
	d.Set("address", ipv6Address.Address)
	d.Set("domain_name", ipv6Address.DomainName)
	d.Set("range_address", ipv6Address.RangeAddress)
	d.Set("address_type", ipv6Address.AddressType)
	d.Set("publish_a", ipv6Address.PublishA)
	d.Set("publish_ptr", ipv6Address.PublishPTR)
	d.Set("class_type", ipv6Address.ClassType)
	d.Set("iaid", ipv6Address.IAID)
	d.Set("range_name", ipv6Address.Range)
	d.Set("ttl", ipv6Address.TTL)
	d.Set("node_name", ipv6Address.NodeName)
	d.Set("unique_id", ipv6Address.UniqueID)
	d.Set("duid", ipv6Address.DUID)
	d.Set("use_mac_address", ipv6Address.UseMACAddress)
	d.Set("mac_address", ipv6Address.MACAddress)
}
