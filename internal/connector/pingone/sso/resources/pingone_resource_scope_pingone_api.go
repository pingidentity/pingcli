package resources

import (
	"fmt"
	"regexp"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneResourceScopePingOneApiResource{}
)

type PingOneResourceScopePingOneApiResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneResourceScopePingOneApiResource
func ResourceScopePingOneApi(clientInfo *connector.ClientInfo) *PingOneResourceScopePingOneApiResource {
	return &PingOneResourceScopePingOneApiResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceScopePingOneApiResource) ResourceType() string {
	return "pingone_resource_scope_pingone_api"
}

func (r *PingOneResourceScopePingOneApiResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	resourceData, err := r.getResourceData()
	if err != nil {
		return nil, err
	}

	for resourceId, resourceName := range resourceData {
		resourceScopeData, err := r.getResourceScopeData(resourceId)
		if err != nil {
			return nil, err
		}

		for resourceScopeId, resourceScopeName := range resourceScopeData {
			commentData := map[string]string{
				"Export Environment ID":           r.clientInfo.PingOneExportEnvironmentID,
				"PingOne API Resource Name":       resourceName,
				"PingOne API Resource Scope ID":   resourceScopeId,
				"PingOne API Resource Scope Name": resourceScopeName,
				"Resource Type":                   r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", resourceName, resourceScopeName),
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, resourceScopeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneResourceScopePingOneApiResource) getResourceData() (map[string]string, error) {
	resourceData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	resourceInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedResourcesInner](iter, "ReadAllResources", "GetResources", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, resourceInner := range resourceInners {
		if resourceInner.Resource != nil {
			resourceId, resourceIdOk := resourceInner.Resource.GetIdOk()
			resourceName, resourceNameOk := resourceInner.Resource.GetNameOk()
			resourceType, resourceTypeOk := resourceInner.Resource.GetTypeOk()

			if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_PINGONE_API {
				resourceData[*resourceId] = *resourceName
			}
		}
	}

	return resourceData, nil
}

func (r *PingOneResourceScopePingOneApiResource) getResourceScopeData(resourceId string) (map[string]string, error) {
	resourceScopeData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, resourceId).Execute()
	resourceScopes, err := pingone.GetManagementAPIObjectsFromIterator[management.ResourceScope](iter, "ReadAllResourceScopes", "GetScopes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, scopePingOneApi := range resourceScopes {
		scopePingOneApiId, scopePingOneApiIdOk := scopePingOneApi.GetIdOk()
		scopePingOneApiName, scopePingOneApiNameOk := scopePingOneApi.GetNameOk()

		if scopePingOneApiIdOk && scopePingOneApiNameOk {
			// Make sure the scope name is in the form of one of the following four patterns
			// p1:read:user, p1:update:user, p1:read:user:{suffix}, or p1:update:user:{suffix}
			// as supported by https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/resource_scope_pingone_api
			re := regexp.MustCompile(`^p1:(read|update):user(|:.+)$`)

			if re.MatchString(*scopePingOneApiName) {
				resourceScopeData[*scopePingOneApiId] = *scopePingOneApiName
			}
		}
	}

	return resourceScopeData, nil
}
