package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneNotificationSettingsResource{}
)

type PingOneNotificationSettingsResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneNotificationSettingsResource
func NotificationSettings(clientInfo *connector.PingOneClientInfo) *PingOneNotificationSettingsResource {
	return &PingOneNotificationSettingsResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneNotificationSettingsResource) ResourceType() string {
	return "pingone_notification_settings"
}

func (r *PingOneNotificationSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportNotificationSettings()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneNotificationSettingsResource) exportNotificationSettings() error {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsSettingsApi.ReadNotificationsSettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadNotificationsSettings", r.ResourceType())
	if err != nil {
		return err
	}

	if response.StatusCode == 204 {
		return common.DataNilError(r.ResourceType(), response)
	}

	r.addImportBlock()

	return nil
}

func (r *PingOneNotificationSettingsResource) addImportBlock() {
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
