# [QIP Resource Record]

Use the `vitalqip_rr` data source to retrieve information for a Resource Record managed by VitalQIP:

* `org_name` - `string`: **required**, Organization Name.
* `rr_id` - `string`: **required**, ID of resource record.
* `owner` - `string`: **optional**, Owner name for the resource record.
* `class_type` - `string`: **optional**, Class type of records pertains to a type of network or software, defaults to IN.
* `rr_type` - `string`: **optional**, Type of resource record:
  * CNAME: Canonical Name
  * A: Host IPv4
  * HINFO: Host Information
  * MX: Mail Exchange
  * NS: Name Server
  * PTR: Pointer
  * TXT: Text
  * WKS: Well Known Services
  * AAAA: Host IPv6
  * AFSDB: Andrew File System
  * MB: Mailbox Name
  * MG: Mail Group
  * MINFO: Mailbox Information
  * MR: Mail Rename
  * IDSN: Integrated Services Digital Network
  * SRV: Server Resource Record
  * X25: PSDN (Public Switched Data Network) address
  * TXT_DKIM: Test DKIM
  * CAA: Certification Authority Authorization
  * TLSA: TLS Authentication
  * NAPTR: Naming Authority Pointer
  * URI: Uniform Resource Identifier (VitalQIP 24)
* `data1` - `string`: **optional**, The data associated with the specific resource record type.
* `data2, data3, data4` - `string`: **optional**, The data associated with the specific resource record type.
* `publishing` - `string`: **optional**, Publishing type: ALWAYS, NEVER, PUSH_ONLY (Default value is ALWAYS.)
* `ttl` - `int`: **optional**, The length of time (in seconds) the name server will hold this information. If no TTL is defined, the value is inherited from the zone.
* `infra_type` - `string`: **optional**, Type of infrastructure: OBJECT, V6ADDRESS, ZONE, V4REVERSEZONE, V6REVERSEZONE, NODE
* `infra_fqdn` - `string`: **optional**, Infrastructure FQDN.
* `infra_addr` - `string`: **optional**, Address of infrastructure. Note: Required if infraType=OBJECT or infraType=V6ADDRESS, and infraFQDN is not specified.
* `is_creating_reverse_zone_rr` - `boolean`: **optional**, If set to true, the PTR resource record will be created in Reverse zone. Otherwise, resource record will be created in Object as normally. Default value is false.
* `optionalAttributeList` - `list`: **optional**, Include udas and groups.
  * `udas` - `set`: **optional**, List of UDAs.
    * `name` - `string`: **optional**, Name of the UDA.
    * `value` - `string`: **optional**, Value of the UDA.
  * `groups` - `set`: **optional**, List of groups UDA.
    * `name` - `string`: **optional**, Name of the group.
    * `udas` - `set`: **optional**, List of UDAs.

### Example of a Resource Record

This example defines a data source of type `vitalqip_rr` with the name `rr_data`, as configured in a Terraform file. By using this data source, you can reference and retrieve information about the specified Resource Record.

```hcl
data "vitalqip_rr" "rr_data" {
	org_name= "Demo"
	rr_id=52
	infra_fqdn="com"
	infra_type="ZONE"
}

output "rr_data_output" {
  value = data.vitalqip_rr.rr_data
}

```