package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateKerberosRealm(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createKerberosRealm,
		DeleteFunc:         deleteKerberosRealm,
		Dependencies:       nil,
		ExportableResource: resources.KerberosRealm(clientInfo),
	}
}

func createKerberosRealm(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createKerberosRealm(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.KerberosRealmsAPI.CreateKerberosRealm(clientInfo.Context)
	clientStruct := client.KerberosRealm{
		ConnectionType:                     utils.Pointer("LOCAL_VALIDATION"),
		Id:                                 utils.Pointer("TestKerberosRealmId"),
		KerberosPassword:                   utils.Pointer("TestPassword1"),
		KerberosRealmName:                  "TestKerberosRealmName",
		KerberosUsername:                   utils.Pointer("TestKerberosUser"),
		RetainPreviousKeysOnPasswordChange: utils.Pointer(true),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateKerberosRealm", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.KerberosRealmName,
	}
}

func deleteKerberosRealm(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.KerberosRealmsAPI.DeleteKerberosRealm(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteKerberosRealm", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
