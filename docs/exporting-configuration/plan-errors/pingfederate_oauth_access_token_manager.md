# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_oauth_access_token_manager)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/oauth_access_token_manager#schema)

## Missing Configuration for Required Attribute - Must set a configuration value for the configuration.tables[0].rows[0].sensitive_fields[Value({"name":"Key","value":<null>})].value attribute as the provider has marked it as required.

**Cause**: Symmetric key values are not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `configuration.tables[0].rows[0].sensitive_fields` field to include an object with `name`=`Key`, and `value` is the symmetric key, in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_oauth_access_token_manager" "my_symmetric_key_app" {
  # ... other configuration parameters

  configuration = {
    fields = [
      # ... other configuration parameters

      {
        name  = "Active Symmetric Encryption Key ID"
        value = "mykey"
      },
      {
        name  = "Active Symmetric Key ID"
        value = "mykey"
      },
    ]
    tables = [
      # ... other configuration parameters

      {
        name = "Symmetric Keys"
        rows = [
          {
            default_row = false
            fields = [
              {
                name  = "Encoding"
                value = "b64u"
              },
              {
                name  = "Key ID"
                value = "mykey"
              },
            ]
            sensitive_fields = [
              {
                name  = "Key"
                value = null # sensitive
              },
            ]
          },
        ]
      },
    ]
  }
}
```

After manual modification (The symmetric key is defined):
```hcl
resource "pingfederate_oauth_access_token_manager" "my_symmetric_key_app" {
  # ... other configuration parameters

  configuration = {
    fields = [
      # ... other configuration parameters

      {
        name  = "Active Symmetric Encryption Key ID"
        value = "mykey"
      },
      {
        name  = "Active Symmetric Key ID"
        value = "mykey"
      },
    ]
    tables = [
      # ... other configuration parameters

      {
        name = "Symmetric Keys"
        rows = [
          {
            default_row = false
            fields = [
              {
                name  = "Encoding"
                value = "b64u"
              },
              {
                name  = "Key ID"
                value = "mykey"
              },
            ]
            sensitive_fields = [
              {
                name  = "Key"
                value = var.my_symmetric_key_mykey
              },
            ]
          },
        ]
      },
    ]
  }
}
```

