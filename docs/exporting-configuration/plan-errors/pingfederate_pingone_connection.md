# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_pingone_connection)

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

After manual modification, using a variable (`credential` is defined):
```hcl
resource "pingfederate_pingone_connection" "my_pingone_environment" {
  # ... other configuration parameters

  credential = var.pingone_credential
  name       = "My PingOne Environment"
}
```

After manual modification, using the PingOne Terraform provider (`credential` is defined):
```hcl
resource "pingone_gateway" "my_awesome_pingfederate_gateway" {
  environment_id = pingone_environment.my_environment.id
  name           = "Advanced Services SSO"
  enabled        = true

  type = "PING_FEDERATE"
}

resource "pingone_gateway_credential" "foo" {
  environment_id = pingone_environment.my_environment.id
  gateway_id     = pingone_gateway.my_awesome_pingfederate_gateway.id
}

resource "pingfederate_pingone_connection" "my_pingone_environment" {
  # ... other configuration parameters

  credential = pingone_gateway_credential.foo.credential
  name       = "My PingOne Environment"
}
```

