package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateNotificationPublisherSettingsResource{}
)

type PingFederateNotificationPublisherSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateNotificationPublisherSettingsResource
func NotificationPublisherSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateNotificationPublisherSettingsResource {
	return &PingFederateNotificationPublisherSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateNotificationPublisherSettingsResource) ResourceType() string {
	return "pingfederate_notification_publisher_settings"
}

func (r *PingFederateNotificationPublisherSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	notificationPublisherSettingsId := "notification_publisher_settings_singleton_id"
	notificationPublisherSettingsName := "Notification Publisher Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       notificationPublisherSettingsName,
		ResourceID:         notificationPublisherSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
