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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneNotificationPolicyResource
func NotificationPolicy(clientInfo *connector.PingOneClientInfo) *PingOneNotificationPolicyResource {
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
			"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
			"Notification Policy ID":   notificationPolicyId,
			"Notification Policy Name": notificationPolicyName,
			"Resource Type":            r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       notificationPolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, notificationPolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneNotificationPolicyResource) getNotificationPolicyData() (map[string]string, error) {
	notificationPolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsPoliciesApi.ReadAllNotificationsPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	notificationPolicies, err := pingone.GetManagementAPIObjectsFromIterator[management.NotificationsPolicy](iter, "ReadAllNotificationsPolicies", "GetNotificationsPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, notificationPolicy := range notificationPolicies {
		notificationPolicyId, notificationPolicyIdOk := notificationPolicy.GetIdOk()
		notificationPolicyName, notificationPolicyNameOk := notificationPolicy.GetNameOk()

		if notificationPolicyIdOk && notificationPolicyNameOk {
			notificationPolicyData[*notificationPolicyId] = *notificationPolicyName
		}
	}

	return notificationPolicyData, nil
}
