package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationPolicyContractResource{}
)

type PingFederateAuthenticationPolicyContractResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationPolicyContractResource
func AuthenticationPolicyContract(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationPolicyContractResource {
	return &PingFederateAuthenticationPolicyContractResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationPolicyContractResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContracts(r.clientInfo.Context).Execute
	apiFunctionName := "GetAuthenticationPolicyContracts"

	authnPolicyContracts, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if authnPolicyContracts == nil {
		l.Error().Msgf("Returned %s() authnPolicyContracts is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	authnPolicyContractsItems, ok := authnPolicyContracts.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() authnPolicyContracts items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authnPolicyContract := range authnPolicyContractsItems {
		authnPolicyContractId, authnPolicyContractIdOk := authnPolicyContract.GetIdOk()
		authnPolicyContractName, authnPolicyContractNameOk := authnPolicyContract.GetNameOk()

		if authnPolicyContractIdOk && authnPolicyContractNameOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authentication Policy Contract Resource ID":   *authnPolicyContractId,
				"Authentication Policy Contract Resource Name": *authnPolicyContractName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authnPolicyContractName,
				ResourceID:         *authnPolicyContractId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationPolicyContractResource) ResourceType() string {
	return "pingfederate_authentication_policy_contract"
}
