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
	_ connector.ExportableResource = &PingOneApplicationRoleAssignmentResource{}
)

type PingOneApplicationRoleAssignmentResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneApplicationRoleAssignmentResource
func ApplicationRoleAssignment(clientInfo *connector.PingOneClientInfo) *PingOneApplicationRoleAssignmentResource {
	return &PingOneApplicationRoleAssignmentResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneApplicationRoleAssignmentResource) ResourceType() string {
	return "pingone_application_role_assignment"
}

func (r *PingOneApplicationRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportApplicationRoleAssignments()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneApplicationRoleAssignmentResource) exportApplicationRoleAssignments() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplications", r.ResourceType())
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

		for _, app := range embedded.GetApplications() {
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

				err := r.exportApplicationRoleAssignmentsByApplication(*appId, *appName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneApplicationRoleAssignmentResource) exportApplicationRoleAssignmentsByApplication(appId, appName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadApplicationRoleAssignments", r.ResourceType())
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

		for _, roleAssignment := range embedded.GetRoleAssignments() {
			roleAssignmentId, roleAssignmentIdOk := roleAssignment.GetIdOk()
			roleAssignmentRole, roleAssignmentRoleOk := roleAssignment.GetRoleOk()
			if roleAssignmentIdOk && roleAssignmentRoleOk {
				roleAssignmentRoleId, roleAssignmentRoleIdOk := roleAssignmentRole.GetIdOk()
				if roleAssignmentRoleIdOk {
					err := r.exportApplicationRoleAssignmentsByRole(appId, appName, *roleAssignmentId, *roleAssignmentRoleId)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneApplicationRoleAssignmentResource) exportApplicationRoleAssignmentsByRole(appId, appName, roleAssignmentId, roleId string) error {
	apiRole, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, roleId).Execute()
	err = common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
	if err != nil {
		return err
	}

	if apiRole != nil {
		apiRoleName, apiRoleNameOk := apiRole.GetNameOk()
		if apiRoleNameOk {
			r.addImportBlock(appId, appName, roleAssignmentId, string(*apiRoleName))
		}
	}

	return nil
}

func (r *PingOneApplicationRoleAssignmentResource) addImportBlock(appId, appName, roleAssignmentId, roleName string) {
	commentData := map[string]string{
		"Application ID":                 appId,
		"Application Name":               appName,
		"Application Role Assignment ID": roleAssignmentId,
		"Application Role Name":          roleName,
		"Export Environment ID":          r.clientInfo.ExportEnvironmentID,
		"Resource Type":                  r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s_%s", appName, roleName, roleAssignmentId),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, roleAssignmentId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
