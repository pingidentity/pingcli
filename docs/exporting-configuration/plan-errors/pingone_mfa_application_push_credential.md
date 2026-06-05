# Terraform Configuration Generation - PingOne Plan Errors (pingone_mfa_application_push_credential)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/mfa_application_push_credential#schema)

## Invalid Attribute Combination - No attribute specified when one (and only one) of [apns.<.fcm,apns.<.apns,apns.<.hms] is required

**Cause**: Push credential values are not exported from PingOne to maintain tenant security.

**Resolution**: Manual modification is required to set one of `apns`, `fcm`, or `hms` in the generated HCL.

**Example**:

Generated configuration:
```hcl
resource "pingone_mfa_application_push_credential" "my_awesome_push_credential" {
  apns           = null
  application_id = "71f7****-****-****-****-********dcd7"
  environment_id = "942b****-****-****-****-********985c"
  fcm            = null
  hms            = null
}
```

After manual modification (using APNS as an example, `apns.key`, `apns.team_id` and `apns.token_signing_key` are defined):
```hcl
resource "pingone_mfa_application_push_credential" "my_awesome_push_credential" {
  apns = {
    key               = var.apns_key
    team_id           = var.apns_team_id
    token_signing_key = var.apns_token_signing_key
  }
  application_id = "71f7****-****-****-****-********dcd7"
  environment_id = "942b****-****-****-****-********985c"
}
```
