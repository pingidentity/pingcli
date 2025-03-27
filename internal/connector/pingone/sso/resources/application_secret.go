// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationSecretResource{}
)

type PingOneApplicationSecretResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneApplicationSecretResource
func ApplicationSecret(clientInfo *connector.ClientInfo) *PingOneApplicationSecretResource {
	return &PingOneApplicationSecretResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationSecretResource) ResourceType() string {
	return "pingone_application_secret"
}

func (r *PingOneApplicationSecretResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for applicationId, applicationName := range applicationData {
		ok, err := r.checkApplicationSecretData(applicationId)
		if err != nil {
			return nil, err
		}

		if !ok {
			continue
		}

		commentData := map[string]string{
			"Application ID":        applicationId,
			"Application Name":      applicationName,
			"Resource Type":         r.ResourceType(),
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_secret", applicationName),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, applicationId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationSecretResource) getApplicationData() (map[string]string, error) {
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

func (r *PingOneApplicationSecretResource) checkApplicationSecretData(applicationId string) (b bool, err error) {
	// The platform enforces that worker apps cannot read their own secret
	// Make sure we can read the secret before adding it to the import blocks
	_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationSecretApi.ReadApplicationSecret(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, applicationId).Execute()
	defer func() {
		cErr := response.Body.Close()
		if cErr != nil {
			err = errors.Join(err, cErr)
		}
	}()

	// If the appId is the same as the worker ID, make sure the API response is a 403 and ignore the error
	if applicationId == r.clientInfo.PingOneApiClientId {
		if response.StatusCode == http.StatusForbidden {
			return false, nil
		} else {
			return false, fmt.Errorf("error: Expected 403 Forbidden response - worker apps cannot read their own secret\n%s Response Code: %s\nResponse Body: %s", "ReadApplicationSecret", response.Status, response.Body)
		}
	}

	// Use output package to warn the user of any errors or non-200 response codes
	// Expected behavior in this case is to skip the resource, and continue exporting the other resources
	if err != nil || response.StatusCode >= 300 || response.StatusCode < 200 {
		output.Warn("Failed to read secret for application", map[string]interface{}{
			"Application ID":    applicationId,
			"API Function Name": "ReadApplicationSecret",
			"Response Code":     response.Status,
			"Response Body":     response.Body,
		})

		return false, nil
	}

	return true, nil
}
