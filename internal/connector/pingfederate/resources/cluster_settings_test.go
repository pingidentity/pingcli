// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func Test_PingFederateClusterSettings(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	resource := resources.ClusterSettings(clientInfo)

	valid, err := resource.ValidPingFederateVersion()
	if err != nil {
		t.Errorf("Error checking version compatibility: %v", err)
	}
	if !valid {
		t.Skipf("'%s' Resource is not supported in the version of PingFederate used. Skipping export test.", resource.ResourceType())
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
