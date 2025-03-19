// Copyright © 2025 Ping Identity Corporation

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateAuthenticationPolicyContractResource
func AuthenticationPolicyContract(clientInfo *connector.ClientInfo) *PingFederateAuthenticationPolicyContractResource {
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

	for authenticationPolicyContractId, authenticationPolicyContractName := range authenticationPolicyContractData {
		commentData := map[string]string{
			"Authentication Policy Contract ID":   authenticationPolicyContractId,
			"Authentication Policy Contract Name": authenticationPolicyContractName,
			"Resource Type":                       r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       authenticationPolicyContractName,
			ResourceID:         authenticationPolicyContractId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationPolicyContractResource) getAuthenticationPolicyContractData() (map[string]string, error) {
	authenticationPolicyContractData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.AuthenticationPolicyContractsAPI.GetAuthenticationPolicyContracts(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetAuthenticationPolicyContracts", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, authenticationPolicyContract := range items {
		authenticationPolicyContractId, authenticationPolicyContractIdOk := authenticationPolicyContract.GetIdOk()
		authenticationPolicyContractName, authenticationPolicyContractNameOk := authenticationPolicyContract.GetNameOk()

		if authenticationPolicyContractIdOk && authenticationPolicyContractNameOk {
			authenticationPolicyContractData[*authenticationPolicyContractId] = *authenticationPolicyContractName
		}
	}

	return authenticationPolicyContractData, nil
}
