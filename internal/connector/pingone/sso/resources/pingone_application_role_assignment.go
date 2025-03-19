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
	_ connector.ExportableResource = &PingOneApplicationRoleAssignmentResource{}
)

type PingOneApplicationRoleAssignmentResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneApplicationRoleAssignmentResource
func ApplicationRoleAssignment(clientInfo *connector.ClientInfo) *PingOneApplicationRoleAssignmentResource {
	return &PingOneApplicationRoleAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationRoleAssignmentResource) ResourceType() string {
	return "pingone_application_role_assignment"
}

func (r *PingOneApplicationRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range applicationData {
		applicationRoleAssignmentData, err := r.getApplicationRoleAssignmentData(appId)
		if err != nil {
			return nil, err
		}

		for roleAssignmentId, roleId := range applicationRoleAssignmentData {
			roleName, roleNameOk, err := r.getRoleName(roleId)
			if err != nil {
				return nil, err
			}
			if !roleNameOk {
				continue
			}

			commentData := map[string]string{
				"Application ID":                 appId,
				"Application Name":               appName,
				"Application Role Assignment ID": roleAssignmentId,
				"Application Role Name":          string(roleName),
				"Export Environment ID":          r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                  r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s_%s", appName, string(roleName), roleAssignmentId),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, appId, roleAssignmentId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationRoleAssignmentResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	applications, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, app := range applications {
		var (
			appId                  *string
			appIdOk                bool
			appName                *string
			appNameOk              bool
			appAccessControlRole   *management.ApplicationAccessControlRole
			appAccessControlRoleOk bool
		)

		switch {
		case app.ApplicationOIDC != nil:
			appId, appIdOk = app.ApplicationOIDC.GetIdOk()
			appName, appNameOk = app.ApplicationOIDC.GetNameOk()
			if app.ApplicationOIDC.AccessControl != nil {
				appAccessControlRole, appAccessControlRoleOk = app.ApplicationOIDC.AccessControl.GetRoleOk()
			}
		case app.ApplicationSAML != nil:
			appId, appIdOk = app.ApplicationSAML.GetIdOk()
			appName, appNameOk = app.ApplicationSAML.GetNameOk()
			if app.ApplicationSAML.AccessControl != nil {
				appAccessControlRole, appAccessControlRoleOk = app.ApplicationSAML.AccessControl.GetRoleOk()
			}
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
			if app.ApplicationExternalLink.AccessControl != nil {
				appAccessControlRole, appAccessControlRoleOk = app.ApplicationExternalLink.AccessControl.GetRoleOk()
			}
		default:
			continue
		}

		if appIdOk && appNameOk && appAccessControlRoleOk {
			if appAccessControlRole.GetType() != management.ENUMAPPLICATIONACCESSCONTROLTYPE_ADMIN_USERS_ONLY {
				continue
			}

			applicationData[*appId] = *appName
		}
	}

	return applicationData, nil
}

func (r *PingOneApplicationRoleAssignmentResource) getApplicationRoleAssignmentData(appId string) (map[string]string, error) {
	applicationRoleAssignmentData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, appId).Execute()
	applicationRoleAssignments, err := pingone.GetManagementAPIObjectsFromIterator[management.RoleAssignment](iter, "ReadApplicationRoleAssignments", "GetRoleAssignments", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, roleAssignment := range applicationRoleAssignments {
		roleAssignmentId, roleAssignmentIdOk := roleAssignment.GetIdOk()
		roleAssignmentRole, roleAssignmentRoleOk := roleAssignment.GetRoleOk()

		if roleAssignmentIdOk && roleAssignmentRoleOk {
			roleAssignmentRoleId, roleAssignmentRoleIdOk := roleAssignmentRole.GetIdOk()

			if roleAssignmentRoleIdOk {
				applicationRoleAssignmentData[*roleAssignmentId] = *roleAssignmentRoleId
			}
		}
	}

	return applicationRoleAssignmentData, nil
}

func (r *PingOneApplicationRoleAssignmentResource) getRoleName(roleId string) (management.EnumRoleName, bool, error) {
	apiRole, resp, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.PingOneContext, roleId).Execute()
	ok, err := common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
	if err != nil {
		return "", false, err
	}
	if !ok {
		return "", false, nil
	}

	if apiRole != nil {
		apiRoleName, apiRoleNameOk := apiRole.GetNameOk()
		if apiRoleNameOk {
			return *apiRoleName, true, nil
		}
	}

	return "", false, fmt.Errorf("unable to get role name for role ID: %s", roleId)
}
