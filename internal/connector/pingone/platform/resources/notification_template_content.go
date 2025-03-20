// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	_ connector.ExportableResource = &PingOneNotificationTemplateContentResource{}
)

type PingOneNotificationTemplateContentResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneNotificationTemplateContentResource
func NotificationTemplateContent(clientInfo *connector.ClientInfo) *PingOneNotificationTemplateContentResource {
	return &PingOneNotificationTemplateContentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneNotificationTemplateContentResource) ResourceType() string {
	return "pingone_notification_template_content"
}

func (r *PingOneNotificationTemplateContentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	notificationTemplateContentData, err := r.getNotificationTemplateContentData()
	if err != nil {
		return nil, err
	}

	for notificationTemplateContentId, notificationTemplateContentName := range notificationTemplateContentData {
		commentData := map[string]string{
			"Notification Template Content ID":   notificationTemplateContentId,
			"Notification Template Content Name": notificationTemplateContentName,
			"Export Environment ID":              r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":                      r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       notificationTemplateContentName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, notificationTemplateContentId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneNotificationTemplateContentResource) getTemplateData() (map[string]string, error) {
	//TODO
}

func (r *PingOneNotificationTemplateContentResource) getNotificationTemplateContentData() (map[string]string, error) {
	notificationTemplateContentData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadAllTemplateContents(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.TemplateContent](iter, "ReadAllTemplateContents", "GetContents", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, notificationTemplateContent := range apiObjs {
		notificationTemplateContentId, notificationTemplateContentIdOk := notificationTemplateContent.GetIdOk()
		notificationTemplateContentName, notificationTemplateContentNameOk := notificationTemplateContent.GetNameOk()

		if notificationTemplateContentIdOk && notificationTemplateContentNameOk {
			notificationTemplateContentData[*notificationTemplateContentId] = *notificationTemplateContentName
		}
	}

	return notificationTemplateContentData, nil
}
