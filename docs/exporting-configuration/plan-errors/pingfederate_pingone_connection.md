# Ping CLI - Exporting Platform Configuration - PingFederate Plan Errors (pingfederate_pingone_connection)

**Documentation**:
- [Terraform Registry - PingFederate pingone_connection](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/pingone_connection#schema)
- [Terraform Registry - PingOne pingone_gateway_credential](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/gateway_credential)

## Missing Configuration for Required Attribute - Must set a configuration value for the credential attribute as the provider has marked it as required

**Cause**: The PingOne credential is not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `credential` field in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_pingone_connection" "my_pingone_environment" {
  # ... other configuration parameters

  credential = null # sensitive
  name       = "My PingOne Environment"
}
```

After manual modification (`credential` is defined):
```hcl
resource "pingfederate_pingone_connection" "my_pingone_environment" {
  # ... other configuration parameters

  credential = var.pingone_credential # see pingone_gateway_credential in the PingOne Terraform provider
  name       = "My PingOne Environment"
}
```

