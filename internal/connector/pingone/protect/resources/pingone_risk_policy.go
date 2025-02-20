package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneRiskPolicyResource{}
)

type PingOneRiskPolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneRiskPolicyResource
func RiskPolicy(clientInfo *connector.PingOneClientInfo) *PingOneRiskPolicyResource {
	return &PingOneRiskPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneRiskPolicyResource) ResourceType() string {
	return "pingone_risk_policy"
}

func (r *PingOneRiskPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ristPolicySetData, err := getRiskPolicySetData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for riskPolicySetId, riskPolicySetName := range ristPolicySetData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
			"Risk Policy ID":        riskPolicySetId,
			"Risk Policy Name":      riskPolicySetName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       riskPolicySetName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, riskPolicySetId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func getRiskPolicySetData(clientInfo *connector.PingOneClientInfo, resourceType string) (map[string]string, error) {
	riskPolicySetData := make(map[string]string)

	iter := clientInfo.ApiClient.RiskAPIClient.RiskPoliciesApi.ReadRiskPolicySets(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()
	riskPolicySets, err := common.GetRiskAPIObjectsFromIterator[risk.RiskPolicySet](iter, "ReadRiskPolicySets", "GetRiskPolicySets", resourceType)
	if err != nil {
		return nil, err
	}

	for _, riskPolicySet := range riskPolicySets {
		riskPolicySetName, riskPolicySetNameOk := riskPolicySet.GetNameOk()
		riskPolicySetId, riskPolicySetIdOk := riskPolicySet.GetIdOk()

		if riskPolicySetIdOk && riskPolicySetNameOk {
			riskPolicySetData[*riskPolicySetId] = *riskPolicySetName
		}
	}

	return riskPolicySetData, nil
}
