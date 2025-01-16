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

	for authenticationSelectorId, authenticationSelectorName := range *authenticationSelectorData {
		commentData := map[string]string{
			"Authentication Selector ID":   authenticationSelectorId,
			"Authentication Selector Name": authenticationSelectorName,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       authenticationSelectorName,
			ResourceID:         authenticationSelectorId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationSelectorResource) getAuthenticationSelectorData() (*map[string]string, error) {
	authenticationSelectorData := make(map[string]string)

	apiObj, response, err := r.clientInfo.ApiClient.AuthenticationSelectorsAPI.GetAuthenticationSelectors(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetAuthenticationSelectors", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, authenticationSelector := range items {
		authenticationSelectorId, authenticationSelectorIdOk := authenticationSelector.GetIdOk()
		authenticationSelectorName, authenticationSelectorNameOk := authenticationSelector.GetNameOk()

		if authenticationSelectorIdOk && authenticationSelectorNameOk {
			authenticationSelectorData[*authenticationSelectorId] = *authenticationSelectorName
		}
	}

	return &authenticationSelectorData, nil
}
