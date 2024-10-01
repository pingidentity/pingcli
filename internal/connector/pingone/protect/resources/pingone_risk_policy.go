package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneRiskPolicyResource{}
)

type PingoneRiskPolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneRiskPolicyResource
func RiskPolicy(clientInfo *connector.PingOneClientInfo) *PingoneRiskPolicyResource {
	return &PingoneRiskPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneRiskPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.RiskAPIClient.RiskPoliciesApi.ReadRiskPolicySets(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadRiskPolicySets"

	embedded, err := common.GetProtectEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, riskPolicySet := range embedded.GetRiskPolicySets() {
		riskPolicySetName, riskPolicySetNameOk := riskPolicySet.GetNameOk()
		riskPolicySetId, riskPolicySetIdOk := riskPolicySet.GetIdOk()

		if riskPolicySetNameOk && riskPolicySetIdOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Risk Policy Name":      *riskPolicySetName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Risk Policy ID":        *riskPolicySetId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *riskPolicySetName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *riskPolicySetId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})

		}
	}

	return &importBlocks, nil
}

func (r *PingoneRiskPolicyResource) ResourceType() string {
	return "pingone_risk_policy"
}
