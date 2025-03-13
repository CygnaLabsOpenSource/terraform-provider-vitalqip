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

func resourceIPv4Address() *schema.Resource {
	return &schema.Resource{
		CreateContext: createIPv4AddressRecord,
		ReadContext:   getIPv4AddressRecord,
		UpdateContext: updateIPv4AddressRecord,
		DeleteContext: deleteIPv4AddressRecord,

		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
				ForceNew:    true,
			},
			"object_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IPv4 address.",
				ForceNew:    true,
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
				ForceNew:    true,
			},
			"object_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Class type of IPv4 object.",
			},
			"expirated_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The date when a reserved object expires and is no longer reserved. The date format for this field is yyyy-mm-dd.",
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
				Description: "Dynamic Configuration of IPv4 object",
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
				Description: "A resource record option: Always, None, Push Only",
			},
			"publish_ptr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "PTR resource record option: Always, None, Push Only",
			},
		},
	}
}

func createIPv4AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	ipv4Address := getIPv4AddressFromResourceData(d)
	log.Println("[DEBUG] Create IPv4 Address: " + fmt.Sprintf("%v", ipv4Address))

	_, err = objMgr.CreateIPv4Address(ipv4Address)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Creation of QIP IPv4 Address failed",
			Detail:   fmt.Sprintf("Creation of Address (%s) failed: %s", ipv4Address.ObjectAddr, err),
		})
		return diags
	}
	d.SetId(ipv4Address.ObjectAddr)
	return getIPv4AddressRecord(ctx, d, m)
}

func getIPv4AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	objectAddr := strings.TrimSpace(d.Get("object_addr").(string))
	query := map[string]string{
		"orgName":        orgName,
		"objectAddr":     objectAddr,
		"addressVersion": "4",
	}
	log.Println("[DEBUG] Get IPv4 Address: " + fmt.Sprintf("%v", query))

	response, err := objMgr.GetIPv4Address(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv4 Address Failed",
			Detail:   fmt.Sprintf("Getting QIP IPv4 Address (%s) failed : %s", objectAddr, err),
		})
		return diags
	}
	flattenIPv4Address(d, response)
	return diags
}

func updateIPv4AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	ipv4Address := getIPv4AddressFromResourceData(d)
	log.Println("[DEBUG] Update IPv4 Address: " + fmt.Sprintf("%v", ipv4Address))

	_, err = objMgr.UpdateIPv4Address(ipv4Address)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Updating of QIP IPv4 Address failed",
			Detail:   fmt.Sprintf("Updating QIP IPv4 Address by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return getIPv4AddressRecord(ctx, d, m)
}

func deleteIPv4AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	objectAddr := strings.TrimSpace(d.Get("object_addr").(string))
	query := map[string]string{
		"orgName":        orgName,
		"objectAddr":     objectAddr,
		"addressVersion": "4",
	}
	log.Println("[DEBUG] Delete IPv4 Address: " + fmt.Sprintf("%v", query))

	err := objMgr.DeleteIPv4Address(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Deletion of QIP IPv4 Address failed",
			Detail:   fmt.Sprintf("Deleting QIP IPv4 Address block by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return diags
}

func getIPv4AddressFromResourceData(d *schema.ResourceData) *en.IPv4Address {

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	objectAddr := strings.TrimSpace(d.Get("object_addr").(string))
	objectName := strings.TrimSpace(d.Get("object_name").(string))
	subnetAddr := strings.TrimSpace(d.Get("subnet_addr").(string))
	objectClass := strings.TrimSpace(d.Get("object_class").(string))
	expiratedDate := strings.TrimSpace(d.Get("expirated_date").(string))
	domainName := strings.TrimSpace(d.Get("domain_name").(string))
	objectDesc := strings.TrimSpace(d.Get("object_desc").(string))
	dynamicConfig := strings.TrimSpace(d.Get("dynamic_config").(string))
	macAddr := strings.TrimSpace(d.Get("mac_addr").(string))
	aTTL := strings.TrimSpace(d.Get("a_ttl").(string))
	ptrTTL := strings.TrimSpace(d.Get("ptr_ttl").(string))
	publishA := strings.TrimSpace(d.Get("publish_a").(string))
	publishPTR := strings.TrimSpace(d.Get("publish_ptr").(string))

	return en.NewIPv4Address(en.IPv4Address{
		OrgName:        orgName,
		ObjectAddr:     objectAddr,
		ObjectName:     objectName,
		SubnetAddr:     subnetAddr,
		ObjectClass:    objectClass,
		ExpiratedDate:  expiratedDate,
		DomainName:     domainName,
		ObjectDesc:     objectDesc,
		DynamicConfig:  dynamicConfig,
		MacAddr:        macAddr,
		ATTL:           aTTL,
		PTRTTL:         ptrTTL,
		PublishA:       publishA,
		PublishPTR:     publishPTR,
		AddressVersion: 4,
	})
}
