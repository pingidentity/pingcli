package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateClusterSettings(t *testing.T) {
	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ClusterSettings(pingfederateClientInfo)

	valid, err := resource.ValidPingFederateVersion()
	if err != nil {
		t.Fatalf("Failed to validate PingFederate version: %v", err)
	}
	if !valid {
		t.Logf("'%s' Resource is not supported in the version of PingFederate used. Skipping tests for export.", resource.ResourceType())
		t.SkipNow()
	}

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: "Cluster Settings",
			ResourceID:   "cluster_settings_singleton_id",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
