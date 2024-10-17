# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_password_credential_validator)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/password_credential_validator#schema)

## Must set a configuration value for the configuration.tables[0].rows[*].sensitive_fields[Value({"name":"Confirm Password","value":<null>})].value attribute as the provider has marked it as required

**Cause**: Simple password credential validator password values are not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `configuration.tables[0].rows[*].sensitive_fields` field to include an object with `name`=`Confirm Password`, and `value` is the simple password to use for that user, in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_password_credential_validator" "simple_username_password_credential_validator" {
  # ... other configuration parameters

  configuration = {
    # ... other configuration parameters
    
    tables = [
      {
        name = "Users"
        rows = [
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = null # sensitive
              },
              {
                name  = "Password"
                value = null # sensitive
              },
            ]
          },
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example2"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = null # sensitive
              },
              {
                name  = "Password"
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

After manual modification (A password for each user is defined):
```hcl
resource "pingfederate_password_credential_validator" "simple_username_password_credential_validator" {
  # ... other configuration parameters

  configuration = {
    # ... other configuration parameters
    
    tables = [
      {
        name = "Users"
        rows = [
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = var.simple_pcv_example_password
              },
              {
                name  = "Password"
                value = var.simple_pcv_example_password
              },
            ]
          },
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example2"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = var.simple_pcv_example2_password
              },
              {
                name  = "Password"
                value = var.simple_pcv_example2_password
              },
            ]
          },
        ]
      },
    ]
  }
}
```

## Must set a configuration value for the configuration.tables[0].rows[*].sensitive_fields[Value({"name":"Password","value":<null>})].value attribute as the provider has marked it as required

**Cause**: Simple password credential validator password values are not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `configuration.tables[0].rows[*].sensitive_fields` field to include an object with `name`=`Password`, and `value` is the simple password to use for that user, in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_password_credential_validator" "simple_username_password_credential_validator" {
  # ... other configuration parameters

  configuration = {
    # ... other configuration parameters
    
    tables = [
      {
        name = "Users"
        rows = [
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = null # sensitive
              },
              {
                name  = "Password"
                value = null # sensitive
              },
            ]
          },
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example2"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = null # sensitive
              },
              {
                name  = "Password"
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

After manual modification (A password for each user is defined):
```hcl
resource "pingfederate_password_credential_validator" "simple_username_password_credential_validator" {
  # ... other configuration parameters

  configuration = {
    # ... other configuration parameters
    
    tables = [
      {
        name = "Users"
        rows = [
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = var.simple_pcv_example_password
              },
              {
                name  = "Password"
                value = var.simple_pcv_example_password
              },
            ]
          },
          {
            default_row = false
            fields = [
              {
                name  = "Relax Password Requirements"
                value = jsonencode(false)
              },
              {
                name  = "Username"
                value = "example2"
              },
            ]
            sensitive_fields = [
              {
                name  = "Confirm Password"
                value = var.simple_pcv_example2_password
              },
              {
                name  = "Password"
                value = var.simple_pcv_example2_password
              },
            ]
          },
        ]
      },
    ]
  }
}
```

