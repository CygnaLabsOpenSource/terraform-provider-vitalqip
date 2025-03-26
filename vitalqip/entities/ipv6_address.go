package entities

import (
	"fmt"
	"strings"
)

type IPv6Address struct {
	ObjBase        `json:"-"`
	OrgName        string `json:"orgName,omitempty"`
	HostName       string `json:"hostName,omitempty"`
	DomainName     string `json:"domainName,omitempty"`
	RangeAddress   string `json:"rangeAddress,omitempty"`
	Address        string `json:"address,omitempty"`
	AddressType    string `json:"addressType,omitempty"`
	PublishA       string `json:"publishA,omitempty"`
	PublishPTR     string `json:"publishPTR,omitempty"`
	ClassType      string `json:"classType,omitempty"`
	IAID           string `json:"iaid,omitempty"`
	FQDN           string `json:"fqdn,omitempty"`
	Range          string `json:"range,omitempty"`
	TTL            int    `json:"ttl,omitempty"`
	NodeName       string `json:"nodeName,omitempty"`
	UniqueID       string `json:"uniqueId,omitempty"`
	DUID           string `json:"duid,omitempty"`
	UseMACAddress  bool   `json:"useMACAddress"`
	MACAddress     string `json:"macAddress,omitempty"`
	ObjectAddr     string `json:"objectAddr,omitempty"`
	AddressVersion int    `json:"addressVersion,omitempty"`
}

func (o IPv6Address) String() string {
	var sb strings.Builder

	sb.WriteString("IPv6Address {\n")
	sb.WriteString(fmt.Sprintf("  HostName: %s\n", o.HostName))
	sb.WriteString(fmt.Sprintf("  DomainName: %s\n", o.DomainName))
	sb.WriteString(fmt.Sprintf("  RangeAddress: %s\n", o.RangeAddress))
	sb.WriteString(fmt.Sprintf("  Address: %s\n", o.Address))
	sb.WriteString(fmt.Sprintf("  AddressType: %s\n", o.AddressType))
	sb.WriteString(fmt.Sprintf("  PublishA: %s\n", o.PublishA))
	sb.WriteString(fmt.Sprintf("  PublishPTR: %s\n", o.PublishPTR))
	sb.WriteString(fmt.Sprintf("  ClassType: %s\n", o.ClassType))
	sb.WriteString(fmt.Sprintf("  IAID: %s\n", o.IAID))
	sb.WriteString(fmt.Sprintf("  Range: %s\n", o.Range))
	sb.WriteString(fmt.Sprintf("  TTL: %d\n", o.TTL))
	sb.WriteString(fmt.Sprintf("  NodeName: %s\n", o.NodeName))
	sb.WriteString(fmt.Sprintf("  UniqueID: %s\n", o.UniqueID))
	sb.WriteString(fmt.Sprintf("  DUID: %s\n", o.DUID))
	sb.WriteString(fmt.Sprintf("  UseMACAddress: %t\n", o.UseMACAddress))
	sb.WriteString(fmt.Sprintf("  MACAddress: %s\n", o.MACAddress))
	sb.WriteString("}\n")

	return sb.String()
}

func NewIPv6Address(sb IPv6Address) *IPv6Address {
	res := sb
	res.objectType = "ipv6_address"
	return &res
}
