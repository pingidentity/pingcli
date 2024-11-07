# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_kerberos_realm)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/kerberos_realm#schema)

## Invalid attribute configuration - kerberos_password is required when connection_type is set to "DIRECT".

**Cause**: The Kerberos password is not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `kerberos_password` field in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_kerberos_realm" "my_kerberos_realm" {
  # ... other configuration parameters

  connection_type                         = "DIRECT"
  kerberos_password                       = null # sensitive
  kerberos_realm_name                     = "My Kerberos Realm"
  kerberos_username                       = "myKerberos"
}
```

After manual modification (`kerberos_password` is defined):
```hcl
resource "pingfederate_kerberos_realm" "my_kerberos_realm" {
  # ... other configuration parameters

  connection_type                         = "DIRECT"
  kerberos_password                       = var.my_kerberos_realm_password
  kerberos_realm_name                     = "My Kerberos Realm"
  kerberos_username                       = "myKerberos"
}
```

