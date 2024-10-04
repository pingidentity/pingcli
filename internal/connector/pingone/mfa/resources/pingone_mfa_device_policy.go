package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneMFADevicePolicyResource{}
)

type PingoneMFADevicePolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneMFADevicePolicyResource
func MFADevicePolicy(clientInfo *connector.PingOneClientInfo) *PingoneMFADevicePolicyResource {
	return &PingoneMFADevicePolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneMFADevicePolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.MFAAPIClient.DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadDeviceAuthenticationPolicies"

	embedded, err := common.GetMFAEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, deviceAuthenticationPolicy := range embedded.GetDeviceAuthenticationPolicies() {
		deviceAuthenticationPolicyName, deviceAuthenticationPolicyNameOk := deviceAuthenticationPolicy.GetNameOk()
		deviceAuthenticationPolicyId, deviceAuthenticationPolicyIdOk := deviceAuthenticationPolicy.GetIdOk()

		if deviceAuthenticationPolicyNameOk && deviceAuthenticationPolicyIdOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"MFA Policy Name":       *deviceAuthenticationPolicyName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"MFA Policy ID":         *deviceAuthenticationPolicyId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *deviceAuthenticationPolicyName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *deviceAuthenticationPolicyId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneMFADevicePolicyResource) ResourceType() string {
	return "pingone_mfa_device_policy"
}
