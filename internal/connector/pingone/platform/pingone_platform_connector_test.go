// Copyright © 2025 Ping Identity Corporation

package platform_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_platform_testable_resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestPlatformTerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name             string
		testableResource *testutils_resource.TestableResource
		ignoredErrors    []string
	}{
		{
			name:             "Agreement",
			testableResource: pingone_platform_testable_resources.Agreement(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AgreementEnable",
			testableResource: pingone_platform_testable_resources.AgreementEnable(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AgreementLocalization",
			testableResource: pingone_platform_testable_resources.AgreementLocalization(t, clientInfo),
			ignoredErrors: []string{
				"Error: Cannot find language by locale en-US",
			},
		},
		{
			name:             "AgreementLocalizationEnable",
			testableResource: pingone_platform_testable_resources.AgreementLocalizationEnable(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "AgreementLocalizationRevision",
			testableResource: pingone_platform_testable_resources.AgreementLocalizationRevision(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "BrandingSettings",
			testableResource: pingone_platform_testable_resources.BrandingSettings(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "BrandingTheme",
			testableResource: pingone_platform_testable_resources.BrandingTheme(t, clientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:             "BrandingThemeDefault",
			testableResource: pingone_platform_testable_resources.BrandingThemeDefault(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Certificate",
			testableResource: pingone_platform_testable_resources.Certificate(t, clientInfo),
			ignoredErrors: []string{
				"Error: Invalid combination of arguments",
			},
		},
		{
			name:             "CustomDomain",
			testableResource: pingone_platform_testable_resources.CustomDomain(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Environment",
			testableResource: pingone_platform_testable_resources.Environment(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Form",
			testableResource: pingone_platform_testable_resources.Form(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "FormsRecaptchaV2",
			testableResource: pingone_platform_testable_resources.FormsRecaptchaV2(t, clientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:             "Gateway",
			testableResource: pingone_platform_testable_resources.Gateway(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "GatewayCredential",
			testableResource: pingone_platform_testable_resources.GatewayCredential(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "GatewayRoleAssignment",
			testableResource: pingone_platform_testable_resources.GatewayRoleAssignment(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "IdentityPropagationPlan",
			testableResource: pingone_platform_testable_resources.IdentityPropagationPlan(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Key",
			testableResource: pingone_platform_testable_resources.Key(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "KeyRotationPolicy",
			testableResource: pingone_platform_testable_resources.KeyRotationPolicy(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Language",
			testableResource: pingone_platform_testable_resources.Language(t, clientInfo),
			ignoredErrors: []string{
				"Error: The language code `pt` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `fr` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `es` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `pl` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `hu` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `cs` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `ru` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `ko` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `th` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `tr` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `fr-CA` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `de` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `it` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `ja` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `zh` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `en` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
				"Error: The language code `nl` is reserved and cannot be imported into this provider.  Please use `pingone_language_override` for system-defined languages instead.",
			},
		},
		{
			name:             "LanguageUpdate",
			testableResource: pingone_platform_testable_resources.LanguageUpdate(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "NotificationPolicy",
			testableResource: pingone_platform_testable_resources.NotificationPolicy(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "NotificationSettings",
			testableResource: pingone_platform_testable_resources.NotificationSettings(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "NotificationSettingsEmail",
			testableResource: pingone_platform_testable_resources.NotificationSettingsEmail(t, clientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:             "NotificationTemplateContent",
			testableResource: pingone_platform_testable_resources.NotificationTemplateContent(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "PhoneDeliverySettings",
			testableResource: pingone_platform_testable_resources.PhoneDeliverySettings(t, clientInfo),
			ignoredErrors: []string{
				"Error: Missing required argument",
			},
		},
		{
			name:             "SystemApplication",
			testableResource: pingone_platform_testable_resources.SystemApplication(t, clientInfo),
			ignoredErrors:    nil,
		},
		// TODO: Currently unable to create a trusted email address via API due to trust email domain verification requirement
		// {
		// 	name:             "TrustedEmailAddress",
		// 	testableResource: pingone_platform_testable_resources.TrustedEmailAddress(t, clientInfo),
		// 	ignoredErrors:    nil,
		// },
		{
			name:             "TrustedEmailDomain",
			testableResource: pingone_platform_testable_resources.TrustedEmailDomain(t, clientInfo),
			ignoredErrors:    nil,
		},
		{
			name:             "Webhook",
			testableResource: pingone_platform_testable_resources.Webhook(t, clientInfo),
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
