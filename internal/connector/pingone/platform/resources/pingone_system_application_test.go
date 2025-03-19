// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestSystemApplicationExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.SystemApplication(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_system_application",
			ResourceName: "PingOne Application Portal",
			ResourceID:   fmt.Sprintf("%s/92a3765c-e135-4afa-8b12-4469672ac8a9", clientInfo.PingOneExportEnvironmentID),
		},
		{
			ResourceType: "pingone_system_application",
			ResourceName: "PingOne Self-Service - MyAccount",
			ResourceID:   fmt.Sprintf("%s/4ce54d01-5138-4c56-8175-4f02f69278f5", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
