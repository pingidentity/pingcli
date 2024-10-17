# Terraform Configuration Generation - PingOne Plan Errors (pingone_notification_settings_email)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/notification_settings_email#schema)

## Missing Configuration for Required Attribute - Must set a configuration value for the password attribute as the provider has marked it as required.

**Cause**: Passwords for email servers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set the `password` field in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_notification_settings_email" "pingone_notification_settings_email" {
  # ... other configuration parameters

  host     = "smtp-example.bxretail.org"
  password = null # sensitive
  username = "test"
}
```

After manual modification (`password` is defined):
```hcl
resource "pingone_notification_settings_email" "pingone_notification_settings_email" {
  # ... other configuration parameters

  host     = "smtp-example.bxretail.org"
  password = var.pingone_notification_settings_email_password
  username = "test"
}
```
