package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFADevicePolicyResource{}
)

type PingOneMFADevicePolicyResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneMFADevicePolicyResource
func MFADevicePolicy(clientInfo *connector.PingOneClientInfo) *PingOneMFADevicePolicyResource {
	return &PingOneMFADevicePolicyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneMFADevicePolicyResource) ResourceType() string {
	return "pingone_mfa_device_policy"
}

func (r *PingOneMFADevicePolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportDevicePolicies()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneMFADevicePolicyResource) exportDevicePolicies() error {
	iter := r.clientInfo.ApiClient.MFAAPIClient.DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadDeviceAuthenticationPolicies", r.ResourceType())
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

		for _, devicePolicy := range embedded.GetDeviceAuthenticationPolicies() {
			devicePolicyId, devicePolicyIdOk := devicePolicy.GetIdOk()
			devicePolicyName, devicePolicyNameOk := devicePolicy.GetNameOk()

			if devicePolicyIdOk && devicePolicyNameOk {
				r.addImportBlock(*devicePolicyId, *devicePolicyName)
			}
		}
	}

	return nil
}

func (r *PingOneMFADevicePolicyResource) addImportBlock(devicePolicyId, devicePolicyName string) {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
