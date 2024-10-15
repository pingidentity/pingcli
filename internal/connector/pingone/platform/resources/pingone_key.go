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

func (r *PingOneKeyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetKeys(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "GetKeys"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, key := range embedded.GetKeys() {
		keyId, keyIdOk := key.GetIdOk()
		keyName, keyNameOk := key.GetNameOk()
		keyUsageType, keyUsageTypeOk := key.GetUsageTypeOk()

		if keyIdOk && keyNameOk && keyUsageTypeOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Key Name":              *keyName,
				"Key Usage Type":        string(*keyUsageType),
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Key ID":                *keyId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", *keyName, *keyUsageType),
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *keyId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneKeyResource) ResourceType() string {
	return "pingone_key"
}
