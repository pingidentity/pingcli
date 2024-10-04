package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingonePhoneDeliverySettingsResource{}
)

type PingonePhoneDeliverySettingsResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingonePhoneDeliverySettingsResource
func PhoneDeliverySettings(clientInfo *connector.PingOneClientInfo) *PingonePhoneDeliverySettingsResource {
	return &PingonePhoneDeliverySettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingonePhoneDeliverySettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.PhoneDeliverySettingsApi.ReadAllPhoneDeliverySettings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllPhoneDeliverySettings"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for index, phoneDeliverySettings := range embedded.GetPhoneDeliverySettings() {
		var (
			phoneDeliverySettingsId     *string
			phoneDeliverySettingsName   *string
			phoneDeliverySettingsIdOk   bool
			phoneDeliverySettingsNameOk bool
		)

		switch {
		case phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsCustom != nil:
			phoneDeliverySettingsId, phoneDeliverySettingsIdOk = phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsCustom.GetIdOk()
			phoneDeliverySettingsName, phoneDeliverySettingsNameOk = phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsCustom.GetNameOk()
		case phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse != nil:
			phoneDeliverySettingsId, phoneDeliverySettingsIdOk = phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse.GetIdOk()
			phoneDeliverySettingsProvider, phoneDeliverySettingProviderOk := phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse.GetProviderOk()
			if phoneDeliverySettingProviderOk {
				switch *phoneDeliverySettingsProvider {
				case management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSPROVIDER_TWILIO:
					twilioName := fmt.Sprintf("CUSTOM_TWILIO_%d", index)
					phoneDeliverySettingsName, phoneDeliverySettingsNameOk = &twilioName, true
				case management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSPROVIDER_SYNIVERSE:
					syniverseName := fmt.Sprintf("CUSTOM_SYNIVERSE_%d", index)
					phoneDeliverySettingsName, phoneDeliverySettingsNameOk = &syniverseName, true
				default:
					continue
				}
			}
		default:
			continue
		}

		if phoneDeliverySettingsIdOk && phoneDeliverySettingsNameOk {
			commentData := map[string]string{
				"Resource Type":                r.ResourceType(),
				"Phone Delivery Settings Name": *phoneDeliverySettingsName,
				"Export Environment ID":        r.clientInfo.ExportEnvironmentID,
				"Phone Delivery Settings ID":   *phoneDeliverySettingsId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *phoneDeliverySettingsName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *phoneDeliverySettingsId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingonePhoneDeliverySettingsResource) ResourceType() string {
	return "pingone_phone_delivery_settings"
}
