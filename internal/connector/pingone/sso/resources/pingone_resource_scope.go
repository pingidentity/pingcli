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
	_ connector.ExportableResource = &PingoneResourceScopeResource{}
)

type PingoneResourceScopeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneResourceScopeResource
func ResourceScope(clientInfo *connector.PingOneClientInfo) *PingoneResourceScopeResource {
	return &PingoneResourceScopeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceScopeResource) ExportAll() (*[]connector.ImportBlock, error) {
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

		if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_CUSTOM {
			apiResourceScopesExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiResourceScopesFunctionName := "ReadAllResourceScopes"

			embeddedResourceScopes, err := common.GetManagementEmbedded(apiResourceScopesExecuteFunc, apiResourceScopesFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, scope := range embeddedResourceScopes.GetScopes() {
				scopeId, scopeIdOk := scope.GetIdOk()
				scopeName, scopeNameOk := scope.GetNameOk()
				if scopeIdOk && scopeNameOk {
					commentData := map[string]string{
						"Resource Type":         r.ResourceType(),
						"Resource Name":         *resourceName,
						"Scope Name":            *scopeName,
						"Export Environment ID": r.clientInfo.ExportEnvironmentID,
						"Resource ID":           *resourceId,
						"Scope ID":              *scopeId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *resourceName, *scopeName),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *resourceId, *scopeId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceScopeResource) ResourceType() string {
	return "pingone_resource_scope"
}
