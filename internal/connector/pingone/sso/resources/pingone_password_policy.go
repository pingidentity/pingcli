package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOnePasswordPolicyResource{}
)

type PingOnePasswordPolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOnePasswordPolicyResource
func PasswordPolicy(clientInfo *connector.PingOneClientInfo) *PingOnePasswordPolicyResource {
	return &PingOnePasswordPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOnePasswordPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.PasswordPoliciesApi.ReadAllPasswordPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllPasswordPolicies"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, passwordPolicy := range embedded.GetPasswordPolicies() {
		passwordPolicyId, passwordPolicyIdOk := passwordPolicy.GetIdOk()
		passwordPolicyName, passwordPolicyNameOk := passwordPolicy.GetNameOk()

		if passwordPolicyIdOk && passwordPolicyNameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Password Policy Name":  *passwordPolicyName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Password Policy ID":    *passwordPolicyId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *passwordPolicyName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *passwordPolicyId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOnePasswordPolicyResource) ResourceType() string {
	return "pingone_password_policy"
}
