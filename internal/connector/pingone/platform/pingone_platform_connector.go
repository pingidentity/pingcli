package platform

import (
	"context"

	pingoneGoClient "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/logger"
)

const (
	serviceName = "pingone-platform"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingOnePlatformConnector{}
	_ connector.Authenticatable = &PingOnePlatformConnector{}
)

type PingOnePlatformConnector struct {
	clientInfo connector.PingOneClientInfo
}

// Utility method for creating a PingOnePlatformConnector
func PlatformConnector(ctx context.Context, apiClient *pingoneGoClient.Client, apiClientId *string, exportEnvironmentID string) *PingOnePlatformConnector {
	return &PingOnePlatformConnector{
		clientInfo: connector.PingOneClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ApiClientId:         apiClientId,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingOnePlatformConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne Platform Resources...")

	exportableResources := []connector.ExportableResource{
		resources.AlertChannel(&c.clientInfo),
		resources.Agreement(&c.clientInfo),
		resources.AgreementEnable(&c.clientInfo),
		resources.AgreementLocalization(&c.clientInfo),
		resources.AgreementLocalizationEnable(&c.clientInfo),
		resources.AgreementLocalizationRevision(&c.clientInfo),
		resources.BrandingSettings(&c.clientInfo),
		resources.BrandingTheme(&c.clientInfo),
		resources.BrandingThemeDefault(&c.clientInfo),
		resources.Certificate(&c.clientInfo),
		resources.CustomDomain(&c.clientInfo),
		resources.Environment(&c.clientInfo),
		resources.Form(&c.clientInfo),
		resources.FormRecaptchaV2(&c.clientInfo),
		resources.Gateway(&c.clientInfo),
		resources.GatewayCredential(&c.clientInfo),
		resources.GatewayRoleAssignment(&c.clientInfo),
		resources.IdentityPropagationPlan(&c.clientInfo),
		resources.Key(&c.clientInfo),
		resources.KeyRotationPolicy(&c.clientInfo),
		resources.Language(&c.clientInfo),
		resources.LanguageUpdate(&c.clientInfo),
		resources.NotificationPolicy(&c.clientInfo),
		resources.NotificationSettings(&c.clientInfo),
		resources.NotificationSettingsEmail(&c.clientInfo),
		resources.NotificationTemplateContent(&c.clientInfo),
		resources.PhoneDeliverySettings(&c.clientInfo),
		resources.SystemApplication(&c.clientInfo),
		resources.TrustedEmailAddress(&c.clientInfo),
		resources.TrustedEmailDomain(&c.clientInfo),
		resources.Webhook(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingOnePlatformConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingOnePlatformConnector) Login() error {
	return nil
}

func (c *PingOnePlatformConnector) Logout() error {
	return nil
}
