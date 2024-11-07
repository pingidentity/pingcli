# Terraform Configuration Generation - PingOne Plan Errors

The following sections describe the actions that must be taken, per resource, to resolve `terraform plan` errors following configuration generation.

If you encounter an error that is not documented, please [raise a new issue](https://github.com/pingidentity/pingcli/issues/new?title=Undocumented%20PingOne%20Config%20Generation%20Error).

## General (Any Resource)

### Reference to undeclared resource - A managed resource "[any]" "[any]" has not been declared in the root module

**Cause**: Terraform configuration has been generated with syntax errors.  This is an issue with the Terraform CLI.

**Resolution**: Upgrade the Terraform CLI to the latest version available and re-generate the HCL configuration.

## Resource Plan Errors

- [pingone_application](pingone_application.md)
- [pingone_branding_theme](pingone_branding_theme.md)
- [pingone_certificate](pingone_certificate.md)
- [pingone_forms_recapcha_v2](pingone_forms_recapcha_v2.md)
- [pingone_gateway](pingone_gateway.md)
- [pingone_identity_provider](pingone_identity_provider.md)
- [pingone_mfa_application_push_credential](pingone_mfa_application_push_credential.md)
- [pingone_notification_settings_email](pingone_notification_settings_email.md)
- [pingone_phone_delivery_settings](pingone_phone_delivery_settings.md)
- [pingone_schema_attribute](pingone_schema_attribute.md)
- [pingone_sign_on_policy_action](pingone_sign_on_policy_action.md)