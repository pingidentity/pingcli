package protect

import (
	"context"

	pingoneGoClient "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingcli/internal/logger"
)

const (
	serviceName = "pingone-protect"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingOneProtectConnector{}
	_ connector.Authenticatable = &PingOneProtectConnector{}
)

type PingOneProtectConnector struct {
	clientInfo connector.PingOneClientInfo
}

// Utility method for creating a PingOneProtectConnector
func ProtectConnector(ctx context.Context, apiClient *pingoneGoClient.Client, apiClientId *string, exportEnvironmentID string) *PingOneProtectConnector {
	return &PingOneProtectConnector{
		clientInfo: connector.PingOneClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ApiClientId:         apiClientId,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingOneProtectConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne MFA Resources...")

	exportableResources := []connector.ExportableResource{
		resources.RiskPolicy(&c.clientInfo),
		resources.RiskPredictor(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingOneProtectConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingOneProtectConnector) Login() error {
	return nil
}

func (c *PingOneProtectConnector) Logout() error {
	return nil
}
