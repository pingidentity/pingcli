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
	_ connector.ExportableResource = &PingOneKeyResource{}
)

type PingOneKeyResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneKeyResource
func Key(clientInfo *connector.ClientInfo) *PingOneKeyResource {
	return &PingOneKeyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneKeyResource) ResourceType() string {
	return "pingone_key"
}

func (r *PingOneKeyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	keyData, err := r.getKeyData()
	if err != nil {
		return nil, err
	}

	for keyId, keyNameAndType := range keyData {
		keyName := keyNameAndType[0]
		keyType := keyNameAndType[1]

		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Key ID":                keyId,
			"Key Name":              keyName,
			"Key Type":              keyType,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", keyName, keyType),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, keyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneKeyResource) getKeyData() (map[string][]string, error) {
	keyData := make(map[string][]string)

	// TODO: Implement pagination once supported in the PingOne Go Client SDK
	entityArray, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.CertificateManagementApi.GetKeys(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()

	ok, err := common.HandleClientResponse(response, err, "GetKeys", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if entityArray == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, key := range embedded.GetKeys() {
		keyId, keyIdOk := key.GetIdOk()
		keyName, keyNameOk := key.GetNameOk()
		keyUsageType, keyUsageTypeOk := key.GetUsageTypeOk()

		if keyIdOk && keyNameOk && keyUsageTypeOk {
			keyData[*keyId] = []string{*keyName, string(*keyUsageType)}
		}
	}

	return keyData, nil
}
