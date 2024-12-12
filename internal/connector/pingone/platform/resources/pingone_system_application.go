package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneSystemApplicationResource{}
)

type PingOneSystemApplicationResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneSystemApplicationResource
func SystemApplication(clientInfo *connector.PingOneClientInfo) *PingOneSystemApplicationResource {
	return &PingOneSystemApplicationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneSystemApplicationResource) ResourceType() string {
	return "pingone_system_application"
}

func (r *PingOneSystemApplicationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getSystemApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range *applicationData {
		commentData := map[string]string{
			"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
			"Resource Type":           r.ResourceType(),
			"System Application ID":   appId,
			"System Application Name": appName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       appName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, appId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneSystemApplicationResource) getSystemApplicationData() (*map[string]string, error) {
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
			case app.ApplicationPingOnePortal != nil:
				appId, appIdOk = app.ApplicationPingOnePortal.GetIdOk()
				appName, appNameOk = app.ApplicationPingOnePortal.GetNameOk()
			case app.ApplicationPingOneSelfService != nil:
				appId, appIdOk = app.ApplicationPingOneSelfService.GetIdOk()
				appName, appNameOk = app.ApplicationPingOneSelfService.GetNameOk()
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
