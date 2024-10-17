# Terraform Configuration Generation - PingOne Plan Errors (pingone_branding_theme)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/branding_theme#schema)

##  Invalid Attribute Combination - 2 attributes specified when one (and only one) of [background_color.<.background_color,background_color.<.use_default_background,background_color.<.background_image] is required

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to ensure only one of `background_color`, `use_default_background` or `background_image` is defined.

**Example**:

Generated configuration:
```hcl
resource "pingone_branding_theme" "my_awesome_theme" {
  # ... other configuration parameters
  
  background_color = null
  background_image = {
    href = "https://uploads.pingone.eu/environments/942b****-****-****-****-********985c/images/image.png"
    id   = "d4a1****-****-****-****-********ba9d"
  }
  use_default_background = false
}
```

After manual modification (`background_color` and `use_default_background` are removed):
```hcl
resource "pingone_branding_theme" "my_awesome_theme" {
  # ... other configuration parameters
  
  background_image = {
    href = "https://uploads.pingone.eu/environments/942b****-****-****-****-********985c/images/image.png"
    id   = "d4a1****-****-****-****-********ba9d"
  }
}
```

