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

func dataSourceIPv6Range() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPv6RangeRead,
		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Range name.",
			},
			"start_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Starting IPv6 address.",
			},
			"range_prefix_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Prefix length of range.",
			},
			"range_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Type of range: DYNAMIC, FIXED, RESERVED, DYNAMIC_TEMPORARY",
			},
			"stand_prim_dhcp_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of standard primary DHCP server.",
			},
			"failover_second_dhcp_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of failover Secondary DHCP server.",
			},
			"opt_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Template option.",
			},
			"address_selection": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "NEXT_AVAILABLE, RANDOM",
			},
			"subnet_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of subnet.",
			},
			"subnet_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Address of subnet.",
			},
			"subnet_prefix_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Prefix length of subnet.",
			},
			"udas": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "List of UDAs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of the UDA.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Value of the UDA.",
						},
					},
				},
			},
			"groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of groups UDA.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of the group.",
						},
						"udas": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of UDAs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Name of the UDA.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Value of the UDA.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIPv6RangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	var err error
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	address := strings.TrimSpace(d.Get("start_address").(string))

	query := map[string]string{
		"orgName": orgName,
		"address": address,
	}
	log.Println("[DEBUG] Get IPv6 Range: " + fmt.Sprintf("%v", query))

	ipv6RangeResponse, err := objMgr.GetIPv6Range(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv6 Range Failed",
			Detail:   fmt.Sprintf("Getting QIP IPv6 Range failed : %s", err),
		})
		return diags
	}

	if ipv6RangeResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API returns a nil/empty QIP IPv6 Range",
			Detail:   "API returns a nil/empty Address response. Getting QIP IPv6 Range failed",
		})
		return diags
	}

	flattenIPv6Range(d, &ipv6RangeResponse.List[0])
	log.Println("[DEBUG] Response Get IPv6 Range: " + fmt.Sprintf("%v", &ipv6RangeResponse.List[0]))
	return diags
}

func flattenIPv6Range(d *schema.ResourceData, ipv6Range *en.IPv6Range) {

	d.SetId(ipv6Range.StartAddress)
	d.Set("name", ipv6Range.Name)
	d.Set("range_prefix_length", ipv6Range.RangePrefixLength)
	d.Set("range_type", ipv6Range.RangeType)
	d.Set("stand_prim_dhcp_server", ipv6Range.StandPrimDHCPServer)
	d.Set("failover_second_dhcp_server", ipv6Range.FailoverSecondDHCPServer)
	d.Set("opt_template", ipv6Range.OptTemplate)
	d.Set("address_selection", ipv6Range.AddressSelection)
	d.Set("subnet_name", ipv6Range.SubnetName)
	d.Set("subnet_address", ipv6Range.SubnetAddress)
	d.Set("subnet_prefix_length", ipv6Range.SubnetPrefixLength)

	var udaList []interface{}
	for _, uda := range ipv6Range.UDAs {
		udaMap := map[string]interface{}{
			"name":  uda.Name,
			"value": uda.Value,
		}
		udaList = append(udaList, udaMap)
	}

	d.Set("udas", udaList)

	var groupList []interface{}
	for _, group := range ipv6Range.Groups {
		var udaList []interface{}
		for _, uda := range group.UDAs {
			udaMap := map[string]interface{}{
				"name":  uda.Name,
				"value": uda.Value,
			}
			udaList = append(udaList, udaMap)
		}

		groupMap := map[string]interface{}{
			"name": group.Name,
			"udas": udaList,
		}
		groupList = append(groupList, groupMap)
	}

	d.Set("groups", groupList)
}
