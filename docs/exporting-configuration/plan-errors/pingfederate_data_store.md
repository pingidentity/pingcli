# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_data_store)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/data_store#nestedatt--ldap_data_store)

## Invalid attribute configuration - 'password' and 'user_dn' must be set together

**Cause**: The data store password is not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `ldap_data_store.password` field in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_data_store" "my_ldap_data_store" {
  # ... other configuration parameters

  ldap_data_store = {
    # ... other configuration parameters

    ldap_type = "PING_DIRECTORY"
    name      = "PingDirectory LDAP Data Store"
    password  = null # sensitive
    user_dn   = "cn=administrator"
  }
}
```

After manual modification (`ldap_data_store.password` is defined):
```hcl
resource "pingfederate_data_store" "my_ldap_data_store" {
  # ... other configuration parameters

  ldap_data_store = {
    # ... other configuration parameters

    ldap_type = "PING_DIRECTORY"
    name      = "PingDirectory LDAP Data Store"
    password  = var.pingdirectory_ldap_data_store
    user_dn   = "cn=administrator"
  }
}
```


