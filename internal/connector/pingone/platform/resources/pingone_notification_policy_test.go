// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestNotificationPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.NotificationPolicy(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_notification_policy",
			ResourceName: "Test",
			ResourceID:   fmt.Sprintf("%s/32cc413d-0ec8-4be9-823c-a9e06f5a5830", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_notification_policy",
			ResourceName: "Default Notification Policy",
			ResourceID:   fmt.Sprintf("%s/54606af4-72a6-4b38-bfb8-75034097af9a", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
