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

func Test_PingFederateNotificationPublisher_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.NotificationPublisher(PingFederateClientInfo)

	notificationPublisherId, notificationPublisherName := createNotificationPublisher(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteNotificationPublisher(t, PingFederateClientInfo, resource.ResourceType(), notificationPublisherId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: notificationPublisherName,
			ResourceID:   notificationPublisherId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createNotificationPublisher(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.NotificationPublishersAPI.CreateNotificationPublisher(clientInfo.Context)
	result := client.NotificationPublisher{
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

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateNotificationPublisher", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteNotificationPublisher(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.NotificationPublishersAPI.DeleteNotificationPublisher(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteNotificationPublisher", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
