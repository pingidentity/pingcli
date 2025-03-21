// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestAlertChannelExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.AlertChannel(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_alert_channel",
			ResourceName: "Test Alert Channel",
			ResourceID:   fmt.Sprintf("%s/6035f516-b474-4941-945a-76514913a74d", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
