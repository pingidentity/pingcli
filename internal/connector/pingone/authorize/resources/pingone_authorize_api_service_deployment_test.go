package resources_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/pingidentity/pingcli/internal/connector"
// 	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
// 	"github.com/pingidentity/pingcli/internal/testing/testutils"
// )

// func TestAuthorizeAPIServiceDeploymentExport(t *testing.T) {
// 	// Get initialized apiClient and resource
// 	PingOneClientInfo := testutils.GetPingOneClientInfo(t)
// 	resource := resources.AuthorizeAPIServiceDeployment(PingOneClientInfo)

// 	// Defined the expected ImportBlocks for the resource
// 	expectedImportBlocks := []connector.ImportBlock{
// 		{
// 			ResourceType: "pingone_authorize_api_service_deployment",
// 			ResourceName: "Test Authorize API Service Deployment",
// 			ResourceID:   fmt.Sprintf("%s/5ae2227f-cb5b-47c3-bb40-440db09a98e6", testutils.GetEnvironmentID()),
// 		},
// 	}

// 	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
// }
