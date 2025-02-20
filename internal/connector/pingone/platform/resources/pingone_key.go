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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneKeyResource
func Key(clientInfo *connector.PingOneClientInfo) *PingOneKeyResource {
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

	keyData, err := getKeyData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for keyId, keyNameAndType := range keyData {
		keyName := keyNameAndType[0]
		keyType := keyNameAndType[1]

		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Key ID":                keyId,
			"Key Name":              keyName,
			"Key Type":              keyType,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", keyName, keyType),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, keyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func getKeyData(clientInfo *connector.PingOneClientInfo, resourceType string) (map[string][]string, error) {
	keyData := make(map[string][]string)

	// TODO: Implement pagination once supported in the PingOne Go Client SDK
	entityArray, response, err := clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetKeys(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()

	ok, err := common.HandleClientResponse(response, err, "GetKeys", resourceType)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if entityArray == nil {
		return nil, fmt.Errorf("failed to export resource '%s'.\n"+
			"PingOne API request for resource '%s' was not successful. response data is nil.\n"+
			"response code: %s\n"+
			"response body: %s",
			resourceType, resourceType, response.Status, response.Body)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		return nil, fmt.Errorf("failed to export resource '%s'.\n"+
			"PingOne API request for resource '%s' was not successful. response data is nil.\n"+
			"response code: %s\n"+
			"response body: %s",
			resourceType, resourceType, response.Status, response.Body)
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
