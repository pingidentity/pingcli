# Terraform Configuration Generation - PingFederate Plan Errors (pingfederate_oauth_client)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingfederate/latest/docs/resources/oauth_client#schema)

## Invalid attribute configuration - client_auth.secret must be defined when client_auth.type is configured to "SECRET"

**Cause**: The OAuth client secret is not exported from PingFederate to maintain tenant security.

**Resolution**: Manual modification is required to set the `client_auth.secret` field in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingfederate_oauth_client" "openid_connect_basic_client_profile" {
  # ... other configuration parameters

  client_auth = {
    # ... other configuration parameters
    
    secret = null # sensitive
    type   = "SECRET"
  }
  client_id = "ac_oic_client"
}
```

After manual modification (`client_auth.secret` is defined):
```hcl
resource "pingfederate_oauth_client" "openid_connect_basic_client_profile" {
  # ... other configuration parameters

  client_auth = {
    # ... other configuration parameters
    
    secret = var.my_oidc_client_secret
    type   = "SECRET"
  }
  client_id = "ac_oic_client"
}
```

