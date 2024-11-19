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

type mfaApplicationPushCredentialImportBlockData struct {
	PushCredentialID          string
	PushCredentialType        string
	NativeOIDCApplicationID   string
	NativeOIDCApplicationName string
}

// Utility method for creating a PingOneMFAApplicationPushCredentialResource
func MFAApplicationPushCredential(clientInfo *connector.PingOneClientInfo) *PingOneMFAApplicationPushCredentialResource {
	return &PingOneMFAApplicationPushCredentialResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMFAApplicationPushCredentialResource) ResourceType() string {
	return "pingone_mfa_application_push_credential"
}

func (r *PingOneMFAApplicationPushCredentialResource) nativeOIDCApplications() (nativeOIDCApplications []*management.ApplicationOIDC, err error) {
	nativeOIDCApplications = []*management.ApplicationOIDC{}

	// Fetch all pingone_application resources that could have pingone_mfa_application_push_credentials
	entityArrayPagedIterator := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	allEmbedded, err := common.GetAllManagementEmbedded(entityArrayPagedIterator, "ReadAllApplications", "pingone_application")
	if err != nil {
		return nil, err
	}

	// pingone_mfa_application_push_credential are for Native applications only
	// Native application authenticate with OIDC only
	for _, embedded := range allEmbedded {
		for _, readOneApplication200Response := range embedded.GetApplications() {
			if readOneApplication200Response.ApplicationOIDC != nil {
				appType, appTypeOk := readOneApplication200Response.ApplicationOIDC.GetTypeOk()
				if appTypeOk && *appType == management.ENUMAPPLICATIONTYPE_NATIVE_APP {
					nativeOIDCApplications = append(nativeOIDCApplications, readOneApplication200Response.ApplicationOIDC)
				}
			}
		}
	}

	return nativeOIDCApplications, nil
}

func (r *PingOneMFAApplicationPushCredentialResource) pushCredentialResponses(nativeOIDCApplicationId *string) (mfaPushCredentialResponses []mfa.MFAPushCredentialResponse, err error) {
	mfaPushCredentialResponses = []mfa.MFAPushCredentialResponse{}

	// Fetch all pingone_mfa_application_push_credentials resources for the given pingone_application
	entityArrayPagedIterator := r.clientInfo.ApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *nativeOIDCApplicationId).Execute()

	allEmbedded, err := common.GetAllMFAEmbedded(entityArrayPagedIterator, "ReadAllMFAPushCredentials", r.ResourceType())
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

func (r *PingOneMFAApplicationPushCredentialResource) importBlockData() (importBlockData []mfaApplicationPushCredentialImportBlockData, err error) {
	importBlockData = []mfaApplicationPushCredentialImportBlockData{}

	nativeOIDCApplications, err := r.nativeOIDCApplications()
	if err != nil {
		return nil, err
	}

	for _, nativeOIDCApplication := range nativeOIDCApplications {
		nativeOIDCApplicationId, nativeOIDCApplicationIdOk := nativeOIDCApplication.GetIdOk()
		nativeOIDCApplicationName, nativeOIDCApplicationNameOk := nativeOIDCApplication.GetNameOk()

		if nativeOIDCApplicationIdOk && nativeOIDCApplicationNameOk {
			mfaPushCredentialResponses, err := r.pushCredentialResponses(nativeOIDCApplicationId)
			if err != nil {
				return nil, err
			}

			for _, mfaPushCredentialResponse := range mfaPushCredentialResponses {
				mfaPushCredentialResponseType, mfaPushCredentialResponseTypeOk := mfaPushCredentialResponse.GetTypeOk()
				mfaPushCredentialResponseId, mfaPushCredentialResponseIdOk := mfaPushCredentialResponse.GetIdOk()
				if mfaPushCredentialResponseTypeOk && mfaPushCredentialResponseIdOk {
					importBlockData = append(importBlockData, mfaApplicationPushCredentialImportBlockData{
						PushCredentialID:          *mfaPushCredentialResponseId,
						PushCredentialType:        string(*mfaPushCredentialResponseType),
						NativeOIDCApplicationID:   *nativeOIDCApplicationId,
						NativeOIDCApplicationName: *nativeOIDCApplicationName,
					})
				}
			}
		}
	}

	return importBlockData, nil
}

func (r *PingOneMFAApplicationPushCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlockData, err := r.importBlockData()
	if err != nil {
		return nil, err
	}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, data := range importBlockData {
		commentData := map[string]string{
			"Export Environment ID":                    r.clientInfo.ExportEnvironmentID,
			"MFA Application Push Credential ID":       data.PushCredentialID,
			"MFA Application Push Credential Type":     data.PushCredentialType,
			"Native Application (OpenID Connect) ID":   data.NativeOIDCApplicationID,
			"Native Application (OpenID Connect) Name": data.NativeOIDCApplicationName,
			"Resource Type":                            r.ResourceType(),
		}

		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", data.NativeOIDCApplicationName, data.PushCredentialType),
			ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, data.NativeOIDCApplicationID, data.PushCredentialID),
			CommentInformation: common.GenerateCommentInformation(commentData),
		})
	}

	return &importBlocks, nil
}
