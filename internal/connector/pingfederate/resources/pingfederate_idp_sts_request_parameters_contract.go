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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIdpStsRequestParametersContractResource
func IdpStsRequestParametersContract(clientInfo *connector.PingFederateClientInfo) *PingFederateIdpStsRequestParametersContractResource {
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

	stsRequestParamContractData, err := r.getStsRequestParamContractData()
	if err != nil {
		return nil, err
	}

	for stsRequestParamContractId, stsRequestParamContractName := range *stsRequestParamContractData {
		commentData := map[string]string{
			"IDP STS Request Parameters Contract ID":   stsRequestParamContractId,
			"IDP STS Request Parameters Contract Name": stsRequestParamContractName,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       stsRequestParamContractName,
			ResourceID:         stsRequestParamContractId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdpStsRequestParametersContractResource) getStsRequestParamContractData() (*map[string]string, error) {
	stsRequestParamContractData := make(map[string]string)

	stsRequestParamContracts, response, err := r.clientInfo.ApiClient.IdpStsRequestParametersContractsAPI.GetStsRequestParamContracts(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetStsRequestParamContracts", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if stsRequestParamContracts == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	stsRequestParamContractsItems, stsRequestParamContractsOk := stsRequestParamContracts.GetItemsOk()
	if !stsRequestParamContractsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, stsRequestParamContract := range stsRequestParamContractsItems {
		stsRequestParamContractId, stsRequestParamContractIdOk := stsRequestParamContract.GetIdOk()
		stsRequestParamContractName, stsRequestParamContractNameOk := stsRequestParamContract.GetNameOk()

		if stsRequestParamContractIdOk && stsRequestParamContractNameOk {
			stsRequestParamContractData[*stsRequestParamContractId] = *stsRequestParamContractName
		}
	}

	return &stsRequestParamContractData, nil
}
