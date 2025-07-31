package vitalqip

import (

	// "regexp"

	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	en "terraform-provider-vitalqip/vitalqip/entities"
	cc "terraform-provider-vitalqip/vitalqip/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceZone() *schema.Resource {
	return &schema.Resource{
		ReadContext: getZoneRecord,
		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
			},
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "ID of zone.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the zone.",
			},
			"default_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "DNS default TTL.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Email contact.",
			},
			"expire_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Expire time.",
			},
			"negative_cache_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "DNS negative cache TTL.",
			},
			"refresh_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "DNS refresh time.",
			},
			"retry_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "DNS retry time.",
			},
			"parent_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Parent address of the split reverse zone.",
			},
			"network_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Network address of the reverse zone.",
			},
			"postfix_zone_extension": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "DNS postfix extension.",
			},
			"prefix_zone_extension": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "DNS prefix extension.",
			},
			"config_private_zone": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Config private zone.",
			},
			"dns_servers": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "List of DNS servers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of DNS Server.",
						},
						"role": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "DNS Server Role. Values: P for Primary or S for Secondary.",
						},
						"secure_update": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Send Secure Updates. Values: true or false.",
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
				Description: "List of UDA groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Name of the UDA group.",
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
			"dns_zone_options": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "DNS zone options.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Group name of zone option.",
						},
						"options": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "List of zone options.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Name of zone option.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Value of zone option.",
									},
									"sub_options": {
										Type:        schema.TypeSet,
										Optional:    true,
										Computed:    true,
										Description: "List of sub options.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Name of sub option.",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Value of sub option.",
												},
											},
										},
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

func getZoneRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	var err error
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	name := strings.TrimSpace(d.Get("name").(string))
	query := map[string]string{
		"orgName":  orgName,
		"zoneName": name,
	}

	log.Println("[DEBUG] Get Zone: " + fmt.Sprintf("%v", query))

	zoneResponse, err := objMgr.GetZone(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP Zone Failed",
			Detail:   fmt.Sprintf("Getting QIP Zone failed : %s", err),
		})
		return diags
	}

	if zoneResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API returns a nil/empty QIP Zone",
			Detail:   "API returns a nil/empty Address response. Getting QIP Zone failed",
		})
		return diags
	}

	flattenZone(d, zoneResponse)
	log.Println("[DEBUG] QIP Zone: " + fmt.Sprintf("%v", zoneResponse))

	return diags
}

func flattenZone(d *schema.ResourceData, zone *en.Zone) {

	d.SetId(strconv.Itoa(zone.ID))
	d.Set("config_private_zone", zone.ConfigPrivateZone)
	d.Set("default_ttl", zone.DefaultTTL)
	d.Set("email", zone.Email)
	d.Set("expire_time", zone.ExpireTime)
	d.Set("negative_cache_ttl", zone.NegativeCacheTTL)
	d.Set("refresh_time", zone.RefreshTime)
	d.Set("retry_time", zone.RetryTime)
	d.Set("parent_address", zone.ParentAddress)
	d.Set("network_address", zone.NetworkAddress)
	d.Set("postfix_zone_extension", zone.PostfixZoneExtension)
	d.Set("prefix_zone_extension", zone.PrefixZoneExtension)
	d.Set("dns_servers", flattenDnsServers(zone.DNSServers))
	d.Set("udas", flattenUdas(zone.UDAs))
	d.Set("groups", flattenGroups(zone.Groups))
	d.Set("dns_zone_options", flattenZoneOptions(zone.DNSZoneOptions))
}

func flattenDnsServers(dnsServers []en.DNSServer) []interface{} {
	var dnsServerList []interface{}
	for _, dnsServer := range dnsServers {
		dnsServerMap := map[string]interface{}{
			"name":          dnsServer.Name,
			"role":          dnsServer.Role,
			"secure_update": dnsServer.SecureUpdate,
		}
		dnsServerList = append(dnsServerList, dnsServerMap)
	}
	return dnsServerList
}

func flattenUdas(udas []en.UDA) []interface{} {
	var udaList []interface{}
	for _, uda := range udas {
		udaMap := map[string]interface{}{
			"name":  uda.Name,
			"value": uda.Value,
		}
		udaList = append(udaList, udaMap)
	}
	return udaList
}

func flattenGroups(groups []en.Group) []interface{} {
	var groupList []interface{}
	for _, group := range groups {
		groupMap := map[string]interface{}{
			"name": group.Name,
			"udas": flattenUdas(group.UDAs),
		}
		groupList = append(groupList, groupMap)
	}
	return groupList
}

func flattenZoneOptions(zoneOptions []en.ZoneOptionSet) []interface{} {
	var result []interface{}
	for _, zos := range zoneOptions {
		item := map[string]interface{}{
			"name":    zos.Name,
			"options": flattenOptions(zos.Options),
		}
		result = append(result, item)
	}
	return result
}

func flattenOptions(options []en.ZoneOption) []interface{} {
	var result []interface{}
	for _, option := range options {
		item := map[string]interface{}{
			"name":        option.Name,
			"value":       option.Value,
			"sub_options": flattenSubOptions(option.SubOptions),
		}
		result = append(result, item)
	}
	return result
}

func flattenSubOptions(subOptions []en.SubOption) []interface{} {
	var result []interface{}
	for _, subOption := range subOptions {
		item := map[string]interface{}{
			"name":  subOption.Name,
			"value": subOption.Value,
		}
		result = append(result, item)
	}
	return result
}
