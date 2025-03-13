# [QIP IPv4 Address]

Use the `vitalqip_ipv4_address` data source to retrieve information for an IPv4 Address managed by VitalQIP:

* `org_name` - `string`: **required**, Organization Name.
* `object_addr` - `string`: **optional**, IPv4 address.
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
* Note: The user must supply `object_addr` or `object_name`. If both are present, `object_addr` will take precedence.

### Example of a IPv4 Address

This example defines a data source of type `vitalqip_ipv4_address` with the name `ipv4_address_data`, as configured in a Terraform file. By using this data source, you can reference and retrieve information about the specified IPv4 Address.

```hcl
data "vitalqip_ipv4_address" "ipv4_address_data" {
  org_name= "Terraform"
  object_addr="10.0.0.3"
}

output "ipv4_address_data_output" {
  value = data.vitalqip_ipv4_address.ipv4_address_data
}

```