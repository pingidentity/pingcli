package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateMetadataUrlResource{}
)

type PingFederateMetadataUrlResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateMetadataUrlResource
func MetadataUrl(clientInfo *connector.ClientInfo) *PingFederateMetadataUrlResource {
	return &PingFederateMetadataUrlResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateMetadataUrlResource) ResourceType() string {
	return "pingfederate_metadata_url"
}

func (r *PingFederateMetadataUrlResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	metadataUrlData, err := r.getMetadataUrlData()
	if err != nil {
		return nil, err
	}

	for metadataUrlId, metadataUrlName := range metadataUrlData {
		commentData := map[string]string{
			"Metadata Url ID":   metadataUrlId,
			"Metadata Url Name": metadataUrlName,
			"Resource Type":     r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       metadataUrlName,
			ResourceID:         metadataUrlId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateMetadataUrlResource) getMetadataUrlData() (map[string]string, error) {
	metadataUrlData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.MetadataUrlsAPI.GetMetadataUrls(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetMetadataUrls", r.ResourceType())
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

	for _, metadataUrl := range items {
		metadataUrlId, metadataUrlIdOk := metadataUrl.GetIdOk()
		metadataUrlName, metadataUrlNameOk := metadataUrl.GetNameOk()

		if metadataUrlIdOk && metadataUrlNameOk {
			metadataUrlData[*metadataUrlId] = *metadataUrlName
		}
	}

	return metadataUrlData, nil
}
