package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFAApplicationPushCredentialResource{}
)

type PingOneMFAApplicationPushCredentialResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneMFAApplicationPushCredentialResource
func MFAApplicationPushCredential(clientInfo *connector.ClientInfo) *PingOneMFAApplicationPushCredentialResource {
	return &PingOneMFAApplicationPushCredentialResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMFAApplicationPushCredentialResource) ResourceType() string {
	return "pingone_mfa_application_push_credential"
}

func (r *PingOneMFAApplicationPushCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getOIDCApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range applicationData {
		pushCredData, err := r.getPushCredentialData(appId)
		if err != nil {
			return nil, err
		}

		for pushCredId, pushCredType := range pushCredData {
			commentData := map[string]string{
				"Export Environment ID":                r.clientInfo.PingOneExportEnvironmentID,
				"MFA Application Push Credential ID":   pushCredId,
				"MFA Application Push Credential Type": pushCredType,
				"Native OIDC Application ID":           appId,
				"Native OIDC Application Name":         appName,
				"Resource Type":                        r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, pushCredType),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, appId, pushCredId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneMFAApplicationPushCredentialResource) getOIDCApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	applications, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		// MFa application push credentials are only for OIDC Native Apps
		if application.ApplicationOIDC != nil {
			applicationId, applicationIdOk := application.ApplicationOIDC.GetIdOk()
			applicationName, applicationNameOk := application.ApplicationOIDC.GetNameOk()
			applicationType, applicationTypeOk := application.ApplicationOIDC.GetTypeOk()

			if applicationIdOk && applicationNameOk && applicationTypeOk {
				if *applicationType == management.ENUMAPPLICATIONTYPE_NATIVE_APP {
					applicationData[*applicationId] = *applicationName
				}
			}
		}
	}

	return applicationData, nil
}

func (r *PingOneMFAApplicationPushCredentialResource) getPushCredentialData(applicationId string) (map[string]string, error) {
	mfaPushCredentialData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, applicationId).Execute()
	mfaPushCredentials, err := pingone.GetMfaAPIObjectsFromIterator[mfa.MFAPushCredentialResponse](iter, "ReadAllMFAPushCredentials", "GetPushCredentials", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, mfaPushCredential := range mfaPushCredentials {
		mfaPushCredentialId, mfaPushCredentialIdOk := mfaPushCredential.GetIdOk()
		mfaPushCredentialType, mfaPushCredentialTypeOk := mfaPushCredential.GetTypeOk()

		if mfaPushCredentialIdOk && mfaPushCredentialTypeOk {
			mfaPushCredentialData[*mfaPushCredentialId] = string(*mfaPushCredentialType)
		}
	}

	return mfaPushCredentialData, nil
}
