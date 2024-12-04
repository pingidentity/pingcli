package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationSecretResource{}
)

type PingOneApplicationSecretResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneApplicationSecretResource
func ApplicationSecret(clientInfo *connector.PingOneClientInfo) *PingOneApplicationSecretResource {
	return &PingOneApplicationSecretResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneApplicationSecretResource) ResourceType() string {
	return "pingone_application_secret"
}

func (r *PingOneApplicationSecretResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportApplicationSecrets()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneApplicationSecretResource) exportApplicationSecrets() error {
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
				err := r.exportApplicationSecretsByApplication(*appId, *appName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneApplicationSecretResource) exportApplicationSecretsByApplication(appId, appName string) error {
	l := logger.Get()

	// The platform enforces that worker apps cannot read their own secret
	// Make sure we can read the secret before adding it to the import blocks
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationSecretApi.ReadApplicationSecret(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()

	// If the appId is the same as the worker ID, make sure the API response is a 403 and ignore the error
	if appId == *r.clientInfo.ApiClientId {
		if response.StatusCode == 403 {
			return nil
		} else {
			return fmt.Errorf("ReadApplicationSecret: Expected response code 403 - worker apps cannot read their own secret, actual response code: %d", response.StatusCode)
		}
	}

	// Use output package to warn the user of any errors or non-200 response codes
	// Expected behavior in this case is to skip the resource, and continue exporting the other resources
	defer response.Body.Close()

	if err != nil {
		l.Warn().Err(err).Msgf("Failed to read secret for application %s. %s Response Code: %s\nResponse Body: %s", appName, "ReadApplicationSecret", response.Status, response.Body)
		return nil
	}

	if response.StatusCode >= 300 {
		l.Warn().Msgf("Failed to read secret for application %s. %s Response Code: %s\nResponse Body: %s", appName, "ReadApplicationSecret", response.Status, response.Body)
		return nil
	}

	r.addImportBlock(appId, appName)

	return nil
}

func (r *PingOneApplicationSecretResource) addImportBlock(appId, appName string) {
	commentData := map[string]string{
		"Application ID":        appId,
		"Application Name":      appName,
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_secret", appName),
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, appId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
