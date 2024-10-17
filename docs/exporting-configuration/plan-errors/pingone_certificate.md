# Terraform Configuration Generation - PingOne Plan Errors (pingone_certificate)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/certificate#schema)

## Invalid combination of arguments - one of `pem_file,pkcs7_file_base64` must be specified

**Cause**: Certificates are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set either `pem_file` or `pkcs7_file_base64` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_certificate" "my_awesome_cert" {
  environment_id    = "942b****-****-****-****-********985c"
  pem_file          = null
  pkcs7_file_base64 = null
  usage_type        = "ENCRYPTION"
}
```

After manual modification (using PEM as an example, `pem_file` is defined):
```hcl
resource "pingone_certificate" "my_awesome_cert" {
  environment_id    = "942b****-****-****-****-********985c"
  pem_file          = file("../path/to/certificate.pem")
  usage_type        = "ENCRYPTION"
}
```


