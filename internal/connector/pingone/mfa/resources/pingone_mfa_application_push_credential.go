package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneMFAApplicationPushCredentialResource{}
)

type PingoneMFAApplicationPushCredentialResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneMFAApplicationPushCredentialResource
func MFAApplicationPushCredential(clientInfo *connector.PingOneClientInfo) *PingoneMFAApplicationPushCredentialResource {
	return &PingoneMFAApplicationPushCredentialResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneMFAApplicationPushCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	// Fetch all pingone_application resources that could have pingone_mfa_application_push_credentials
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
		}

		if appIdOk && appNameOk {
			// Fetch all pingone_mfa_application_push_credentials for each application
			apiExecuteFunc := r.clientInfo.ApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiFunctionName := "ReadAllMFAPushCredentials"

			embedded, err := common.GetMFAEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, mfaPushCredentialResponse := range embedded.GetPushCredentials() {
				mfaPushCredentialResponseType, mfaPushCredentialResponseTypeOk := mfaPushCredentialResponse.GetTypeOk()
				mfaPushCredentialResponseId, mfaPushCredentialResponseIdOk := mfaPushCredentialResponse.GetIdOk()

				if mfaPushCredentialResponseTypeOk && mfaPushCredentialResponseIdOk {
					commentData := map[string]string{
						"Resource Type":            r.ResourceType(),
						"Application Name":         *appName,
						"MFA Push Credential Type": string(*mfaPushCredentialResponseType),
						"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
						"Application ID":           *appId,
						"MFA Push Credential ID":   *mfaPushCredentialResponseId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *appName, *mfaPushCredentialResponseType),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *mfaPushCredentialResponseId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneMFAApplicationPushCredentialResource) ResourceType() string {
	return "pingone_mfa_application_push_credential"
}
