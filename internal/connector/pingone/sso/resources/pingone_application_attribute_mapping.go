package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationAttributeMappingResource{}
)

type PingOneApplicationAttributeMappingResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneApplicationAttributeMappingResource
func ApplicationAttributeMapping(clientInfo *connector.PingOneClientInfo) *PingOneApplicationAttributeMappingResource {
	return &PingOneApplicationAttributeMappingResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneApplicationAttributeMappingResource) ResourceType() string {
	return "pingone_application_attribute_mapping"
}

func (r *PingOneApplicationAttributeMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportApplicationAttributeMappings()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneApplicationAttributeMappingResource) exportApplicationAttributeMappings() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplications", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, app := range embedded.GetApplications() {
			var (
				appId     *string
				appIdOk   bool
				appName   *string
				appNameOk bool
			)

			switch {
			case app.ApplicationOIDC != nil:
				appId, appIdOk = app.ApplicationOIDC.GetIdOk()
				appName, appNameOk = app.ApplicationOIDC.GetNameOk()
			case app.ApplicationSAML != nil:
				appId, appIdOk = app.ApplicationSAML.GetIdOk()
				appName, appNameOk = app.ApplicationSAML.GetNameOk()
			default:
				continue
			}

			if appIdOk && appNameOk {
				err := r.exportApplicationAttributeMappingsByApplication(*appId, *appName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneApplicationAttributeMappingResource) exportApplicationAttributeMappingsByApplication(appId, appName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplicationAttributeMappings", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, attributeMapping := range embedded.GetAttributes() {
			if attributeMapping.ApplicationAttributeMapping != nil {
				attributeMappingId, attributeMappingIdOk := attributeMapping.ApplicationAttributeMapping.GetIdOk()
				attributeMappingName, attributeMappingNameOk := attributeMapping.ApplicationAttributeMapping.GetNameOk()

				if attributeMappingIdOk && attributeMappingNameOk {
					r.addImportBlock(appId, appName, *attributeMappingId, *attributeMappingName)
				}
			}
		}
	}

	return nil
}

func (r *PingOneApplicationAttributeMappingResource) addImportBlock(appId, appName, attributeMappingId, attributeMappingName string) {
	commentData := map[string]string{
		"Application ID":         appId,
		"Application Name":       appName,
		"Attribute Mapping ID":   attributeMappingId,
		"Attribute Mapping Name": attributeMappingName,
		"Export Environment ID":  r.clientInfo.ExportEnvironmentID,
		"Resource Type":          r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", appName, attributeMappingName),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, attributeMappingId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
