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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneResourceScopeResource
func ResourceScope(clientInfo *connector.PingOneClientInfo) *PingOneResourceScopeResource {
	return &PingOneResourceScopeResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneResourceScopeResource) ResourceType() string {
	return "pingone_resource_scope"
}

func (r *PingOneResourceScopeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportResourceScopes()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneResourceScopeResource) exportResourceScopes() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllResources", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, resourceInner := range embedded.GetResources() {
			if resourceInner.Resource != nil {
				resourceId, resourceIdOk := resourceInner.Resource.GetIdOk()
				resourceName, resourceNameOk := resourceInner.Resource.GetNameOk()
				resourceType, resourceTypeOk := resourceInner.Resource.GetTypeOk()

				if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_CUSTOM {
					err := r.exportResourceScopesByResource(*resourceId, *resourceName)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneResourceScopeResource) exportResourceScopesByResource(resourceId, resourceName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, resourceId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllResourceScopes", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, scope := range embedded.GetScopes() {
			scopeId, scopeIdOk := scope.GetIdOk()
			scopeName, scopeNameOk := scope.GetNameOk()
			if scopeIdOk && scopeNameOk {
				r.addImportBlock(resourceId, resourceName, *scopeId, *scopeName)
			}
		}
	}

	return nil
}

func (r *PingOneResourceScopeResource) addImportBlock(resourceId, resourceName, scopeId, scopeName string) {
	commentData := map[string]string{
		"Custom Resource ID":         resourceId,
		"Custom Resource Name":       resourceName,
		"Custom Resource Scope ID":   scopeId,
		"Custom Resource Scope Name": scopeName,
		"Export Environment ID":      r.clientInfo.ExportEnvironmentID,
		"Resource Type":              r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", resourceName, scopeName),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, resourceId, scopeId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
