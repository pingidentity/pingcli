package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
)

func TestPingFederateNotificationPublisherExport(t *testing.T) {
	// Get initialized apiClient and resource
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.NotificationPublisher(PingFederateClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingfederate_notification_publisher",
			ResourceName: "exampleSmtpPublisher",
			ResourceID:   "exampleSmtpPublisher",
		},
		{
			ResourceType: "pingfederate_notification_publisher",
			ResourceName: "exampleSmtpPublisher2",
			ResourceID:   "exampleSmtpPublisher2",
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
