package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeDecisionEndpointResource{}
)

type PingoneAuthorizeDecisionEndpointResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingoneAuthorizeDecisionEndpointResource
func AuthorizeDecisionEndpoint(clientInfo *connector.ClientInfo) *PingoneAuthorizeDecisionEndpointResource {
	return &PingoneAuthorizeDecisionEndpointResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeDecisionEndpointResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	DecisionEndpointData, err := r.getDecisionEndpointData()
	if err != nil {
		return nil, err
	}

	for decisionEndpointId, decisionEndpointName := range DecisionEndpointData {
		commentData := map[string]string{
			"Export Environment ID":  r.clientInfo.PingOneExportEnvironmentID,
			"Decision Endpoint ID":   decisionEndpointId,
			"Decision Endpoint Name": decisionEndpointName,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       decisionEndpointName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, decisionEndpointId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeDecisionEndpointResource) getDecisionEndpointData() (map[string]string, error) {
	decisionEndpointData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.PolicyDecisionManagementApi.ReadAllDecisionEndpoints(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	decisionEndpoints, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.DecisionEndpoint](iter, "ReadAllDecisionEndpoints", "GetDecisionEndpoints", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, decisionEndpoint := range decisionEndpoints {
		decisionEndpointId, decisionEndpointIdOk := decisionEndpoint.GetIdOk()
		decisionEndpointName, decisionEndpointNameOk := decisionEndpoint.GetNameOk()

		if decisionEndpointIdOk && decisionEndpointNameOk {
			decisionEndpointData[*decisionEndpointId] = *decisionEndpointName
		}
	}

	return decisionEndpointData, nil
}

func (r *PingoneAuthorizeDecisionEndpointResource) ResourceType() string {
	return "pingone_authorize_decision_endpoint"
}
