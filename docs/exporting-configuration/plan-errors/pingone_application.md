# Terraform Configuration Generation - PingOne Plan Errors (pingone_application)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/application#nestedatt--saml_options)

## Invalid Attribute Value Match - Attribute saml_options.type value must be one of: ["WEB_APP" "CUSTOM_APP"], got: "TEMPLATE_APP"

**Cause**: Template applications are not supported in the PingOne provider version being used.

**Resolution**: Upgrade the PingOne Terraform provider version.  Further details can be found at https://github.com/pingidentity/terraform-provider-pingone/issues/841

