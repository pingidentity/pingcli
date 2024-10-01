package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

func TestKeyExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.Key(PingOneClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_key",
			ResourceName: "PingOne SSO Certificate for PingFederate Terraform Provider environment_ENCRYPTION",
			ResourceID:   fmt.Sprintf("%s/46a2d7ad-27ee-4743-92ce-aac279a4358a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_key",
			ResourceName: "test_SIGNING",
			ResourceID:   fmt.Sprintf("%s/619bad1d-c884-47c5-99d7-a998bc317791", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_key",
			ResourceName: "PingOne SSO Certificate for PingFederate Terraform Provider environment_SIGNING",
			ResourceID:   fmt.Sprintf("%s/702d1a27-10e9-40cc-ba73-d0274a2c97d2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_key",
			ResourceName: "common name_SIGNING",
			ResourceID:   fmt.Sprintf("%s/7d16daa9-f7eb-405f-b130-6567fe9d918f", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
