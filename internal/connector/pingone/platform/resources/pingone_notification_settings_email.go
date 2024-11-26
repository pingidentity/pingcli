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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneNotificationSettingsEmailResource
func NotificationSettingsEmail(clientInfo *connector.PingOneClientInfo) *PingOneNotificationSettingsEmailResource {
	return &PingOneNotificationSettingsEmailResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneNotificationSettingsEmailResource) ResourceType() string {
	return "pingone_notification_settings_email"
}

func (r *PingOneNotificationSettingsEmailResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportNotificationSettingsEmails()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneNotificationSettingsEmailResource) exportNotificationSettingsEmails() error {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsSettingsSMTPApi.ReadEmailNotificationsSettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadEmailNotificationsSettings", r.ResourceType())
	if err != nil {
		return err
	}

	if response.StatusCode == 204 {
		return common.DataNilError(r.ResourceType(), response)
	}

	r.addImportBlock()

	return nil
}

func (r *PingOneNotificationSettingsEmailResource) addImportBlock() {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
