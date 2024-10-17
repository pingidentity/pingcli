# Terraform Configuration Generation - PingOne Plan Errors (pingone_forms_recaptcha_v2)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/forms_recaptcha_v2#schema)

## Missing Configuration for Required Attribute - Must set a configuration value for the secret_key attribute as the provider has marked it as required

**Cause**: The reCaptcha v2 secret key is not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `secret_key` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_forms_recaptcha_v2" "my_awesome_recaptcha_configuration" {
  environment_id = "942b****-****-****-****-********985c"
  secret_key     = null # sensitive
  site_key       = "6L****************-******************hp"
}
```

After manual modification (`secret_key` is defined):
```hcl
resource "pingone_forms_recaptcha_v2" "my_awesome_recaptcha_configuration" {
  environment_id = "942b****-****-****-****-********985c"
  secret_key     = var.my_awesome_recaptcha_configuration_secret_key
  site_key       = "6L****************-******************hp"
}
```
