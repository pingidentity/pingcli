package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkAttributeResource{}
)

type PingoneAuthorizeTrustFrameworkAttributeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkAttributeResource
func AuthorizeTrustFrameworkAttribute(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkAttributeResource {
	return &PingoneAuthorizeTrustFrameworkAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorAttributesApi.ListAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ListAttributes"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authorizationAttribute := range embedded.GetAuthorizationAttributes() {
		authorizationAttributeName, authorizationAttributeNameOk := authorizationAttribute.GetNameOk()
		authorizationAttributeId, authorizationAttributeIdOk := authorizationAttribute.GetIdOk()

		if authorizationAttributeNameOk && authorizationAttributeIdOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authorize Trust Framework Attribute Name": *authorizationAttributeName,
				"Export Environment ID":                    r.clientInfo.ExportEnvironmentID,
				"Authorize Trust Framework Attribute ID":   *authorizationAttributeId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authorizationAttributeName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *authorizationAttributeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkAttributeResource) ResourceType() string {
	return "pingone_authorize_trust_framework_attribute"
}
