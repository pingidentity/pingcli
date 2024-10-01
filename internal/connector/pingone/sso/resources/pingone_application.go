package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationResource{}
)

type PingoneApplicationResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneApplicationResource
func Application(clientInfo *connector.PingOneClientInfo) *PingoneApplicationResource {
	return &PingoneApplicationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationResource) ExportAll() (*[]connector.ImportBlock, error) {
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
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Application Name":      *appName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Application ID":        *appId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *appName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *appId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationResource) ResourceType() string {
	return "pingone_application"
}
