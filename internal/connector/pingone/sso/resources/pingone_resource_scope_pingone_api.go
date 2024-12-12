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

	for resourceId, resourceName := range *resourceData {
		resourceScopeData, err := r.getResourceScopeData(resourceId)
		if err != nil {
			return nil, err
		}

		for resourceScopeId, resourceScopeName := range *resourceScopeData {
			commentData := map[string]string{
				"Export Environment ID":           r.clientInfo.ExportEnvironmentID,
				"PingOne API Resource Name":       resourceName,
				"PingOne API Resource Scope ID":   resourceScopeId,
				"PingOne API Resource Scope Name": resourceScopeName,
				"Resource Type":                   r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", resourceName, resourceScopeName),
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, resourceScopeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneResourceScopePingOneApiResource) getResourceData() (*map[string]string, error) {
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

				if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_PINGONE_API {
					resourceData[*resourceId] = *resourceName
				}
			}
		}
	}

	return &resourceData, nil
}

func (r *PingOneResourceScopePingOneApiResource) getResourceScopeData(resourceId string) (*map[string]string, error) {
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

		for _, scopePingOneApi := range embedded.GetScopes() {
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
	}

	return &resourceScopeData, nil
}
