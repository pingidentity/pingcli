// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIdpStsRequestParametersContractResource{}
)

type PingFederateIdpStsRequestParametersContractResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateIdpStsRequestParametersContractResource
func IdpStsRequestParametersContract(clientInfo *connector.ClientInfo) *PingFederateIdpStsRequestParametersContractResource {
	return &PingFederateIdpStsRequestParametersContractResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIdpStsRequestParametersContractResource) ResourceType() string {
	return "pingfederate_idp_sts_request_parameters_contract"
}

func (r *PingFederateIdpStsRequestParametersContractResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	idpStsRequestParametersContractData, err := r.getIdpStsRequestParametersContractData()
	if err != nil {
		return nil, err
	}

	for idpStsRequestParametersContractId, idpStsRequestParametersContractName := range idpStsRequestParametersContractData {
		commentData := map[string]string{
			"Idp Sts Request Parameters Contract ID":   idpStsRequestParametersContractId,
			"Idp Sts Request Parameters Contract Name": idpStsRequestParametersContractName,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       idpStsRequestParametersContractName,
			ResourceID:         idpStsRequestParametersContractId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdpStsRequestParametersContractResource) getIdpStsRequestParametersContractData() (map[string]string, error) {
	idpStsRequestParametersContractData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.IdpStsRequestParametersContractsAPI.GetStsRequestParamContracts(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetStsRequestParamContracts", r.ResourceType())
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

	for _, idpStsRequestParametersContract := range items {
		idpStsRequestParametersContractId, idpStsRequestParametersContractIdOk := idpStsRequestParametersContract.GetIdOk()
		idpStsRequestParametersContractName, idpStsRequestParametersContractNameOk := idpStsRequestParametersContract.GetNameOk()

		if idpStsRequestParametersContractIdOk && idpStsRequestParametersContractNameOk {
			idpStsRequestParametersContractData[*idpStsRequestParametersContractId] = *idpStsRequestParametersContractName
		}
	}

	return idpStsRequestParametersContractData, nil
}
