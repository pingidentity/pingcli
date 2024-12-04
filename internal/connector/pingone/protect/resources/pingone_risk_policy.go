package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneRiskPolicyResource{}
)

type PingOneRiskPolicyResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneRiskPolicyResource
func RiskPolicy(clientInfo *connector.PingOneClientInfo) *PingOneRiskPolicyResource {
	return &PingOneRiskPolicyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneRiskPolicyResource) ResourceType() string {
	return "pingone_risk_policy"
}

func (r *PingOneRiskPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportPolicies()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneRiskPolicyResource) exportPolicies() error {
	iter := r.clientInfo.ApiClient.RiskAPIClient.RiskPoliciesApi.ReadRiskPolicySets(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadRiskPolicySets", r.ResourceType())
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

		for _, riskPolicySet := range embedded.GetRiskPolicySets() {
			riskPolicySetName, riskPolicySetNameOk := riskPolicySet.GetNameOk()
			riskPolicySetId, riskPolicySetIdOk := riskPolicySet.GetIdOk()

			if riskPolicySetIdOk && riskPolicySetNameOk {
				r.addImportBlock(*riskPolicySetId, *riskPolicySetName)
			}
		}
	}

	return nil
}

func (r *PingOneRiskPolicyResource) addImportBlock(riskPolicySetId, riskPolicySetName string) {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
