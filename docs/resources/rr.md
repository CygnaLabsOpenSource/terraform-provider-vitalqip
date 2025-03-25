# Resource: [QIP Resource Record]

###  Descriptions
The `vitalqip_rr` resource is designed to manage and associate a Resource Record with an organizational structure within VitalQIP.

### Parameters
The following list describes the parameters you can define in the Resource Record of the record:

* `org_name` - `string`: **required**, Organization Name.
* `rr_id` - `string`: **optional**, ID of resource record.
* `owner` - `string`: **required**, Owner name for the resource record.
* `class_type` - `string`: **optional**, Class type of records pertains to a type of network or software, defaults to IN.
* `rr_type` - `string`: **required**, Type of resource record:
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
  * TXT_DKIM: Text DKIM
  * CAA: Certification Authority Authorization
  * TLSA: TLS Authentication
  * NAPTR: Naming Authority Pointer
  * URI: Uniform Resource Identifier (VitalQIP 24)
* `data1` - `string`: **required**, The data associated with the specific resource record type.
* `data2, data3, data4` - `string`: **optional**, The data associated with the specific resource record type.
* `publishing` - `string`: **optional**, Publishing type: ALWAYS, NEVER, PUSH_ONLY (Default value is ALWAYS.)
* `ttl` - `int`: **optional**, The length of time (in seconds) the name server will hold this information. If no TTL is defined, the value is inherited from the zone.
* `infra_type` - `string`: **required**, Type of infrastructure: OBJECT, V6ADDRESS, ZONE, V4REVERSEZONE, V6REVERSEZONE, NODE
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

### ⚠️ Force replacement fields
The following fields after changes will require deleting and recreating the resource:
* `org_name`
* `infra_type`
* `infra_fqdn`
* `infra_addr`

> **WARNING**: Changing the above fields will result in the current resource being deleted and a new one created. Because the resource record modification API does not support changes to them. Make sure you back up your data and understand the impact before making changes.

## How to use
First define `resource` in the .tf file.<br>
`Resource Record` example:

```hcl
resource "vitalqip_rr" "rr_resource" {
  // required parameters
	org_name= "Demo"
	owner="owner"
	rr_type="A"
	data1="6.6.6.6"
	infra_type="ZONE"

  // optional parameters
  infra_fqdn="com"
  class_type="IN"
  publishing="ALWAYS"
  ttl=123
  is_creating_reverse_zone_rr=false
	optional_attribute_list {
    udas {
      name  = "att1"
      value = "true"
    }
    udas {
      name  = "att2"
      value = "test1"
    }
    groups {
      name = "group"
      udas {
        name  = "att1"
        value = "false"
      }
      udas {
        name  = "att2"
        value = "test2"
      }
    }
  }
}


output "rr_resource_output" {
  value = resource.vitalqip_rr.rr_resource
}
```

Then run
```bash
terraform apply
```