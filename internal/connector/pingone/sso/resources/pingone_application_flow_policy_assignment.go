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
	_ connector.ExportableResource = &PingOneApplicationFlowPolicyAssignmentResource{}
)

type PingOneApplicationFlowPolicyAssignmentResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneApplicationFlowPolicyAssignmentResource
func ApplicationFlowPolicyAssignment(clientInfo *connector.ClientInfo) *PingOneApplicationFlowPolicyAssignmentResource {
	return &PingOneApplicationFlowPolicyAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) ResourceType() string {
	return "pingone_application_flow_policy_assignment"
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range applicationData {
		flowPolicyAssignmentData, err := r.getFlowPolicyAssignmentData(appId)
		if err != nil {
			return nil, err
		}

		for flowPolicyAssignmentId, flowPolicyId := range flowPolicyAssignmentData {
			flowPolicyName, flowPolicyNameOk, err := r.getFlowPolicyName(flowPolicyId)
			if err != nil {
				return nil, err
			}
			if !flowPolicyNameOk {
				continue
			}

			commentData := map[string]string{
				"Application ID":            appId,
				"Application Name":          appName,
				"Export Environment ID":     r.clientInfo.PingOneExportEnvironmentID,
				"Flow Policy Assignment ID": flowPolicyAssignmentId,
				"Flow Policy Name":          flowPolicyName,
				"Resource Type":             r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, flowPolicyName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, appId, flowPolicyAssignmentId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	applications, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, app := range applications {
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
			applicationData[*appId] = *appName
		}
	}

	return applicationData, nil
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) getFlowPolicyAssignmentData(appId string) (map[string]string, error) {
	flowPolicyAssignmentData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, appId).Execute()
	flowPolicyAssignments, err := pingone.GetManagementAPIObjectsFromIterator[management.FlowPolicyAssignment](iter, "ReadAllFlowPolicyAssignments", "GetFlowPolicyAssignments", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, flowPolicyAssignment := range flowPolicyAssignments {
		flowPolicyAssignmentId, flowPolicyAssignmentIdOk := flowPolicyAssignment.GetIdOk()
		flowPolicyAssignmentFlowPolicy, flowPolicyAssignmentFlowPolicyOk := flowPolicyAssignment.GetFlowPolicyOk()

		if flowPolicyAssignmentIdOk && flowPolicyAssignmentFlowPolicyOk {
			flowPolicyId, flowPolicyIdOk := flowPolicyAssignmentFlowPolicy.GetIdOk()

			if flowPolicyIdOk {
				flowPolicyAssignmentData[*flowPolicyAssignmentId] = *flowPolicyId
			}
		}
	}

	return flowPolicyAssignmentData, nil
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) getFlowPolicyName(flowPolicyId string) (string, bool, error) {
	flowPolicy, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.FlowPoliciesApi.ReadOneFlowPolicy(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, flowPolicyId).Execute()

	ok, err := common.HandleClientResponse(response, err, "ReadOneFlowPolicy", r.ResourceType())
	if err != nil {
		return "", false, err
	}
	if !ok {
		return "", false, nil
	}

	if flowPolicy != nil {
		flowPolicyName, flowPolicyNameOk := flowPolicy.GetNameOk()

		if flowPolicyNameOk {
			return *flowPolicyName, true, nil
		}
	}

	return "", false, fmt.Errorf("unable to get Flow Policy Name for Flow Policy ID: %s", flowPolicyId)
}
