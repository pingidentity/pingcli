package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateKerberosRealm_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.KerberosRealm(PingFederateClientInfo)

	kerberosRealmId, kerberosRealmName := createKerberosRealm(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKerberosRealm(t, PingFederateClientInfo, resource.ResourceType(), kerberosRealmId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: kerberosRealmName,
			ResourceID:   kerberosRealmId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createKerberosRealm(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.KerberosRealmsAPI.CreateKerberosRealm(clientInfo.Context)
	result := client.KerberosRealm{}
	result.Id = utils.Pointer("TestKerberosRealmId")
	result.KerberosRealmName = "TestKerberosRealmName"

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateKerberosRealm", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.KerberosRealmName
}

func deleteKerberosRealm(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.KerberosRealmsAPI.DeleteKerberosRealm(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteKerberosRealm", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
