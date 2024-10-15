package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneSchemaAttributeResource{}
)

type PingOneSchemaAttributeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneSchemaAttributeResource
func SchemaAttribute(clientInfo *connector.PingOneClientInfo) *PingOneSchemaAttributeResource {
	return &PingOneSchemaAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneSchemaAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteSchemaFunc := r.clientInfo.ApiClient.ManagementAPIClient.SchemasApi.ReadAllSchemas(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiSchemaFunctionName := "ReadAllSchemas"

	embedded, err := common.GetManagementEmbedded(apiExecuteSchemaFunc, apiSchemaFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, schema := range embedded.GetSchemas() {
		schemaId, schemaIdOk := schema.GetIdOk()
		schemaName, schemaNameOk := schema.GetNameOk()
		if schemaIdOk && schemaNameOk {
			apiExecuteSchemaAttributeFunc := r.clientInfo.ApiClient.ManagementAPIClient.SchemasApi.ReadAllSchemaAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *schemaId).Execute
			apiSchemaAttributeFunctionName := "ReadAllSchemaAttributes"

			schemaEmbedded, err := common.GetManagementEmbedded(apiExecuteSchemaAttributeFunc, apiSchemaAttributeFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, schemaAttribute := range schemaEmbedded.GetAttributes() {
				schemaAttributeId, schemaAttributeIdOk := schemaAttribute.SchemaAttribute.GetIdOk()
				schemaAttributeName, schemaAttributeNameOk := schemaAttribute.SchemaAttribute.GetNameOk()
				if schemaAttributeIdOk && schemaAttributeNameOk {
					commentData := map[string]string{
						"Resource Type":         r.ResourceType(),
						"Schema Name":           *schemaName,
						"Schema Attribute Name": *schemaAttributeName,
						"Export Environment ID": r.clientInfo.ExportEnvironmentID,
						"Schema ID":             *schemaId,
						"Schema Attribute ID":   *schemaAttributeId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *schemaName, *schemaAttributeName),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *schemaId, *schemaAttributeId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneSchemaAttributeResource) ResourceType() string {
	return "pingone_schema_attribute"
}
