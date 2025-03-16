package mfa_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestMFATerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:     "MFAApplicationPushCredential",
			resource: resources.MFAApplicationPushCredential(clientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:          "MFAFido2Policy",
			resource:      resources.MFAFido2Policy(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "MFADevicePolicy",
			resource:      resources.MFADevicePolicy(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "MFASettings",
			resource:      resources.MFASettings(clientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
