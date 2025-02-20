package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource{}
)

type PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateKerberosRealmResource
func KeypairsOauthOpenidConnectAdditionalKeySet(clientInfo *connector.PingFederateClientInfo) *PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource {
	return &PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource) ResourceType() string {
	return "pingfederate_keypairs_oauth_openid_connect_additional_key_set"
}

func (r *PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	keySetData, err := r.getKeySetData()
	if err != nil {
		return nil, err
	}

	for keySetId, keySetName := range keySetData {
		commentData := map[string]string{
			"Keypairs OAuth OpenID Connect Addition Key Set ID":   keySetId,
			"Keypairs OAuth OpenID Connect Addition Key Set Name": keySetName,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       keySetName,
			ResourceID:         keySetId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource) getKeySetData() (map[string]string, error) {
	keySetData := make(map[string]string)

	keySets, response, err := r.clientInfo.ApiClient.KeyPairsOauthOpenIdConnectAPI.GetKeySets(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetKeySets", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if keySets == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	keySetsItems, keySetsItemsOk := keySets.GetItemsOk()
	if !keySetsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, keySet := range keySetsItems {
		keySetId, keySetIdOk := keySet.GetIdOk()
		keySetName, keySetNameOk := keySet.GetNameOk()

		if keySetIdOk && keySetNameOk {
			keySetData[*keySetId] = *keySetName
		}
	}

	return keySetData, nil
}
