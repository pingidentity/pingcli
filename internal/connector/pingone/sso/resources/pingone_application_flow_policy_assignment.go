package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationFlowPolicyAssignmentResource{}
)

type PingOneApplicationFlowPolicyAssignmentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneApplicationFlowPolicyAssignmentResource
func ApplicationFlowPolicyAssignment(clientInfo *connector.PingOneClientInfo) *PingOneApplicationFlowPolicyAssignmentResource {
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

	for appId, appName := range *applicationData {
		flowPolicyAssignmentData, err := r.getFlowPolicyAssignmentData(appId)
		if err != nil {
			return nil, err
		}

		for flowPolicyAssignmentId, flowPolicyId := range *flowPolicyAssignmentData {
			flowPolicyName, err := r.getFlowPolicyName(flowPolicyId)
			if err != nil {
				return nil, err
			}

			commentData := map[string]string{
				"Application ID":            appId,
				"Application Name":          appName,
				"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
				"Flow Policy Assignment ID": flowPolicyAssignmentId,
				"Flow Policy Name":          *flowPolicyName,
				"Resource Type":             r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, *flowPolicyName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, flowPolicyAssignmentId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) getApplicationData() (*map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllApplications", r.ResourceType())
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
				applicationData[*appId] = *appName
			}
		}
	}

	return &applicationData, nil
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) getFlowPolicyAssignmentData(appId string) (*map[string]string, error) {
	flowPolicyAssignmentData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllFlowPolicyAssignments", r.ResourceType())
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

		for _, flowPolicyAssignment := range embedded.GetFlowPolicyAssignments() {
			flowPolicyAssignmentId, flowPolicyAssignmentIdOk := flowPolicyAssignment.GetIdOk()
			flowPolicyAssignmentFlowPolicy, flowPolicyAssignmentFlowPolicyOk := flowPolicyAssignment.GetFlowPolicyOk()

			if flowPolicyAssignmentIdOk && flowPolicyAssignmentFlowPolicyOk {
				flowPolicyId, flowPolicyIdOk := flowPolicyAssignmentFlowPolicy.GetIdOk()

				if flowPolicyIdOk {
					flowPolicyAssignmentData[*flowPolicyAssignmentId] = *flowPolicyId
				}
			}
		}
	}

	return &flowPolicyAssignmentData, nil
}

func (r *PingOneApplicationFlowPolicyAssignmentResource) getFlowPolicyName(flowPolicyId string) (*string, error) {
	flowPolicy, response, err := r.clientInfo.ApiClient.ManagementAPIClient.FlowPoliciesApi.ReadOneFlowPolicy(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, flowPolicyId).Execute()

	err = common.HandleClientResponse(response, err, "ReadOneFlowPolicy", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if flowPolicy != nil {
		flowPolicyName, flowPolicyNameOk := flowPolicy.GetNameOk()

		if flowPolicyNameOk {
			return flowPolicyName, nil
		}
	}

	return nil, fmt.Errorf("Unable to get Flow Policy Name for Flow Policy ID: %s", flowPolicyId)
}
