package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateDataStore_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.DataStore(PingFederateClientInfo)

	dataStoreId, dataStoreType := createDataStore(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteDataStore(t, PingFederateClientInfo, resource.ResourceType(), dataStoreId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", dataStoreType, dataStoreId),
			ResourceID:   dataStoreId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createDataStore(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.DataStoresAPI.CreateDataStore(clientInfo.Context)
	result := client.DataStoreAggregation{}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateDataStore", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	var (
		dataStoreId     *string
		dataStoreIdOk   bool
		dataStoreType   *string
		dataStoreTypeOk bool
	)

	switch {
	case resource.CustomDataStore != nil:
		dataStoreId, dataStoreIdOk = resource.CustomDataStore.GetIdOk()
		dataStoreType, dataStoreTypeOk = resource.CustomDataStore.GetTypeOk()
	case resource.JdbcDataStore != nil:
		dataStoreId, dataStoreIdOk = resource.JdbcDataStore.GetIdOk()
		dataStoreType, dataStoreTypeOk = resource.JdbcDataStore.GetTypeOk()
	case resource.LdapDataStore != nil:
		dataStoreId, dataStoreIdOk = resource.LdapDataStore.GetIdOk()
		dataStoreType, dataStoreTypeOk = resource.LdapDataStore.GetTypeOk()
	case resource.PingOneLdapGatewayDataStore != nil:
		dataStoreId, dataStoreIdOk = resource.PingOneLdapGatewayDataStore.GetIdOk()
		dataStoreType, dataStoreTypeOk = resource.PingOneLdapGatewayDataStore.GetTypeOk()
	}

	if !dataStoreIdOk || !dataStoreTypeOk {
		t.Fatalf("Failed to get data store id and type")
	}

	return *dataStoreId, *dataStoreType
}

func deleteDataStore(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.DataStoresAPI.DeleteDataStore(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteDataStore", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
