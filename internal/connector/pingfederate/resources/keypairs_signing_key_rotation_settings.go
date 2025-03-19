// Copyright © 2025 Ping Identity Corporation

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateKeypairsSigningKeyRotationSettingsResource
func KeypairsSigningKeyRotationSettings(clientInfo *connector.ClientInfo) *PingFederateKeypairsSigningKeyRotationSettingsResource {
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

	for keypairsSigningKeyId, keypairsSigningKeyInfo := range keypairsSigningKeyData {
		ok, err := r.checkKeypairsSigningKeyRotationSettingsData(keypairsSigningKeyId)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		keypairsSigningKeyIssuerDn := keypairsSigningKeyInfo[0]
		keypairsSigningKeySerialNumber := keypairsSigningKeyInfo[1]

		commentData := map[string]string{
			"Keypairs Signing Key ID":            keypairsSigningKeyId,
			"Keypairs Signing Key Issuer DN":     keypairsSigningKeyIssuerDn,
			"Keypairs Signing Key Serial Number": keypairsSigningKeySerialNumber,
			"Resource Type":                      r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s_rotation_settings", keypairsSigningKeyIssuerDn, keypairsSigningKeySerialNumber),
			ResourceID:         keypairsSigningKeyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateKeypairsSigningKeyRotationSettingsResource) getKeypairsSigningKeyData() (map[string][]string, error) {
	signingKeyPairData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.KeyPairsSigningAPI.GetSigningKeyPairs(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetSigningKeyPairs", r.ResourceType())
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

	for _, signingKeyPair := range items {
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

	return signingKeyPairData, nil
}

func (r *PingFederateKeypairsSigningKeyRotationSettingsResource) checkKeypairsSigningKeyRotationSettingsData(id string) (bool, error) {
	_, response, err := r.clientInfo.PingFederateApiClient.KeyPairsSigningAPI.GetRotationSettings(r.clientInfo.PingFederateContext, id).Execute()
	return common.CheckSingletonResource(response, err, "GetRotationSettings", r.ResourceType())
}
