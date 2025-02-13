# Resource: [QIP IPv6 Address]

###  Descriptions
The `vitalqip_ipv6_address` resource is designed to manage and associate an IPv6 address with an organizational structure within VitalQIP.

### Parameters
The following list describes the parameters you can define in the resource IPv6 Address of the record:

* `org_name` - `string`: **required**, Organization Name.
* `host_name` - `string`: **required**, Hostname of address.
* `address` - `string`: **required**, IPv6 address.
* `domain_name` - `string`: **required**, Domain name of address.
* `range_address` - `string`: **required**, Address range.
* `address_type` - `string`: **required**, Type of address.
* `publish_a` - `string`: **required**, AAAA resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL
* `publish_ptr` - `string`: **required**, PTR resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL
* `class_type` - `string`: **required**, Class of object.
* `fqdn` - `string`: **optional**, Fully qualified domain name.
* `iaid` - `string`: **optional**, Identity Association Identifier (IAID).
* `range_name` - `string`: **optional**, Name of range.
* `ttl` - `int`: **optional**, Time to live.
* `node_name` - `string`: **optional**, Name of node.
* `unique_id` - `string`: **optional**, Unique ID.
* `duid` - `string`: **optional**, DUID. Note: IPv6 address using DUID will be dynamic updated to DHCPv6 Server.
* `use_mac_address` - `boolean`: **optional**, True: use MAC address, False: use DUID value instead.
* `mac_address` - `string`: **optional**, MAC address.

### ⚠️ Force replacement fields
The following fields after changes will require deleting and recreating the resource:
* `org_name`
* `address`
* `range_address`
* `range_name`

> **WARNING**: Changing the above fields will result in the current resource being deleted and a new one created. Because the IPv6 address modification API does not support changes to them. Make sure you back up your data and understand the impact before making changes.

## How to use
First define `resource` in the .tf file.<br>
`IPv6 Address` example:

```hcl
resource "vitalqip_ipv6_address" "ipv6_address_resource" {
  // required parameters
  org_name= "Terraform"
  host_name="v6obj"
  address="2000::5"
  domain_name="com"
  range_address="2000::/112"
  address_type="MANUAL"
  publish_a="ALWAYS"
  publish_ptr="ALWAYS"
  class_type="Workstation"
  
  // optional parameters
  fqdn="v6obj.com"
  iaid="10"
  range_name="range1"
  ttl="100"
  node_name="v6obj.com"
  unique_id="123"
  //duid="11-12-22-33-31-11"
  use_mac_address=true
  mac_address="11-12-22-33-31-11"
}

output "ipv6_address_resource_output" {
  value = resource.vitalqip_ipv6_address.ipv6_address_resource
}
```

Then run
```bash
terraform apply
```