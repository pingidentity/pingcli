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
	_ connector.ExportableResource = &PingOneIdentityPropagationPlanResource{}
)

type PingOneIdentityPropagationPlanResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneIdentityPropagationPlanResource
func IdentityPropagationPlan(clientInfo *connector.ClientInfo) *PingOneIdentityPropagationPlanResource {
	return &PingOneIdentityPropagationPlanResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneIdentityPropagationPlanResource) ResourceType() string {
	return "pingone_identity_propagation_plan"
}

func (r *PingOneIdentityPropagationPlanResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	planData, err := r.getIdentityPropagationPlanData()
	if err != nil {
		return nil, err
	}

	for planId, planName := range planData {
		commentData := map[string]string{
			"Export Environment ID":          r.clientInfo.ExportEnvironmentID,
			"Identity Propagation Plan ID":   planId,
			"Identity Propagation Plan Name": planName,
			"Resource Type":                  r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       planName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, planId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneIdentityPropagationPlanResource) getIdentityPropagationPlanData() (map[string]string, error) {
	identityPropagationPlanData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.IdentityPropagationPlansApi.ReadAllPlans(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	identityPropagationPlans, err := pingone.GetManagementAPIObjectsFromIterator[management.IdentityPropagationPlan](iter, "ReadAllPlans", "GetPlans", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, identityPropagationPlan := range identityPropagationPlans {
		identityPropagationPlanId, identityPropagationPlanIdOk := identityPropagationPlan.GetIdOk()
		identityPropagationPlanName, identityPropagationPlanNameOk := identityPropagationPlan.GetNameOk()

		if identityPropagationPlanIdOk && identityPropagationPlanNameOk {
			identityPropagationPlanData[*identityPropagationPlanId] = *identityPropagationPlanName
		}
	}

	return identityPropagationPlanData, nil
}
