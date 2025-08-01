# Resource: [QIP Zone]

###  Descriptions
The `vitalqip_zone` resource is designed to manage and associate a Zone with an organizational structure within VitalQIP.

### Parameters
The following list describes the parameters you can define in the Zone of the record:

* `org_name` - `string`: **required**, Organization Name.
* `id` - `string`: **optional**, ID of zone.
* `name` - `string`: **required**, Name of the zone.
* `default_ttl` - `int`: **required**, DNS default TTL.
* `email` - `string`: **required**, Email contact.
* `expire_time` - `int`: **required**, Expire time.
* `negative_cache_ttl` - `int`: **required**, DNS negative cache TTL.
* `refresh_time` - `int`: **required**, DNS refresh time.
* `retry_time` - `int`: **required**, DNS retry time.
* `parent_address` - `string`: **optional**, Parent address of the split reverse zone.
* `network_address` - `string`: **optional**, Network address of the reverse zone.
* `postfix_zone_extension` - `string`: **optional**, DNS postfix extension.
* `prefix_zone_extension` - `string`: **optional**, DNS prefix extension.
* `config_private_zone` - `boolean`: **optional**, Config private zone.
* `dns_servers` - `set`: **optional**, List of DNS servers.
    * `name` - `string`: **optional**, Name of DNS Server.
    * `role` - `string`: **optional**, DNS Server Role. Values: P for Primary or S for Secondary.
    * `secure_update` - `boolean`: **optional**, Send Secure Updates. Values: true or false.
* `udas` - `set`: **optional**, List of UDAs.
    * `name` - `string`: **optional**, Name of the UDA.
    * `value` - `string`: **optional**, Value of the UDA.
* `groups` - `set`: **optional**, List of UDA groups.
    * `name` - `string`: **optional**, Name of the UDA group.
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

### ⚠️ Force replacement fields
The following fields after changes will require deleting and recreating the resource:
* `org_name`

> **WARNING**: Changing the above fields will result in the current resource being deleted and a new one created. Because the resource record modification API does not support changes to them. Make sure you back up your data and understand the impact before making changes.

## How to use
First define `zone_resource` in the .tf file.<br>
`Zone` example:

```hcl
resource "vitalqip_zone" "zone_resource" {
  // required parameters
  org_name = "Terraform"
  name = "test.com"
  negative_cache_ttl = 600
  expire_time = 604800
  email = "example@gmail.com"
  retry_time = 3600
  default_ttl = 86400
  refresh_time = 21600

  // optional parameters
  udas {
      name  = "uda"
      value = "true"
  }
  udas {
      name  = "ipv4"
	  value = "6.6.8.8"
  }
  groups {
      name = "udg"
      udas {
        name  = "uda"
        value = "false"
      }
      udas {
        name  = "string"
        value = "false"
      }
  }
  dns_servers {
      name = "sv1.example.com"
      role =  "P"
      secure_update = false
  }
  dns_zone_options {
    name = "DNS 6.4 Options"

    options {
      name  = "DNS RPZ (Firewall)"
      value = ""

      sub_options {
        name  = "max-policy-ttl"
        value = "Use Server Value"
      }

      sub_options {
        name  = "policy"
        value = "Use Server Value"
      }

      sub_options {
        name  = "recursive-only"
        value = "Use Server Value"
      }
    }

    options {
      name  = "DNSSEC enabled zone"
      value = "False"

      sub_options {
        name  = "Auto-DNSSEC"
        value = "Maintain"
      }

      sub_options {
        name  = "Inline-signing"
        value = "Yes"
      }
    }

    options {
      name  = "Import External Updates"
      value = "True"

      sub_options {
        name  = "A (Host IPV4)"
        value = "True"
      }

      sub_options {
        name  = "AAAA (Host IPV6)"
        value = "True"
      }

      sub_options {
        name  = "CNAME (Canonical Name)"
        value = "True"
      }

      sub_options {
        name  = "PTR (Pointer)"
        value = "False"
      }

      sub_options {
        name  = "SRV (Server Resource Record)"
        value = "False"
      }

      sub_options {
        name  = "TXT (Text)"
        value = "True"
      }
    }

    options {
      name  = "allow-notify"
      value = "Use Server Value"

      sub_options {
        name  = "ACL Templates"
        value = ""
      }

      sub_options {
        name  = "other"
        value = ""
      }
    }

    options {
      name  = "allow-query"
      value = "Any"

      sub_options {
        name  = "ACL Templates"
        value = ""
      }

      sub_options {
        name  = "other"
        value = ""
      }
    }

    options {
      name  = "allow-transfer"
      value = "Any"

      sub_options {
        name  = "ACL Templates"
        value = ""
      }

      sub_options {
        name  = "other"
        value = ""
      }
    }

    options {
      name  = "allow-update"
      value = "Use Server Value"

      sub_options {
        name  = "ACL Templates"
        value = ""
      }

      sub_options {
        name  = "other"
        value = ""
      }
    }

    options {
      name  = "forwarders"
      value = ""

      sub_options {
        name  = "forward"
        value = ""
      }
    }

    options {
      name  = "notify"
      value = "Yes"

      sub_options {
        name  = "also-notify"
        value = ""
      }
    }

    options {
      name  = "zone block of named.conf"
      value = ""
    }
  }
}

output "vitalqip_zone_output" {
  value = resource.vitalqip_zone.zone_resource
}
```

Then run
```bash
terraform apply
```