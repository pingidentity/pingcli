package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationSelectorResource{}
)

type PingFederateAuthenticationSelectorResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationSelectorResource
func AuthenticationSelector(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationSelectorResource {
	return &PingFederateAuthenticationSelectorResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationSelectorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthenticationSelectorsAPI.GetAuthenticationSelectors(r.clientInfo.Context).Execute
	apiFunctionName := "GetAuthenticationSelectors"

	authnSelectors, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if authnSelectors == nil {
		l.Error().Msgf("Returned %s() authnSelectors is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	authnSelectorsItems, ok := authnSelectors.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() authnSelectors items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authnSelector := range authnSelectorsItems {
		authnSelectorId, authnSelectorIdOk := authnSelector.GetIdOk()
		authnSelectorName, authnSelectorNameOk := authnSelector.GetNameOk()

		if authnSelectorIdOk && authnSelectorNameOk {
			commentData := map[string]string{
				"Resource Type":                         r.ResourceType(),
				"Authentication Selector Resource ID":   *authnSelectorId,
				"Authentication Selector Resource Name": *authnSelectorName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authnSelectorName,
				ResourceID:         *authnSelectorId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationSelectorResource) ResourceType() string {
	return "pingfederate_authentication_selector"
}
