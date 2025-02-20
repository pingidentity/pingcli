package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFADevicePolicyResource{}
)

type PingOneMFADevicePolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneMFADevicePolicyResource
func MFADevicePolicy(clientInfo *connector.PingOneClientInfo) *PingOneMFADevicePolicyResource {
	return &PingOneMFADevicePolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMFADevicePolicyResource) ResourceType() string {
	return "pingone_mfa_device_policy"
}

func (r *PingOneMFADevicePolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	deviceAuthPolicyData, err := r.getDeviceAuthPolicyData()
	if err != nil {
		return nil, err
	}

	for devicePolicyId, devicePolicyName := range deviceAuthPolicyData {
		commentData := map[string]string{
			"Export Environment ID":  r.clientInfo.ExportEnvironmentID,
			"MFA Device Policy ID":   devicePolicyId,
			"MFA Device Policy Name": devicePolicyName,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       devicePolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, devicePolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneMFADevicePolicyResource) getDeviceAuthPolicyData() (map[string]string, error) {
	deviceAuthPolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.MFAAPIClient.DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	deviceAuthPolicies, err := common.GetMfaAPIObjectsFromIterator[mfa.DeviceAuthenticationPolicy](iter, "ReadDeviceAuthenticationPolicies", "GetDeviceAuthenticationPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, devicePolicy := range deviceAuthPolicies {
		devicePolicyId, devicePolicyIdOk := devicePolicy.GetIdOk()
		devicePolicyName, devicePolicyNameOk := devicePolicy.GetNameOk()

		if devicePolicyIdOk && devicePolicyNameOk {
			deviceAuthPolicyData[*devicePolicyId] = *devicePolicyName
		}
	}

	return deviceAuthPolicyData, nil
}
