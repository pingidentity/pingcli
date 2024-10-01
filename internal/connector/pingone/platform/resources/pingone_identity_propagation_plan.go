package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneIdentityPropagationPlanResource{}
)

type PingoneIdentityPropagationPlanResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneIdentityPropagationPlanResource
func IdentityPropagationPlan(clientInfo *connector.PingOneClientInfo) *PingoneIdentityPropagationPlanResource {
	return &PingoneIdentityPropagationPlanResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneIdentityPropagationPlanResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.IdentityPropagationPlansApi.ReadAllPlans(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllPlans"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, identityPropagationPlan := range embedded.GetPlans() {
		identityPropagationPlanId, identityPropagationPlanIdOk := identityPropagationPlan.GetIdOk()
		identityPropagationPlanName, identityPropagationPlanNameOk := identityPropagationPlan.GetNameOk()

		if identityPropagationPlanIdOk && identityPropagationPlanNameOk {
			commentData := map[string]string{
				"Resource Type":                  r.ResourceType(),
				"Identity Propagation Plan Name": *identityPropagationPlanName,
				"Export Environment ID":          r.clientInfo.ExportEnvironmentID,
				"Identity Propagation Plan ID":   *identityPropagationPlanId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *identityPropagationPlanName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *identityPropagationPlanId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneIdentityPropagationPlanResource) ResourceType() string {
	return "pingone_identity_propagation_plan"
}
