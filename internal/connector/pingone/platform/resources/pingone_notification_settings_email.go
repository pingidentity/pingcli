package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneNotificationSettingsEmailResource{}
)

type PingOneNotificationSettingsEmailResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneNotificationSettingsEmailResource
func NotificationSettingsEmail(clientInfo *connector.ClientInfo) *PingOneNotificationSettingsEmailResource {
	return &PingOneNotificationSettingsEmailResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneNotificationSettingsEmailResource) ResourceType() string {
	return "pingone_notification_settings_email"
}

func (r *PingOneNotificationSettingsEmailResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ok, err := r.checkNotificationSettingsEmailData()
	if err != nil {
		return nil, err
	}
	if !ok {
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.PingOneExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneNotificationSettingsEmailResource) checkNotificationSettingsEmailData() (bool, error) {
	_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsSettingsSMTPApi.ReadEmailNotificationsSettings(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	return common.CheckSingletonResource(response, err, "ReadEmailNotificationsSettings", r.ResourceType())
}
