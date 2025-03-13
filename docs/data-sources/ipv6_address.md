# [QIP IPv6 Address]

Use the `vitalqip_ipv6_address` data source to retrieve information for an IPv6 Address managed by VitalQIP:

* `org_name` - `string`: **required**, Organization Name.
* `host_name` - `string`: **optional**, Hostname of address.
* `address` - `string`: **optional**, IPv6 address.
* `domain_name` - `string`: **optional**, Domain name of address.
* `range_address` - `string`: **optional**, Address range.
* `address_type` - `string`: **optional**, Type of address.
* `publish_a` - `string`: **optional**, AAAA resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL
* `publish_ptr` - `string`: **optional**, PTR resource record option: ALWAYS, NEVER, PUSH_ONLY, EXTERNAL
* `class_type` - `string`: **optional**, Class of object.
* `iaid` - `string`: **optional**, Identity Association Identifier (IAID).
* `range_name` - `string`: **optional**, Name of range.
* `ttl` - `int`: **optional**, Time to live.
* `node_name` - `string`: **optional**, Name of node.
* `unique_id` - `string`: **optional**, Unique ID.
* `duid` - `string`: **optional**, DUID. Note: IPv6 address using DUID will be dynamic updated to DHCPv6 Server.
* `use_mac_address` - `boolean`: **optional**, True: use MAC address, False: use DUID value instead.
* `mac_address` - `string`: **optional**, MAC address.
* Note: The user must supply `address` or `host_name`. If both are present, `address` will take precedence.

### Example of a IPv6 Address

This example defines a data source of type `vitalqip_ipv6_address` with the name `ipv6_address_data`, as configured in a Terraform file. By using this data source, you can reference and retrieve information about the specified IPv6 Address.

```hcl
data "vitalqip_ipv6_address" "ipv6_address_data" {
  org_name= "Terraform"
  address="2000::5"
}

output "ipv6_address_data_output" {
  value = data.vitalqip_ipv6_address.ipv6_address_data
}

```