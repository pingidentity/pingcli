package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneNotificationSettingsResource{}
)

type PingOneNotificationSettingsResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneNotificationSettingsResource
func NotificationSettings(clientInfo *connector.PingOneClientInfo) *PingOneNotificationSettingsResource {
	return &PingOneNotificationSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneNotificationSettingsResource) ResourceType() string {
	return "pingone_notification_settings"
}

func (r *PingOneNotificationSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ok, err := r.checkNotificationSettingsData()
	if err != nil {
		return nil, err
	}
	if !ok {
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneNotificationSettingsResource) checkNotificationSettingsData() (bool, error) {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsSettingsApi.ReadNotificationsSettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	return pingone.CheckSingletonResource(response, err, "ReadNotificationsSettings", r.ResourceType())
}
