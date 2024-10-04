package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationResourceGrantResource{}
)

type PingoneApplicationResourceGrantResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneApplicationResourceGrantResource
func ApplicationResourceGrant(clientInfo *connector.PingOneClientInfo) *PingoneApplicationResourceGrantResource {
	return &PingoneApplicationResourceGrantResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationResourceGrantResource) ExportAll() (*[]connector.ImportBlock, error) {
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
			apiExecutePoliciesFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.ReadAllApplicationGrants(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiApplicationGrantFunctionName := "ReadAllApplicationGrants"

			applicationEmbedded, err := common.GetManagementEmbedded(apiExecutePoliciesFunc, apiApplicationGrantFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, grant := range applicationEmbedded.GetGrants() {
				grantId, grantIdOk := grant.GetIdOk()
				grantResource, grantResourceOk := grant.GetResourceOk()

				var (
					grantResourceId   *string
					grantResourceIdOk bool
				)

				if grantResourceOk {
					grantResourceId, grantResourceIdOk = grantResource.GetIdOk()
				}

				if grantIdOk && grantResourceOk && grantResourceIdOk {
					resource, response, err := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadOneResource(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *grantResourceId).Execute()
					err = common.HandleClientResponse(response, err, "ReadOneResource", r.ResourceType())
					if err != nil {
						return nil, err
					}

					if resource != nil {
						resourceName, resourceNameOk := resource.GetNameOk()
						if resourceNameOk {
							commentData := map[string]string{
								"Resource Type":         r.ResourceType(),
								"Application Name":      *appName,
								"Resource Name":         *resourceName,
								"Export Environment ID": r.clientInfo.ExportEnvironmentID,
								"Application ID":        *appId,
								"Resource Grant ID":     *grantId,
							}

							importBlocks = append(importBlocks, connector.ImportBlock{
								ResourceType:       r.ResourceType(),
								ResourceName:       fmt.Sprintf("%s_%s", *appName, *resourceName),
								ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *grantId),
								CommentInformation: common.GenerateCommentInformation(commentData),
							})
						}
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationResourceGrantResource) ResourceType() string {
	return "pingone_application_resource_grant"
}
