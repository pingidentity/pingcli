package resources

import (
	"fmt"

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

func (r *PingFederatePasswordCredentialValidatorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.PasswordCredentialValidatorsAPI.GetPasswordCredentialValidators(r.clientInfo.Context).Execute
	apiFunctionName := "GetPasswordCredentialValidators"

	passwordCredentialValidators, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if passwordCredentialValidators == nil {
		l.Error().Msgf("Returned %s() passwordCredentialValidators is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	passwordCredentialValidatorsItems, ok := passwordCredentialValidators.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() passwordCredentialValidators items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, passwordCredentialValidator := range passwordCredentialValidatorsItems {
		passwordCredentialValidatorId, passwordCredentialValidatorIdOk := passwordCredentialValidator.GetIdOk()
		passwordCredentialValidatorName, passwordCredentialValidatorNameOk := passwordCredentialValidator.GetNameOk()

		if passwordCredentialValidatorIdOk && passwordCredentialValidatorNameOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Password Credential Validator Resource ID":   *passwordCredentialValidatorId,
				"Password Credential Validator Resource Name": *passwordCredentialValidatorName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *passwordCredentialValidatorName,
				ResourceID:         *passwordCredentialValidatorId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederatePasswordCredentialValidatorResource) ResourceType() string {
	return "pingfederate_password_credential_validator"
}
