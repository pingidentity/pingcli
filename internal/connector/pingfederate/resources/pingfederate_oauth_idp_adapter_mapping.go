package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthIdpAdapterMappingResource{}
)

type PingFederateOAuthIdpAdapterMappingResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthIdpAdapterMappingResource
func OAuthIdpAdapterMapping(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthIdpAdapterMappingResource {
	return &PingFederateOAuthIdpAdapterMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthIdpAdapterMappingResource) ResourceType() string {
	return "pingfederate_oauth_idp_adapter_mapping"
}

func (r *PingFederateOAuthIdpAdapterMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	idpAdapterMappingData, err := r.getIdpAdapterMappingData()
	if err != nil {
		return nil, err
	}

	for _, idpAdapterMappingId := range idpAdapterMappingData {
		commentData := map[string]string{
			"OAuth IDP Adapter Mapping ID": idpAdapterMappingId,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_mapping", idpAdapterMappingId),
			ResourceID:         idpAdapterMappingId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthIdpAdapterMappingResource) getIdpAdapterMappingData() ([]string, error) {
	idpAdapterMappingData := []string{}

	idpAdapterMappings, response, err := r.clientInfo.ApiClient.OauthIdpAdapterMappingsAPI.GetIdpAdapterMappings(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetIdpAdapterMappings", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if idpAdapterMappings == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	idpAdapterMappingsItems, idpAdapterMappingsItemsOk := idpAdapterMappings.GetItemsOk()
	if !idpAdapterMappingsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, idpAdapterMapping := range idpAdapterMappingsItems {
		idpAdapterMappingId, idpAdapterMappingIdOk := idpAdapterMapping.GetIdOk()

		if idpAdapterMappingIdOk {
			idpAdapterMappingData = append(idpAdapterMappingData, *idpAdapterMappingId)
		}
	}

	return idpAdapterMappingData, nil
}
