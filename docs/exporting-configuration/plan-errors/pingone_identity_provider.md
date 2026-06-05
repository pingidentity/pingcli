# Terraform Configuration Generation - PingOne Plan Errors (pingone_identity_provider)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/identity_provider#schema)

## Missing Configuration for Required Attribute - Must set a configuration value for the amazon.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `amazon.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "amazon" {
  # ... other configuration parameters

  amazon = {
    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`amazon.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "amazon" {
  # ... other configuration parameters

  amazon = {
    client_id     = "********"
    client_secret = var.identity_provider_amazon_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the apple.client_secret_signing_key attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `apple.client_secret_signing_key` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "apple" {
  # ... other configuration parameters

  apple = {
    client_id                 = "********"
    client_secret_signing_key = null # sensitive
    key_id                    = "********"
    team_id                   = "********"
  }
}
```

After manual modification (`apple.client_secret_signing_key` is defined):
```hcl
resource "pingone_identity_provider" "apple" {
  # ... other configuration parameters

  apple = {
    client_id                 = "********"
    client_secret_signing_key = var.identity_provider_apple_client_secret
    key_id                    = "********"
    team_id                   = "********"
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the facebook.app_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `facebook.app_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "facebook" {
  # ... other configuration parameters

  facebook = {
    app_id     = "********"
    app_secret = null # sensitive
  }
}
```

After manual modification (`facebook.app_secret` is defined):
```hcl
resource "pingone_identity_provider" "facebook" {
  # ... other configuration parameters

  facebook = {
    app_id     = "********"
    app_secret = var.identity_provider_facebook_app_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the github.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `github.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "github" {
  # ... other configuration parameters

  github = {
    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`github.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "github" {
  # ... other configuration parameters

  github = {
    client_id     = "********"
    client_secret = var.identity_provider_github_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the google.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `google.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "google" {
  # ... other configuration parameters

  google = {
    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`google.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "google" {
  # ... other configuration parameters

  google = {
    client_id     = "********"
    client_secret = var.identity_provider_google_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the linkedin.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `linkedin.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "linkedin" {
  # ... other configuration parameters

  linkedin = {
    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`linkedin.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "linkedin" {
  # ... other configuration parameters

  linkedin = {
    client_id     = "********"
    client_secret = var.identity_provider_linkedin_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the microsoft.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `microsoft.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "microsoft" {
  # ... other configuration parameters

  microsoft = {
    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`microsoft.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "microsoft" {
  # ... other configuration parameters

  microsoft = {
    client_id     = "********"
    client_secret = var.identity_provider_microsoft_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the openid_connect.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `openid_connect.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "openid_connect" {
  # ... other configuration parameters

  openid_connect = {
    # ... other configuration parameters

    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`openid_connect.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "openid_connect" {
  # ... other configuration parameters

  openid_connect = {
    # ... other configuration parameters

    client_id     = "********"
    client_secret = var.identity_provider_openid_connect_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the paypal.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `paypal.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "paypal" {
  # ... other configuration parameters

  paypal = {
    # ... other configuration parameters

    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`paypal.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "paypal" {
  # ... other configuration parameters

  paypal = {
    # ... other configuration parameters

    client_id     = "********"
    client_secret = var.identity_provider_paypal_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the twitter.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `twitter.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "twitter" {
  # ... other configuration parameters

  twitter = {
    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`twitter.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "twitter" {
  # ... other configuration parameters

  twitter = {
    client_id     = "********"
    client_secret = var.identity_provider_twitter_client_secret
  }
}
```

## Missing Configuration for Required Attribute - Must set a configuration value for the yahoo.client_secret attribute as the provider has marked it as required.

**Cause**: Client secrets for external identity providers are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set `yahoo.client_secret` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_identity_provider" "yahoo" {
  # ... other configuration parameters

  yahoo = {
    client_id     = "********"
    client_secret = null # sensitive
  }
}
```

After manual modification (`yahoo.client_secret` is defined):
```hcl
resource "pingone_identity_provider" "yahoo" {
  # ... other configuration parameters

  yahoo = {
    client_id     = "********"
    client_secret = var.identity_provider_yahoo_client_secret
  }
}
```
