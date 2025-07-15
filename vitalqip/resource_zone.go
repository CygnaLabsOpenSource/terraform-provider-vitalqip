package vitalqip

import (
	"context"
	"fmt"
	"log"
	"strconv"

	// "strconv"

	"strings"
	en "terraform-provider-vitalqip/vitalqip/entities"
	cc "terraform-provider-vitalqip/vitalqip/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: createZoneRecord,
		ReadContext:   getZoneRecord,
		UpdateContext: updateZoneRecord,
		DeleteContext: deleteZoneRecord,

		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
				Required:    true,
				Description: "DNS default TTL.",
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Email contact.",
			},
			"expire_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Expire time.",
			},
			"negative_cache_ttl": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "DNS negative cache TTL.",
			},
			"refresh_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "DNS refresh time.",
			},
			"retry_time": {
				Type:        schema.TypeInt,
				Required:    true,
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

func createZoneRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	zone := getZoneFromResourceData(d)

	log.Println("[DEBUG] Create Zone: " + fmt.Sprintf("%v", zone))

	zoneAdded, err := objMgr.CreateZone(zone)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Creation of QIP Zone",
			Detail:   fmt.Sprintf("Creation of QIP Zone failed: %s", err),
		})
		return diags
	}
	d.SetId(strconv.Itoa(zoneAdded.ID))

	return getZoneRecord(ctx, d, m)
}

func updateZoneRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	zoneUpdate := getZoneUpdateFromResourceData(d)
	log.Println("[DEBUG] Update Zone: " + fmt.Sprintf("%v", zoneUpdate))

	err = objMgr.UpdateZone(zoneUpdate)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Updating of QIP Zone failed",
			Detail:   fmt.Sprintf("Updating QIP Zone by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}

	return getZoneRecord(ctx, d, m)
}

func deleteZoneRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	zoneName := strings.TrimSpace(d.Get("name").(string))
	query := map[string]string{
		"orgName":  orgName,
		"zoneName": zoneName,
	}
	log.Println("[DEBUG] Delete Zone: " + fmt.Sprintf("%v", query))

	err := objMgr.DeleteZone(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Deletion of QIP Zone failed",
			Detail:   fmt.Sprintf("Deleting QIP Zone by name (%s) failed : %s", zoneName, err),
		})
		return diags
	}
	return diags
}

func getZoneFromResourceData(d *schema.ResourceData) *en.Zone {

	zone := &en.Zone{}
	zone.OrgName = strings.TrimSpace(d.Get("org_name").(string))
	zone.Name = strings.TrimSpace(d.Get("name").(string))
	zone.ConfigPrivateZone = d.Get("config_private_zone").(bool)
	zone.Email = strings.TrimSpace(d.Get("email").(string))
	zone.DefaultTTL = d.Get("default_ttl").(int)
	zone.ExpireTime = d.Get("expire_time").(int)
	zone.NegativeCacheTTL = d.Get("negative_cache_ttl").(int)
	zone.RefreshTime = d.Get("refresh_time").(int)
	zone.RetryTime = d.Get("retry_time").(int)
	zone.ParentAddress = strings.TrimSpace(d.Get("parent_address").(string))
	zone.NetworkAddress = strings.TrimSpace(d.Get("network_address").(string))
	zone.PostfixZoneExtension = strings.TrimSpace(d.Get("postfix_zone_extension").(string))
	zone.PrefixZoneExtension = strings.TrimSpace(d.Get("prefix_zone_extension").(string))
	zone.DNSServers = parseDnsServers(d.Get("dns_servers").(*schema.Set))
	zone.UDAs = parseUDAs(d.Get("udas").(*schema.Set))
	zone.Groups = parseGroups(d.Get("groups").(*schema.Set))
	zone.DNSZoneOptions = parseZoneOptions(d.Get("dns_zone_options").(*schema.Set))

	return zone
}

func getZoneUpdateFromResourceData(d *schema.ResourceData) *en.ZoneUpdate {

	zone := getZoneFromResourceData(d)
	oldName, newName := d.GetChange("name")

	zoneUpdate := &en.ZoneUpdate{
		Zone:    *zone,
		NewName: strings.TrimSpace(newName.(string)),
	}
	zoneUpdate.Name = strings.TrimSpace(oldName.(string))
	return zoneUpdate
}

func parseDnsServers(data *schema.Set) []en.DNSServer {
	var dnsServers []en.DNSServer

	for _, item := range data.List() {
		serverData := item.(map[string]interface{})

		server := en.DNSServer{
			Name:         strings.TrimSpace(serverData["name"].(string)),
			Role:         strings.TrimSpace(serverData["role"].(string)),
			SecureUpdate: serverData["secure_update"].(bool),
		}
		dnsServers = append(dnsServers, server)
	}
	return dnsServers
}

func parseZoneOptions(data *schema.Set) []en.ZoneOptionSet {
	var zoneOptions []en.ZoneOptionSet

	for _, item := range data.List() {
		optionSetData := item.(map[string]interface{})

		optionSet := en.ZoneOptionSet{
			Name: strings.TrimSpace(optionSetData["name"].(string)),
		}

		if options, ok := optionSetData["options"].(*schema.Set); ok {
			for _, optItem := range options.List() {
				optionData := optItem.(map[string]interface{})

				option := en.ZoneOption{
					Name:  strings.TrimSpace(optionData["name"].(string)),
					Value: strings.TrimSpace(optionData["value"].(string)),
				}

				if subOptions, ok := optionData["sub_options"].(*schema.Set); ok {
					for _, subOptItem := range subOptions.List() {
						subOptionData := subOptItem.(map[string]interface{})

						subOption := en.SubOption{
							Name:  strings.TrimSpace(subOptionData["name"].(string)),
							Value: strings.TrimSpace(subOptionData["value"].(string)),
						}
						option.SubOptions = append(option.SubOptions, subOption)
					}
				}
				optionSet.Options = append(optionSet.Options, option)
			}
		}
		zoneOptions = append(zoneOptions, optionSet)
	}
	log.Println("[DEBUG] zone option ne: " + fmt.Sprintf("%v", zoneOptions))
	return zoneOptions
}
