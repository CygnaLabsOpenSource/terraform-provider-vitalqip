# Resource: [QIP IPv6 Range]

###  Descriptions
The `vitalqip_ipv6_range` resource is designed to manage and associate an IPv6 Range with an organizational structure within VitalQIP.

### Parameters
The following list describes the parameters you can define in the resource IPv6 Range of the record:

* `org_name` - `string`: **required**, Organization Name.
* `name` - `string`: **required**, Range name.
* `start_address` - `string`: **required**, Starting IPv6 address.
* `new_start_address` - `string`: **optional**, New starting IPv6 address.
* `range_prefix_length` - `int`: **required**, Prefix length of range.
* `range_type` - `string`: **required**, Type of range: DYNAMIC, FIXED, RESERVED, DYNAMIC_TEMPORARY
* `stand_prim_dhcp_server` - `string`: **optional**, Name of standard primary DHCP server.
* `failover_second_dhcp_server` - `string`: **optional**, Name of failover Secondary DHCP server.
* `opt_template` - `string`: **optional**, Template option.
* `address_selection` - `string`: **required**, NEXT_AVAILABLE, RANDOM
* `subnet_name` - `string`: **optional**, Name of subnet.
* `subnet_address` - `string`: **required**, Address of subnet.
* `subnet_prefix_length` - `int`: **required**, Prefix length of subnet.
* `udas` - `set`: **optional**, List of UDAs.
  * `name` - `string`: **optional**, Name of the UDA.
  * `value` - `string`: **optional**, Value of the UDA.
* `groups` - `set`: **optional**, List of groups UDA.
  * `name` - `string`: **optional**, Name of the group.
  * `udas` - `set`: **optional**, List of UDAs.
* `dhcpParams` - `object`: **optional**, DHCP parameters.
  * `information_refresh_time` - `string`: **optional**, This option specifies an upper bound for how long a client should wait before refreshing information retrieved from DHCPv6.
  * `preferred_life_time` - `string`: **optional**, This option specifies the IPv6 prefix, which constitutes an agreement about the length of time over which the requesting router is allowed to use the prefix.
  * `rebinding_time` - `string`: **optional**, This option specifies the time interval from address assignment until the client transitions to the REBINDING state. The value is in seconds and is specified as a 32-bit unsigned integer.
  * `renewal_time` - `string`: **optional**, This option specifies the time interval from address assignment until the client transitions to the RENEWING state. The value is seconds and is specified as a 32-bit unsigned integer.
  * `valid_life_time` - `string`: **optional**, This option specifies the time duration for which an address remains in the valid state (i.e., the time till the address gets invalid). The valid lifetime must be greater than or equal to the preferred lifetime. When the valid lifetime expires, the address becomes invalid.
  * `bcmcs_server_address_list` - `string`: **optional**, This option specifies a list of IPv6 addresses indicating SIP outbound proxy servers available to the client. The servers are listed in order of preference.
  * `bcmcs_server_domain_name_list` - `string`: **optional**, This option specifies the domain names of the SIP outbound proxy servers for the client to use. The list of domain names may contain the domain name of the access provider and its partner networks that also offer broadcast and multicast service.
  * `dns_recursive_name_server` - `string`: **optional**, This option provides a list of one or more IPv6 addresses of DNS recursive name servers to which a client's DNS resolver might send DNS queries.
  * `domain_search_list` - `string`: **optional**, This option specifies the domain search list that the client needs to use when resolving hostname with DNS. It does not apply to other name resolution mechanisms.
  * `nis_servers` - `string`: **optional**, This option provides a list of one or more IPv6 addresses of NIS servers available to the client.
  * `nisp_domain_name` - `string`: **optional**, This option is used by the server to convey the NIS and domain name information to the client.
  * `pana_authentication_agents` - `string`: **optional**, This option defines a DHCPv6 option that carries a list of 128-bit (binary) IPv6 addresses indicating one or more PANA Authentication Agents (PAA) available to the PANA client.
  * `posix_time_zone` - `string`: **optional**, This option specifies the POSIX style timezone.
  * `sip_servers_domain_name_list` - `string`: **optional**, This option specifies domain names in messages that are expressed in terms of a sequence of labels. Each label is represented as a one octet length field followed by that number of octets. The total length of a domain name is restricted to 255 octets or less.
  * `sip_servers_ipv6_address_list` - `string`: **optional**, This option specifies a list of IPv6 addresses indicating SIP outbound proxy servers available to the client.
  * `sntp_servers` - `string`: **optional**, This option (Simple Network Time Protocol servers) provides a list of one or more IPv6 addresses of SNTP servers available to the client for synchronization. The clients use these SNTP servers to synchronize their system time to that of the standard time servers.
  * `tzdb_time_zone` - `string`: **optional**, This option specifies the TZTB timezone.
  * `vendor_options` - `string`: **optional**, This option is used by clients and servers to exchange vendor-specific information.
### ⚠️ Force replacement fields
The following fields after changes will require deleting and recreating the resource:
* `org_name`
* `start_address`
* `subnet_name`
* `subnet_address`
* `subnet_prefix_length`

> **WARNING**: Changing the above fields will result in the current resource being deleted and a new one created. Because the IPv6 range modification API does not support changes to them. Make sure you back up your data and understand the impact before making changes.

## How to use
First define `resource` in the .tf file.<br>
`IPv6 Range` example:

```hcl
resource "vitalqip_ipv6_range" "ipv6_range_resource" {
  // required parameters
  org_name= "Terraform"
  name="range"
  start_address="2000::4:0"
  range_prefix_length=112
  range_type="DYNAMIC"
  address_selection="NEXT_AVAILABLE"
  subnet_prefix_length=60
  subnet_address="2000::"

  // optional parameters
  new_start_address="2000::4:0"
  stand_prim_dhcp_server="dhcpv6.com"
  failover_second_dhcp_server="seconddhcpv6.com"
  opt_template="opt"
  subnet_name="subnet"
  dhcp_params {
	information_refresh_time        = "3600"
    preferred_life_time             = "7200"
    rebinding_time                  = "5400"
    renewal_time                    = "3600"
    valid_life_time                 = "86400"
	  bcmcs_server_address_list   = "192.168.0.1"
	  bcmcs_server_domain_name_list = "example.com"
	  dns_recursive_name_server   = "8.8.8.8"
	  domain_search_list          = "example.com"
	  nis_servers                 = "192.168.0.2"
	  nisp_domain_name            = "nis.example.com"
	  pana_authentication_agents  = "pana_agent1"
	  posix_time_zone             = "PST"
	  sip_servers_domain_name_list = "sip.example.com"
	  sip_servers_ipv6_address_list = "2001:db8::1"
	  sntp_servers                = "time.example.com"
	  tzdb_time_zone              = "America/Los_Angeles"
	  vendor_options             = "option1"
  }
  udas {
    name  = "att1"
    value = "false"
  }
  
  udas {
    name  = "att2"
    value = "text"
  }
  
  groups {
	name="group1"
	udas {
		name  = "att1"
		value = "true"
	}
	udas {
		name  = "att2"
		value = "text"
	}
  }
}

output "ipv6_range_resource_output" {
  value = resource.vitalqip_ipv6_range.ipv6_range_resource
}
```

Then run
```bash
terraform apply
```