package entities

import (
	"fmt"
	"strings"
)

type IPv6Subnet struct {
	ObjBase            `json:"-"`
	OrgName            string `json:"orgName,omitempty"`
	AddressVersion     int    `json:"addressVersion,omitempty"`
	SubnetAddress      string `json:"subnetAddress,omitempty"`
	PoolName           string `json:"poolName,omitempty"`
	BlockName          string `json:"blockName,omitempty"`
	BlockAddress       string `json:"blockAddress,omitempty"`
	PrefixLength       int    `json:"prefixLength,omitempty"`
	SubnetPrefixLength int    `json:"subnetPrefixLength,omitempty"`
	SubnetName         string `json:"subnetName,omitempty"`
	CreateSubnet       string `json:"createSubnet,omitempty"`
	AlgorithmType      string `json:"algorithmType,omitempty"`
	CreateReverseZone  bool   `json:"createReverseZone"`
}

func (s IPv6Subnet) String() string {
	var sb strings.Builder
	sb.WriteString("IPv6Subnet {\n")
	sb.WriteString(fmt.Sprintf("  OrgName: %s\n", s.OrgName))
	sb.WriteString(fmt.Sprintf("  AddressVersion: %d\n", s.AddressVersion))
	sb.WriteString(fmt.Sprintf("  SubnetAddress: %s\n", s.SubnetAddress))
	sb.WriteString(fmt.Sprintf("  PoolName: %s\n", s.PoolName))
	sb.WriteString(fmt.Sprintf("  BlockName: %s\n", s.BlockName))
	sb.WriteString(fmt.Sprintf("  BlockAddress: %s\n", s.BlockAddress))
	sb.WriteString(fmt.Sprintf("  PrefixLength: %d\n", s.PrefixLength))
	sb.WriteString(fmt.Sprintf("  SubnetPrefixLength: %d\n", s.SubnetPrefixLength))
	sb.WriteString(fmt.Sprintf("  SubnetName: %s\n", s.SubnetName))
	sb.WriteString(fmt.Sprintf("  CreateSubnet: %s\n", s.CreateSubnet))
	sb.WriteString(fmt.Sprintf("  AlgorithmType: %s\n", s.AlgorithmType))
	sb.WriteString(fmt.Sprintf("  CreateReverseZone: %t\n", s.CreateReverseZone))
	sb.WriteString("}\n")
	return sb.String()
}

func NewIPv6Subnet(sb IPv6Subnet) *IPv6Subnet {
	res := sb
	res.objectType = "ipv6_subnet"
	return &res
}

type IPv6SubnetModify struct {
	ObjBase        `json:"-"`
	OrgName        string `json:"orgName,omitempty"`
	AddressVersion int    `json:"addressVersion,omitempty"`
	SubnetAddress  string `json:"subnetAddress,omitempty"`
	PrefixLength   int    `json:"prefixLength,omitempty"`
	SubnetName     string `json:"name,omitempty"`
}

func NewIPv6SubnetModify(sb IPv6SubnetModify) *IPv6SubnetModify {
	res := sb
	res.objectType = "ipv6_subnet_modify"
	return &res
}

type IPv6SubnetGet struct {
	ObjBase            `json:"-"`
	SubnetAddress      string `json:"address,omitempty"`
	PoolName           string `json:"poolName,omitempty"`
	BlockName          string `json:"blockName,omitempty"`
	SubnetPrefixLength int    `json:"prefixLength,omitempty"`
	SubnetName         string `json:"name,omitempty"`
}

func (s IPv6SubnetGet) String() string {
	var sb strings.Builder
	sb.WriteString("IPv6SubnetGet {\n")
	sb.WriteString(fmt.Sprintf("  SubnetAddress: %s\n", s.SubnetAddress))
	sb.WriteString(fmt.Sprintf("  PoolName: %s\n", s.PoolName))
	sb.WriteString(fmt.Sprintf("  BlockName: %s\n", s.BlockName))
	sb.WriteString(fmt.Sprintf("  SubnetPrefixLength: %d\n", s.SubnetPrefixLength))
	sb.WriteString(fmt.Sprintf("  SubnetName: %s\n", s.SubnetName))
	sb.WriteString("}\n")
	return sb.String()
}

func NewIPv6SubnetGet(sb IPv6SubnetGet) *IPv6SubnetGet {
	res := sb
	res.objectType = "ipv6_subnet_get"
	return &res
}
