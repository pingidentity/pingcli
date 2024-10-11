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

func (r *PingOneApplicationAttributeMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteApplicationsFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiApplicationFunctionName := "ReadAllApplications"

	embedded, err := common.GetManagementEmbedded(apiExecuteApplicationsFunc, apiApplicationFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

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
			apiExecuteAttributeMappingFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiAttributeMappingFunctionName := "ReadAllApplicationAttributeMappings"

			attributeMappingsEmbedded, err := common.GetManagementEmbedded(apiExecuteAttributeMappingFunc, apiAttributeMappingFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, attributeMapping := range attributeMappingsEmbedded.GetAttributes() {
				if attributeMapping.ApplicationAttributeMapping == nil {
					continue
				}

				attributeMappingId, attributeMappingIdOk := attributeMapping.ApplicationAttributeMapping.GetIdOk()
				attributeMappingName, attributeMappingNameOk := attributeMapping.ApplicationAttributeMapping.GetNameOk()

				if attributeMappingIdOk && attributeMappingNameOk {
					commentData := map[string]string{
						"Resource Type":          r.ResourceType(),
						"Application Name":       *appName,
						"Attribute Mapping Name": *attributeMappingName,
						"Export Environment ID":  r.clientInfo.ExportEnvironmentID,
						"Application ID":         *appId,
						"Attribute Mapping ID":   *attributeMappingId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *appName, *attributeMappingName),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *attributeMappingId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationAttributeMappingResource) ResourceType() string {
	return "pingone_application_attribute_mapping"
}
