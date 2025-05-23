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
	_ connector.ExportableResource = &PingOneNotificationPolicyResource{}
)

type PingOneNotificationPolicyResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneNotificationPolicyResource
func NotificationPolicy(clientInfo *connector.ClientInfo) *PingOneNotificationPolicyResource {
	return &PingOneNotificationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneNotificationPolicyResource) ResourceType() string {
	return "pingone_notification_policy"
}

func (r *PingOneNotificationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	notificationPolicyData, err := r.getNotificationPolicyData()
	if err != nil {
		return nil, err
	}

	for notificationPolicyId, notificationPolicyName := range notificationPolicyData {
		commentData := map[string]string{
			"Notification Policy ID":   notificationPolicyId,
			"Notification Policy Name": notificationPolicyName,
			"Export Environment ID":    r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":            r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       notificationPolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, notificationPolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneNotificationPolicyResource) getNotificationPolicyData() (map[string]string, error) {
	notificationPolicyData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsPoliciesApi.ReadAllNotificationsPolicies(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.NotificationsPolicy](iter, "ReadAllNotificationsPolicies", "GetNotificationsPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, notificationPolicy := range apiObjs {
		notificationPolicyId, notificationPolicyIdOk := notificationPolicy.GetIdOk()
		notificationPolicyName, notificationPolicyNameOk := notificationPolicy.GetNameOk()

		if notificationPolicyIdOk && notificationPolicyNameOk {
			notificationPolicyData[*notificationPolicyId] = *notificationPolicyName
		}
	}

	return notificationPolicyData, nil
}
