package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederatePasswordCredentialValidatorResource{}
)

type PingFederatePasswordCredentialValidatorResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederatePasswordCredentialValidatorResource
func PasswordCredentialValidator(clientInfo *connector.ClientInfo) *PingFederatePasswordCredentialValidatorResource {
	return &PingFederatePasswordCredentialValidatorResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederatePasswordCredentialValidatorResource) ResourceType() string {
	return "pingfederate_password_credential_validator"
}

func (r *PingFederatePasswordCredentialValidatorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	passwordCredentialValidatorData, err := r.getPasswordCredentialValidatorData()
	if err != nil {
		return nil, err
	}

	for passwordCredentialValidatorId, passwordCredentialValidatorName := range passwordCredentialValidatorData {
		commentData := map[string]string{
			"Password Credential Validator ID":   passwordCredentialValidatorId,
			"Password Credential Validator Name": passwordCredentialValidatorName,
			"Resource Type":                      r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       passwordCredentialValidatorName,
			ResourceID:         passwordCredentialValidatorId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederatePasswordCredentialValidatorResource) getPasswordCredentialValidatorData() (map[string]string, error) {
	passwordCredentialValidatorData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.PasswordCredentialValidatorsAPI.GetPasswordCredentialValidators(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetPasswordCredentialValidators", r.ResourceType())
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

	for _, passwordCredentialValidator := range items {
		passwordCredentialValidatorId, passwordCredentialValidatorIdOk := passwordCredentialValidator.GetIdOk()
		passwordCredentialValidatorName, passwordCredentialValidatorNameOk := passwordCredentialValidator.GetNameOk()

		if passwordCredentialValidatorIdOk && passwordCredentialValidatorNameOk {
			passwordCredentialValidatorData[*passwordCredentialValidatorId] = *passwordCredentialValidatorName
		}
	}

	return passwordCredentialValidatorData, nil
}
