# Terraform Configuration Generation - PingFederate Plan Errors

The following sections describe the actions that must be taken, per resource, to resolve `terraform plan` errors following configuration generation.

If you encounter an error that is not documented, please [raise a new issue](https://github.com/pingidentity/pingcli/issues/new?title=Undocumented%20PingFederate%20Config%20Generation%20Error).

## General (Any Resource)

### Reference to undeclared resource - A managed resource "[any]" "[any]" has not been declared in the root module

**Cause**: Terraform configuration has been generated with syntax errors.  This is an issue with the Terraform CLI.

**Resolution**: Upgrade the Terraform CLI to the latest version available and re-generate the HCL configuration.

## Resource Plan Errors

- [pingfederate_certificate_ca](pingfederate_certificate_ca.md)
- [pingfederate_data_store](pingfederate_data_store.md)
- [pingfederate_idp_adapter](pingfederate_idp_adapter.md)
- [pingfederate_kerberos_realm](pingfederate_kerberos_realm.md)
- [pingfederate_oauth_access_token_manager](pingfederate_oauth_access_token_manager.md)
- [pingfederate_oauth_client](pingfederate_oauth_client.md)
- [pingfederate_password_credential_validator](pingfederate_password_credential_validator.md)
- [pingfederate_pingone_connection](pingfederate_pingone_connection.md)
