package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFAApplicationPushCredentialResource{}
)

type PingOneMFAApplicationPushCredentialResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneMFAApplicationPushCredentialResource
func MFAApplicationPushCredential(clientInfo *connector.PingOneClientInfo) *PingOneMFAApplicationPushCredentialResource {
	return &PingOneMFAApplicationPushCredentialResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMFAApplicationPushCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	nativeOIDCApplications, err := r.getNativeOIDCApplications()
	if err != nil {
		return nil, err
	}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}
	for _, nativeOIDCApplication := range nativeOIDCApplications {
		nativeOIDCApplicationId, nativeOIDCApplicationIdOk := nativeOIDCApplication.GetIdOk()
		nativeOIDCApplicationName, nativeOIDCApplicationNameOk := nativeOIDCApplication.GetNameOk()

		if nativeOIDCApplicationIdOk && nativeOIDCApplicationNameOk {
			mfaPushCredentialResponses, err := r.getMFAPushCredentials(nativeOIDCApplicationId)
			if err != nil {
				return nil, err
			}

			for _, mfaPushCredentialResponse := range mfaPushCredentialResponses {
				mfaPushCredentialResponseType, mfaPushCredentialResponseTypeOk := mfaPushCredentialResponse.GetTypeOk()
				mfaPushCredentialResponseId, mfaPushCredentialResponseIdOk := mfaPushCredentialResponse.GetIdOk()

				if mfaPushCredentialResponseTypeOk && mfaPushCredentialResponseIdOk {
					commentData := map[string]string{

						"Export Environment ID":                    r.clientInfo.ExportEnvironmentID,
						"MFA Push Credential ID":                   *mfaPushCredentialResponseId,
						"MFA Push Credential Type":                 string(*mfaPushCredentialResponseType),
						"Native Application (OpenID Connect) ID":   *nativeOIDCApplicationId,
						"Native Application (OpenID Connect) Name": *nativeOIDCApplicationName,
						"Resource Type":                            r.ResourceType(),
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *nativeOIDCApplicationName, *mfaPushCredentialResponseType),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *nativeOIDCApplicationId, *mfaPushCredentialResponseId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneMFAApplicationPushCredentialResource) ResourceType() string {
	return "pingone_mfa_application_push_credential"
}

func (r *PingOneMFAApplicationPushCredentialResource) getNativeOIDCApplications() (nativeOIDCApplications []*management.ApplicationOIDC, err error) {
	nativeOIDCApplications = []*management.ApplicationOIDC{}

	// Fetch all pingone_application resources that could have pingone_mfa_application_push_credentials
	applicationsPagedIterator := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	allEmbedded, err := common.GetAllManagementEmbedded(applicationsPagedIterator, "ReadAllApplications", "pingone_application")
	if err != nil {
		return nil, err
	}

	// pingone_mfa_application_push_credential are for Native applications only
	// Native application authenticate with OIDC only
	for _, embedded := range allEmbedded {
		for _, app := range embedded.GetApplications() {
			if app.ApplicationOIDC != nil {
				appType, appTypeOk := app.ApplicationOIDC.GetTypeOk()
				if appTypeOk && *appType == management.ENUMAPPLICATIONTYPE_NATIVE_APP {
					nativeOIDCApplications = append(nativeOIDCApplications, app.ApplicationOIDC)
				}
			}
		}
	}

	return nativeOIDCApplications, nil
}

func (r *PingOneMFAApplicationPushCredentialResource) getMFAPushCredentials(nativeOIDCApplicationId *string) (mfaPushCredentialResponses []mfa.MFAPushCredentialResponse, err error) {
	mfaPushCredentialResponses = []mfa.MFAPushCredentialResponse{}

	// Fetch all pingone_mfa_application_push_credentials resources for the given pingone_application
	mfaApplicationPushCredentialsPagedIterator := r.clientInfo.ApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *nativeOIDCApplicationId).Execute()

	allEmbedded, err := common.GetAllMFAEmbedded(mfaApplicationPushCredentialsPagedIterator, "ReadAllMFAPushCredentials", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, embedded := range allEmbedded {
		for _, mfaPushCredentialResponse := range embedded.GetPushCredentials() {
			mfaPushCredentialResponses = append(mfaPushCredentialResponses, mfaPushCredentialResponse)
		}
	}

	return mfaPushCredentialResponses, nil
}
