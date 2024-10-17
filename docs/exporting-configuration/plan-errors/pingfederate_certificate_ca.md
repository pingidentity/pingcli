# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_certificate_ca)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/certificate_ca#schema)

## Invalid Attribute Value Length - Attribute file_data string length must be at least 1, got: 0

**Cause**: The CA file data is not exported.

**Resolution**: Manual modification is required to set the `file_data` field in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_certificate_ca" "my_awesome_certificate_ca" {
  ca_id           = "7zz3****************5fnja"
  crypto_provider = null
  file_data       = ""
}
```

After manual modification (`file_data` is defined):
```hcl
resource "pingfederate_certificate_ca" "my_awesome_certificate_ca" {
  ca_id           = "7zz3****************5fnja"
  crypto_provider = null
  file_data       = filebase64("my_ca.pem")
}
```



