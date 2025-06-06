// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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

	for applicationId, applicationName := range applicationData {
		applicationFlowPolicyAssignmentData, err := r.getApplicationFlowPolicyAssignmentData(applicationId)
		if err != nil {
			return nil, err
		}

		for applicationFlowPolicyAssignmentId, applicationFlowPolicyAssignmentFlowPolicyId := range applicationFlowPolicyAssignmentData {
			flowPolicyName, flowPolicyNameOk, err := r.getFlowPolicyName(applicationFlowPolicyAssignmentFlowPolicyId)
			if err != nil {
				return nil, err
			}

			if !flowPolicyNameOk {
				continue
			}

			commentData := map[string]string{
				"Application ID":                                      applicationId,
				"Application Name":                                    applicationName,
				"Application Flow Policy Assignment ID":               applicationFlowPolicyAssignmentId,
				"Application Flow Policy Assignment Flow Policy Name": flowPolicyName,
				"Export Environment ID":                               r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                                       r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", applicationName, flowPolicyName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, applicationId, applicationFlowPolicyAssignmentId),
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
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		var (
			applicationId     *string
			applicationIdOk   bool
			applicationName   *string
			applicationNameOk bool
		)

		switch {
		case innerObj.ApplicationOIDC != nil:
			applicationId, applicationIdOk = innerObj.ApplicationOIDC.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationOIDC.GetNameOk()
		case innerObj.ApplicationSAML != nil:
			applicationId, applicationIdOk = innerObj.ApplicationSAML.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationSAML.GetNameOk()
		case innerObj.ApplicationExternalLink != nil:
			applicationId, applicationIdOk = innerObj.ApplicationExternalLink.GetIdOk()
			applicationName, applicationNameOk = innerObj.ApplicationExternalLink.GetNameOk()
		default:
			continue
		}

		if applicationIdOk && applicationNameOk {
			applicationData[*applicationId] = *applicationName
		}
	}

	return applicationData, nil
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) getApplicationFlowPolicyAssignmentData(applicationId string) (map[string]string, error) {
	applicationFlowPolicyAssignmentData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, applicationId).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.FlowPolicyAssignment](iter, "ReadAllFlowPolicyAssignments", "GetFlowPolicyAssignments", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationFlowPolicyAssignment := range apiObjs {
		applicationFlowPolicyAssignmentId, applicationFlowPolicyAssignmentIdOk := applicationFlowPolicyAssignment.GetIdOk()
		applicationFlowPolicyAssignmentFlowPolicy, applicationFlowPolicyAssignmentFlowPolicyOk := applicationFlowPolicyAssignment.GetFlowPolicyOk()

		if applicationFlowPolicyAssignmentIdOk && applicationFlowPolicyAssignmentFlowPolicyOk {
			applicationFlowPolicyAssignmentFlowPolicyId, applicationFlowPolicyAssignmentFlowPolicyIdOk := applicationFlowPolicyAssignmentFlowPolicy.GetIdOk()

			if applicationFlowPolicyAssignmentFlowPolicyIdOk {
				applicationFlowPolicyAssignmentData[*applicationFlowPolicyAssignmentId] = *applicationFlowPolicyAssignmentFlowPolicyId
			}
		}
	}

	return applicationFlowPolicyAssignmentData, nil
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
