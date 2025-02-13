# [QIP IPv6 Range]

Use the `vitalqip_ipv6_range` data source to retrieve information for an IPv6 Range managed by VitalQIP:

* `org_name` - `string`: **required**, Organization Name.
* `name` - `string`: **optional**, Range name.
* `start_address` - `string`: **required**, Starting IPv6 address.
* `range_prefix_length` - `int`: **optional**, Prefix length of range.
* `range_type` - `string`: **optional**, Type of range: DYNAMIC, FIXED, RESERVED, DYNAMIC_TEMPORARY
* `stand_prim_dhcp_server` - `string`: **optional**, Name of standard primary DHCP server.
* `failover_second_dhcp_server` - `string`: **optional**, Name of failover Secondary DHCP server.
* `opt_template` - `string`: **optional**, Template option.
* `address_selection` - `string`: **optional**, NEXT_AVAILABLE, RANDOM
* `subnet_name` - `string`: **optional**, Name of subnet.
* `subnet_address` - `string`: **optional**, Address of subnet.
* `subnet_prefix_length` - `int`: **optional**, Prefix length of subnet.
* `udas` - `set`: **optional**, List of UDAs.
  * `name` - `string`: **optional**, Name of the UDA.
  * `value` - `string`: **optional**, Value of the UDA.
* `groups` - `set`: **optional**, List of groups UDA.
  * `name` - `string`: **optional**, Name of the group.
  * `udas` - `set`: **optional**, List of UDAs.

### Example of a IPv6 Range

This example defines a data source of type `vitalqip_ipv6_range` with the name `ipv6_range_data`, as configured in a Terraform file. By using this data source, you can reference and retrieve information about the specified IPv6 Range.

```hcl
data "vitalqip_ipv6_range" "ipv6_range_data" {
  org_name= "Demo"
  start_address="2000::"
}

output "ipv6_range_data_output" {
  value = data.vitalqip_ipv6_range.ipv6_range_data
}

```