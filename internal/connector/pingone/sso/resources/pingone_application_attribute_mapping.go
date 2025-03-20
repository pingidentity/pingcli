// Copyright © 2025 Ping Identity Corporation

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
	_ connector.ExportableResource = &PingOneApplicationAttributeMappingResource{}
)

type PingOneApplicationAttributeMappingResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneApplicationAttributeMappingResource
func ApplicationAttributeMapping(clientInfo *connector.ClientInfo) *PingOneApplicationAttributeMappingResource {
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

	for appId, appName := range applicationData {
		applicationAttributeMappingData, err := r.getApplicationAttributeMappingData(appId)
		if err != nil {
			return nil, err
		}

		for attributeMappingId, attributeMappingName := range applicationAttributeMappingData {
			commentData := map[string]string{
				"Application ID":         appId,
				"Application Name":       appName,
				"Attribute Mapping ID":   attributeMappingId,
				"Attribute Mapping Name": attributeMappingName,
				"Export Environment ID":  r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":          r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, attributeMappingName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, appId, attributeMappingId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationAttributeMappingResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	applications, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, app := range applications {
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

	return applicationData, nil
}

func (r *PingOneApplicationAttributeMappingResource) getApplicationAttributeMappingData(appId string) (map[string]string, error) {
	applicationAttributeMappingData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, appId).Execute()
	attributeMappingInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedAttributesInner](iter, "ReadAllApplicationAttributeMappings", "GetAttributes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, attributeMappingInner := range attributeMappingInners {
		if attributeMappingInner.ApplicationAttributeMapping != nil {
			attributeMappingId, attributeMappingIdOk := attributeMappingInner.ApplicationAttributeMapping.GetIdOk()
			attributeMappingName, attributeMappingNameOk := attributeMappingInner.ApplicationAttributeMapping.GetNameOk()

			if attributeMappingIdOk && attributeMappingNameOk {
				applicationAttributeMappingData[*attributeMappingId] = *attributeMappingName
			}
		}
	}

	return applicationAttributeMappingData, nil
}
