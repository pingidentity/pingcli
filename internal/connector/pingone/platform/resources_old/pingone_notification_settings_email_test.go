// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestNotificationSettingsEmailExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.NotificationSettingsEmail(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_notification_settings_email",
			ResourceName: "pingone_notification_settings_email",
			ResourceID:   clientInfo.PingOneExportEnvironmentID,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
