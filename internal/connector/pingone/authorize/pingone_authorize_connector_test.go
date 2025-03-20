// Copyright © 2025 Ping Identity Corporation

package authorize_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestAuthorizeTerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "AuthorizeAPIService",
			resource:      resources.AuthorizeAPIService(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeAPIServiceDeployment",
			resource:      resources.AuthorizeAPIServiceDeployment(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeAPIServiceOperation",
			resource:      resources.AuthorizeAPIServiceOperation(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeApplicationRole",
			resource:      resources.AuthorizeApplicationRole(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeApplicationRolePermission",
			resource:      resources.AuthorizeApplicationRolePermission(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeDecisionEndpoint",
			resource:      resources.AuthorizeDecisionEndpoint(clientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
