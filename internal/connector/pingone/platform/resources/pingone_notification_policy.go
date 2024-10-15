package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
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

func (r *PingOneNotificationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsPoliciesApi.ReadAllNotificationsPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllNotificationsPolicies"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, notificationPolicy := range embedded.GetNotificationsPolicies() {
		notificationPolicyId, notificationPolicyIdOk := notificationPolicy.GetIdOk()
		notificationPolicyName, notificationPolicyNameOk := notificationPolicy.GetNameOk()

		if notificationPolicyIdOk && notificationPolicyNameOk {
			commentData := map[string]string{
				"Resource Type":            r.ResourceType(),
				"Notification Policy Name": *notificationPolicyName,
				"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
				"Notification Policy ID":   *notificationPolicyId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *notificationPolicyName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *notificationPolicyId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneNotificationPolicyResource) ResourceType() string {
	return "pingone_notification_policy"
}
