// Copyright © 2025 Ping Identity Corporation

package sso_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestSSOTerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "Application",
			resource:      resources.Application(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationAttributeMapping",
			resource:      resources.ApplicationAttributeMapping(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationFlowPolicyAssignment",
			resource:      resources.ApplicationFlowPolicyAssignment(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationResourceGrant",
			resource:      resources.ApplicationResourceGrant(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationRoleAssignment",
			resource:      resources.ApplicationRoleAssignment(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationSecret",
			resource:      resources.ApplicationSecret(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationSignOnPolicyAssignment",
			resource:      resources.ApplicationSignOnPolicyAssignment(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Group",
			resource:      resources.Group(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GroupNesting",
			resource:      resources.GroupNesting(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GroupRoleAssignment",
			resource:      resources.GroupRoleAssignment(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityProvider",
			resource:      resources.IdentityProvider(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityProviderAttribute",
			resource:      resources.IdentityProviderAttribute(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PasswordPolicy",
			resource:      resources.PasswordPolicy(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Population",
			resource:      resources.Population(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PopulationDefault",
			resource:      resources.PopulationDefault(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PopulationDefaultIdp",
			resource:      resources.PopulationDefaultIdp(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Resource",
			resource:      resources.Resource(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceAttribute",
			resource:      resources.ResourceAttribute(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScope",
			resource:      resources.ResourceScope(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScopeOpenId",
			resource:      resources.ResourceScopeOpenId(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScopePingOneApi",
			resource:      resources.ResourceScopePingOneApi(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceSecret",
			resource:      resources.ResourceSecret(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "SchemaAttribute",
			resource: resources.SchemaAttribute(clientInfo),
			ignoredErrors: []string{
				"Error: Data Loss Protection",
			},
		},
		{
			name:          "SignOnPolicy",
			resource:      resources.SignOnPolicy(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "SignOnPolicyAction",
			resource: resources.SignOnPolicyAction(clientInfo),
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
