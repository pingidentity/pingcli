// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeApplicationResourcePermissionResource{}
)

type PingoneAuthorizeApplicationResourcePermissionResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingoneAuthorizeApplicationResourcePermissionResource
func AuthorizeApplicationResourcePermission(clientInfo *connector.ClientInfo) *PingoneAuthorizeApplicationResourcePermissionResource {
	return &PingoneAuthorizeApplicationResourcePermissionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeApplicationResourcePermissionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationResourceData, err := r.getApplicationResourceData()
	if err != nil {
		return nil, err
	}

	for appResourceId, appResourceName := range applicationResourceData {
		appResourcePermissionData, err := r.getApplicationResourcePermissionData(appResourceId)
		if err != nil {
			return nil, err
		}

		for appResourcePermissionId, appResourcePermissionKey := range appResourcePermissionData {
			commentData := map[string]string{
				"Application Resource ID":             appResourceId,
				"Application Resource Name":           appResourceName,
				"Application Resource Permission ID":  appResourcePermissionId,
				"Application Resource Permission Key": appResourcePermissionKey,
				"Export Environment ID":               r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                       r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       appResourcePermissionKey,
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, appResourceId, appResourcePermissionId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeApplicationResourcePermissionResource) getApplicationResourceData() (map[string]string, error) {
	applicationResourceData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationResourcesApi.ReadApplicationResources(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	applicationResources, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.ApplicationResource](iter, "ReadApplicationResources", "GetResources", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationResource := range applicationResources {
		applicationResourceId, applicationResourceIdOk := applicationResource.GetIdOk()
		applicationResourceName, applicationResourceNameOk := applicationResource.GetNameOk()

		if applicationResourceIdOk && applicationResourceNameOk {
			applicationResourceData[*applicationResourceId] = *applicationResourceName
		}
	}

	return applicationResourceData, nil
}

func (r *PingoneAuthorizeApplicationResourcePermissionResource) getApplicationResourcePermissionData(appResourceId string) (map[string]string, error) {
	applicationResourcePermissionData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationResourcePermissionsApi.ReadApplicationPermissions(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, appResourceId).Execute()
	applicationResourcePermissions, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.EntityArrayEmbeddedPermissionsInner](iter, "ReadApplicationPermissions", "GetPermissions", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationResourcePermission := range applicationResourcePermissions {

		var (
			applicationResourcePermissionId    *string
			applicationResourcePermissionIdOk  bool
			applicationResourcePermissionKey   *string
			applicationResourcePermissionKeyOk bool
		)

		switch t := applicationResourcePermission.GetActualInstance().(type) {
		case *authorize.ApplicationResourcePermission:
			applicationResourcePermissionId, applicationResourcePermissionIdOk = t.GetIdOk()
		case *authorize.ApplicationRolePermission:
			applicationResourcePermissionId, applicationResourcePermissionIdOk = t.GetIdOk()
			applicationResourcePermissionKey, applicationResourcePermissionKeyOk = t.GetKeyOk()
		default:
			continue
		}

		if applicationResourcePermissionIdOk && applicationResourcePermissionKeyOk {
			applicationResourcePermissionData[*applicationResourcePermissionId] = *applicationResourcePermissionKey
		}

		if applicationResourcePermissionIdOk && !applicationResourcePermissionKeyOk {
			applicationResourcePermissionData[*applicationResourcePermissionId] = *applicationResourcePermissionId
		}
	}

	return applicationResourcePermissionData, nil
}

func (r *PingoneAuthorizeApplicationResourcePermissionResource) ResourceType() string {
	return "pingone_application_resource_permission"
}
