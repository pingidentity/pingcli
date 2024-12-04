package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFAApplicationPushCredentialResource{}
)

type PingOneMFAApplicationPushCredentialResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneMFAApplicationPushCredentialResource
func MFAApplicationPushCredential(clientInfo *connector.PingOneClientInfo) *PingOneMFAApplicationPushCredentialResource {
	return &PingOneMFAApplicationPushCredentialResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneMFAApplicationPushCredentialResource) ResourceType() string {
	return "pingone_mfa_application_push_credential"
}

func (r *PingOneMFAApplicationPushCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportPushCreds()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneMFAApplicationPushCredentialResource) exportPushCreds() error {
	// Fetch all pingone_application resources that could have pingone_mfa_application_push_credentials
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplications", "pingone_application")
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, app := range embedded.GetApplications() {
			// MFa application push credentials are only for OIDC Native Apps
			if app.ApplicationOIDC != nil {
				appId, appIdOk := app.ApplicationOIDC.GetIdOk()
				appName, appNameOk := app.ApplicationOIDC.GetNameOk()
				appType, appTypeOk := app.ApplicationOIDC.GetTypeOk()

				if appIdOk && appNameOk && appTypeOk {
					if *appType == management.ENUMAPPLICATIONTYPE_NATIVE_APP {
						err := r.exportPushCredsByApp(*appId, *appName)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneMFAApplicationPushCredentialResource) exportPushCredsByApp(appId, appName string) error {
	iter := r.clientInfo.ApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllMFAPushCredentials", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, pushCred := range embedded.GetPushCredentials() {
			pushCredId, pushCredIdOk := pushCred.GetIdOk()
			pushCredType, pushCredTypeOk := pushCred.GetTypeOk()

			if pushCredIdOk && pushCredTypeOk {
				r.addImportBlock(appId, appName, *pushCredId, string(*pushCredType))
			}
		}
	}

	return nil
}

func (r *PingOneMFAApplicationPushCredentialResource) addImportBlock(appId, appName, pushCredId, pushCredType string) {
	commentData := map[string]string{
		"Export Environment ID":                r.clientInfo.ExportEnvironmentID,
		"MFA Application Push Credential ID":   pushCredId,
		"MFA Application Push Credential Type": pushCredType,
		"Native OIDC Application ID":           appId,
		"Native OIDC Application Name":         appName,
		"Resource Type":                        r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", appName, pushCredType),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, pushCredId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
