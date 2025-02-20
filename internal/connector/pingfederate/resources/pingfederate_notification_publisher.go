package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateNotificationPublisherResource{}
)

type PingFederateNotificationPublisherResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateNotificationPublisherResource
func NotificationPublisher(clientInfo *connector.PingFederateClientInfo) *PingFederateNotificationPublisherResource {
	return &PingFederateNotificationPublisherResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateNotificationPublisherResource) ResourceType() string {
	return "pingfederate_notification_publisher"
}

func (r *PingFederateNotificationPublisherResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	notificationPublisherData, err := r.getNotificationPublisherData()
	if err != nil {
		return nil, err
	}

	for notificationPublisherId, notificationPublisherName := range notificationPublisherData {
		commentData := map[string]string{
			"Notification Publisher ID":   notificationPublisherId,
			"Notification Publisher Name": notificationPublisherName,
			"Resource Type":               r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       notificationPublisherName,
			ResourceID:         notificationPublisherId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateNotificationPublisherResource) getNotificationPublisherData() (map[string]string, error) {
	notificationPublisherData := make(map[string]string)

	notificationPublishers, response, err := r.clientInfo.ApiClient.NotificationPublishersAPI.GetNotificationPublishers(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetNotificationPublishers", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if notificationPublishers == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	notificationPublishersItems, notificationPublishersItemsOk := notificationPublishers.GetItemsOk()
	if !notificationPublishersItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, notificationPublisher := range notificationPublishersItems {
		notificationPublisherId, notificationPublisherIdOk := notificationPublisher.GetIdOk()
		notificationPublisherName, notificationPublisherNameOk := notificationPublisher.GetNameOk()

		if notificationPublisherIdOk && notificationPublisherNameOk {
			notificationPublisherData[*notificationPublisherId] = *notificationPublisherName
		}
	}

	return notificationPublisherData, nil
}
