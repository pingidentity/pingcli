// Copyright © 2025 Ping Identity Corporation

package authorize

import (
	"context"

	pingoneGoClient "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/logger"
)

const (
	serviceName = "pingone-authorize"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingoneAuthorizeConnector{}
	_ connector.Authenticatable = &PingoneAuthorizeConnector{}
)

type PingoneAuthorizeConnector struct {
	clientInfo connector.ClientInfo
}

// Utility method for creating a PingoneAuthorizeConnector
func AuthorizeConnector(ctx context.Context, apiClient *pingoneGoClient.Client, apiClientId *string, exportEnvironmentID string) *PingoneAuthorizeConnector {
	return &PingoneAuthorizeConnector{
		clientInfo: connector.ClientInfo{
			PingOneContext:             ctx,
			PingOneApiClient:           apiClient,
			PingOneApiClientId:         *apiClientId,
			PingOneExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingoneAuthorizeConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne Authorize Resources...")

	exportableResources := []connector.ExportableResource{
		resources.AuthorizeAPIService(&c.clientInfo),
		resources.AuthorizeAPIServiceDeployment(&c.clientInfo),
		resources.AuthorizeAPIServiceOperation(&c.clientInfo),
		resources.ApplicationResource(&c.clientInfo),
		resources.AuthorizeApplicationResourcePermission(&c.clientInfo),
		resources.AuthorizeApplicationRole(&c.clientInfo),
		resources.AuthorizeApplicationRolePermission(&c.clientInfo),
		resources.AuthorizeDecisionEndpoint(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingoneAuthorizeConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingoneAuthorizeConnector) Login() error {
	return nil
}

func (c *PingoneAuthorizeConnector) Logout() error {
	return nil
}
