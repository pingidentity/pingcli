package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
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

	for schemaId, schemaName := range schemaData {
		schemaAttributeData, err := r.getSchemaAttributeData(schemaId)
		if err != nil {
			return nil, err
		}

		for schemaAttributeId, schemaAttributeName := range schemaAttributeData {
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

func (r *PingOneSchemaAttributeResource) getSchemaData() (map[string]string, error) {
	schemaData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.SchemasApi.ReadAllSchemas(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	schemas, err := pingone.GetManagementAPIObjectsFromIterator[management.Schema](iter, "ReadAllSchemas", "GetSchemas", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, schema := range schemas {
		schemaId, schemaIdOk := schema.GetIdOk()
		schemaName, schemaNameOk := schema.GetNameOk()
		if schemaIdOk && schemaNameOk {
			schemaData[*schemaId] = *schemaName
		}
	}

	return schemaData, nil
}

func (r *PingOneSchemaAttributeResource) getSchemaAttributeData(schemaId string) (map[string]string, error) {
	schemaAttributeData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.SchemasApi.ReadAllSchemaAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, schemaId).Execute()
	attributeInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedAttributesInner](iter, "ReadAllSchemaAttributes", "GetAttributes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, attributeInner := range attributeInners {
		if attributeInner.SchemaAttribute != nil {
			schemaAttributeId, schemaAttributeIdOk := attributeInner.SchemaAttribute.GetIdOk()
			schemaAttributeName, schemaAttributeNameOk := attributeInner.SchemaAttribute.GetNameOk()
			if schemaAttributeIdOk && schemaAttributeNameOk {
				schemaAttributeData[*schemaAttributeId] = *schemaAttributeName
			}
		}
	}

	return schemaAttributeData, nil
}
