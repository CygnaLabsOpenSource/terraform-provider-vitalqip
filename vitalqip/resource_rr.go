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

func resourceRR() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRRRecord,
		ReadContext:   getRRRecord,
		UpdateContext: updateRRRecord,
		DeleteContext: deleteRRRecord,

		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
				ForceNew:    true,
			},
			"rr_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "ID of resource record.",
			},
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
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
				Required:    true,
				Description: "Type of resource record.",
			},
			"data1": {
				Type:        schema.TypeString,
				Required:    true,
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
				Required:    true,
				Description: "Type of infrastructure.",
				ForceNew:    true,
			},
			"infra_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Infrastructure FQDN.",
				ForceNew:    true,
			},
			"infra_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Address of infrastructure. Note: Required if infraType=OBJECT or infraType=V6ADDRESS, and infraFQDN is not specified.",
				ForceNew:    true,
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

func createRRRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	rr := getRRFromResourceData(d)

	log.Println("[DEBUG] Create Resource Record: " + fmt.Sprintf("%v", rr))

	rrAdded, err := objMgr.CreateRR(rr)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Creation of QIP Resource Record failed",
			Detail:   fmt.Sprintf("Creation of QIP Resource Record failed: %s", err),
		})
		return diags
	}
	d.SetId(strconv.Itoa(rrAdded.RRID))
	d.Set("rr_id", rrAdded.RRID)

	return getRRRecord(ctx, d, m)
}

func getRRRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	var err error
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	query := map[string]string{
		"orgName":   orgName,
		"pageSize":  "0",
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

func updateRRRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	rrUpdate := getRRUpdateFromResourceData(d)
	log.Println("[DEBUG] Update Resource Record: " + fmt.Sprintf("%v", rrUpdate))

	_, err = objMgr.UpdateRR(rrUpdate)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Updating of QIP Resource Record failed",
			Detail:   fmt.Sprintf("Updating QIP Resource Record by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return getRRRecord(ctx, d, m)
}

func deleteRRRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	rrId := d.Get("rr_id").(int)
	query := map[string]string{
		"orgName": orgName,
		"rrId":    strconv.Itoa(rrId),
	}
	log.Println("[DEBUG] Delete Resource Record: " + fmt.Sprintf("%v", query))

	err := objMgr.DeleteRR(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Deletion of QIP Resource Record failed",
			Detail:   fmt.Sprintf("Deleting QIP Resource Record by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return diags
}

func getRRFromResourceData(d *schema.ResourceData) *en.RR {

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	owner := strings.TrimSpace(d.Get("owner").(string))
	classType := strings.TrimSpace(d.Get("class_type").(string))
	rrType := strings.TrimSpace(d.Get("rr_type").(string))
	data1 := strings.TrimSpace(d.Get("data1").(string))
	data2 := strings.TrimSpace(d.Get("data2").(string))
	data3 := strings.TrimSpace(d.Get("data3").(string))
	data4 := strings.TrimSpace(d.Get("data4").(string))
	publishing := strings.TrimSpace(d.Get("publishing").(string))
	ttl := d.Get("ttl").(int)
	infraType := strings.TrimSpace(d.Get("infra_type").(string))
	infraFQDN := strings.TrimSpace(d.Get("infra_fqdn").(string))
	infraAddr := strings.TrimSpace(d.Get("infra_addr").(string))
	isCreatingReverseZoneRR := d.Get("is_creating_reverse_zone_rr").(bool)

	rr := en.NewRR(en.RR{
		OrgName:                 orgName,
		Owner:                   owner,
		ClassType:               classType,
		RrType:                  rrType,
		Data1:                   data1,
		Data2:                   data2,
		Data3:                   data3,
		Data4:                   data4,
		Publishing:              publishing,
		TTL:                     ttl,
		InfraType:               infraType,
		InfraFQDN:               infraFQDN,
		InfraAddr:               infraAddr,
		IsCreatingReverseZoneRR: isCreatingReverseZoneRR,
	})

	optionalAttrRaw := d.Get("optional_attribute_list").([]interface{})
	rr.OptionalAttributeList = parseOptionalAttributes(optionalAttrRaw)

	return rr
}

func getRRUpdateFromResourceData(d *schema.ResourceData) *en.RRUpdate {

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	owner := strings.TrimSpace(d.Get("owner").(string))
	classType := strings.TrimSpace(d.Get("class_type").(string))
	rrType := strings.TrimSpace(d.Get("rr_type").(string))
	data1 := strings.TrimSpace(d.Get("data1").(string))
	data2 := strings.TrimSpace(d.Get("data2").(string))
	data3 := strings.TrimSpace(d.Get("data3").(string))
	data4 := strings.TrimSpace(d.Get("data4").(string))
	publishing := strings.TrimSpace(d.Get("publishing").(string))
	ttl := d.Get("ttl").(int)
	isCreatingReverseZoneRR := d.Get("is_creating_reverse_zone_rr").(bool)

	rrUpdate := en.NewRRUpdate(en.RRUpdate{
		OrgName: orgName,
		RRID:    d.Get("rr_id").(int),
		UpdateFields: en.UpdateFields{
			Owner:                   owner,
			ClassType:               classType,
			RRType:                  rrType,
			Data1:                   data1,
			Data2:                   data2,
			Data3:                   data3,
			Data4:                   data4,
			Publishing:              publishing,
			TTL:                     strconv.Itoa(ttl),
			IsCreatingReverseZoneRR: isCreatingReverseZoneRR,
		},
	})

	optionalAttrRaw := d.Get("optional_attribute_list").([]interface{})
	rrUpdate.UpdateFields.OptionalAttributeList = parseOptionalAttributes(optionalAttrRaw)

	return rrUpdate
}

func parseUDAs(udasRaw *schema.Set) []en.UDA {
	var udas []en.UDA
	for _, udaRaw := range udasRaw.List() {
		udaMap := udaRaw.(map[string]interface{})
		udas = append(udas, en.UDA{
			Name:  udaMap["name"].(string),
			Value: udaMap["value"].(string),
		})
	}
	return udas
}

func parseGroups(groupsRaw *schema.Set) []en.Group {
	var groups []en.Group
	for _, groupRaw := range groupsRaw.List() {
		groupMap := groupRaw.(map[string]interface{})
		group := en.Group{
			Name: groupMap["name"].(string),
		}

		if groupUDAsRaw, ok := groupMap["udas"].(*schema.Set); ok {
			group.UDAs = parseUDAs(groupUDAsRaw)
		}

		groups = append(groups, group)
	}
	return groups
}

func parseOptionalAttributes(optionalAttrRaw []interface{}) en.OptionalAttributeList {
	var optionalAttributeList en.OptionalAttributeList

	if len(optionalAttrRaw) > 0 {
		optionalAttrMap := optionalAttrRaw[0].(map[string]interface{})

		if udasRaw, ok := optionalAttrMap["udas"].(*schema.Set); ok {
			optionalAttributeList.UDAs = parseUDAs(udasRaw)
		}

		if groupsRaw, ok := optionalAttrMap["groups"].(*schema.Set); ok {
			optionalAttributeList.Groups = parseGroups(groupsRaw)
		}
	}
	return optionalAttributeList
}
