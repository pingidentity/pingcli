# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_idp_adapter)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/idp_adapter#schema)

## Missing Configuration for Required Attribute - Must set a configuration value for the configuration.sensitive_fields[Value({"name":"API Key","value":<null>})].value attribute as the provider has marked it as required.

**Cause**: The DaVinci adapter's API key is not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `configuration.sensitive_fields` field to include an object with `name`=`API Key`, and `value` is the API key, in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_idp_adapter" "my_davinci_adapter" {
  # ... other configuration parameters

  adapter_id = "myDaVinciAdapter"
  
  configuration = {
    # ... other configuration parameters

    fields = [
       # ... other configuration parameters

      {
        name  = "API Request Timeout"
        value = jsonencode(5000)
      },
      {
        name  = "Additional Properties Attribute"
        value = "additionalProperties"
      },
    ]
    sensitive_fields = [
      {
        name  = "API Key"
        value = null # sensitive
      },
    ]
  }
}
```

After manual modification (The DaVinci API key is defined):
```hcl
resource "pingfederate_idp_adapter" "my_davinci_adapter" {
  # ... other configuration parameters

  adapter_id = "myDaVinciAdapter"
  
  configuration = {
    # ... other configuration parameters

    fields = [
       # ... other configuration parameters

      {
        name  = "API Request Timeout"
        value = jsonencode(5000)
      },
      {
        name  = "Additional Properties Attribute"
        value = "additionalProperties"
      },
    ]
    sensitive_fields = [
      {
        name  = "API Key"
        value = var.my_davinci_adapter_api_key
      },
    ]
  }
}
```


