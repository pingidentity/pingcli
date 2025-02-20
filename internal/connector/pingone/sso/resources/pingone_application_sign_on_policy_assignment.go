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

func (r *PingOneApplicationSignOnPolicyAssignmentResource) ResourceType() string {
	return "pingone_application_sign_on_policy_assignment"
}

func (r *PingOneApplicationSignOnPolicyAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range applicationData {
		signOnPolicyAssignmentData, err := r.getApplicationSignOnPolicyAssignmentData(appId)
		if err != nil {
			return nil, err
		}

		for signOnPolicyAssignmentId, signOnPolicyId := range signOnPolicyAssignmentData {
			signOnPolicyName, signOnPolicyNameOk, err := r.getSignOnPolicyName(signOnPolicyId)
			if err != nil {
				return nil, err
			}
			if !signOnPolicyNameOk {
				continue
			}

			commentData := map[string]string{
				"Resource Type":    r.ResourceType(),
				"Application ID":   appId,
				"Application Name": appName,
				"Application Sign-On Policy Assignment ID": signOnPolicyAssignmentId,
				"Application Sign-On Policy Name":          signOnPolicyName,
				"Export Environment ID":                    r.clientInfo.ExportEnvironmentID,
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appName, signOnPolicyName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appId, signOnPolicyAssignmentId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationSignOnPolicyAssignmentResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	applications, err := common.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
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

func (r *PingOneApplicationSignOnPolicyAssignmentResource) getApplicationSignOnPolicyAssignmentData(appId string) (map[string]string, error) {
	signOnPolicyAssignmentData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationSignOnPolicyAssignmentsApi.ReadAllSignOnPolicyAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()
	signOnPolicyAssignments, err := common.GetManagementAPIObjectsFromIterator[management.SignOnPolicyAssignment](iter, "ReadAllSignOnPolicyAssignments", "GetSignOnPolicyAssignments", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, signOnPolicyAssignment := range signOnPolicyAssignments {
		signOnPolicyAssignmentId, signOnPolicyAssignmentIdOk := signOnPolicyAssignment.GetIdOk()
		signOnPolicyAssignmentSignOnPolicy, signOnPolicyAssignmentSignOnPolicyOk := signOnPolicyAssignment.GetSignOnPolicyOk()

		if signOnPolicyAssignmentIdOk && signOnPolicyAssignmentSignOnPolicyOk {
			signOnPolicyAssignmentSignOnPolicyId, signOnPolicyAssignmentSignOnPolicyIdOk := signOnPolicyAssignmentSignOnPolicy.GetIdOk()

			if signOnPolicyAssignmentSignOnPolicyIdOk {
				signOnPolicyAssignmentData[*signOnPolicyAssignmentId] = *signOnPolicyAssignmentSignOnPolicyId
			}
		}
	}

	return signOnPolicyAssignmentData, nil
}

func (r *PingOneApplicationSignOnPolicyAssignmentResource) getSignOnPolicyName(signOnPolicyId string) (string, bool, error) {
	signOnPolicy, response, err := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadOneSignOnPolicy(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, signOnPolicyId).Execute()
	ok, err := common.HandleClientResponse(response, err, "ReadOneSignOnPolicy", r.ResourceType())
	if err != nil {
		return "", false, err
	}
	if !ok {
		return "", false, nil
	}

	if signOnPolicy != nil {
		signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

		if signOnPolicyNameOk {
			return *signOnPolicyName, true, nil
		}
	}

	return "", false, fmt.Errorf("Unable to get sign-on policy name for sign-on policy ID: %s", signOnPolicyId)
}
