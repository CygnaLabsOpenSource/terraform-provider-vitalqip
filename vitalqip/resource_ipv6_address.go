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

func resourceIPv6Address() *schema.Resource {
	return &schema.Resource{
		CreateContext: createIPv6AddressRecord,
		ReadContext:   getIPv6AddressRecord,
		UpdateContext: updateIPv6AddressRecord,
		DeleteContext: deleteIPv6AddressRecord,

		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization Name.",
				ForceNew:    true,
			},
			"host_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Hostname of address.",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IPv6 address.",
				ForceNew:    true,
			},
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name of address.",
			},
			"range_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Address range.",
				ForceNew:    true,
			},
			"address_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of address.",
			},
			"publish_a": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AAAA resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL",
			},
			"publish_ptr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PTR resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL",
			},
			"class_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Class of object.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name.",
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
				ForceNew:    true,
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
				Description: "True: use MAC address, False: use DUID value instead",
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

func createIPv6AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	ipv6Address := getIPv6AddressFromResourceData(d)
	log.Println("[DEBUG] Create IPv6 Address: " + fmt.Sprintf("%v", ipv6Address))

	_, err = objMgr.CreateIPv6Address(ipv6Address)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Creation of QIP IPv6 Address failed",
			Detail:   fmt.Sprintf("Creation of Address (%s) failed: %s", ipv6Address.Address, err),
		})
		return diags
	}

	d.SetId(ipv6Address.Address)
	return getIPv6AddressRecord(ctx, d, m)
}

func getIPv6AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	address := strings.TrimSpace(d.Get("address").(string))
	query := map[string]string{
		"orgName":        orgName,
		"objectAddr":     address,
		"addressVersion": "6",
	}
	log.Println("[DEBUG] Get IPv6 Address: " + fmt.Sprintf("%v", query))

	response, err := objMgr.GetIPv6Address(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv6 Address Failed",
			Detail:   fmt.Sprintf("Getting QIP IPv6 Address (%s) failed : %s", address, err),
		})
		return diags
	}
	flattenIPv6Address(d, response)
	return diags
}

func updateIPv6AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	ipv6Address := getIPv6AddressFromResourceData(d)
	ipv6Address.ObjectAddr = ipv6Address.Address
	log.Println("[DEBUG] Update IPv6 Address: " + fmt.Sprintf("%v", ipv6Address))

	_, err = objMgr.UpdateIPv6Address(ipv6Address)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Updating of QIP IPv6 Address failed",
			Detail:   fmt.Sprintf("Updating QIP IPv6 Address by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return getIPv6AddressRecord(ctx, d, m)
}

func deleteIPv6AddressRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics
	orgName := strings.TrimSpace(d.Get("org_name").(string))
	address := strings.TrimSpace(d.Get("address").(string))
	query := map[string]string{
		"orgName":        orgName,
		"objectAddr":     address,
		"addressVersion": "6",
	}
	log.Println("[DEBUG] Delete IPv6 Address: " + fmt.Sprintf("%v", query))

	err := objMgr.DeleteIPv6Address(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Deletion of QIP IPv6 Address failed",
			Detail:   fmt.Sprintf("Deleting QIP IPv6 Address block by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}
	return diags
}

func getIPv6AddressFromResourceData(d *schema.ResourceData) *en.IPv6Address {

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	hostName := strings.TrimSpace(d.Get("host_name").(string))
	address := strings.TrimSpace(d.Get("address").(string))
	domainName := strings.TrimSpace(d.Get("domain_name").(string))
	rangeAddress := strings.TrimSpace(d.Get("range_address").(string))
	fqdn := strings.TrimSpace(d.Get("fqdn").(string))
	addressType := strings.TrimSpace(d.Get("address_type").(string))
	publishA := strings.TrimSpace(d.Get("publish_a").(string))
	publishPTR := strings.TrimSpace(d.Get("publish_ptr").(string))
	classType := strings.TrimSpace(d.Get("class_type").(string))
	iaid := strings.TrimSpace(d.Get("iaid").(string))
	rangeName := strings.TrimSpace(d.Get("range_name").(string))
	ttl := d.Get("ttl").(int)
	nodeName := strings.TrimSpace(d.Get("node_name").(string))
	uniqueId := strings.TrimSpace(d.Get("unique_id").(string))
	duid := strings.TrimSpace(d.Get("duid").(string))
	useMACAddress := d.Get("use_mac_address").(bool)
	macAddress := strings.TrimSpace(d.Get("mac_address").(string))

	return en.NewIPv6Address(en.IPv6Address{
		OrgName:        orgName,
		HostName:       hostName,
		Address:        address,
		DomainName:     domainName,
		RangeAddress:   rangeAddress,
		FQDN:           fqdn,
		AddressType:    addressType,
		PublishA:       publishA,
		PublishPTR:     publishPTR,
		ClassType:      classType,
		IAID:           iaid,
		Range:          rangeName,
		TTL:            ttl,
		NodeName:       nodeName,
		UniqueID:       uniqueId,
		DUID:           duid,
		UseMACAddress:  useMACAddress,
		MACAddress:     macAddress,
		AddressVersion: 6,
	})
}
