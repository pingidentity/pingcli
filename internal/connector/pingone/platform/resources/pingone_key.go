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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneKeyResource
func Key(clientInfo *connector.PingOneClientInfo) *PingOneKeyResource {
	return &PingOneKeyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneKeyResource) ResourceType() string {
	return "pingone_key"
}

func (r *PingOneKeyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportKeys()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneKeyResource) exportKeys() error {
	// TODO: Implement pagination once supported in the PingOne Go Client SDK
	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetKeys(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	err = common.HandleClientResponse(response, err, "GetKeys", r.ResourceType())
	if err != nil {
		return err
	}

	if entityArray == nil {
		return common.DataNilError(r.ResourceType(), response)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		return common.DataNilError(r.ResourceType(), response)
	}

	for _, key := range embedded.GetKeys() {
		keyId, keyIdOk := key.GetIdOk()
		keyName, keyNameOk := key.GetNameOk()
		keyUsageType, keyUsageTypeOk := key.GetUsageTypeOk()

		if keyIdOk && keyNameOk && keyUsageTypeOk {
			r.addImportBlock(*keyId, *keyName, string(*keyUsageType))
		}
	}

	return nil
}

func (r *PingOneKeyResource) addImportBlock(keyId, keyName, keyUsageType string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Key ID":                keyId,
		"Key Name":              keyName,
		"Key Usage Type":        keyUsageType,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", keyName, keyUsageType),
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, keyId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
