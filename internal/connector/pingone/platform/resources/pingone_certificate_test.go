package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestCertificateExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.Certificate(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_certificate",
			ResourceName: "common name",
			ResourceID:   fmt.Sprintf("%s/b9eb2b6e-381e-4b1c-86d3-096d951787f4", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_certificate",
			ResourceName: "terraform",
			ResourceID:   fmt.Sprintf("%s/fa8f15d6-1c62-4db1-920e-d22f6dd68ba8", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
