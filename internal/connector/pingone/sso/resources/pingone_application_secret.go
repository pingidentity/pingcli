package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationSecretResource{}
)

type PingoneApplicationSecretResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneApplicationSecretResource
func ApplicationSecret(clientInfo *connector.PingOneClientInfo) *PingoneApplicationSecretResource {
	return &PingoneApplicationSecretResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationSecretResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllApplications"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
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
			// The platform enforces that worker apps cannot read their own secret
			// Make sure we can read the secret before adding it to the import blocks
			_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationSecretApi.ReadApplicationSecret(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute()

			// If the appId is the same as the worker ID, make sure the API response is a 403 and ignore the error
			if *appId == *r.clientInfo.ApiClientId {
				if response.StatusCode == 403 {
					continue
				} else {
					return nil, fmt.Errorf("ReadApplicationSecret: Expected response code 403 - worker apps cannot read their own secret, actual response code: %d", response.StatusCode)
				}
			}

			// Use output package to warn the user of any errors or non-200 response codes
			// Expected behavior in this case is to skip the resource, and continue exporting the other resources
			defer response.Body.Close()

			if err != nil {
				l.Warn().Err(err).Msgf("Failed to read secret for application %s. %s Response Code: %s\nResponse Body: %s", *appName, apiFunctionName, response.Status, response.Body)
				continue
			}

			if response.StatusCode >= 300 {
				l.Warn().Msgf("Failed to read secret for application %s. %s Response Code: %s\nResponse Body: %s", *appName, apiFunctionName, response.Status, response.Body)
				continue
			}

			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Application Name":      *appName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Application ID":        *appId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_secret", *appName),
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *appId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationSecretResource) ResourceType() string {
	return "pingone_application_secret"
}
