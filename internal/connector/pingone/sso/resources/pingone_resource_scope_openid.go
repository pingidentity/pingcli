package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceScopeOpenIdResource{}
)

type PingoneResourceScopeOpenIdResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneResourceScopeOpenIdResource
func ResourceScopeOpenId(clientInfo *connector.PingOneClientInfo) *PingoneResourceScopeOpenIdResource {
	return &PingoneResourceScopeOpenIdResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceScopeOpenIdResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllResources"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, resourceInner := range embedded.GetResources() {
		resource := resourceInner.Resource
		resourceId, resourceIdOk := resource.GetIdOk()
		resourceName, resourceNameOk := resource.GetNameOk()
		resourceType, resourceTypeOk := resource.GetTypeOk()

		if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_OPENID_CONNECT {
			apiResourceScopeOpenIdsExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiResourceScopeOpenIdsFunctionName := "ReadAllResourceScopes"

			embeddedResourceScopeOpenIds, err := common.GetManagementEmbedded(apiResourceScopeOpenIdsExecuteFunc, apiResourceScopeOpenIdsFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, scopeOpenId := range embeddedResourceScopeOpenIds.GetScopes() {
				scopeOpenIdId, scopeOpenIdIdOk := scopeOpenId.GetIdOk()
				scopeOpenIdName, scopeOpenIdNameOk := scopeOpenId.GetNameOk()
				if scopeOpenIdIdOk && scopeOpenIdNameOk {
					commentData := map[string]string{
						"Resource Type":            r.ResourceType(),
						"Resource Name":            *resourceName,
						"Scope OpenID Name":        *scopeOpenIdName,
						"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
						"Resource Scope OpenID ID": *scopeOpenIdId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *resourceName, *scopeOpenIdName),
						ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *scopeOpenIdId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceScopeOpenIdResource) ResourceType() string {
	return "pingone_resource_scope_openid"
}
