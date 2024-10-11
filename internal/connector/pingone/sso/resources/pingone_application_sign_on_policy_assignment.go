package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationSignOnPolicyAssignmentResource{}
)

type PingOneApplicationSignOnPolicyAssignmentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneApplicationSignOnPolicyAssignmentResource
func ApplicationSignOnPolicyAssignment(clientInfo *connector.PingOneClientInfo) *PingOneApplicationSignOnPolicyAssignmentResource {
	return &PingOneApplicationSignOnPolicyAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationSignOnPolicyAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteApplicationsFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiApplicationFunctionName := "ReadAllApplications"

	embedded, err := common.GetManagementEmbedded(apiExecuteApplicationsFunc, apiApplicationFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, app := range embedded.GetApplications() {
		var (
			appId     *string
			appIdOk   bool
			appName   *string
			appNameOk bool
		)

		switch {
		case app.ApplicationOIDC != nil:
			appId, appIdOk = app.ApplicationOIDC.GetIdOk()
			appName, appNameOk = app.ApplicationOIDC.GetNameOk()
		case app.ApplicationSAML != nil:
			appId, appIdOk = app.ApplicationSAML.GetIdOk()
			appName, appNameOk = app.ApplicationSAML.GetNameOk()
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			apiExecutePoliciesFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationSignOnPolicyAssignmentsApi.ReadAllSignOnPolicyAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiPolicyFunctionName := "ReadAllSignOnPolicyAssignments"

			signOnPolicyEmbedded, err := common.GetManagementEmbedded(apiExecutePoliciesFunc, apiPolicyFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, signOnPolicyAssignment := range signOnPolicyEmbedded.GetSignOnPolicyAssignments() {
				signOnPolicyAssignmentId, signOnPolicyAssignmentIdOk := signOnPolicyAssignment.GetIdOk()
				signOnPolicyAssignmentSignOnPolicy, signOnPolicyAssignmentSignOnPolicyOk := signOnPolicyAssignment.GetSignOnPolicyOk()

				var (
					signOnPolicyAssignmentSignOnPolicyId   *string
					signOnPolicyAssignmentSignOnPolicyIdOk bool
				)

				if signOnPolicyAssignmentSignOnPolicyOk {
					signOnPolicyAssignmentSignOnPolicyId, signOnPolicyAssignmentSignOnPolicyIdOk = signOnPolicyAssignmentSignOnPolicy.GetIdOk()
				}

				if signOnPolicyAssignmentIdOk && signOnPolicyAssignmentSignOnPolicyOk && signOnPolicyAssignmentSignOnPolicyIdOk {
					signOnPolicy, response, err := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadOneSignOnPolicy(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *signOnPolicyAssignmentSignOnPolicyId).Execute()
					err = common.HandleClientResponse(response, err, "ReadOneSignOnPolicy", r.ResourceType())
					if err != nil {
						return nil, err
					}

					if signOnPolicy != nil {
						signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

						if signOnPolicyNameOk {
							commentData := map[string]string{
								"Resource Type":                r.ResourceType(),
								"Application Name":             *appName,
								"Sign On Policy Name":          *signOnPolicyName,
								"Export Environment ID":        r.clientInfo.ExportEnvironmentID,
								"Application ID":               *appId,
								"Sign On Policy Assignment ID": *signOnPolicyAssignmentId,
							}

							importBlocks = append(importBlocks, connector.ImportBlock{
								ResourceType:       r.ResourceType(),
								ResourceName:       fmt.Sprintf("%s_%s", *appName, *signOnPolicyName),
								ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *signOnPolicyAssignmentId),
								CommentInformation: common.GenerateCommentInformation(commentData),
							})
						}
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationSignOnPolicyAssignmentResource) ResourceType() string {
	return "pingone_application_sign_on_policy_assignment"
}
