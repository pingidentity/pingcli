package resources

import (
	"fmt"
	"regexp"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneResourceScopePingOneApiResource{}
)

type PingOneResourceScopePingOneApiResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneResourceScopePingOneApiResource
func ResourceScopePingOneApi(clientInfo *connector.PingOneClientInfo) *PingOneResourceScopePingOneApiResource {
	return &PingOneResourceScopePingOneApiResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceScopePingOneApiResource) ExportAll() (*[]connector.ImportBlock, error) {
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

		if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_PINGONE_API {
			apiResourceScopePingOneApisExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiResourceScopePingOneApisFunctionName := "ReadAllResourceScopes"

			embeddedResourceScopePingOneApis, err := common.GetManagementEmbedded(apiResourceScopePingOneApisExecuteFunc, apiResourceScopePingOneApisFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, scopePingOneApi := range embeddedResourceScopePingOneApis.GetScopes() {
				scopePingOneApiId, scopePingOneApiIdOk := scopePingOneApi.GetIdOk()
				scopePingOneApiName, scopePingOneApiNameOk := scopePingOneApi.GetNameOk()

				// Make sure the scope name is in the form of one of the following four patterns
				// p1:read:user, p1:update:user, p1:read:user:{suffix}, or p1:update:user:{suffix}
				// as supported by https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/resource_scope_pingone_api
				var scopeMatch bool
				if scopePingOneApiNameOk {
					re := regexp.MustCompile(`p1:(read|update):user($|(:.+))`)
					scopeMatch = re.MatchString(*scopePingOneApiName)
				}

				if scopeMatch && scopePingOneApiIdOk && scopePingOneApiNameOk {
					commentData := map[string]string{
						"Resource Type":                r.ResourceType(),
						"Resource Name":                *resourceName,
						"Scope PingOneApi Name":        *scopePingOneApiName,
						"Export Environment ID":        r.clientInfo.ExportEnvironmentID,
						"Resource Scope PingOneApi ID": *scopePingOneApiId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *resourceName, *scopePingOneApiName),
						ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *scopePingOneApiId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneResourceScopePingOneApiResource) ResourceType() string {
	return "pingone_resource_scope_pingone_api"
}
