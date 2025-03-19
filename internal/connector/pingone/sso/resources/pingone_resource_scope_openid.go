// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneResourceScopeOpenIdResource{}
)

type PingOneResourceScopeOpenIdResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneResourceScopeOpenIdResource
func ResourceScopeOpenId(clientInfo *connector.ClientInfo) *PingOneResourceScopeOpenIdResource {
	return &PingOneResourceScopeOpenIdResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceScopeOpenIdResource) ResourceType() string {
	return "pingone_resource_scope_openid"
}

func (r *PingOneResourceScopeOpenIdResource) ExportAll() (*[]connector.ImportBlock, error) {
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
				"Export Environment ID":              r.clientInfo.PingOneExportEnvironmentID,
				"OpenID Connect Resource Name":       resourceName,
				"OpenID Connect Resource Scope ID":   resourceScopeId,
				"OpenID Connect Resource Scope Name": resourceScopeName,
				"Resource Type":                      r.ResourceType(),
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

func (r *PingOneResourceScopeOpenIdResource) getResourceData() (map[string]string, error) {
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

			if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_OPENID_CONNECT {
				resourceData[*resourceId] = *resourceName
			}
		}
	}

	return resourceData, nil
}

func (r *PingOneResourceScopeOpenIdResource) getResourceScopeData(resourceId string) (map[string]string, error) {
	resourceScopeData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, resourceId).Execute()
	resourceScopes, err := pingone.GetManagementAPIObjectsFromIterator[management.ResourceScope](iter, "ReadAllResourceScopes", "GetScopes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, scopeOpenId := range resourceScopes {
		scopeOpenIdId, scopeOpenIdIdOk := scopeOpenId.GetIdOk()
		scopeOpenIdName, scopeOpenIdNameOk := scopeOpenId.GetNameOk()

		if scopeOpenIdIdOk && scopeOpenIdNameOk {
			resourceScopeData[*scopeOpenIdId] = *scopeOpenIdName
		}
	}

	return resourceScopeData, nil
}
