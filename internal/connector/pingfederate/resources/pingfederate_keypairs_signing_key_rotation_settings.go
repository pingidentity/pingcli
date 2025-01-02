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

	signingKeyPairData, err := r.getSigningKeyPairData()
	if err != nil {
		return nil, err
	}

	for signingKeyPairId, signingKeyPairInfo := range *signingKeyPairData {
		signingKeyPairIssuerDN := signingKeyPairInfo[0]
		signingKeyPairSerialNumber := signingKeyPairInfo[1]

		commentData := map[string]string{
			"Signing Keypair ID":            signingKeyPairId,
			"Signing Keypair Issuer DN":     signingKeyPairIssuerDN,
			"Signing Keypair Serial Number": signingKeyPairSerialNumber,
			"Resource Type":                 r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s_rotation_settings", signingKeyPairIssuerDN, signingKeyPairSerialNumber),
			ResourceID:         signingKeyPairId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateKeypairsSigningKeyRotationSettingsResource) getSigningKeyPairData() (*map[string][]string, error) {
	signingKeyPairData := make(map[string][]string)

	signingKeyPairs, response, err := r.clientInfo.ApiClient.KeyPairsSigningAPI.GetSigningKeyPairs(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetSigningKeyPairs", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if signingKeyPairs == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	signingKeyPairsItems, signingKeyPairsItemsOk := signingKeyPairs.GetItemsOk()
	if !signingKeyPairsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, signingKeyPair := range signingKeyPairsItems {
		_, signingKeyPairRotationSettingsOk := signingKeyPair.GetRotationSettingsOk()

		if signingKeyPairRotationSettingsOk {
			signingKeyPairId, signingKeyPairIdOk := signingKeyPair.GetIdOk()
			signingKeyPairIssuerDN, signingKeyPairIssuerDNOk := signingKeyPair.GetIssuerDNOk()
			signingKeyPairSerialNumber, signingKeyPairSerialNumberOk := signingKeyPair.GetSerialNumberOk()

			if signingKeyPairIdOk && signingKeyPairIssuerDNOk && signingKeyPairSerialNumberOk {
				signingKeyPairData[*signingKeyPairId] = []string{*signingKeyPairIssuerDN, *signingKeyPairSerialNumber}
			}
		}
	}

	return &signingKeyPairData, nil
}
