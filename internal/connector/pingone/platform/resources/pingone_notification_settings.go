package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneNotificationSettingsResource{}
)

type PingoneNotificationSettingsResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneNotificationSettingsResource
func NotificationSettings(clientInfo *connector.PingOneClientInfo) *PingoneNotificationSettingsResource {
	return &PingoneNotificationSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneNotificationSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsSettingsApi.ReadNotificationsSettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadNotificationsSettings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	if response.StatusCode == 204 {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       "notification_settings",
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingoneNotificationSettingsResource) ResourceType() string {
	return "pingone_notification_settings"
}
