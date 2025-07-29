package entities

import (
	"fmt"
	"strings"
)

type Zone struct {
	ObjBase              `json:"-"`
	ID                   int             `json:"id,omitempty"`
	ConfigPrivateZone    bool            `json:"configPrivateZone,omitempty"`
	Name                 string          `json:"name,omitempty"`
	OrgName              string          `json:"orgName,omitempty"`
	Email                string          `json:"email,omitempty"`
	DefaultTTL           int             `json:"defaultTtl,omitempty"`
	ExpireTime           int             `json:"expireTime,omitempty"`
	NegativeCacheTTL     int             `json:"negativeCacheTtl,omitempty"`
	RefreshTime          int             `json:"refreshTime,omitempty"`
	RetryTime            int             `json:"retryTime,omitempty"`
	ParentAddress        string          `json:"parentAddress,omitempty"`
	NetworkAddress       string          `json:"networkAddress,omitempty"`
	PostfixZoneExtension string          `json:"postfixZoneExtension,omitempty"`
	PrefixZoneExtension  string          `json:"prefixZoneExtension,omitempty"`
	DNSServers           []DNSServer     `json:"dnsServers,omitempty"`
	DNSZoneOptions       []ZoneOptionSet `json:"dnsZoneOptions,omitempty"`
	UDAs                 []UDA           `json:"udas,omitempty"`
	Groups               []Group         `json:"groups,omitempty"`
}

type DNSServer struct {
	Name         string `json:"name"`
	Role         string `json:"role"`
	SecureUpdate bool   `json:"secureUpdate"`
}

type ZoneOptionSet struct {
	Name    string       `json:"name"`
	Options []ZoneOption `json:"options"`
}

type ZoneOption struct {
	Name       string      `json:"name"`
	Value      string      `json:"value,omitempty"`
	SubOptions []SubOption `json:"subOptions,omitempty"`
}

type SubOption struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

func (zone Zone) String() string {
	var sb strings.Builder

	sb.WriteString("Zone {\n")
	sb.WriteString(fmt.Sprintf("  ID: %d\n", zone.ID))
	sb.WriteString(fmt.Sprintf("  ConfigPrivateZone: %t\n", zone.ConfigPrivateZone))
	sb.WriteString(fmt.Sprintf("  Name: %s\n", zone.Name))
	sb.WriteString(fmt.Sprintf("  OrgName: %s\n", zone.OrgName))
	sb.WriteString(fmt.Sprintf("  Email: %s\n", zone.Email))
	sb.WriteString(fmt.Sprintf("  DefaultTTL: %d\n", zone.DefaultTTL))
	sb.WriteString(fmt.Sprintf("  ExpireTime: %d\n", zone.ExpireTime))
	sb.WriteString(fmt.Sprintf("  NegativeCacheTTL: %d\n", zone.NegativeCacheTTL))
	sb.WriteString(fmt.Sprintf("  RefreshTime: %d\n", zone.RefreshTime))
	sb.WriteString(fmt.Sprintf("  RetryTime: %d\n", zone.RetryTime))
	sb.WriteString(fmt.Sprintf("  ParentAddress: %s\n", zone.ParentAddress))
	sb.WriteString(fmt.Sprintf("  NetworkAddress: %s\n", zone.NetworkAddress))
	sb.WriteString(fmt.Sprintf("  PostfixZoneExtension: %s\n", zone.PostfixZoneExtension))
	sb.WriteString(fmt.Sprintf("  PrefixZoneExtension: %s\n", zone.PrefixZoneExtension))

	sb.WriteString("  DNSServers:\n")
	for _, dns := range zone.DNSServers {
		sb.WriteString(fmt.Sprintf("    - Name: %s, Role: %s, SecureUpdate: %t\n", dns.Name, dns.Role, dns.SecureUpdate))
	}

	sb.WriteString("  DNSZoneOptions:\n")
	for _, zos := range zone.DNSZoneOptions {
		sb.WriteString(fmt.Sprintf("    - Name: %s\n", zos.Name))
		for _, opt := range zos.Options {
			sb.WriteString(fmt.Sprintf("      - Option Name: %s, Value: %s\n", opt.Name, opt.Value))
			for _, sub := range opt.SubOptions {
				sb.WriteString(fmt.Sprintf("        - SubOption Name: %s, Value: %s\n", sub.Name, sub.Value))
			}
		}
	}

	sb.WriteString("  UDAs:\n")
	for _, uda := range zone.UDAs {
		sb.WriteString(fmt.Sprintf("    - %v\n", uda))
	}

	sb.WriteString("  Groups:\n")
	for _, g := range zone.Groups {
		sb.WriteString(fmt.Sprintf("    - %v\n", g))
	}

	sb.WriteString("}")

	return sb.String()
}

func (zoneUpdate ZoneUpdate) String() string {
	zoneString := zoneUpdate.Zone.String()
	return fmt.Sprintf("%s, NewName: %s", zoneString, zoneUpdate.NewName)
}

type ZoneUpdate struct {
	Zone
	NewName string `json:"newName,omitempty"`
}

func NewZone(sb Zone) *Zone {
	res := sb
	res.objectType = "zone"
	return &res
}
