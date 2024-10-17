# Ping CLI - Exporting Platform Configuration

## Resolving Terraform Plan Errors

When generating Terraform HCL configuration, errors on `terraform plan` are expected. Reasons for plan errors include:

- Certain field values are not retrievable from the Ping system.  This might be because values are sensitive (secret) and are not retrievable to maintain tenant security.  In these cases, manual adjustment is needed to ensure these values are defined in generated HCL.
- Ambiguity in the retrieved configuration from the Ping system.  In these cases, the intention of the configuration cannot be accurately determined and requires manual correction.
- Limitations with Terraform's `terraform plan --generate-config-out` command action. Limitations are described in further detail on Terraform's developer documentation, [Generating Configuration](https://developer.hashicorp.com/terraform/language/import/generating-configuration)

The following documents describe the actions that must be taken, per provider, to resolve the various `terraform plan` errors following configuration generation.

- [PingFederate Terraform Provider](./plan-errors/pingfederate.md)
- [PingOne Terraform Provider](./plan-errors/pingone.md)

If you encounter an error that is not documented, please [raise a new issue](https://github.com/pingidentity/pingcli/issues/new?title=Undocumented%20Config%20Generation%20Error).
