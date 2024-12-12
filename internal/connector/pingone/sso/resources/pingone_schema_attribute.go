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

func (r *PingOneSchemaAttributeResource) ResourceType() string {
	return "pingone_schema_attribute"
}

func (r *PingOneSchemaAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	schemaData, err := r.getSchemaData()
	if err != nil {
		return nil, err
	}

	for schemaId, schemaName := range *schemaData {
		schemaAttributeData, err := r.getSchemaAttributeData(schemaId)
		if err != nil {
			return nil, err
		}

		for schemaAttributeId, schemaAttributeName := range *schemaAttributeData {
			commentData := map[string]string{
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Resource Type":         r.ResourceType(),
				"Schema Attribute ID":   schemaAttributeId,
				"Schema Attribute Name": schemaAttributeName,
				"Schema ID":             schemaId,
				"Schema Name":           schemaName,
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", schemaName, schemaAttributeName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, schemaId, schemaAttributeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneSchemaAttributeResource) getSchemaData() (*map[string]string, error) {
	schemaData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.SchemasApi.ReadAllSchemas(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllSchemas", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, schema := range embedded.GetSchemas() {
			schemaId, schemaIdOk := schema.GetIdOk()
			schemaName, schemaNameOk := schema.GetNameOk()
			if schemaIdOk && schemaNameOk {
				schemaData[*schemaId] = *schemaName
			}
		}
	}

	return &schemaData, nil
}

func (r *PingOneSchemaAttributeResource) getSchemaAttributeData(schemaId string) (*map[string]string, error) {
	schemaAttributeData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.SchemasApi.ReadAllSchemaAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, schemaId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllSchemaAttributes", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, schemaAttribute := range embedded.GetAttributes() {
			schemaAttributeId, schemaAttributeIdOk := schemaAttribute.SchemaAttribute.GetIdOk()
			schemaAttributeName, schemaAttributeNameOk := schemaAttribute.SchemaAttribute.GetNameOk()
			if schemaAttributeIdOk && schemaAttributeNameOk {
				schemaAttributeData[*schemaAttributeId] = *schemaAttributeName
			}
		}
	}

	return &schemaAttributeData, nil
}
