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

func (r *PingOnePasswordPolicyResource) ResourceType() string {
	return "pingone_password_policy"
}

func (r *PingOnePasswordPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	passwordPolicyData, err := r.getPasswordPolicyData()
	if err != nil {
		return nil, err
	}

	for passwordPolicyId, passwordPolicyName := range passwordPolicyData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Password Policy ID":    passwordPolicyId,
			"Password Policy Name":  passwordPolicyName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       passwordPolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, passwordPolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOnePasswordPolicyResource) getPasswordPolicyData() (map[string]string, error) {
	passwordPolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.PasswordPoliciesApi.ReadAllPasswordPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	passwordPolicies, err := pingone.GetManagementAPIObjectsFromIterator[management.PasswordPolicy](iter, "ReadAllPasswordPolicies", "GetPasswordPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, passwordPolicy := range passwordPolicies {
		passwordPolicyId, passwordPolicyIdOk := passwordPolicy.GetIdOk()
		passwordPolicyName, passwordPolicyNameOk := passwordPolicy.GetNameOk()

		if passwordPolicyIdOk && passwordPolicyNameOk {
			passwordPolicyData[*passwordPolicyId] = *passwordPolicyName
		}
	}

	return passwordPolicyData, nil
}
