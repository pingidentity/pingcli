// Copyright © 2025 Ping Identity Corporation

package authorize_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_authorize_testable_resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestAuthorizeTerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name             string
		testableResource *testutils_resource.TestableResource
		ignoredErrors    []string
	}{
		{
			name:             "ApplicationResource",
			testableResource: pingone_authorize_testable_resources.ApplicationResource(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ApplicationResourcePermission",
			testableResource: pingone_authorize_testable_resources.ApplicationResourcePermission(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AuthorizeAPIService",
			testableResource: pingone_authorize_testable_resources.AuthorizeApiService(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AuthorizeAPIServiceDeployment",
			testableResource: pingone_authorize_testable_resources.AuthorizeApiServiceDeployment(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AuthorizeAPIServiceOperation",
			testableResource: pingone_authorize_testable_resources.AuthorizeApiServiceOperation(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AuthorizeApplicationRole",
			testableResource: pingone_authorize_testable_resources.AuthorizeApplicationRole(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AuthorizeApplicationRolePermission",
			testableResource: pingone_authorize_testable_resources.AuthorizeApplicationRolePermission(t, clientInfo),
			ignoredErrors:    nil,
		},
		// TODO: Remove after completion of TRIAGE-26607
		// {
		// 	name:             "AuthorizeDecisionEndpoint",
		// 	testableResource: pingone_authorize_testable_resources.AuthorizeDecisionEndpoint(t, clientInfo),
		// 	ignoredErrors:    nil,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testableResource.CreateResource(t)
			defer tc.testableResource.DeleteResource(t)

			testutils_terraform.ValidateTerraformPlan(t, tc.testableResource.ExportableResource, tc.ignoredErrors)
		})
	}
}
