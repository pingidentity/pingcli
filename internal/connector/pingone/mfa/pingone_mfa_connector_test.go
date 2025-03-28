// Copyright © 2025 Ping Identity Corporation

package mfa_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_mfa_testable_resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestMFATerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name             string
		testableResource *testutils_resource.TestableResource
		ignoredErrors    []string
	}{
		{
			name:             "MFAApplicationPushCredential",
			testableResource: pingone_mfa_testable_resources.MfaApplicationPushCredential(t, clientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:             "MFAFido2Policy",
			testableResource: pingone_mfa_testable_resources.MfaFido2Policy(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "MFADevicePolicy",
			testableResource: pingone_mfa_testable_resources.MfaDevicePolicy(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "MFASettings",
			testableResource: pingone_mfa_testable_resources.MfaSettings(t, clientInfo),
			ignoredErrors:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testableResource.CreateResource(t)
			defer tc.testableResource.DeleteResource(t)

			testutils_terraform.ValidateTerraformPlan(t, tc.testableResource.ExportableResource, tc.ignoredErrors)
		})
	}
}
