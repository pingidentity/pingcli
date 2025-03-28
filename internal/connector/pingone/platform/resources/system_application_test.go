// Copyright © 2025 Ping Identity Corporation

package resources_test

import (
	"fmt"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_platform_testable_resources"
)

func TestSystemApplicationExport(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_platform_testable_resources.SystemApplication(t, clientInfo)

	expectedImportBlocks := []connector.ImportBlock{}

	// Every environment has pre-configured system applications, so we need to get them all.
	iter := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", tr.ExportableResource.ResourceType())
	if err != nil {
		t.Fatalf("Failed to get system applications: %v", err)
	}
	for _, application := range apiObjs {
		var (
			applicationId     *string
			applicationIdOk   bool
			applicationName   *string
			applicationNameOk bool
		)

		switch {
		case application.ApplicationPingOnePortal != nil:
			applicationId, applicationIdOk = application.ApplicationPingOnePortal.GetIdOk()
			applicationName, applicationNameOk = application.ApplicationPingOnePortal.GetNameOk()
		case application.ApplicationPingOneSelfService != nil:
			applicationId, applicationIdOk = application.ApplicationPingOneSelfService.GetIdOk()
			applicationName, applicationNameOk = application.ApplicationPingOneSelfService.GetNameOk()
		default:
			continue
		}

		if applicationIdOk && applicationNameOk {
			expectedImportBlocks = append(expectedImportBlocks, connector.ImportBlock{
				ResourceType: tr.ExportableResource.ResourceType(),
				ResourceName: *applicationName,
				ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, *applicationId),
			})
		}
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
