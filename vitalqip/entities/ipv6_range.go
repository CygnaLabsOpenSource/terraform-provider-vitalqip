package entities

import (
	"fmt"
	"strings"
)

type DHCPParams struct {
	InformationRefreshTime    string `json:"informationRefreshTime,omitempty"`
	PreferredLifeTime         string `json:"preferredLifeTime,omitempty"`
	RebindingTime             string `json:"rebindingTime,omitempty"`
	RenewalTime               string `json:"renewalTime,omitempty"`
	ValidLifeTime             string `json:"validLifeTime,omitempty"`
	BCMCSAddressList          string `json:"bcmcsServerAddressList,omitempty"`
	BCMCSDomainNameList       string `json:"bcmcsServerDomainNameList,omitempty"`
	DNSRecursiveNameServer    string `json:"dnsRecursiveNameServer,omitempty"`
	DomainSearchList          string `json:"domainSearchList,omitempty"`
	NISServers                string `json:"nisServers,omitempty"`
	NISPDomainName            string `json:"nispDomainName,omitempty"`
	PanaAuthenticationAgents  string `json:"panaAuthenticationAgents,omitempty"`
	POSIXTimeZone             string `json:"posixTimeZone,omitempty"`
	SIPServersDomainNameList  string `json:"sipServersDomainNameList,omitempty"`
	SIPServersIPv6AddressList string `json:"sipServersIpv6AddressList,omitempty"`
	SNTPServers               string `json:"sntpServers,omitempty"`
	TZDBTimeZone              string `json:"tzdbTimeZone,omitempty"`
	VendorOptions             string `json:"vendorOptions,omitempty"`
}

type UDA struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Group struct {
	Name string `json:"name,omitempty"`
	UDAs []UDA  `json:"udas,omitempty"`
}

type IPv6Range struct {
	ObjBase                  `json:"-"`
	OrgName                  string     `json:"orgName,omitempty"`
	Name                     string     `json:"name,omitempty"`
	StartAddress             string     `json:"startAddress,omitempty"`
	NewStartAddress          string     `json:"newStartAddress,omitempty"`
	RangePrefixLength        int        `json:"rangePrefixLength,omitempty"`
	RangeType                string     `json:"rangeType,omitempty"`
	StandPrimDHCPServer      string     `json:"standPrimDHCPServer,omitempty"`
	FailoverSecondDHCPServer string     `json:"failoverSecondDHCPServer,omitempty"`
	OptTemplate              string     `json:"optTemplate,omitempty"`
	AddressSelection         string     `json:"addressSelection,omitempty"`
	SubnetName               string     `json:"subnetName,omitempty"`
	SubnetAddress            string     `json:"subnetAddress,omitempty"`
	SubnetPrefixLength       int        `json:"subnetPrefixLength,omitempty"`
	DHCPParams               DHCPParams `json:"dhcpParams,omitempty"`
	UDAs                     []UDA      `json:"udas,omitempty"`
	Groups                   []Group    `json:"groups,omitempty"`
}

func (ipv6Range IPv6Range) String() string {
	var sb strings.Builder

	sb.WriteString("IPv6Range {\n")
	sb.WriteString(fmt.Sprintf("  Name: %s\n", ipv6Range.Name))
	sb.WriteString(fmt.Sprintf("  OrgName: %s\n", ipv6Range.OrgName))
	sb.WriteString(fmt.Sprintf("  StartAddress: %s\n", ipv6Range.StartAddress))
	sb.WriteString(fmt.Sprintf("  NewStartAddress: %s\n", ipv6Range.NewStartAddress))
	sb.WriteString(fmt.Sprintf("  RangePrefixLength: %d\n", ipv6Range.RangePrefixLength))
	sb.WriteString(fmt.Sprintf("  RangeType: %s\n", ipv6Range.RangeType))
	sb.WriteString(fmt.Sprintf("  StandPrimDHCPServer: %s\n", ipv6Range.StandPrimDHCPServer))
	sb.WriteString(fmt.Sprintf("  FailoverSecondDHCPServer: %s\n", ipv6Range.FailoverSecondDHCPServer))
	sb.WriteString(fmt.Sprintf("  OptTemplate: %s\n", ipv6Range.OptTemplate))
	sb.WriteString(fmt.Sprintf("  AddressSelection: %s\n", ipv6Range.AddressSelection))
	sb.WriteString(fmt.Sprintf("  SubnetName: %s\n", ipv6Range.SubnetName))
	sb.WriteString(fmt.Sprintf("  SubnetAddress: %s\n", ipv6Range.SubnetAddress))
	sb.WriteString(fmt.Sprintf("  SubnetPrefixLength: %d\n", ipv6Range.SubnetPrefixLength))

	sb.WriteString("  DHCPParams: {\n")
	sb.WriteString(fmt.Sprintf("    %+v\n", ipv6Range.DHCPParams))
	sb.WriteString("  }\n")

	sb.WriteString("  UDAs: [\n")
	for _, uda := range ipv6Range.UDAs {
		sb.WriteString(fmt.Sprintf("    %+v\n", uda))
	}
	sb.WriteString("  ]\n")

	sb.WriteString("  Groups: [\n")
	for _, group := range ipv6Range.Groups {
		sb.WriteString(fmt.Sprintf("    Group Name: %s\n", group.Name))
		sb.WriteString("    UDAs: [\n")
		for _, uda := range group.UDAs {
			sb.WriteString(fmt.Sprintf("      %+v\n", uda))
		}
		sb.WriteString("    ]\n")
	}
	sb.WriteString("  ]\n")

	sb.WriteString("}\n")

	return sb.String()
}

func NewIPv6Range(sb IPv6Range) *IPv6Range {
	res := sb
	res.objectType = "ipv6_range"
	return &res
}

type IPv6RangeResponse struct {
	List []IPv6Range `json:"list"`
}
