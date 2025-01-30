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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSecretManagerResource
func SecretManager(clientInfo *connector.ClientInfo) *PingFederateSecretManagerResource {
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

	for secretManagerId, secretManagerName := range *secretManagerData {
		commentData := map[string]string{
			"Secret Manager ID":   secretManagerId,
			"Secret Manager Name": secretManagerName,
			"Resource Type":       r.ResourceType(),
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

func (r *PingFederateSecretManagerResource) getSecretManagerData() (*map[string]string, error) {
	secretManagerData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.SecretManagersAPI.GetSecretManagers(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetSecretManagers", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, secretManager := range items {
		secretManagerId, secretManagerIdOk := secretManager.GetIdOk()
		secretManagerName, secretManagerNameOk := secretManager.GetNameOk()

		if secretManagerIdOk && secretManagerNameOk {
			secretManagerData[*secretManagerId] = *secretManagerName
		}
	}

	return &secretManagerData, nil
}
