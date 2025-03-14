package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationApiApplicationResource{}
)

type PingFederateAuthenticationApiApplicationResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateAuthenticationApiApplicationResource
func AuthenticationApiApplication(clientInfo *connector.ClientInfo) *PingFederateAuthenticationApiApplicationResource {
	return &PingFederateAuthenticationApiApplicationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationApiApplicationResource) ResourceType() string {
	return "pingfederate_authentication_api_application"
}

func (r *PingFederateAuthenticationApiApplicationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	authenticationApiApplicationData, err := r.getAuthenticationApiApplicationData()
	if err != nil {
		return nil, err
	}

	for authenticationApiApplicationId, authenticationApiApplicationName := range authenticationApiApplicationData {
		commentData := map[string]string{
			"Authentication Api Application ID":   authenticationApiApplicationId,
			"Authentication Api Application Name": authenticationApiApplicationName,
			"Resource Type":                       r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       authenticationApiApplicationName,
			ResourceID:         authenticationApiApplicationId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationApiApplicationResource) getAuthenticationApiApplicationData() (map[string]string, error) {
	authenticationApiApplicationData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.AuthenticationApiAPI.GetAuthenticationApiApplications(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetAuthenticationApiApplications", r.ResourceType())
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

	for _, authenticationApiApplication := range items {
		authenticationApiApplicationId, authenticationApiApplicationIdOk := authenticationApiApplication.GetIdOk()
		authenticationApiApplicationName, authenticationApiApplicationNameOk := authenticationApiApplication.GetNameOk()

		if authenticationApiApplicationIdOk && authenticationApiApplicationNameOk {
			authenticationApiApplicationData[*authenticationApiApplicationId] = *authenticationApiApplicationName
		}
	}

	return authenticationApiApplicationData, nil
}
