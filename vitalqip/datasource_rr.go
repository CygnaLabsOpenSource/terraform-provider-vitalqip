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

func dataSourceRR() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRRRead,
		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
			},
			"rr_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "ID of resource record.",
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Owner name for the resource record.",
			},
			"class_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Class type of records pertains to a type of network or software, defaults to IN.",
			},
			"rr_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Type of resource record.",
			},
			"data1": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The data associated with the specific resource record type.",
			},
			"data2": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The data associated with the specific resource record type.",
			},
			"data3": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The data associated with the specific resource record type.",
			},
			"data4": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The data associated with the specific resource record type.",
			},
			"publishing": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Publishing type.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The length of time (in seconds) the name server will hold this information. If no TTL is defined, the value is inherited from the zone.",
			},
			"infra_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Type of infrastructure.",
			},
			"infra_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Infrastructure FQDN.",
			},
			"infra_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Address of infrastructure. Note: Required if infraType=OBJECT or infraType=V6ADDRESS, and infraFQDN is not specified.",
			},
			"is_creating_reverse_zone_rr": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "If set to true, the PTR resource record will be created in Reverse zone. Otherwise, resource record will be created in Object as normally. Default value is false.",
			},
			"optional_attribute_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Include udas and groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
		},
	}
}

func dataSourceRRRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	var err error
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	query := map[string]string{
		"orgName":   orgName,
		"pageSizge": "0",
		"pageIndex": "1",
	}

	infraFQDN := strings.TrimSpace(d.Get("infra_fqdn").(string))
	infraAddr := strings.TrimSpace(d.Get("infra_addr").(string))
	if infraFQDN == "" && infraAddr == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP Resource Record Failed",
			Detail:   "Missing infra_fqdn and infra_addr field",
		})
		return diags
	}

	if infraFQDN != "" {
		query["name"] = infraFQDN
	}

	if infraAddr != "" {
		query["address"] = infraAddr
	}

	infraType := strings.TrimSpace(d.Get("infra_type").(string))
	if infraType != "" {
		query["type"] = infraType
	}

	log.Println("[DEBUG] Get Resource Record: " + fmt.Sprintf("%v", query))

	rrResponse, err := objMgr.GetRR(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP Resource Record Failed",
			Detail:   fmt.Sprintf("Getting QIP Resource Record failed : %s", err),
		})
		return diags
	}

	if rrResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API returns a nil/empty QIP Resource Record",
			Detail:   "API returns a nil/empty Address response. Getting QIP Resource Record failed",
		})
		return diags
	}

	rrID := d.Get("rr_id").(int)
	rr, found := getRecordByID(*rrResponse, rrID)
	if found && rr != nil {
		log.Println("[DEBUG] Response Get Resource Record: " + fmt.Sprintf("%v", rr))
		flattenRR(d, rr)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP Resource Record Failed",
			Detail:   "Getting QIP Resource Record Not Found",
		})
	}
	return diags
}

func flattenRR(d *schema.ResourceData, rr *en.RR) {

	d.SetId(strconv.Itoa(rr.RRID))
	d.Set("owner", rr.Owner)
	d.Set("class_type", rr.ClassType)
	d.Set("rr_type", rr.RrType)
	d.Set("data1", rr.Data1)
	d.Set("data2", rr.Data2)
	d.Set("data3", rr.Data3)
	d.Set("data4", rr.Data4)
	d.Set("publishing", rr.Publishing)
	d.Set("ttl", rr.TTL)
	d.Set("infra_type", rr.InfraType)
	d.Set("infra_fqdn", rr.InfraFQDN)
	d.Set("infra_addr", rr.InfraAddr)
	d.Set("is_creating_reverse_zone_rr", rr.IsCreatingReverseZoneRR)

	var udaList []interface{}
	for _, uda := range rr.OptionalAttributeList.UDAs {
		udaMap := map[string]interface{}{
			"name":  uda.Name,
			"value": uda.Value,
		}
		udaList = append(udaList, udaMap)
	}

	var groupList []interface{}
	for _, group := range rr.OptionalAttributeList.Groups {
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

	optionalAttributeList := []interface{}{
		map[string]interface{}{
			"udas":   udaList,
			"groups": groupList,
		},
	}

	d.Set("optional_attribute_list", optionalAttributeList)
}

func getRecordByID(response en.RRResponse, rrId int) (*en.RR, bool) {
	for _, record := range response.List {
		if record.RRID == rrId {
			return &record, true
		}
	}
	return nil, false
}
