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

type mfaDevicePolicyImportBlockData struct {
	ID   string
	Name string
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

func (r *PingOneMFADevicePolicyResource) deviceAuthenticationPolicies() (deviceAuthenticationPolicies []mfa.DeviceAuthenticationPolicy, err error) {
	deviceAuthenticationPolicies = []mfa.DeviceAuthenticationPolicy{}

	// Fetch all pingone_mfa_application_push_credentials resources for the given pingone_application
	entityArrayPagedIterator := r.clientInfo.ApiClient.MFAAPIClient.DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	allEmbedded, err := common.GetAllMFAEmbedded(entityArrayPagedIterator, "ReadDeviceAuthenticationPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, embedded := range allEmbedded {
		for _, deviceAuthenticationPolicy := range embedded.GetDeviceAuthenticationPolicies() {
			deviceAuthenticationPolicies = append(deviceAuthenticationPolicies, deviceAuthenticationPolicy)
		}
	}

	return deviceAuthenticationPolicies, nil
}

func (r *PingOneMFADevicePolicyResource) importBlockData() (importBlockData []mfaDevicePolicyImportBlockData, err error) {
	importBlockData = []mfaDevicePolicyImportBlockData{}

	deviceAuthenticationPolicies, err := r.deviceAuthenticationPolicies()
	if err != nil {
		return nil, err
	}

	for _, deviceAuthenticationPolicy := range deviceAuthenticationPolicies {
		deviceAuthenticationPolicyName, deviceAuthenticationPolicyNameOk := deviceAuthenticationPolicy.GetNameOk()
		deviceAuthenticationPolicyId, deviceAuthenticationPolicyIdOk := deviceAuthenticationPolicy.GetIdOk()

		if deviceAuthenticationPolicyNameOk && deviceAuthenticationPolicyIdOk {
			importBlockData = append(importBlockData, mfaDevicePolicyImportBlockData{
				ID:   *deviceAuthenticationPolicyId,
				Name: *deviceAuthenticationPolicyName,
			})
		}
	}

	return importBlockData, nil
}

func (r *PingOneMFADevicePolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
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
			"Export Environment ID":  r.clientInfo.ExportEnvironmentID,
			"MFA Device Policy ID":   data.ID,
			"MFA Device Policy Name": data.Name,
			"Resource Type":          r.ResourceType(),
		}

		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       data.Name,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, data.ID),
			CommentInformation: common.GenerateCommentInformation(commentData),
		})
	}

	return &importBlocks, nil
}
