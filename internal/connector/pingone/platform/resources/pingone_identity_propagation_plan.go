package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneIdentityPropagationPlanResource{}
)

type PingOneIdentityPropagationPlanResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneIdentityPropagationPlanResource
func IdentityPropagationPlan(clientInfo *connector.PingOneClientInfo) *PingOneIdentityPropagationPlanResource {
	return &PingOneIdentityPropagationPlanResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneIdentityPropagationPlanResource) ResourceType() string {
	return "pingone_identity_propagation_plan"
}

func (r *PingOneIdentityPropagationPlanResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportIdentityPropagationPlans()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneIdentityPropagationPlanResource) exportIdentityPropagationPlans() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.IdentityPropagationPlansApi.ReadAllPlans(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllPlans", r.ResourceType())
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

		for _, identityPropagationPlan := range embedded.GetPlans() {
			identityPropagationPlanId, identityPropagationPlanIdOk := identityPropagationPlan.GetIdOk()
			identityPropagationPlanName, identityPropagationPlanNameOk := identityPropagationPlan.GetNameOk()

			if identityPropagationPlanIdOk && identityPropagationPlanNameOk {
				r.addImportBlock(*identityPropagationPlanId, *identityPropagationPlanName)
			}
		}
	}

	return nil
}

func (r *PingOneIdentityPropagationPlanResource) addImportBlock(identityPropagationPlanId, identityPropagationPlanName string) {
	commentData := map[string]string{
		"Export Environment ID":          r.clientInfo.ExportEnvironmentID,
		"Identity Propagation Plan ID":   identityPropagationPlanId,
		"Identity Propagation Plan Name": identityPropagationPlanName,
		"Resource Type":                  r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       identityPropagationPlanName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, identityPropagationPlanId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
