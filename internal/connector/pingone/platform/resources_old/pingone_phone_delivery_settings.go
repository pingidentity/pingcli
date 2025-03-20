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
	_ connector.ExportableResource = &PingOnePhoneDeliverySettingsResource{}
)

type PingOnePhoneDeliverySettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOnePhoneDeliverySettingsResource
func PhoneDeliverySettings(clientInfo *connector.ClientInfo) *PingOnePhoneDeliverySettingsResource {
	return &PingOnePhoneDeliverySettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOnePhoneDeliverySettingsResource) ResourceType() string {
	return "pingone_phone_delivery_settings"
}

func (r *PingOnePhoneDeliverySettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	phoneDeliverySettingsData, err := r.getPhoneDeliverySettingsData()
	if err != nil {
		return nil, err
	}

	for phoneDeliverySettingsId, phoneDeliverySettingsName := range phoneDeliverySettingsData {
		commentData := map[string]string{
			"Export Environment ID":        r.clientInfo.PingOneExportEnvironmentID,
			"Phone Delivery Settings ID":   phoneDeliverySettingsId,
			"Phone Delivery Settings Name": phoneDeliverySettingsName,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       phoneDeliverySettingsName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, phoneDeliverySettingsId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOnePhoneDeliverySettingsResource) getPhoneDeliverySettingsData() (map[string]string, error) {
	phoneDeliverySettingsData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.PhoneDeliverySettingsApi.ReadAllPhoneDeliverySettings(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	phoneDeliverySettings, err := pingone.GetManagementAPIObjectsFromIterator[management.NotificationsSettingsPhoneDeliverySettings](iter, "ReadAllPhoneDeliverySettings", "GetPhoneDeliverySettings", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, phoneDeliverySettings := range phoneDeliverySettings {
		var (
			phoneDeliverySettingsId     *string
			phoneDeliverySettingsIdOk   bool
			phoneDeliverySettingsName   string
			phoneDeliverySettingsNameOk bool
		)

		switch {
		case phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsCustom != nil:
			phoneDeliverySettingsId, phoneDeliverySettingsIdOk = phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsCustom.GetIdOk()
			if phoneDeliverySettingsIdOk {
				phoneDeliverySettingsName, phoneDeliverySettingsNameOk = fmt.Sprintf("provider_custom_%s", *phoneDeliverySettingsId), true
			}
		case phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse != nil:
			phoneDeliverySettingsId, phoneDeliverySettingsIdOk = phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse.GetIdOk()
			phoneDeliverySettingsProvider, phoneDeliverySettingProviderOk := phoneDeliverySettings.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse.GetProviderOk()
			if phoneDeliverySettingsIdOk && phoneDeliverySettingProviderOk {
				switch *phoneDeliverySettingsProvider {
				case management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSPROVIDER_TWILIO:
					phoneDeliverySettingsName, phoneDeliverySettingsNameOk = fmt.Sprintf("provider_twilio_%s", *phoneDeliverySettingsId), true
				case management.ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSPROVIDER_SYNIVERSE:
					phoneDeliverySettingsName, phoneDeliverySettingsNameOk = fmt.Sprintf("provider_syniverse_%s", *phoneDeliverySettingsId), true
				default:
					continue
				}
			}
		default:
			continue
		}

		if phoneDeliverySettingsIdOk && phoneDeliverySettingsNameOk {
			phoneDeliverySettingsData[*phoneDeliverySettingsId] = phoneDeliverySettingsName
		}
	}

	return phoneDeliverySettingsData, nil
}
