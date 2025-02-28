package resources

import (
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

func (r *PingFederateAuthenticationPolicyContractResource) ResourceType() string {
	return "pingfederate_authentication_policy_contract"
}

func (r *PingFederateAuthenticationPolicyContractResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	authenticationPolicyContractData, err := r.getAuthenticationPolicyContractData()
	if err != nil {
		return nil, err
	}

	for authnPolicyContractId, authnPolicyContractName := range authenticationPolicyContractData {
		commentData := map[string]string{
			"Authentication Policy Contract ID":   authnPolicyContractId,
			"Authentication Policy Contract Name": authnPolicyContractName,
			"Resource Type":                       r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       authnPolicyContractName,
			ResourceID:         authnPolicyContractId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationPolicyContractResource) getAuthenticationPolicyContractData() (map[string]string, error) {
	authenticationPolicyContractData := make(map[string]string)

	authnPolicyContracts, response, err := r.clientInfo.ApiClient.AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContracts(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetAuthenticationPolicyContracts", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if authnPolicyContracts == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	authnPolicyContractsItems, authnPolicyContractsItemsOk := authnPolicyContracts.GetItemsOk()
	if !authnPolicyContractsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, authnPolicyContract := range authnPolicyContractsItems {
		authnPolicyContractId, authnPolicyContractIdOk := authnPolicyContract.GetIdOk()
		authnPolicyContractName, authnPolicyContractNameOk := authnPolicyContract.GetNameOk()

		if authnPolicyContractIdOk && authnPolicyContractNameOk {
			authenticationPolicyContractData[*authnPolicyContractId] = *authnPolicyContractName
		}
	}

	return authenticationPolicyContractData, nil
}
