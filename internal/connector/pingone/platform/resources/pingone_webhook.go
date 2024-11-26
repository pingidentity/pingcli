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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneWebhookResource
func Webhook(clientInfo *connector.PingOneClientInfo) *PingOneWebhookResource {
	return &PingOneWebhookResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneWebhookResource) ResourceType() string {
	return "pingone_webhook"
}

func (r *PingOneWebhookResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportWebhooks()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneWebhookResource) exportWebhooks() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.SubscriptionsWebhooksApi.ReadAllSubscriptions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllSubscriptions", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, subscription := range embedded.GetSubscriptions() {
			subscriptionId, subscriptionIdOk := subscription.GetIdOk()
			subscriptionName, subscriptionNameOk := subscription.GetNameOk()

			if subscriptionIdOk && subscriptionNameOk {
				r.addImportBlock(*subscriptionId, *subscriptionName)
			}
		}
	}

	return nil
}

func (r *PingOneWebhookResource) addImportBlock(subscriptionId, subscriptionName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
		"Webhook ID":            subscriptionId,
		"Webhook Name":          subscriptionName,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       subscriptionName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, subscriptionId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
