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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneApplicationAttributeMappingResource
func ApplicationAttributeMapping(clientInfo *connector.PingOneClientInfo) *PingOneApplicationAttributeMappingResource {
	return &PingOneApplicationAttributeMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationAttributeMappingResource) ResourceType() string {
	return "pingone_application_attribute_mapping"
}

func (r *PingOneApplicationAttributeMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range *applicationData {
		applicationAttributeMappingData, err := r.getApplicationAttributeMappingData(appId)
		if err != nil {
			return nil, err
		}

		for attributeMappingId, attributeMappingName := range *applicationAttributeMappingData {
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

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationAttributeMappingResource) getApplicationData() (*map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplications", r.ResourceType())
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
				applicationData[*appId] = *appName
			}
		}
	}

	return &applicationData, nil
}

func (r *PingOneApplicationAttributeMappingResource) getApplicationAttributeMappingData(appId string) (*map[string]string, error) {
	applicationAttributeMappingData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplicationAttributeMappings", r.ResourceType())
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

		for _, attributeMapping := range embedded.GetAttributes() {
			if attributeMapping.ApplicationAttributeMapping != nil {
				attributeMappingId, attributeMappingIdOk := attributeMapping.ApplicationAttributeMapping.GetIdOk()
				attributeMappingName, attributeMappingNameOk := attributeMapping.ApplicationAttributeMapping.GetNameOk()

				if attributeMappingIdOk && attributeMappingNameOk {
					applicationAttributeMappingData[*attributeMappingId] = *attributeMappingName
				}
			}
		}
	}

	return &applicationAttributeMappingData, nil
}
