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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationApiApplicationResource
func AuthenticationApiApplication(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationApiApplicationResource {
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

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range *applicationData {
		commentData := map[string]string{
			"Authentication API Application ID":   appId,
			"Authentication API Application Name": appName,
			"Resource Type":                       r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       appName,
			ResourceID:         appId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationApiApplicationResource) getApplicationData() (*map[string]string, error) {
	applicationData := make(map[string]string)

	authnApiApplications, response, err := r.clientInfo.ApiClient.AuthenticationApiAPI.GetAuthenticationApiApplications(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetAuthenticationApiApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if authnApiApplications == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	authnApiApplicationsItems, authnApiApplicationsItemsOk := authnApiApplications.GetItemsOk()
	if !authnApiApplicationsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, authnApiApplication := range authnApiApplicationsItems {
		authnApiApplicationId, authnApiApplicationIdOk := authnApiApplication.GetIdOk()
		authnApiApplicationName, authnApiApplicationNameOk := authnApiApplication.GetNameOk()

		if authnApiApplicationIdOk && authnApiApplicationNameOk {
			applicationData[*authnApiApplicationId] = *authnApiApplicationName
		}
	}

	return &applicationData, nil
}
