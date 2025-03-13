# Resource: [QIP IPv4 Address]

###  Descriptions
The `vitalqip_ipv4_address` resource is designed to manage and associate an IPv4 address with an organizational structure within VitalQIP.

### Parameters
The following list describes the parameters you can define in the resource IPv4 Address of the record:

* `org_name` - `string`: **required**, Organization Name.
* `object_addr` - `string`: **required**, IPv4 address.
* `object_name` - `string`: **optional**, Name of IPv4 object.
* `subnet_addr` - `string`: **optional**, Subnet of IPv4 object.
* `object_class` - `string`: **optional**, Class type of IPv4 object:
  -  Workstation
  -  X-terminal
  -  PC
  -  Printer
  -  Server
  -  Router
  -  Bridge
  -  Wiring_HUB
  -  Switch
  -  Gateway
  -  Undefined
  -  Others
  -  Any user-defined object class name.
* `expirated_date` - `string`: **optional**, The date when a reserved object expires and is no longer reserved.The date format for this field is yyyy-mm-dd.
* `domain_name` - `string`: **optional**, Domain of IPv4 object.
* `object_desc` - `string`: **optional**, Description of IPv4 object.
* `dynamic_config` - `string`: **optional**, Dynamic Configuration of object include:
  -  Blank - Static object (This is the default status when the user skips this field).
  -  None - Dynamic (none)
  -  M-DHCP - Manual DHCP
  -  A-DHCP - Automatic DHCP
  -  D-DHCP - Dynamic DHCP
  -  M-BOOTP - Manual Bootp
  -  A-BOOTP - Automatic Bootp
  -  Reserved - To reserve an object (the IP must be unused).
* `mac_addr` - `string`: **optional**, MAC address of IPv4 object. Note: Exclude colons (:).
* `a_ttl` - `string`: **optional**, Time to live value of default A resource record.
* `ptr_ttl` - `string`: **optional**, Time to live value of default PTR resource record.
* `publish_a` - `string`: **optional**, A resource record option: Always, None, Push Only.
* `publish_ptr` - `string`: **optional**, PTR resource record option: Always, None, Push Only.

### ⚠️ Force replacement fields
The following fields after changes will require deleting and recreating the resource:
* `org_name`
* `object_addr`
* `subnet_addr`

> **WARNING**: Changing the above fields will result in the current resource being deleted and a new one created. Because the IPv4 address modification API does not support changes to them. Make sure you back up your data and understand the impact before making changes.

## How to use
First define `resource` in the .tf file.<br>
`IPv4 Address` example:

```hcl
resource "vitalqip_ipv4_address" "ipv4_address_resource" {
  // required parameters
  org_name= "Terraform"
  object_addr = "10.0.0.5"
  
  // optional parameters
  dynamic_config = "Static"
  object_class = "Workstation"
  object_desc = "desc"
  object_name = "obj5"
  subnet_addr = "10.0.0.0"
  domain_name = "com"
  expirated_date = "2024-10-10"
  mac_addr = "111222333444"
  a_ttl = "-1"
  ptr_ttl = "-1"
  publish_a = "Always"
  publish_ptr = "Always"
}

output "ipv4_address_resource_output" {
  value = resource.vitalqip_ipv4_address.ipv4_address_resource
}
```

Then run
```bash
terraform apply
```