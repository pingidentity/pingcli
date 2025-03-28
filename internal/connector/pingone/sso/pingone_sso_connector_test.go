// Copyright © 2025 Ping Identity Corporation

package sso_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_sso_testable_resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestSSOTerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name             string
		testableResource *testutils_resource.TestableResource
		ignoredErrors    []string
	}{
		{
			name:             "Application",
			testableResource: pingone_sso_testable_resources.ApplicationDeviceAuthorization(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ApplicationAttributeMapping",
			testableResource: pingone_sso_testable_resources.ApplicationAttributeMapping(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ApplicationFlowPolicyAssignment",
			testableResource: pingone_sso_testable_resources.ApplicationFlowPolicyAssignment(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ApplicationResourceGrant",
			testableResource: pingone_sso_testable_resources.ApplicationResourceGrant(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ApplicationRoleAssignment",
			testableResource: pingone_sso_testable_resources.ApplicationRoleAssignment(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ApplicationSecret",
			testableResource: pingone_sso_testable_resources.ApplicationSecret(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ApplicationSignOnPolicyAssignment",
			testableResource: pingone_sso_testable_resources.ApplicationSignOnPolicyAssignment(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Group",
			testableResource: pingone_sso_testable_resources.Group(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "GroupNesting",
			testableResource: pingone_sso_testable_resources.GroupNesting(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "GroupRoleAssignment",
			testableResource: pingone_sso_testable_resources.GroupRoleAssignment(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "IdentityProvider",
			testableResource: pingone_sso_testable_resources.IdentityProvider(t, clientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:             "IdentityProviderAttribute",
			testableResource: pingone_sso_testable_resources.IdentityProviderAttribute(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PasswordPolicy",
			testableResource: pingone_sso_testable_resources.PasswordPolicy(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Population",
			testableResource: pingone_sso_testable_resources.Population(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PopulationDefault",
			testableResource: pingone_sso_testable_resources.PopulationDefault(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PopulationDefaultIdp",
			testableResource: pingone_sso_testable_resources.PopulationDefaultIdentityProvider(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Resource",
			testableResource: pingone_sso_testable_resources.Resource(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ResourceAttribute",
			testableResource: pingone_sso_testable_resources.ResourceAttribute(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ResourceScope",
			testableResource: pingone_sso_testable_resources.ResourceScope(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ResourceScopeOpenId",
			testableResource: pingone_sso_testable_resources.ResourceScopeOpenId(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "ResourceScopePingOneApi",
			testableResource: pingone_sso_testable_resources.ResourceScopePingOneApi(t, clientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Match",
			},
		},
		{
			name:             "ResourceSecret",
			testableResource: pingone_sso_testable_resources.ResourceSecret(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "SchemaAttribute",
			testableResource: pingone_sso_testable_resources.SchemaAttribute(t, clientInfo),
			ignoredErrors: []string{
				"Error: Data Loss Protection",
			},
		},
		// TODO: Re-enable test after compleition of TRIAGE-26632
		// {
		// 	name:             "SignOnPolicy",
		// 	testableResource: pingone_sso_testable_resources.SignOnPolicy(t, clientInfo),
		// 	ignoredErrors:    nil,
		// },
		// TODO: Re-enable test after compleition of TRIAGE-26632
		// {
		// 	name:             "SignOnPolicyAction",
		// 	testableResource: pingone_sso_testable_resources.SignOnPolicyAction(t, clientInfo),
		// 	ignoredErrors: []string{
		// 		"Error: Conflicting configuration arguments",
		// 	},
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
