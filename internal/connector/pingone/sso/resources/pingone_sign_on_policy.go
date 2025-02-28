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
	_ connector.ExportableResource = &PingOneSignOnPolicyResource{}
)

type PingOneSignOnPolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneSignOnPolicyResource
func SignOnPolicy(clientInfo *connector.PingOneClientInfo) *PingOneSignOnPolicyResource {
	return &PingOneSignOnPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneSignOnPolicyResource) ResourceType() string {
	return "pingone_sign_on_policy"
}

func (r *PingOneSignOnPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	signOnPolicyData, err := r.getSignOnPolicyData()
	if err != nil {
		return nil, err
	}

	for signOnPolicyId, signOnPolicyName := range signOnPolicyData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
			"Sign-On Policy ID":     signOnPolicyId,
			"Sign-On Policy Name":   signOnPolicyName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       signOnPolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, signOnPolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneSignOnPolicyResource) getSignOnPolicyData() (map[string]string, error) {
	signOnPolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadAllSignOnPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	signOnPolicies, err := pingone.GetManagementAPIObjectsFromIterator[management.SignOnPolicy](iter, "ReadAllSignOnPolicies", "GetSignOnPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, signOnPolicy := range signOnPolicies {
		signOnPolicyId, signOnPolicyIdOk := signOnPolicy.GetIdOk()
		signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

		if signOnPolicyIdOk && signOnPolicyNameOk {
			signOnPolicyData[*signOnPolicyId] = *signOnPolicyName
		}
	}

	return signOnPolicyData, nil
}
