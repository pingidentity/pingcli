package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneNotificationSettingsEmailResource{}
)

type PingoneNotificationSettingsEmailResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneNotificationSettingsEmailResource
func NotificationSettingsEmail(clientInfo *connector.PingOneClientInfo) *PingoneNotificationSettingsEmailResource {
	return &PingoneNotificationSettingsEmailResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneNotificationSettingsEmailResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	emailNotificationSettings, response, err := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsSettingsSMTPApi.ReadEmailNotificationsSettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadEmailNotificationsSettings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	if emailNotificationSettings == nil {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &importBlocks, nil
	}

	if response.StatusCode == 204 {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &importBlocks, nil
	}

	emailNotificationSettingsEnv, emailNotificationSettingsEnvOk := emailNotificationSettings.GetEnvironmentOk()
	var (
		emailNotificationSettingsEnvID   *string
		emailNotificationSettingsEnvIDOk bool
	)

	if emailNotificationSettingsEnvOk {
		emailNotificationSettingsEnvID, emailNotificationSettingsEnvIDOk = emailNotificationSettingsEnv.GetIdOk()
	}

	if !emailNotificationSettingsEnvOk || !emailNotificationSettingsEnvIDOk || emailNotificationSettingsEnvID == nil {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       "pingone_notification_settings_email",
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingoneNotificationSettingsEmailResource) ResourceType() string {
	return "pingone_notification_settings_email"
}
