// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_sso_testable_resources"
)

func Test_PopulationDefault(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_sso_testable_resources.PopulationDefault(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	populationTr := tr.Dependencies[0]

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_population_default", populationTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   clientInfo.PingOneExportEnvironmentID,
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
