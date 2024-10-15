package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIDPAdapterResource{}
)

type PingFederateIDPAdapterResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIDPAdapterResource
func IDPAdapter(clientInfo *connector.PingFederateClientInfo) *PingFederateIDPAdapterResource {
	return &PingFederateIDPAdapterResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIDPAdapterResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.IdpAdaptersAPI.GetIdpAdapters(r.clientInfo.Context).Execute
	apiFunctionName := "GetIdpAdapters"

	idpAdapters, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if idpAdapters == nil {
		l.Error().Msgf("Returned %s() idpAdapters is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	idpAdaptersItems, ok := idpAdapters.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() idpAdapters items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, idpAdapter := range idpAdaptersItems {
		idpAdapterId, idpAdapterIdOk := idpAdapter.GetIdOk()
		idpAdapterName, idpAdapterNameOk := idpAdapter.GetNameOk()

		if idpAdapterIdOk && idpAdapterNameOk {
			commentData := map[string]string{
				"Resource Type":             r.ResourceType(),
				"IDP Adapter Resource ID":   *idpAdapterId,
				"IDP Adapter Resource Name": *idpAdapterName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *idpAdapterName,
				ResourceID:         *idpAdapterId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateIDPAdapterResource) ResourceType() string {
	return "pingfederate_idp_adapter"
}
