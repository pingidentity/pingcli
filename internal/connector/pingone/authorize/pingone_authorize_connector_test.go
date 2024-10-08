package authorize_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_terraform"
)

func TestAuthorizeTerraformPlan(t *testing.T) {
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "AuthorizeAPIService",
			resource:      resources.AuthorizeAPIService(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeAPIServiceDeployment",
			resource:      resources.AuthorizeAPIServiceDeployment(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeAPIServiceOperation",
			resource:      resources.AuthorizeAPIServiceOperation(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeApplicationRole",
			resource:      resources.AuthorizeApplicationRole(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeApplicationRolePermission",
			resource:      resources.AuthorizeApplicationRolePermission(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeDecisionEndpoint",
			resource:      resources.AuthorizeDecisionEndpoint(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizePolicyManagementPolicy",
			resource:      resources.AuthorizePolicyManagementPolicy(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizePolicyManagementRule",
			resource:      resources.AuthorizePolicyManagementRule(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizePolicyManagementStatement",
			resource:      resources.AuthorizePolicyManagementStatement(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeTrustFrameworkAttribute",
			resource:      resources.AuthorizeTrustFrameworkAttribute(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeTrustFrameworkCondition",
			resource:      resources.AuthorizeTrustFrameworkCondition(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeTrustFrameworkProcessor",
			resource:      resources.AuthorizeTrustFrameworkProcessor(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AuthorizeTrustFrameworkService",
			resource:      resources.AuthorizeTrustFrameworkService(PingOneClientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
