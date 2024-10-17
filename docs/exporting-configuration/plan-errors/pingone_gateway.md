# Terraform Configuration Generation - PingOne Plan Errors (pingone_gateway)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/gateway#schema)

## Invalid Attribute Combination - Attribute "bind_password" must be specified when "[bind_dn|connection_security|follow_referrals|servers|user_types|validate_tls_certificates|vendor]" is specified

**Cause**: The LDAP bind password is not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `bind_password` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_gateway" "my_pingdirectory" {
  # ... other configuration parameters

  bind_dn             = "cn=administrator"
  bind_password       = null # sensitive
  connection_security = "TLS"
  follow_referrals    = false
  servers             = ["my-directory:636"]
  type                = "LDAP"
  user_types          = {
    # ... other configuration parameters
  }
  validate_tls_certificates = true
  vendor                    = "PingDirectory"
}
```

After manual modification (`bind_password` is defined):
```hcl
resource "pingone_gateway" "my_pingdirectory" {
  # ... other configuration parameters

  bind_dn             = "cn=administrator"
  bind_password       = var.my_directory_pingdirectory_bind_dn
  connection_security = "TLS"
  follow_referrals    = false
  servers             = ["my-directory:636"]
  type                = "LDAP"
  user_types          = {
    # ... other configuration parameters
  }
  validate_tls_certificates = true
  vendor                    = "PingDirectory"
}
```

## Missing required argument - The argument bind_password is required because type is configured as: "LDAP"

**Cause**: The LDAP bind password is not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `bind_password` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_gateway" "my_pingdirectory" {
  # ... other configuration parameters

  bind_dn             = "cn=administrator"
  bind_password       = null # sensitive
  connection_security = "TLS"
  follow_referrals    = false
  servers             = ["my-directory:636"]
  type                = "LDAP"
  user_types          = {
    # ... other configuration parameters
  }
  validate_tls_certificates = true
  vendor                    = "PingDirectory"
}
```

After manual modification (`bind_password` is defined):
```hcl
resource "pingone_gateway" "my_pingdirectory" {
  # ... other configuration parameters

  bind_dn             = "cn=administrator"
  bind_password       = var.my_directory_pingdirectory_bind_dn
  connection_security = "TLS"
  follow_referrals    = false
  servers             = ["my-directory:636"]
  type                = "LDAP"
  user_types          = {
    # ... other configuration parameters
  }
  validate_tls_certificates = true
  vendor                    = "PingDirectory"
}
```
