// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	_ connector.ExportableResource = &PingOneMfaApplicationPushCredentialResource{}
)

type PingOneMfaApplicationPushCredentialResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneMfaApplicationPushCredentialResource
func MfaApplicationPushCredential(clientInfo *connector.ClientInfo) *PingOneMfaApplicationPushCredentialResource {
	return &PingOneMfaApplicationPushCredentialResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMfaApplicationPushCredentialResource) ResourceType() string {
	return "pingone_mfa_application_push_credential"
}

func (r *PingOneMfaApplicationPushCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for applicationId, applicationName := range applicationData {
		mfaApplicationPushCredentialData, err := r.getMfaApplicationPushCredentialData(applicationId)
		if err != nil {
			return nil, err
		}

		for mfaApplicationPushCredentialId, mfaApplicationPushCredentialType := range mfaApplicationPushCredentialData {
			commentData := map[string]string{
				"Application ID":                       applicationId,
				"Application Name":                     applicationName,
				"Mfa Application Push Credential ID":   mfaApplicationPushCredentialId,
				"Mfa Application Push Credential Type": mfaApplicationPushCredentialType,
				"Export Environment ID":                r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                        r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", applicationName, mfaApplicationPushCredentialType),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, applicationId, mfaApplicationPushCredentialId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneMfaApplicationPushCredentialResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		// MFa application push credentials are only for OIDC Native Apps
		if innerObj.ApplicationOIDC != nil {
			applicationId, applicationIdOk := innerObj.ApplicationOIDC.GetIdOk()
			applicationName, applicationNameOk := innerObj.ApplicationOIDC.GetNameOk()
			applicationType, applicationTypeOk := innerObj.ApplicationOIDC.GetTypeOk()

			if applicationIdOk && applicationNameOk && applicationTypeOk {
				if *applicationType == management.ENUMAPPLICATIONTYPE_NATIVE_APP {
					applicationData[*applicationId] = *applicationName
				}
			}
		}
	}

	return applicationData, nil
}

func (r *PingOneMfaApplicationPushCredentialResource) getMfaApplicationPushCredentialData(applicationId string) (map[string]string, error) {
	mfaApplicationPushCredentialData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, applicationId).Execute()
	apiObjs, err := pingone.GetMfaAPIObjectsFromIterator[mfa.MFAPushCredentialResponse](iter, "ReadAllMFAPushCredentials", "GetPushCredentials", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, mfaApplicationPushCredential := range apiObjs {
		mfaApplicationPushCredentialId, mfaApplicationPushCredentialIdOk := mfaApplicationPushCredential.GetIdOk()
		mfaApplicationPushCredentialType, mfaApplicationPushCredentialTypeOk := mfaApplicationPushCredential.GetTypeOk()

		if mfaApplicationPushCredentialIdOk && mfaApplicationPushCredentialTypeOk {
			mfaApplicationPushCredentialData[*mfaApplicationPushCredentialId] = string(*mfaApplicationPushCredentialType)
		}
	}

	return mfaApplicationPushCredentialData, nil
}
