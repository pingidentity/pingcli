// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestLanguageExport(t *testing.T) {
	// Get initialized apiClient and resource
	clientInfo := testutils.GetClientInfo(t)
	resource := resources.Language(clientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_language",
			ResourceName: "Afar",
			ResourceID:   fmt.Sprintf("%s/c5f97303-c5c1-459a-b8ee-3fcb183eb52a", clientInfo.PingOneExportEnvironmentID),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
