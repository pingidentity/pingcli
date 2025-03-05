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
	_ connector.ExportableResource = &PingOneAlertChannelResource{}
)

type PingOneAlertChannelResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneAlertChannelResource
func AlertChannel(clientInfo *connector.PingOneClientInfo) *PingOneAlertChannelResource {
	return &PingOneAlertChannelResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneAlertChannelResource) ResourceType() string {
	return "pingone_alert_channel"
}

func (r *PingOneAlertChannelResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	alertChannelData, err := r.getAlertChannelData()
	if err != nil {
		return nil, err
	}

	for alertChannelId, alertChannelName := range alertChannelData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Alert Channel ID":      alertChannelId,
			"Alert Channel Name":    alertChannelName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       alertChannelName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, alertChannelId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneAlertChannelResource) getAlertChannelData() (map[string]string, error) {
	alertChannelData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.AlertingApi.ReadAllAlertChannels(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	alertChannels, err := pingone.GetManagementAPIObjectsFromIterator[management.AlertChannel](iter, "ReadAllAlertChannels", "GetAlertChannels", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, alertChannel := range alertChannels {
		alertChannelId, alertChannelIdOk := alertChannel.GetIdOk()
		alertChannelName, alertChannelNameOk := alertChannel.GetAlertNameOk()

		if alertChannelIdOk && alertChannelNameOk {
			alertChannelData[*alertChannelId] = *alertChannelName
		}
	}

	return alertChannelData, nil
}
