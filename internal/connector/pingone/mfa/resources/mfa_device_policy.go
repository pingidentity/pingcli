// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMfaDevicePolicyResource{}
)

type PingOneMfaDevicePolicyResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneMfaDevicePolicyResource
func MfaDevicePolicy(clientInfo *connector.ClientInfo) *PingOneMfaDevicePolicyResource {
	return &PingOneMfaDevicePolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMfaDevicePolicyResource) ResourceType() string {
	return "pingone_mfa_device_policy"
}

func (r *PingOneMfaDevicePolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	mfaDevicePolicyData, err := r.getMfaDevicePolicyData()
	if err != nil {
		return nil, err
	}

	for mfaDevicePolicyId, mfaDevicePolicyName := range mfaDevicePolicyData {
		commentData := map[string]string{
			"Mfa Device Policy ID":   mfaDevicePolicyId,
			"Mfa Device Policy Name": mfaDevicePolicyName,
			"Export Environment ID":  r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       mfaDevicePolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, mfaDevicePolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneMfaDevicePolicyResource) getMfaDevicePolicyData() (map[string]string, error) {
	mfaDevicePolicyData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.MFAAPIClient.DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetMfaAPIObjectsFromIterator[mfa.DeviceAuthenticationPolicy](iter, "ReadDeviceAuthenticationPolicies", "GetDeviceAuthenticationPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, mfaDevicePolicy := range apiObjs {
		mfaDevicePolicyId, mfaDevicePolicyIdOk := mfaDevicePolicy.GetIdOk()
		mfaDevicePolicyName, mfaDevicePolicyNameOk := mfaDevicePolicy.GetNameOk()

		if mfaDevicePolicyIdOk && mfaDevicePolicyNameOk {
			mfaDevicePolicyData[*mfaDevicePolicyId] = *mfaDevicePolicyName
		}
	}

	return mfaDevicePolicyData, nil
}
