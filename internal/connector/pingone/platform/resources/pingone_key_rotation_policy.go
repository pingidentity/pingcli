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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneKeyRotationPolicyResource
func KeyRotationPolicy(clientInfo *connector.PingOneClientInfo) *PingOneKeyRotationPolicyResource {
	return &PingOneKeyRotationPolicyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneKeyRotationPolicyResource) ResourceType() string {
	return "pingone_key_rotation_policy"
}

func (r *PingOneKeyRotationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportKeyRotationPolicies()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneKeyRotationPolicyResource) exportKeyRotationPolicies() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.KeyRotationPoliciesApi.GetKeyRotationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "GetKeyRotationPolicies", r.ResourceType())
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

		for _, keyRotationPolicy := range embedded.GetKeyRotationPolicies() {
			keyRotationPolicyId, keyRotationPolicyIdOk := keyRotationPolicy.GetIdOk()
			keyRotationPolicyName, keyRotationPolicyNameOk := keyRotationPolicy.GetNameOk()

			if keyRotationPolicyIdOk && keyRotationPolicyNameOk {
				r.addImportBlock(*keyRotationPolicyId, *keyRotationPolicyName)
			}
		}
	}

	return nil
}

func (r *PingOneKeyRotationPolicyResource) addImportBlock(keyRotationPolicyId, keyRotationPolicyName string) {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
