// Copyright © 2025 Ping Identity Corporation

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
	_ connector.ExportableResource = &PingOneWebhookResource{}
)

type PingOneWebhookResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneWebhookResource
func Webhook(clientInfo *connector.ClientInfo) *PingOneWebhookResource {
	return &PingOneWebhookResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneWebhookResource) ResourceType() string {
	return "pingone_webhook"
}

func (r *PingOneWebhookResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	subscriptionData, err := r.getSubscriptionData()
	if err != nil {
		return nil, err
	}

	for subscriptionId, subscriptionName := range subscriptionData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
			"Webhook ID":            subscriptionId,
			"Webhook Name":          subscriptionName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       subscriptionName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, subscriptionId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneWebhookResource) getSubscriptionData() (map[string]string, error) {
	subscriptionData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.SubscriptionsWebhooksApi.ReadAllSubscriptions(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	subscriptions, err := pingone.GetManagementAPIObjectsFromIterator[management.Subscription](iter, "ReadAllSubscriptions", "GetSubscriptions", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, subscription := range subscriptions {
		subscriptionId, subscriptionIdOk := subscription.GetIdOk()
		subscriptionName, subscriptionNameOk := subscription.GetNameOk()

		if subscriptionIdOk && subscriptionNameOk {
			subscriptionData[*subscriptionId] = *subscriptionName
		}
	}

	return subscriptionData, nil
}
