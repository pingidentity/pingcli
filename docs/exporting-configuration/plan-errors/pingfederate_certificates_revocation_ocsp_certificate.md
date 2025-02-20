# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_certificates_revocation_ocsp_certificate)

**Documentation**:
- [Terraform Registry - PingFederate certificates_revocation_ocsp_certificate](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/certificates_revocation_ocsp_certificate#schema)

## Missing Configuration for Required Attribute - Must set a configuration value for the file_data attribute as the provider has marked it as required.

**Cause**: The certificate file data is not exported from the PingFederate API.

**Resolution**: Manual modification is required to set the `file_data` field in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_certificates_revocation_ocsp_certificate" "my_certificate_revocation_ocsp_certificate" {
  # ... other configuration parameters

  file_data = null
}
```

After manual modification (`file_data` is defined):
```hcl
resource "pingfederate_certificates_revocation_ocsp_certificate" "my_certificate_revocation_ocsp_certificate" {
  # ... other configuration parameters

  file_data = filebase64("my_certifcate.pem")
}
```
