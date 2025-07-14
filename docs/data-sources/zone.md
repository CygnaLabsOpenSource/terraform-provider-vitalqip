# [QIP Zone]

Use the `vitalqip_zone` data source to retrieve information for a Zone managed by VitalQIP:

* `org_name` - `string`: **required**, Organization Name.
* `id` - `string`: **optional**, ID of zone.
* `name` - `string`: **required**, Name of the zone.
* `default_ttl` - `int`: **optional**, DNS default TTL.
* `email` - `string`: **optional**, Email contact.
* `expire_time` - `int`: **optional**, Expire time.
* `negative_cache_ttl` - `int`: **optional**, DNS negative cache TTL.
* `refresh_time` - `int`: **optional**, DNS refresh time.
* `retry_time` - `int`: **optional**, DNS retry time.
* `parent_address` - `string`: **optional**, Parent address of the splitted reverse zone.
* `network_address` - `string`: **optional**, Network address of the reverse zone.
* `postfix_zone_extension` - `string`: **optional**, DNS postfix extension.
* `prefix_zone_extension` - `string`: **optional**, DNS prefix extension.
* `config_private_zone` - `boolean`: **optional**, Config private zone.
* `dns_servers` - `set`: **optional**, List of DNS servers.
    * `name` - `string`: **optional**, Name of DNS Server.
    * `role` - `string`: **optional**, Primary/Secondary DNS Server: With P is for Primary and S is for Secondary
    * `secure_update` - `boolean`: **optional**, true is turn on the flag of sending secure update false is not.
* `udas` - `set`: **optional**, List of UDAs.
    * `name` - `string`: **optional**, Name of the UDA.
    * `value` - `string`: **optional**, Value of the UDA.
* `groups` - `set`: **optional**, List of groups UDA.
    * `name` - `string`: **optional**, Name of the group.
    * `udas` - `set`: **optional**, List of UDAs.
        * `name` - `string`: **optional**, Name of the UDA.
        * `value` - `string`: **optional**, Value of the UDA.
* `dns_zone_options` - `set`: **optional**, DNS zone options.
    * `name` - `string`: **optional**, Group name of zone option.
    * `options` - `set`: **optional**, List of zone options.
        * `name` - `string`: **optional**, Name of zone option.
        * `value` - `string`: **optional**, Value of zone option.
        * `sub_options` - `set`: **optional**, List of sub options.
            * `name` - `string`: **optional**, Name of sub option.
            * `value` - `string`: **optional**, Value of sub option.

### Example of a Resource Record

This example defines a data source of type `vitalqip_zone` with the name `zone_data`, as configured in a Terraform file. By using this data source, you can reference and retrieve information about the specified Zone.

```hcl
data "vitalqip_zone" "zone_data" {
  org_name = "Terraform"
  name = "test.com"
}

output "vitalqip_zone_output" {
  value = data.vitalqip_zone.zone_data
}
```