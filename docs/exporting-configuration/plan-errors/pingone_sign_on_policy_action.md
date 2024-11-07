# Terraform Configuration Generation - PingOne Plan Errors (pingone_sign_on_policy_action)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/sign_on_policy_action)

## Conflicting configuration arguments - "conditions.0.anonymous_network_detected": conflicts with [identifier_first|login]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `conditions.0.anonymous_network_detected` value from the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    anonymous_network_detected      = false
    last_sign_on_older_than_seconds = 604800
  }

  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

After manual modification (`conditions.anonymous_network_detected` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds = 604800
  }
  
  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

## Conflicting configuration arguments - "conditions.0.anonymous_network_detected_allowed_cidr": conflicts with [identifier_first|login]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `conditions.0.anonymous_network_detected_allowed_cidr` value from the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    anonymous_network_detected_allowed_cidr = []
    last_sign_on_older_than_seconds         = 604800
  }

  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

After manual modification (`conditions.anonymous_network_detected_allowed_cidr` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds = 604800
  }
  
  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

## Conflicting configuration arguments - "conditions.0.geovelocity_anomaly_detected": conflicts with [identifier_first|login]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `conditions.0.geovelocity_anomaly_detected` value from the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    geovelocity_anomaly_detected    = false
    last_sign_on_older_than_seconds = 604800
  }

  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

After manual modification (`conditions.geovelocity_anomaly_detected` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds = 604800
  }
  
  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

## Conflicting configuration arguments - "conditions.0.ip_out_of_range_cidr": conflicts with [identifier_first|login]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `conditions.0.ip_out_of_range_cidr` value from the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    ip_out_of_range_cidr            = []
    last_sign_on_older_than_seconds = 604800
  }

  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

After manual modification (`conditions.ip_out_of_range_cidr` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds = 604800
  }
  
  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

## Conflicting configuration arguments - "conditions.0.ip_reputation_high_risk": conflicts with [identifier_first|login]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `conditions.0.ip_reputation_high_risk` value from the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    ip_reputation_high_risk         = false
    last_sign_on_older_than_seconds = 604800
  }

  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

After manual modification (`conditions.ip_reputation_high_risk` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds = 604800
  }
  
  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

## Conflicting configuration arguments - "conditions.0.last_sign_on_older_than_seconds_mfa": conflicts with [identifier_first|login]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `conditions.0.last_sign_on_older_than_seconds_mfa` value from the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds     = 604800
    last_sign_on_older_than_seconds_mfa = 0
  }

  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

After manual modification (`conditions.last_sign_on_older_than_seconds_mfa` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds = 604800
  }
  
  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

## Conflicting configuration arguments - "conditions.0.last_sign_on_older_than_seconds": conflicts with conditions.0.last_sign_on_older_than_seconds_mfa

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to ensure only one of `conditions.0.last_sign_on_older_than_seconds` or `conditions.0.last_sign_on_older_than_seconds_mfa` is set in the generated configuration.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds     = 604800
    last_sign_on_older_than_seconds_mfa = 0
  }

  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

After manual modification (`conditions.last_sign_on_older_than_seconds_mfa` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_identifier_first_action" {
  # ... other configuration parameters

  conditions {
    # ... other configuration parameters

    last_sign_on_older_than_seconds = 604800
  }
  
  identifier_first {
    # ... other configuration parameters

    recovery_enabled = true
    discovery_rule {
      attribute_contains_text = "@pingidentity.com"
      identity_provider_id    = "ad3a****-****-****-****-********ef83"
    }
  }
}
```

## Conflicting configuration arguments - "enforce_lockout_for_identity_providers": conflicts with [mfa|progressive_profiling]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `enforce_lockout_for_identity_providers` value from the generated HCL.

**Example**:

```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters

  enforce_lockout_for_identity_providers = false
  
  progressive_profiling {
    # ... other configuration parameters

    prompt_text = "For the best experience, we need a couple things from you."
  }
}
```

After manual modification (`enforce_lockout_for_identity_providers` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters

  progressive_profiling {
    # ... other configuration parameters

    prompt_text = "For the best experience, we need a couple things from you."
  }
}
```

## Conflicting configuration arguments - "registration_confirm_user_attributes": conflicts with [mfa|progressive_profiling]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `registration_confirm_user_attributes` value from the generated HCL.

**Example**:

```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters

  registration_confirm_user_attributes = false
  
  progressive_profiling {
    # ... other configuration parameters

    prompt_text = "For the best experience, we need a couple things from you."
  }
}
```

After manual modification (`registration_confirm_user_attributes` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters

  progressive_profiling {
    # ... other configuration parameters

    prompt_text = "For the best experience, we need a couple things from you."
  }
}
```

## Conflicting configuration arguments - "social_provider_ids": conflicts with [mfa|progressive_profiling]

**Cause**: Due to a [Terraform configuration generation limitation](https://developer.hashicorp.com/terraform/language/import/generating-configuration#conflicting-resource-arguments), conflicting parameters are included in the generated HCL.

**Resolution**: Manual modification is required to remove the `social_provider_ids` value from the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters

  social_provider_ids = []
  
  progressive_profiling {
    # ... other configuration parameters

    prompt_text = "For the best experience, we need a couple things from you."
  }
}
```

After manual modification (`social_provider_ids` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters

  progressive_profiling {
    # ... other configuration parameters

    prompt_text = "For the best experience, we need a couple things from you."
  }
}
```

## expected last_sign_on_older_than_seconds_mfa to be at least (1), got 0

**Cause**: The `last_sign_on_older_than_seconds_mfa` value is not set in PingOne, and has been exported incorrectly as `0`.

**Resolution**: Manual modification is required to remove the `last_sign_on_older_than_seconds_mfa` value from the generated HCL, or define a new value greater than `0`.

**Example**:

Generated configuration:
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters
  
  conditions {
    # ... other configuration parameters

    anonymous_network_detected              = true
    geovelocity_anomaly_detected            = true
    ip_reputation_high_risk                 = true
    last_sign_on_older_than_seconds_mfa     = 0
  }
  mfa {
    # ... other configuration parameters

    device_sign_on_policy_id = "7266****-****-****-****-********a5a9"
    no_device_mode           = "BLOCK"
  }
}
```

After manual modification (`last_sign_on_older_than_seconds_mfa` is removed):
```hcl
resource "pingone_sign_on_policy_action" "my_awesome_progressive_profiling_action" {
  # ... other configuration parameters
  
  conditions {
    # ... other configuration parameters

    anonymous_network_detected              = true
    geovelocity_anomaly_detected            = true
    ip_reputation_high_risk                 = true
  }
  mfa {
    # ... other configuration parameters
    
    device_sign_on_policy_id = "7266****-****-****-****-********a5a9"
    no_device_mode           = "BLOCK"
  }
}
```
