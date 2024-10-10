package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeDecisionEndpointResource{}
)

type PingoneAuthorizeDecisionEndpointResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeDecisionEndpointResource
func AuthorizeDecisionEndpoint(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeDecisionEndpointResource {
	return &PingoneAuthorizeDecisionEndpointResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeDecisionEndpointResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.PolicyDecisionManagementApi.ReadAllDecisionEndpoints(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllDecisionEndpoints"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, decisionEndpoint := range embedded.GetDecisionEndpoints() {
		decisionEndpointName, decisionEndpointNameOk := decisionEndpoint.GetNameOk()
		decisionEndpointId, decisionEndpointIdOk := decisionEndpoint.GetIdOk()

		if decisionEndpointNameOk && decisionEndpointIdOk {
			commentData := map[string]string{
				"Resource Type":                    r.ResourceType(),
				"Authorize Decision Endpoint Name": *decisionEndpointName,
				"Export Environment ID":            r.clientInfo.ExportEnvironmentID,
				"Authorize Decision Endpoint ID":   *decisionEndpointId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *decisionEndpointName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *decisionEndpointId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeDecisionEndpointResource) ResourceType() string {
	return "pingone_authorize_decision_endpoint"
}
