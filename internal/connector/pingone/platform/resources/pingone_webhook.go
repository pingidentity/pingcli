package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneWebhookResource{}
)

type PingOneWebhookResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneWebhookResource
func Webhook(clientInfo *connector.PingOneClientInfo) *PingOneWebhookResource {
	return &PingOneWebhookResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneWebhookResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.SubscriptionsWebhooksApi.ReadAllSubscriptions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllSubscriptions"

	usersEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, subscription := range usersEmbedded.GetSubscriptions() {
		subscriptionId, subscriptionIdOk := subscription.GetIdOk()
		subscriptionName, subscriptionNameOk := subscription.GetNameOk()

		if subscriptionIdOk && subscriptionNameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Webhook Name":          *subscriptionName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Webhook ID":            *subscriptionId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *subscriptionName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *subscriptionId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneWebhookResource) ResourceType() string {
	return "pingone_webhook"
}
