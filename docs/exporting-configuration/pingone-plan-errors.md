# Ping CLI - Exporting Platform Configuration - PingOne Plan Errors

The following sections describe the actions that must be taken, per resource, to resolve `terraform plan` errors following configuration generation.

If you encounter an error that is not documented, please [raise a new issue](https://github.com/pingidentity/pingcli/issues/new?title=Undocumented%20PingOne%20Config%20Generation%20Error).

## Resource: pingone_application

### Attribute saml_options.type value must be one of: ["WEB_APP" "CUSTOM_APP"], got: "TEMPLATE_APP"

**Cause**: Template applications are not supported in the PingOne provider version used to run `terraform plan`.

**Resolution**: Upgrade the PingOne Terraform provider version.  Further details can be found at https://github.com/pingidentity/terraform-provider-pingone/issues/841

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/application#nestedatt--saml_options)

## Resource: pingone_branding_theme

### 2 attributes specified when one (and only one) of [background_color.<.background_color,background_color.<.use_default_background,background_color.<.background_image] is required

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to ensure only one of `background_color`, `use_default_background` or `background_image` is defined.

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/branding_theme#schema)

## Resource: pingone_certificate

### one of `pem_file,pkcs7_file_base64` must be specified

**Cause**: Certificates are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set either `pem_file` or `pkcs7_file_base64` in the generated HCL.

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/certificate#schema)

## Resource: pingone_forms_recaptcha_v2

### Must set a configuration value for the secret_key attribute as the provider has marked it as required

**Cause**: The reCaptcha v2 secret key is not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `secret_key` in the generated HCL.

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/forms_recaptcha_v2#schema)

## Resource: pingone_mfa_application_push_credential

### No attribute specified when one (and only one) of [apns.<.fcm,apns.<.apns,apns.<.hms] is required

**Cause**: Push credential values are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set one of `apns`, `fcm`, or `hms` in the generated HCL.

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/mfa_application_push_credential#schema)

## Resource: pingone_notification_settings_email

### Must set a configuration value for the password attribute as the provider has marked it as required.

**Cause**: Passwords for email servers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set the `password` field in the generated HCL.

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/notification_settings_email#schema)

## Resource: pingone_phone_delivery_settings

### The argument provider_custom.authentication.password is required because provider_custom.authentication.method is configured as: "BASIC"

**Cause**: Password fields are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set the `provider_custom.authentication.password` value in the generated HCL.

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/phone_delivery_settings#password)

