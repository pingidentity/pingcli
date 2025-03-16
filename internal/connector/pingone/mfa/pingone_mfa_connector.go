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
	_ connector.Exportable      = &PingOneMFAConnector{}
	_ connector.Authenticatable = &PingOneMFAConnector{}
)

type PingOneMFAConnector struct {
	clientInfo connector.ClientInfo
}

// Utility method for creating a PingOneMFAConnector
func MFAConnector(ctx context.Context, apiClient *pingoneGoClient.Client, apiClientId *string, exportEnvironmentID string) *PingOneMFAConnector {
	return &PingOneMFAConnector{
		clientInfo: connector.ClientInfo{
			PingOneContext:             ctx,
			PingOneApiClient:           apiClient,
			PingOneApiClientId:         *apiClientId,
			PingOneExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingOneMFAConnector) Export(format, outputDir string, overwriteExport bool) error {
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

func (c *PingOneMFAConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingOneMFAConnector) Login() error {
	return nil
}

func (c *PingOneMFAConnector) Logout() error {
	return nil
}
