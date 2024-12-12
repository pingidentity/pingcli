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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederatePasswordCredentialValidatorResource
func PasswordCredentialValidator(clientInfo *connector.PingFederateClientInfo) *PingFederatePasswordCredentialValidatorResource {
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

	for passwordCredentialValidatorId, passwordCredentialValidatorName := range *passwordCredentialValidatorData {
		commentData := map[string]string{
			"Password Credential Validator Resource ID":   passwordCredentialValidatorId,
			"Password Credential Validator Resource Name": passwordCredentialValidatorName,
			"Resource Type": r.ResourceType(),
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

func (r *PingFederatePasswordCredentialValidatorResource) getPasswordCredentialValidatorData() (*map[string]string, error) {
	passwordCredentialValidatorData := make(map[string]string)

	passwordCredentialValidators, response, err := r.clientInfo.ApiClient.PasswordCredentialValidatorsAPI.GetPasswordCredentialValidators(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetPasswordCredentialValidators", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if passwordCredentialValidators == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	passwordCredentialValidatorsItems, ok := passwordCredentialValidators.GetItemsOk()
	if !ok {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, passwordCredentialValidator := range passwordCredentialValidatorsItems {
		passwordCredentialValidatorId, passwordCredentialValidatorIdOk := passwordCredentialValidator.GetIdOk()
		passwordCredentialValidatorName, passwordCredentialValidatorNameOk := passwordCredentialValidator.GetNameOk()

		if passwordCredentialValidatorIdOk && passwordCredentialValidatorNameOk {
			passwordCredentialValidatorData[*passwordCredentialValidatorId] = *passwordCredentialValidatorName
		}
	}

	return &passwordCredentialValidatorData, nil
}
