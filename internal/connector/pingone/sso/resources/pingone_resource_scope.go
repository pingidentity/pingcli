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
	_ connector.ExportableResource = &PingOneResourceScopeResource{}
)

type PingOneResourceScopeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneResourceScopeResource
func ResourceScope(clientInfo *connector.PingOneClientInfo) *PingOneResourceScopeResource {
	return &PingOneResourceScopeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceScopeResource) ResourceType() string {
	return "pingone_resource_scope"
}

func (r *PingOneResourceScopeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	resourceData, err := r.getResourceData()
	if err != nil {
		return nil, err
	}

	for resourceId, resourceName := range *resourceData {
		resourceScopeData, err := r.getResourceScopeData(resourceId)
		if err != nil {
			return nil, err
		}

		for resourceScopeId, resourceScopeName := range *resourceScopeData {
			commentData := map[string]string{
				"Custom Resource ID":         resourceId,
				"Custom Resource Name":       resourceName,
				"Custom Resource Scope ID":   resourceScopeId,
				"Custom Resource Scope Name": resourceScopeName,
				"Export Environment ID":      r.clientInfo.ExportEnvironmentID,
				"Resource Type":              r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", resourceName, resourceScopeName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, resourceId, resourceScopeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneResourceScopeResource) getResourceData() (*map[string]string, error) {
	resourceData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllResources", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, resourceInner := range embedded.GetResources() {
			if resourceInner.Resource != nil {
				resourceId, resourceIdOk := resourceInner.Resource.GetIdOk()
				resourceName, resourceNameOk := resourceInner.Resource.GetNameOk()
				resourceType, resourceTypeOk := resourceInner.Resource.GetTypeOk()

				if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_CUSTOM {
					resourceData[*resourceId] = *resourceName
				}
			}
		}
	}

	return &resourceData, nil
}

func (r *PingOneResourceScopeResource) getResourceScopeData(resourceId string) (*map[string]string, error) {
	resourceScopeData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, resourceId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllResourceScopes", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, scope := range embedded.GetScopes() {
			scopeId, scopeIdOk := scope.GetIdOk()
			scopeName, scopeNameOk := scope.GetNameOk()
			if scopeIdOk && scopeNameOk {
				resourceScopeData[*scopeId] = *scopeName
			}
		}
	}

	return &resourceScopeData, nil
}
