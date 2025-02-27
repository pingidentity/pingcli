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
	_ connector.ExportableResource = &PingOneApplicationResourceResource{}
)

type PingOneApplicationResourceResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneApplicationResourceResource
func ApplicationResource(clientInfo *connector.PingOneClientInfo) *PingOneApplicationResourceResource {
	return &PingOneApplicationResourceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationResourceResource) ResourceType() string {
	return "pingone_application_resource"
}

func (r *PingOneApplicationResourceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range applicationData {
		applicationResourceData, err := r.getApplicationResourceData(appId)
		if err != nil {
			return nil, err
		}

		for resourceId, resourceName := range applicationResourceData {
			commentData := map[string]string{
				"Application ID":            appId,
				"Application Name":          appName,
				"Application Resource ID":   resourceId,
				"Application Resource Name": resourceName,
				"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
				"Resource Type":             r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, resourceName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, resourceId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationResourceResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
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
		default:
			continue
		}

		if appIdOk && appNameOk {
			applicationData[*appId] = *appName
		}
	}

	return applicationData, nil
}

func (r *PingOneApplicationResourceResource) getApplicationResourceData(appId string) (map[string]string, error) {
	applicationResourceData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationResourcesApi.ReadAllApplicationResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()
	applicationResources, err := pingone.GetManagementAPIObjectsFromIterator[management.ResourceApplicationResource](iter, "ReadAllApplicationResources", "GetAttributes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationResource := range applicationResources {
		resourceId, resourceIdOk := applicationResource.GetIdOk()
		resourceName, resourceNameOk := applicationResource.GetNameOk()

		if resourceIdOk && resourceNameOk {
			applicationResourceData[*resourceId] = *resourceName
		}
	}

	return applicationResourceData, nil
}
