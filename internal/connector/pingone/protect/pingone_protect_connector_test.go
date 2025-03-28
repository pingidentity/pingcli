// Copyright © 2025 Ping Identity Corporation

package protect_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_protect_testable_resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestProtectTerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name             string
		testableResource *testutils_resource.TestableResource
		ignoredErrors    []string
	}{
		{
			name:             "RiskPolicy",
			testableResource: pingone_protect_testable_resources.RiskPolicy(t, clientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:             "RiskPredictor",
			testableResource: pingone_protect_testable_resources.RiskPredictor(t, clientInfo),
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
