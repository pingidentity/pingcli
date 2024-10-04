package mfa

import (
	"context"

	pingoneGoClient "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingcli/internal/logger"
)

const (
	serviceName = "pingone-mfa"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingoneMFAConnector{}
	_ connector.Authenticatable = &PingoneMFAConnector{}
)

type PingoneMFAConnector struct {
	clientInfo connector.PingOneClientInfo
}

// Utility method for creating a PingoneMFAConnector
func MFAConnector(ctx context.Context, apiClient *pingoneGoClient.Client, apiClientId *string, exportEnvironmentID string) *PingoneMFAConnector {
	return &PingoneMFAConnector{
		clientInfo: connector.PingOneClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ApiClientId:         apiClientId,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingoneMFAConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne MFA Resources...")

	exportableResources := []connector.ExportableResource{
		resources.MFAApplicationPushCredential(&c.clientInfo),
		resources.MFAFido2Policy(&c.clientInfo),
		resources.MFADevicePolicy(&c.clientInfo),
		resources.MFASettings(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingoneMFAConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingoneMFAConnector) Login() error {
	return nil
}

func (c *PingoneMFAConnector) Logout() error {
	return nil
}
