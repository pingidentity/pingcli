package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestLanguageUpdateExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.LanguageUpdate(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_language_update",
			ResourceName: "French_update",
			ResourceID:   fmt.Sprintf("%s/3f8a2e14-0ace-41db-a92d-74b3b7913ffe", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language_update",
			ResourceName: "English_update",
			ResourceID:   fmt.Sprintf("%s/88c78fb2-9d74-41e3-a1d8-a9f729a2b463", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
