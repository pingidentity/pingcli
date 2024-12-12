package platform_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_terraform"
)

func TestPlatformTerraformPlan(t *testing.T) {
	PingOneClientInfo := testutils.GetPingOneClientInfo(t)

	testutils_terraform.InitPingOneTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "Agreement",
			resource:      resources.Agreement(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementEnable",
			resource:      resources.AgreementEnable(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalization",
			resource:      resources.AgreementLocalization(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalizationEnable",
			resource:      resources.AgreementLocalizationEnable(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalizationRevision",
			resource:      resources.AgreementLocalizationRevision(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "BrandingSettings",
			resource:      resources.BrandingSettings(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "BrandingTheme",
			resource: resources.BrandingTheme(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:          "BrandingThemeDefault",
			resource:      resources.BrandingThemeDefault(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "Certificate",
			resource: resources.Certificate(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Invalid combination of arguments",
			},
		},
		{
			name:          "CustomDomain",
			resource:      resources.CustomDomain(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Environment",
			resource:      resources.Environment(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Form",
			resource:      resources.Form(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "FormRecaptchaV2",
			resource: resources.FormRecaptchaV2(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:     "Gateway",
			resource: resources.Gateway(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
				"Error: Missing required argument",
			},
		},
		{
			name:          "GatewayCredential",
			resource:      resources.GatewayCredential(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GatewayRoleAssignment",
			resource:      resources.GatewayRoleAssignment(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityPropagationPlan",
			resource:      resources.IdentityPropagationPlan(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Key",
			resource:      resources.Key(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "KeyRotationPolicy",
			resource:      resources.KeyRotationPolicy(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Language",
			resource:      resources.Language(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "LanguageUpdate",
			resource:      resources.LanguageUpdate(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "NotificationPolicy",
			resource:      resources.NotificationPolicy(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "NotificationSettings",
			resource:      resources.NotificationSettings(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "NotificationSettingsEmail",
			resource: resources.NotificationSettingsEmail(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		// TODO remove this skip dependent upon STAGING-25369
		// {
		// 	name:          "NotificationTemplateContent",
		// 	resource:      resources.NotificationTemplateContent(PingOneClientInfo),
		// 	ignoredErrors: nil,
		// },
		{
			name:     "PhoneDeliverySettings",
			resource: resources.PhoneDeliverySettings(PingOneClientInfo),
			ignoredErrors: []string{
				"Error: Missing required argument",
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:          "SystemApplication",
			resource:      resources.SystemApplication(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "TrustedEmailAddress",
			resource:      resources.TrustedEmailAddress(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "TrustedEmailDomain",
			resource:      resources.TrustedEmailDomain(PingOneClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Webhook",
			resource:      resources.Webhook(PingOneClientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
