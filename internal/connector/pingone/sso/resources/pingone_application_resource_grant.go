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
	_ connector.ExportableResource = &PingOneApplicationResourceGrantResource{}
)

type PingOneApplicationResourceGrantResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneApplicationResourceGrantResource
func ApplicationResourceGrant(clientInfo *connector.ClientInfo) *PingOneApplicationResourceGrantResource {
	return &PingOneApplicationResourceGrantResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationResourceGrantResource) ResourceType() string {
	return "pingone_application_resource_grant"
}

func (r *PingOneApplicationResourceGrantResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range applicationData {
		applicationGrantData, err := r.getApplicationGrantData(appId)
		if err != nil {
			return nil, err
		}

		for grantId, grantResourceId := range applicationGrantData {
			resourceName, resourceNameOk, err := r.getGrantResourceName(grantResourceId)
			if err != nil {
				return nil, err
			}
			if !resourceNameOk {
				continue
			}

			commentData := map[string]string{
				"Application ID":                appId,
				"Application Name":              appName,
				"Application Resource Grant ID": grantId,
				"Application Resource Name":     resourceName,
				"Export Environment ID":         r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                 r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, resourceName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, appId, grantId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationResourceGrantResource) getApplicationData() (map[string]string, error) {
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
		case app.ApplicationPingOnePortal != nil:
			appId, appIdOk = app.ApplicationPingOnePortal.GetIdOk()
			appName, appNameOk = app.ApplicationPingOnePortal.GetNameOk()
		case app.ApplicationPingOneSelfService != nil:
			appId, appIdOk = app.ApplicationPingOneSelfService.GetIdOk()
			appName, appNameOk = app.ApplicationPingOneSelfService.GetNameOk()
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			applicationData[*appId] = *appName
		}
	}

	return applicationData, nil
}

func (r *PingOneApplicationResourceGrantResource) getApplicationGrantData(appId string) (map[string]string, error) {
	applicationGrantData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.ReadAllApplicationGrants(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, appId).Execute()
	applicationGrants, err := pingone.GetManagementAPIObjectsFromIterator[management.ApplicationResourceGrant](iter, "ReadAllApplicationGrants", "GetGrants", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, grant := range applicationGrants {
		grantId, grantIdOk := grant.GetIdOk()
		grantResource, grantResourceOk := grant.GetResourceOk()

		if grantIdOk && grantResourceOk {
			grantResourceId, grantResourceIdOk := grantResource.GetIdOk()

			if grantResourceIdOk {
				applicationGrantData[*grantId] = *grantResourceId
			}
		}
	}

	return applicationGrantData, nil
}

func (r *PingOneApplicationResourceGrantResource) getGrantResourceName(grantResourceId string) (string, bool, error) {
	resource, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, grantResourceId).Execute()
	ok, err := common.HandleClientResponse(response, err, "ReadOneResource", r.ResourceType())
	if err != nil {
		return "", false, err
	}
	if !ok {
		return "", false, nil
	}

	if resource != nil {
		resourceName, resourceNameOk := resource.GetNameOk()
		if resourceNameOk {
			return *resourceName, true, nil
		}
	}

	return "", false, fmt.Errorf("unable to get resource name for grant resource ID: %s", grantResourceId)
}
