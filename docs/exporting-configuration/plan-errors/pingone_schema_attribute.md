# Terraform Configuration Generation - PingOne Plan Errors (pingone_schema_attribute)

**Documentation**:
- [Terraform Registry](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/schema_attribute)

## Data Loss Protection - This field is immutable and cannot be changed once defined

**Cause**: Terraform is looking to make a replacement change to a schema attribute, which will cause data to be lost.  Data loss protections are invoked.

**Resolution**: Manual modification is required to remove the resources from the generated HCL, or ensure that state is synchronised with the target platform.
