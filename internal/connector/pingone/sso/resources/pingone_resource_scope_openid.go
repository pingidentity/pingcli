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
	_ connector.ExportableResource = &PingOneResourceScopeOpenIdResource{}
)

type PingOneResourceScopeOpenIdResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneResourceScopeOpenIdResource
func ResourceScopeOpenId(clientInfo *connector.PingOneClientInfo) *PingOneResourceScopeOpenIdResource {
	return &PingOneResourceScopeOpenIdResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneResourceScopeOpenIdResource) ResourceType() string {
	return "pingone_resource_scope_openid"
}

func (r *PingOneResourceScopeOpenIdResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportResourceScopeOpenIds()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneResourceScopeOpenIdResource) exportResourceScopeOpenIds() error {
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

				if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_OPENID_CONNECT {
					err := r.exportResourceScopeOpenIdByResource(*resourceId, *resourceName)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneResourceScopeOpenIdResource) exportResourceScopeOpenIdByResource(resourceId, resourceName string) error {
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

		for _, scopeOpenId := range embedded.GetScopes() {
			scopeOpenIdId, scopeOpenIdIdOk := scopeOpenId.GetIdOk()
			scopeOpenIdName, scopeOpenIdNameOk := scopeOpenId.GetNameOk()
			if scopeOpenIdIdOk && scopeOpenIdNameOk {
				r.addImportBlock(resourceName, *scopeOpenIdId, *scopeOpenIdName)
			}
		}
	}

	return nil
}

func (r *PingOneResourceScopeOpenIdResource) addImportBlock(resourceName, scopeOpenIdId, scopeOpenIdName string) {
	commentData := map[string]string{
		"Export Environment ID":              r.clientInfo.ExportEnvironmentID,
		"OpenID Connect Resource Name":       resourceName,
		"OpenID Connect Resource Scope ID":   scopeOpenIdId,
		"OpenID Connect Resource Scope Name": scopeOpenIdName,
		"Resource Type":                      r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", resourceName, scopeOpenIdName),
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, scopeOpenIdId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
