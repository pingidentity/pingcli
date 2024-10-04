package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationFlowPolicyAssignmentResource{}
)

type PingoneApplicationFlowPolicyAssignmentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneApplicationFlowPolicyAssignmentResource
func ApplicationFlowPolicyAssignment(clientInfo *connector.PingOneClientInfo) *PingoneApplicationFlowPolicyAssignmentResource {
	return &PingoneApplicationFlowPolicyAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationFlowPolicyAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
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
			apiExecutePoliciesFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiPolicyFunctionName := "ReadAllFlowPolicyAssignments"

			policyEmbedded, err := common.GetManagementEmbedded(apiExecutePoliciesFunc, apiPolicyFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, flowPolicyAssignment := range policyEmbedded.GetFlowPolicyAssignments() {
				flowPolicyAssignmentId, flowPolicyAssignmentIdOk := flowPolicyAssignment.GetIdOk()
				flowPolicyAssignmentFlowPolicy, flowPolicyAssignmentFlowPolicyOk := flowPolicyAssignment.GetFlowPolicyOk()

				var (
					flowPolicyId   *string
					flowPolicyIdOk bool
				)

				if flowPolicyAssignmentFlowPolicyOk {
					flowPolicyId, flowPolicyIdOk = flowPolicyAssignmentFlowPolicy.GetIdOk()
				}

				if flowPolicyAssignmentIdOk && flowPolicyAssignmentFlowPolicyOk && flowPolicyIdOk {
					flowPolicy, response, err := r.clientInfo.ApiClient.ManagementAPIClient.FlowPoliciesApi.ReadOneFlowPolicy(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *flowPolicyId).Execute()
					err = common.HandleClientResponse(response, err, "ReadOneFlowPolicy", r.ResourceType())
					if err != nil {
						return nil, err
					}

					if flowPolicy != nil {
						flowPolicyName, flowPolicyNameOk := flowPolicy.GetNameOk()
						if flowPolicyNameOk {
							commentData := map[string]string{
								"Resource Type":             r.ResourceType(),
								"Application Name":          *appName,
								"Flow Policy Name":          *flowPolicyName,
								"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
								"Application ID":            *appId,
								"Flow Policy Assignment ID": *flowPolicyAssignmentId,
							}

							importBlocks = append(importBlocks, connector.ImportBlock{
								ResourceType:       r.ResourceType(),
								ResourceName:       fmt.Sprintf("%s_%s", *appName, *flowPolicyName),
								ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *flowPolicyAssignmentId),
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

func (r *PingoneApplicationFlowPolicyAssignmentResource) ResourceType() string {
	return "pingone_application_flow_policy_assignment"
}
