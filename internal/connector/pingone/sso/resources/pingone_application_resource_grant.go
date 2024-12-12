package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationResourceGrantResource{}
)

type PingOneApplicationResourceGrantResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneApplicationResourceGrantResource
func ApplicationResourceGrant(clientInfo *connector.PingOneClientInfo) *PingOneApplicationResourceGrantResource {
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

	for appId, appName := range *applicationData {
		applicationGrantData, err := r.getApplicationGrantData(appId)
		if err != nil {
			return nil, err
		}

		for grantId, grantResourceId := range *applicationGrantData {
			resourceName, err := r.getGrantResourceName(grantResourceId)
			if err != nil {
				return nil, err
			}

			commentData := map[string]string{
				"Application ID":                appId,
				"Application Name":              appName,
				"Application Resource Grant ID": grantId,
				"Application Resource Name":     *resourceName,
				"Export Environment ID":         r.clientInfo.ExportEnvironmentID,
				"Resource Type":                 r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, *resourceName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, grantId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationResourceGrantResource) getApplicationData() (*map[string]string, error) {
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
	}

	return &applicationData, nil
}

func (r *PingOneApplicationResourceGrantResource) getApplicationGrantData(appId string) (*map[string]string, error) {
	applicationGrantData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.ReadAllApplicationGrants(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplicationGrants", r.ResourceType())
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

		for _, grant := range embedded.GetGrants() {
			grantId, grantIdOk := grant.GetIdOk()
			grantResource, grantResourceOk := grant.GetResourceOk()

			if grantIdOk && grantResourceOk {
				grantResourceId, grantResourceIdOk := grantResource.GetIdOk()

				if grantResourceIdOk {
					applicationGrantData[*grantId] = *grantResourceId
				}
			}
		}
	}

	return &applicationGrantData, nil
}

func (r *PingOneApplicationResourceGrantResource) getGrantResourceName(grantResourceId string) (*string, error) {
	resource, response, err := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, grantResourceId).Execute()
	err = common.HandleClientResponse(response, err, "ReadOneResource", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if resource != nil {
		resourceName, resourceNameOk := resource.GetNameOk()
		if resourceNameOk {
			return resourceName, nil
		}
	}

	return nil, fmt.Errorf("Unable to get resource name for grant resource ID: %s", grantResourceId)
}
