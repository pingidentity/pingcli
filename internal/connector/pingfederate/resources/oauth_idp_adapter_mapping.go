package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthIdpAdapterMappingResource{}
)

type PingFederateOauthIdpAdapterMappingResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthIdpAdapterMappingResource
func OauthIdpAdapterMapping(clientInfo *connector.ClientInfo) *PingFederateOauthIdpAdapterMappingResource {
	return &PingFederateOauthIdpAdapterMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthIdpAdapterMappingResource) ResourceType() string {
	return "pingfederate_oauth_idp_adapter_mapping"
}

func (r *PingFederateOauthIdpAdapterMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthIdpAdapterMappingData, err := r.getOauthIdpAdapterMappingData()
	if err != nil {
		return nil, err
	}

	for _, oauthIdpAdapterMappingId := range oauthIdpAdapterMappingData {
		commentData := map[string]string{
			"Oauth Idp Adapter Mapping ID": oauthIdpAdapterMappingId,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_mapping", oauthIdpAdapterMappingId),
			ResourceID:         oauthIdpAdapterMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOauthIdpAdapterMappingResource) getOauthIdpAdapterMappingData() ([]string, error) {
	oauthIdpAdapterMappingData := []string{}

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthIdpAdapterMappingsAPI.GetIdpAdapterMappings(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetIdpAdapterMappings", r.ResourceType())
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

	for _, oauthIdpAdapterMapping := range items {
		oauthIdpAdapterMappingId, oauthIdpAdapterMappingIdOk := oauthIdpAdapterMapping.GetIdOk()

		if oauthIdpAdapterMappingIdOk {
			oauthIdpAdapterMappingData = append(oauthIdpAdapterMappingData, *oauthIdpAdapterMappingId)
		}
	}

	return oauthIdpAdapterMappingData, nil
}
