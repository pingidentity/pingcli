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
	_ connector.ExportableResource = &PingOneSystemApplicationResource{}
)

type PingOneSystemApplicationResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneSystemApplicationResource
func SystemApplication(clientInfo *connector.ClientInfo) *PingOneSystemApplicationResource {
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

	for appId, appName := range applicationData {
		commentData := map[string]string{
			"Export Environment ID":   r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":           r.ResourceType(),
			"System Application ID":   appId,
			"System Application Name": appName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       appName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, appId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneSystemApplicationResource) getSystemApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}
	for _, application := range apiObjs {
		var (
			applicationId     *string
			applicationIdOk   bool
			applicationName   *string
			applicationNameOk bool
		)

		switch {
		case application.ApplicationPingOnePortal != nil:
			applicationId, applicationIdOk = application.ApplicationPingOnePortal.GetIdOk()
			applicationName, applicationNameOk = application.ApplicationPingOnePortal.GetNameOk()
		case application.ApplicationPingOneSelfService != nil:
			applicationId, applicationIdOk = application.ApplicationPingOneSelfService.GetIdOk()
			applicationName, applicationNameOk = application.ApplicationPingOneSelfService.GetNameOk()
		default:
			continue
		}

		if applicationIdOk && applicationNameOk {
			applicationData[*applicationId] = *applicationName
		}
	}

	return applicationData, nil
}
