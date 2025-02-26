package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSecretManagerResource{}
)

type PingFederateSecretManagerResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSecretManagerResource
func SecretManager(clientInfo *connector.PingFederateClientInfo) *PingFederateSecretManagerResource {
	return &PingFederateSecretManagerResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSecretManagerResource) ResourceType() string {
	return "pingfederate_secret_manager"
}

func (r *PingFederateSecretManagerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	secretManagerData, err := r.getSecretManagerData()
	if err != nil {
		return nil, err
	}

	for secretManagerId, secretManagerName := range secretManagerData {
		commentData := map[string]string{
			"Resource Type":       r.ResourceType(),
			"Secret Manager ID":   secretManagerId,
			"Secret Manager Name": secretManagerName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       secretManagerName,
			ResourceID:         secretManagerId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSecretManagerResource) getSecretManagerData() (map[string]string, error) {
	secretManagerData := make(map[string]string)

	secretManagers, response, err := r.clientInfo.ApiClient.SecretManagersAPI.GetSecretManagers(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetSecretManagers", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if secretManagers == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	secretManagersItems, secretManagersItemsOk := secretManagers.GetItemsOk()
	if !secretManagersItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, secretManager := range secretManagersItems {
		secretManagerId, secretManagerIdOk := secretManager.GetIdOk()
		secretManagerName, secretManagerNameOk := secretManager.GetNameOk()

		if secretManagerIdOk && secretManagerNameOk {
			secretManagerData[*secretManagerId] = *secretManagerName
		}
	}

	return secretManagerData, nil
}
