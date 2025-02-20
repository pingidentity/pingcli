package resources

import (
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

func (r *PingFederateAuthenticationSelectorResource) ResourceType() string {
	return "pingfederate_authentication_selector"
}

func (r *PingFederateAuthenticationSelectorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	authenticationSelectorData, err := r.getAuthenticationSelectorData()
	if err != nil {
		return nil, err
	}

	for authnSelectorId, authnSelectorName := range authenticationSelectorData {
		commentData := map[string]string{
			"Authentication Selector ID":   authnSelectorId,
			"Authentication Selector Name": authnSelectorName,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       authnSelectorName,
			ResourceID:         authnSelectorId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationSelectorResource) getAuthenticationSelectorData() (map[string]string, error) {
	authenticationSelectorData := make(map[string]string)

	authnSelectors, response, err := r.clientInfo.ApiClient.AuthenticationSelectorsAPI.GetAuthenticationSelectors(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetAuthenticationSelectors", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if authnSelectors == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	authnSelectorsItems, authnSelectorsItemsOk := authnSelectors.GetItemsOk()
	if !authnSelectorsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, authnSelector := range authnSelectorsItems {
		authnSelectorId, authnSelectorIdOk := authnSelector.GetIdOk()
		authnSelectorName, authnSelectorNameOk := authnSelector.GetNameOk()

		if authnSelectorIdOk && authnSelectorNameOk {
			authenticationSelectorData[*authnSelectorId] = *authnSelectorName
		}
	}

	return authenticationSelectorData, nil
}
