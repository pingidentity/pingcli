// Copyright © 2025 Ping Identity Corporation

package platform_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestPlatformTerraformPlan(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "Agreement",
			resource:      resources.Agreement(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementEnable",
			resource:      resources.AgreementEnable(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalization",
			resource:      resources.AgreementLocalization(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalizationEnable",
			resource:      resources.AgreementLocalizationEnable(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalizationRevision",
			resource:      resources.AgreementLocalizationRevision(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "BrandingSettings",
			resource:      resources.BrandingSettings(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "BrandingTheme",
			resource: resources.BrandingTheme(clientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:          "BrandingThemeDefault",
			resource:      resources.BrandingThemeDefault(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "Certificate",
			resource: resources.Certificate(clientInfo),
			ignoredErrors: []string{
				"Error: Invalid combination of arguments",
			},
		},
		{
			name:          "CustomDomain",
			resource:      resources.CustomDomain(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Environment",
			resource:      resources.Environment(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Form",
			resource:      resources.Form(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "FormRecaptchaV2",
			resource: resources.FormRecaptchaV2(clientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:     "Gateway",
			resource: resources.Gateway(clientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
				"Error: Missing required argument",
			},
		},
		{
			name:          "GatewayCredential",
			resource:      resources.GatewayCredential(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GatewayRoleAssignment",
			resource:      resources.GatewayRoleAssignment(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityPropagationPlan",
			resource:      resources.IdentityPropagationPlan(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Key",
			resource:      resources.Key(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "KeyRotationPolicy",
			resource:      resources.KeyRotationPolicy(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Language",
			resource:      resources.Language(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "LanguageUpdate",
			resource:      resources.LanguageUpdate(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "NotificationPolicy",
			resource:      resources.NotificationPolicy(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "NotificationSettings",
			resource:      resources.NotificationSettings(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "NotificationSettingsEmail",
			resource: resources.NotificationSettingsEmail(clientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		// TODO remove this skip dependent upon STAGING-25369
		// {
		// 	name:          "NotificationTemplateContent",
		// 	resource:      resources.NotificationTemplateContent(clientInfo),
		// 	ignoredErrors: nil,
		// },
		{
			name:     "PhoneDeliverySettings",
			resource: resources.PhoneDeliverySettings(clientInfo),
			ignoredErrors: []string{
				"Error: Missing required argument",
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:          "SystemApplication",
			resource:      resources.SystemApplication(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "TrustedEmailAddress",
			resource:      resources.TrustedEmailAddress(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "TrustedEmailDomain",
			resource:      resources.TrustedEmailDomain(clientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Webhook",
			resource:      resources.Webhook(clientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
