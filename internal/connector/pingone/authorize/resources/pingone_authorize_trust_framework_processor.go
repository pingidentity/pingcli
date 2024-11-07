package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkProcessorResource{}
)

type PingoneAuthorizeTrustFrameworkProcessorResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkProcessorResource
func AuthorizeTrustFrameworkProcessor(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkProcessorResource {
	return &PingoneAuthorizeTrustFrameworkProcessorResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkProcessorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorProcessorsApi.ListProcessors(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ListProcessors"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authorizationProcessor := range embedded.GetAuthorizationProcessors() {
		authorizationProcessorName, authorizationProcessorNameOk := authorizationProcessor.GetFullNameOk()
		authorizationProcessorId, authorizationProcessorIdOk := authorizationProcessor.GetIdOk()

		if authorizationProcessorNameOk && authorizationProcessorIdOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authorize Trust Framework Processor Name": *authorizationProcessorName,
				"Export Environment ID":                    r.clientInfo.ExportEnvironmentID,
				"Authorize Trust Framework Processor ID":   *authorizationProcessorId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authorizationProcessorName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *authorizationProcessorId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkProcessorResource) ResourceType() string {
	return "pingone_authorize_trust_framework_processor"
}
