package entities

import (
	"fmt"
	"strings"
)

type RR struct {
	ObjBase                 `json:"-"`
	OrgName                 string                `json:"orgName,omitempty"`
	RRID                    int                   `json:"rrId,omitempty"`
	Owner                   string                `json:"owner,omitempty"`
	ClassType               string                `json:"classType,omitempty"`
	RrType                  string                `json:"rrType,omitempty"`
	Data1                   string                `json:"data1,omitempty"`
	Publishing              string                `json:"publishing,omitempty"`
	TTL                     int                   `json:"ttl,omitempty"`
	InfraType               string                `json:"infraType,omitempty"`
	InfraFQDN               string                `json:"infraFQDN,omitempty"`
	Data2                   string                `json:"data2,omitempty"`
	Data3                   string                `json:"data3,omitempty"`
	Data4                   string                `json:"data4,omitempty"`
	InfraAddr               string                `json:"infraAddr,omitempty"`
	IsCreatingReverseZoneRR bool                  `json:"isCreatingReverseZoneRR,omitempty"`
	IsDefaultRR             bool                  `json:"isDefaultRR,omitempty"`
	OptionalAttributeList   OptionalAttributeList `json:"optionalAttributeList,omitempty"`
}

type OptionalAttributeList struct {
	UDAs   []UDA   `json:"udas,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}

func (rr RR) String() string {
	var sb strings.Builder

	sb.WriteString("RR {\n")
	sb.WriteString(fmt.Sprintf("  OrgName: %s\n", rr.OrgName))
	sb.WriteString(fmt.Sprintf("  RrID: %d\n", rr.RRID))
	sb.WriteString(fmt.Sprintf("  Owner: %s\n", rr.Owner))
	sb.WriteString(fmt.Sprintf("  ClassType: %s\n", rr.ClassType))
	sb.WriteString(fmt.Sprintf("  RrType: %s\n", rr.RrType))
	sb.WriteString(fmt.Sprintf("  Data1: %s\n", rr.Data1))
	sb.WriteString(fmt.Sprintf("  Publishing: %s\n", rr.Publishing))
	sb.WriteString(fmt.Sprintf("  TTL: %d\n", rr.TTL))
	sb.WriteString(fmt.Sprintf("  InfraType: %s\n", rr.InfraType))
	sb.WriteString(fmt.Sprintf("  InfraFQDN: %s\n", rr.InfraFQDN))
	sb.WriteString(fmt.Sprintf("  Data2: %s\n", rr.Data2))
	sb.WriteString(fmt.Sprintf("  Data3: %s\n", rr.Data3))
	sb.WriteString(fmt.Sprintf("  Data4: %s\n", rr.Data4))
	sb.WriteString(fmt.Sprintf("  InfraAddr: %s\n", rr.InfraAddr))
	sb.WriteString(fmt.Sprintf("  IsCreatingReverseZoneRR: %t\n", rr.IsCreatingReverseZoneRR))
	sb.WriteString(fmt.Sprintf("  IsDefaultRR: %t\n", rr.IsDefaultRR))
	sb.WriteString(fmt.Sprintf("  OptionalAttributeList: %s\n", rr.OptionalAttributeList))
	sb.WriteString("}\n")

	return sb.String()
}

func NewRR(sb RR) *RR {
	res := sb
	res.objectType = "rr"
	return &res
}

type RRResponse struct {
	List []RR `json:"list"`
}

type UpdateFields struct {
	Owner                   string                `json:"owner,omitempty"`
	ClassType               string                `json:"classType,omitempty"`
	RRType                  string                `json:"rrType,omitempty"`
	Data1                   string                `json:"data1,omitempty"`
	Publishing              string                `json:"publishing,omitempty"`
	TTL                     string                `json:"ttl,omitempty"`
	Data2                   string                `json:"data2,omitempty"`
	Data3                   string                `json:"data3,omitempty"`
	Data4                   string                `json:"data4,omitempty"`
	IsCreatingReverseZoneRR bool                  `json:"isCreatingReverseZoneRR"`
	OptionalAttributeList   OptionalAttributeList `json:"optionalAttributeList,omitempty"`
}

type RRUpdate struct {
	ObjBase      `json:"-"`
	OrgName      string       `json:"orgName,omitempty"`
	RRID         int          `json:"rrId"`
	UpdateFields UpdateFields `json:"updateFields"`
}

func NewRRUpdate(sb RRUpdate) *RRUpdate {
	res := sb
	res.objectType = "rr_update"
	return &res
}
