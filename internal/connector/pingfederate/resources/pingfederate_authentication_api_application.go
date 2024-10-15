package resources

import (
	"fmt"

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

func (r *PingFederateAuthenticationApiApplicationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthenticationApiAPI.GetAuthenticationApiApplications(r.clientInfo.Context).Execute
	apiFunctionName := "GetAuthenticationApiApplications"

	authnApiApplications, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if authnApiApplications == nil {
		l.Error().Msgf("Returned %s() authnApiApplications is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	authnApiApplicationsItems, ok := authnApiApplications.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() authnApiApplications items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authnApiApplication := range authnApiApplicationsItems {
		authnApiApplicationId, authnApiApplicationIdOk := authnApiApplication.GetIdOk()
		authnApiApplicationName, authnApiApplicationNameOk := authnApiApplication.GetNameOk()

		if authnApiApplicationIdOk && authnApiApplicationNameOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authentication API Application Resource ID":   *authnApiApplicationId,
				"Authentication API Application Resource Name": *authnApiApplicationName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authnApiApplicationName,
				ResourceID:         *authnApiApplicationId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationApiApplicationResource) ResourceType() string {
	return "pingfederate_authentication_api_application"
}
