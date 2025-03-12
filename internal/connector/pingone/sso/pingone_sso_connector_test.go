package sso_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestSSOTerraformPlan(t *testing.T) {
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "Application",
			resource:      resources.Application(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationAttributeMapping",
			resource:      resources.ApplicationAttributeMapping(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationFlowPolicyAssignment",
			resource:      resources.ApplicationFlowPolicyAssignment(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationResourceGrant",
			resource:      resources.ApplicationResourceGrant(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationRoleAssignment",
			resource:      resources.ApplicationRoleAssignment(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationSecret",
			resource:      resources.ApplicationSecret(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationSignOnPolicyAssignment",
			resource:      resources.ApplicationSignOnPolicyAssignment(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Group",
			resource:      resources.Group(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GroupNesting",
			resource:      resources.GroupNesting(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GroupRoleAssignment",
			resource:      resources.GroupRoleAssignment(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityProvider",
			resource:      resources.IdentityProvider(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityProviderAttribute",
			resource:      resources.IdentityProviderAttribute(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PasswordPolicy",
			resource:      resources.PasswordPolicy(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Population",
			resource:      resources.Population(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PopulationDefault",
			resource:      resources.PopulationDefault(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PopulationDefaultIdp",
			resource:      resources.PopulationDefaultIdp(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Resource",
			resource:      resources.Resource(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceAttribute",
			resource:      resources.ResourceAttribute(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScope",
			resource:      resources.ResourceScope(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScopeOpenId",
			resource:      resources.ResourceScopeOpenId(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScopePingOneApi",
			resource:      resources.ResourceScopePingOneApi(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "SchemaAttribute",
			resource: resources.SchemaAttribute(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Data Loss Protection",
			},
		},
		{
			name:          "SignOnPolicy",
			resource:      resources.SignOnPolicy(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "SignOnPolicyAction",
			resource: resources.SignOnPolicyAction(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Conflicting configuration arguments",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
