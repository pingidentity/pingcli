package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateNotificationPublisher(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createNotificationPublisher,
		DeleteFunc:         deleteNotificationPublisher,
		Dependencies:       nil,
		ExportableResource: resources.NotificationPublisher(clientInfo),
	}
}

func createNotificationPublisher(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createNotificationPublisher(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.NotificationPublishersAPI.CreateNotificationPublisher(clientInfo.PingFederateContext)
	clientStruct := client.NotificationPublisher{
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "From Address",
					Value: utils.Pointer("test@example.com"),
				},
				{
					Name:  "Email Server",
					Value: utils.Pointer("example.com"),
				},
			},
		},
		Id:   "TestNotificationPublisherId",
		Name: "TestNotificationPublisherName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.email.SmtpNotificationPlugin",
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateNotificationPublisher", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteNotificationPublisher(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.NotificationPublishersAPI.DeleteNotificationPublisher(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteNotificationPublisher", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
