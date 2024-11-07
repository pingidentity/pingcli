package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkServiceResource{}
)

type PingoneAuthorizeTrustFrameworkServiceResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkServiceResource
func AuthorizeTrustFrameworkService(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkServiceResource {
	return &PingoneAuthorizeTrustFrameworkServiceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkServiceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorServicesApi.ListServices(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ListServices"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authorizationService := range embedded.GetAuthorizationServices() {

		var (
			authorizationServiceId     *string
			authorizationServiceIdOk   bool
			authorizationServiceName   *string
			authorizationServiceNameOk bool
		)

		switch t := authorizationService.GetActualInstance().(type) {
		case *authorize.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO:
			authorizationServiceId, authorizationServiceIdOk = t.GetIdOk()
			authorizationServiceName, authorizationServiceNameOk = t.GetFullNameOk()
		case *authorize.AuthorizeEditorDataServicesHttpServiceDefinitionDTO:
			authorizationServiceId, authorizationServiceIdOk = t.GetIdOk()
			authorizationServiceName, authorizationServiceNameOk = t.GetFullNameOk()
		case *authorize.AuthorizeEditorDataServicesNoneServiceDefinitionDTO:
			authorizationServiceId, authorizationServiceIdOk = t.GetIdOk()
			authorizationServiceName, authorizationServiceNameOk = t.GetFullNameOk()
		default:
			continue
		}

		if authorizationServiceNameOk && authorizationServiceIdOk {
			commentData := map[string]string{
				"Resource Type":                          r.ResourceType(),
				"Authorize Trust Framework Service Name": *authorizationServiceName,
				"Export Environment ID":                  r.clientInfo.ExportEnvironmentID,
				"Authorize Trust Framework Service ID":   *authorizationServiceId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authorizationServiceName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *authorizationServiceId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkServiceResource) ResourceType() string {
	return "pingone_authorize_trust_framework_service"
}
