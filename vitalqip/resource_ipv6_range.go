package vitalqip

import (
	"context"
	"fmt"
	"log"

	// "strconv"

	"strings"
	en "terraform-provider-vitalqip/vitalqip/entities"
	cc "terraform-provider-vitalqip/vitalqip/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPv6Range() *schema.Resource {
	return &schema.Resource{
		CreateContext: createIPv6RangeRecord,
		ReadContext:   getIPv6RangeRecord,
		UpdateContext: updateIPv6RangeRecord,
		DeleteContext: deleteIPv6RangeRecord,

		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Range name.",
			},
			"start_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Starting IPv6 Range.",
				ForceNew:    true,
			},
			"new_start_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "New starting IPv6 address.",
			},
			"range_prefix_length": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Prefix length of range.",
			},
			"range_type": {
				Type:        schema.TypeString,
				Required:    true,
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
				Required:    true,
				Description: "NEXT_AVAILABLE, RANDOM",
			},
			"subnet_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of subnet.",
				ForceNew:    true,
			},
			"subnet_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Address of subnet.",
				ForceNew:    true,
			},
			"subnet_prefix_length": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Prefix length of subnet.",
				ForceNew:    true,
			},
			"dhcp_params": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"information_refresh_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"preferred_life_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"rebinding_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"renewal_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"valid_life_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bcmcs_server_address_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bcmcs_server_domain_name_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"dns_recursive_name_server": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain_search_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"nis_servers": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"nisp_domain_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"pana_authentication_agents": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"posix_time_zone": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sip_servers_domain_name_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sip_servers_ipv6_address_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sntp_servers": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tzdb_time_zone": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vendor_options": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
				Computed:    true,
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
					},
				},
			},
		},
	}
}

func createIPv6RangeRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	ipv6Range := getIPv6RangeFromResourceData(d)
	newStartAddress := strings.TrimSpace(d.Get("new_start_address").(string))
	if newStartAddress != "" {
		ipv6Range.StartAddress = newStartAddress
	}
	log.Println("[DEBUG] Create IPv6 Range: " + fmt.Sprintf("%v", ipv6Range))

	_, err = objMgr.CreateIPv6Range(ipv6Range)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Creation of QIP IPv6 Range failed",
			Detail:   fmt.Sprintf("Creation of QIP IPv6 Range (%s) failed: %s", ipv6Range.StartAddress, err),
		})
		return diags
	}

	d.SetId(ipv6Range.StartAddress)
	return getIPv6RangeRecord(ctx, d, m)
}

func getIPv6RangeRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
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

func updateIPv6RangeRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	ipv6Range := getIPv6RangeFromResourceData(d)
	log.Println("[DEBUG] Update IPv6 Range: " + fmt.Sprintf("%v", ipv6Range))

	_, err = objMgr.UpdateIPv6Range(ipv6Range)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Updating of QIP IPv6 Range failed",
			Detail:   fmt.Sprintf("Updating QIP IPv6 Range by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return getIPv6RangeRecord(ctx, d, m)
}

func deleteIPv6RangeRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	address := strings.TrimSpace(d.Get("start_address").(string))
	query := map[string]string{
		"orgName": orgName,
		"address": address,
	}
	log.Println("[DEBUG] Delete IPv6 Range: " + fmt.Sprintf("%v", query))

	err := objMgr.DeleteIPv6Range(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Deletion of QIP IPv6 Range failed",
			Detail:   fmt.Sprintf("Deleting QIP IPv6 Range by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return diags
}

func getIPv6RangeFromResourceData(d *schema.ResourceData) *en.IPv6Range {

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	name := strings.TrimSpace(d.Get("name").(string))
	startAddress := strings.TrimSpace(d.Get("start_address").(string))
	newStartAddress := strings.TrimSpace(d.Get("new_start_address").(string))
	rangePrefixLength := d.Get("range_prefix_length").(int)
	rangeType := strings.TrimSpace(d.Get("range_type").(string))
	standPrimDHCPServer := strings.TrimSpace(d.Get("stand_prim_dhcp_server").(string))
	failoverSecondDHCPServer := strings.TrimSpace(d.Get("failover_second_dhcp_server").(string))
	optTemplate := strings.TrimSpace(d.Get("opt_template").(string))
	addressSelection := strings.TrimSpace(d.Get("address_selection").(string))
	subnetName := strings.TrimSpace(d.Get("subnet_name").(string))
	subnetAddress := strings.TrimSpace(d.Get("subnet_address").(string))
	subnetPrefixLength := d.Get("subnet_prefix_length").(int)

	ipv6Range := en.NewIPv6Range(en.IPv6Range{
		OrgName:                  orgName,
		Name:                     name,
		StartAddress:             startAddress,
		NewStartAddress:          newStartAddress,
		RangePrefixLength:        rangePrefixLength,
		RangeType:                rangeType,
		StandPrimDHCPServer:      standPrimDHCPServer,
		FailoverSecondDHCPServer: failoverSecondDHCPServer,
		OptTemplate:              optTemplate,
		AddressSelection:         addressSelection,
		SubnetName:               subnetName,
		SubnetAddress:            subnetAddress,
		SubnetPrefixLength:       subnetPrefixLength,
	})

	if dhcpParamsList, ok := d.Get("dhcp_params").([]interface{}); ok && len(dhcpParamsList) > 0 {
		dhcpParamsMap := dhcpParamsList[0].(map[string]interface{})
		dhcpParams := en.DHCPParams{
			InformationRefreshTime:    dhcpParamsMap["information_refresh_time"].(string),
			PreferredLifeTime:         dhcpParamsMap["preferred_life_time"].(string),
			RebindingTime:             dhcpParamsMap["rebinding_time"].(string),
			RenewalTime:               dhcpParamsMap["renewal_time"].(string),
			ValidLifeTime:             dhcpParamsMap["valid_life_time"].(string),
			BCMCSAddressList:          dhcpParamsMap["bcmcs_server_address_list"].(string),
			BCMCSDomainNameList:       dhcpParamsMap["bcmcs_server_domain_name_list"].(string),
			DNSRecursiveNameServer:    dhcpParamsMap["dns_recursive_name_server"].(string),
			DomainSearchList:          dhcpParamsMap["domain_search_list"].(string),
			NISServers:                dhcpParamsMap["nis_servers"].(string),
			NISPDomainName:            dhcpParamsMap["nisp_domain_name"].(string),
			PanaAuthenticationAgents:  dhcpParamsMap["pana_authentication_agents"].(string),
			POSIXTimeZone:             dhcpParamsMap["posix_time_zone"].(string),
			SIPServersDomainNameList:  dhcpParamsMap["sip_servers_domain_name_list"].(string),
			SIPServersIPv6AddressList: dhcpParamsMap["sip_servers_ipv6_address_list"].(string),
			SNTPServers:               dhcpParamsMap["sntp_servers"].(string),
			TZDBTimeZone:              dhcpParamsMap["tzdb_time_zone"].(string),
			VendorOptions:             dhcpParamsMap["vendor_options"].(string),
		}
		ipv6Range.DHCPParams = dhcpParams

		log.Printf("DHCP Params: %+v", dhcpParams)
	} else {
		log.Println("No DHCP Params provided")
	}

	if udasSet, ok := d.Get("udas").(*schema.Set); ok && udasSet.Len() > 0 {
		for _, uda := range udasSet.List() {
			if udaMap, ok := uda.(map[string]interface{}); ok {
				name := ""
				value := ""

				if n, ok := udaMap["name"].(string); ok {
					name = n
				}
				if v, ok := udaMap["value"].(string); ok {
					value = v
				}

				ipv6Range.UDAs = append(ipv6Range.UDAs, en.UDA{
					Name:  name,
					Value: value,
				})
			}
		}
	}

	if groupsSet, ok := d.Get("groups").(*schema.Set); ok && groupsSet.Len() > 0 {
		for _, group := range groupsSet.List() {
			if groupMap, ok := group.(map[string]interface{}); ok {
				name := ""
				if n, ok := groupMap["name"].(string); ok {
					name = n
				}

				var udas []en.UDA
				if udasSet, ok := groupMap["udas"].(*schema.Set); ok && udasSet.Len() > 0 {
					for _, uda := range udasSet.List() {
						if udaMap, ok := uda.(map[string]interface{}); ok {
							udaName := ""
							udaValue := ""

							if n, ok := udaMap["name"].(string); ok {
								udaName = n
							}
							if v, ok := udaMap["value"].(string); ok {
								udaValue = v
							}
							udas = append(udas, en.UDA{
								Name:  udaName,
								Value: udaValue,
							})
						}
					}
				}

				ipv6Range.Groups = append(ipv6Range.Groups, en.Group{
					Name: name,
					UDAs: udas,
				})
			}
		}
	}

	return ipv6Range
}
