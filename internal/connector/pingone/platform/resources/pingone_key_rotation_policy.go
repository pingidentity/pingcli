package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
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

func (r *PingOneKeyRotationPolicyResource) ResourceType() string {
	return "pingone_key_rotation_policy"
}

func (r *PingOneKeyRotationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	keyRotationPolicyData, err := r.getKeyRotationPolicyData()
	if err != nil {
		return nil, err
	}

	for keyRotationPolicyId, keyRotationPolicyName := range keyRotationPolicyData {
		commentData := map[string]string{
			"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
			"Key Rotation Policy ID":   keyRotationPolicyId,
			"Key Rotation Policy Name": keyRotationPolicyName,
			"Resource Type":            r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       keyRotationPolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, keyRotationPolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneKeyRotationPolicyResource) getKeyRotationPolicyData() (map[string]string, error) {
	keyRotationPolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.KeyRotationPoliciesApi.GetKeyRotationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	keyRotationPolicies, err := pingone.GetManagementAPIObjectsFromIterator[management.KeyRotationPolicy](iter, "GetKeyRotationPolicies", "GetKeyRotationPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, keyRotationPolicy := range keyRotationPolicies {
		keyRotationPolicyId, keyRotationPolicyIdOk := keyRotationPolicy.GetIdOk()
		keyRotationPolicyName, keyRotationPolicyNameOk := keyRotationPolicy.GetNameOk()

		if keyRotationPolicyIdOk && keyRotationPolicyNameOk {
			keyRotationPolicyData[*keyRotationPolicyId] = *keyRotationPolicyName
		}
	}

	return keyRotationPolicyData, nil
}
