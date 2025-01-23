package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateKeypairsSigningKeyRotationSettingsResource{}
)

type PingFederateKeypairsSigningKeyRotationSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateKeypairsSigningKeyRotationSettingsResource
func KeypairsSigningKeyRotationSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateKeypairsSigningKeyRotationSettingsResource {
	return &PingFederateKeypairsSigningKeyRotationSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateKeypairsSigningKeyRotationSettingsResource) ResourceType() string {
	return "pingfederate_keypairs_signing_key_rotation_settings"
}

func (r *PingFederateKeypairsSigningKeyRotationSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	keypairsSigningKeyData, err := r.getKeypairsSigningKeyData()
	if err != nil {
		return nil, err
	}

	for keyPairViewId, keyPairViewInfo := range *keypairsSigningKeyData {
		keyPairViewIssuerDn := keyPairViewInfo[0]
		keyPairViewSerialNumber := keyPairViewInfo[1]

		commentData := map[string]string{
			"Keypairs Signing Key  ID":           keyPairViewId,
			"Keypairs Signing Key Issuer DN":     keyPairViewIssuerDn,
			"Keypairs Signing Key Serial Number": keyPairViewSerialNumber,
			"Resource Type":                      r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s_rotation_settings", keyPairViewIssuerDn, keyPairViewSerialNumber),
			ResourceID:         keyPairViewId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateKeypairsSigningKeyRotationSettingsResource) getKeypairsSigningKeyData() (*map[string][]string, error) {
	keypairsSigningKeyData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.ApiClient.KeyPairsSigningAPI.GetSigningKeyPairs(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetSigningKeyPairs", r.ResourceType())
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

	for _, keyPairView := range items {
		keyPairViewId, keyPairViewIdOk := keyPairView.GetIdOk()
		keyPairViewIssuerDn, keyPairViewIssuerDnOk := keyPairView.GetIssuerDNOk()
		keyPairViewSerialNumber, keyPairViewSerialNumberOk := keyPairView.GetSerialNumberOk()

		if keyPairViewIdOk && keyPairViewIssuerDnOk && keyPairViewSerialNumberOk {
			keypairsSigningKeyData[*keyPairViewId] = []string{*keyPairViewIssuerDn, *keyPairViewSerialNumber}
		}
	}

	return &keypairsSigningKeyData, nil
}
