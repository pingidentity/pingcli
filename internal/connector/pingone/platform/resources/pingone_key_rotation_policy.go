package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneKeyRotationPolicyResource{}
)

type PingOneKeyRotationPolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneKeyRotationPolicyResource
func KeyRotationPolicy(clientInfo *connector.PingOneClientInfo) *PingOneKeyRotationPolicyResource {
	return &PingOneKeyRotationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneKeyRotationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.KeyRotationPoliciesApi.GetKeyRotationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "GetKeyRotationPolicies"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, keyRotationPolicy := range embedded.GetKeyRotationPolicies() {
		keyRotationPolicyId, keyRotationPolicyIdOk := keyRotationPolicy.GetIdOk()
		keyRotationPolicyName, keyRotationPolicyNameOk := keyRotationPolicy.GetNameOk()

		if keyRotationPolicyIdOk && keyRotationPolicyNameOk {
			commentData := map[string]string{
				"Resource Type":            r.ResourceType(),
				"Key Rotation Policy Name": *keyRotationPolicyName,
				"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
				"Key Rotation Policy ID":   *keyRotationPolicyId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *keyRotationPolicyName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *keyRotationPolicyId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneKeyRotationPolicyResource) ResourceType() string {
	return "pingone_key_rotation_policy"
}
