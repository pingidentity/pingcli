# Terraform Configuration Generation - PingOne Plan Errors (pingone_phone_delivery_settings)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/phone_delivery_settings#password)

## Missing required argument - The argument provider_custom.authentication.password is required because provider_custom.authentication.method is configured as: "BASIC"

**Cause**: Password fields are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set the `provider_custom.authentication.password` value in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_phone_delivery_settings" "my_awesome_phone_delivery_settings" {
  # ... other configuration parameters

  provider_custom = {
    # ... other configuration parameters

    authentication = {
      auth_token = null # sensitive
      method     = "BASIC"
      password   = null # sensitive
      username   = "myusername"
    }
  }
}
```

After manual modification (`provider_custom.authentication.password` is defined):
```hcl
resource "pingone_phone_delivery_settings" "my_awesome_phone_delivery_settings" {
  # ... other configuration parameters

  provider_custom = {
    # ... other configuration parameters

    authentication = {
      method     = "BASIC"
      password   = var.my_phone_delivery_settings_password
      username   = "myusername"
    }
  }
}
```
