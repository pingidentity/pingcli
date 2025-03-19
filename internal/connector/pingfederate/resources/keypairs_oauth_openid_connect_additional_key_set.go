// Copyright © 2025 Ping Identity Corporation

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource
func KeypairsOauthOpenidConnectAdditionalKeySet(clientInfo *connector.ClientInfo) *PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource {
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

	keypairsOauthOpenidConnectAdditionalKeySetData, err := r.getKeypairsOauthOpenidConnectAdditionalKeySetData()
	if err != nil {
		return nil, err
	}

	for keypairsOauthOpenidConnectAdditionalKeySetId, keypairsOauthOpenidConnectAdditionalKeySetName := range keypairsOauthOpenidConnectAdditionalKeySetData {
		commentData := map[string]string{
			"Keypairs Oauth Openid Connect Additional Key Set ID":   keypairsOauthOpenidConnectAdditionalKeySetId,
			"Keypairs Oauth Openid Connect Additional Key Set Name": keypairsOauthOpenidConnectAdditionalKeySetName,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       keypairsOauthOpenidConnectAdditionalKeySetName,
			ResourceID:         keypairsOauthOpenidConnectAdditionalKeySetId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateKeypairsOauthOpenidConnectAdditionalKeySetResource) getKeypairsOauthOpenidConnectAdditionalKeySetData() (map[string]string, error) {
	keypairsOauthOpenidConnectAdditionalKeySetData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.KeyPairsOauthOpenIdConnectAPI.GetKeySets(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetKeySets", r.ResourceType())
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

	for _, keypairsOauthOpenidConnectAdditionalKeySet := range items {
		keypairsOauthOpenidConnectAdditionalKeySetId, keypairsOauthOpenidConnectAdditionalKeySetIdOk := keypairsOauthOpenidConnectAdditionalKeySet.GetIdOk()
		keypairsOauthOpenidConnectAdditionalKeySetName, keypairsOauthOpenidConnectAdditionalKeySetNameOk := keypairsOauthOpenidConnectAdditionalKeySet.GetNameOk()

		if keypairsOauthOpenidConnectAdditionalKeySetIdOk && keypairsOauthOpenidConnectAdditionalKeySetNameOk {
			keypairsOauthOpenidConnectAdditionalKeySetData[*keypairsOauthOpenidConnectAdditionalKeySetId] = *keypairsOauthOpenidConnectAdditionalKeySetName
		}
	}

	return keypairsOauthOpenidConnectAdditionalKeySetData, nil
}
