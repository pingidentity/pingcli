// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_platform_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func NotificationSettingsEmail(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         nil,
		DeleteFunc:         nil,
		Dependencies:       nil,
		ExportableResource: resources.NotificationSettingsEmail(clientInfo),
	}
}
