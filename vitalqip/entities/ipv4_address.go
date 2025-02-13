package entities

import (
	"fmt"
	"strings"
)

type IPv4Address struct {
	ObjBase        `json:"-"`
	OrgName        string `json:"orgName,omitempty"`
	ObjectAddr     string `json:"objectAddr,omitempty"`
	SubnetAddr     string `json:"subnetAddr,omitempty"`
	ObjectName     string `json:"objectName,omitempty"`
	ObjectClass    string `json:"objectClass,omitempty"`
	ExpiratedDate  string `json:"expiratedDate,omitempty"`
	DomainName     string `json:"domainName,omitempty"`
	ObjectDesc     string `json:"objectDesc,omitempty"`
	DynamicConfig  string `json:"dynamicConfig,omitempty"`
	MacAddr        string `json:"macAddr,omitempty"`
	ATTL           string `json:"aTTL,omitempty"`
	PTRTTL         string `json:"ptrTTL,omitempty"`
	PublishA       string `json:"publishA,omitempty"`
	PublishPTR     string `json:"publishPTR,omitempty"`
	AddressVersion int    `json:"addressVersion"`
}

func (o IPv4Address) String() string {
	var sb strings.Builder

	sb.WriteString("IPv4Address {\n")
	sb.WriteString(fmt.Sprintf("  ObjectAddr: %s\n", o.ObjectAddr))
	sb.WriteString(fmt.Sprintf("  SubnetAddr: %s\n", o.SubnetAddr))
	sb.WriteString(fmt.Sprintf("  ObjectName: %s\n", o.ObjectName))
	sb.WriteString(fmt.Sprintf("  ObjectClass: %s\n", o.ObjectClass))
	sb.WriteString(fmt.Sprintf("  ExpiratedDate: %s\n", o.ExpiratedDate))
	sb.WriteString(fmt.Sprintf("  DomainName: %s\n", o.DomainName))
	sb.WriteString(fmt.Sprintf("  ObjectDesc: %s\n", o.ObjectDesc))
	sb.WriteString(fmt.Sprintf("  DynamicConfig: %s\n", o.DynamicConfig))
	sb.WriteString(fmt.Sprintf("  MacAddr: %s\n", o.MacAddr))
	sb.WriteString(fmt.Sprintf("  ATTL: %s\n", o.ATTL))
	sb.WriteString(fmt.Sprintf("  PTRTTL: %s\n", o.PTRTTL))
	sb.WriteString(fmt.Sprintf("  PublishA: %s\n", o.PublishA))
	sb.WriteString(fmt.Sprintf("  PublishPTR: %s\n", o.PublishPTR))
	sb.WriteString("}\n")

	return sb.String()
}

func NewIPv4Address(sb IPv4Address) *IPv4Address {
	res := sb
	res.objectType = "ipv4_address"
	return &res
}
