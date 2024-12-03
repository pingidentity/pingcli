package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationResource{}
)

type PingOneApplicationResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneApplicationResource
func Application(clientInfo *connector.PingOneClientInfo) *PingOneApplicationResource {
	return &PingOneApplicationResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneApplicationResource) ResourceType() string {
	return "pingone_application"
}

func (r *PingOneApplicationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportApplications()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneApplicationResource) exportApplications() error {
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
				r.addImportBlock(*appId, *appName)
			}
		}
	}

	return nil
}

func (r *PingOneApplicationResource) addImportBlock(appId, appName string) {
	commentData := map[string]string{
		"Application ID":        appId,
		"Application Name":      appName,
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       appName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, appId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
