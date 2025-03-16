package sso

import (
	"context"

	pingoneGoClient "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/logger"
)

const (
	serviceName = "pingone-sso"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingOneSSOConnector{}
	_ connector.Authenticatable = &PingOneSSOConnector{}
)

type PingOneSSOConnector struct {
	clientInfo connector.ClientInfo
}

// Utility method for creating a PingOneSSOConnector
func SSOConnector(ctx context.Context, apiClient *pingoneGoClient.Client, apiClientId *string, exportEnvironmentID string) *PingOneSSOConnector {
	return &PingOneSSOConnector{
		clientInfo: connector.ClientInfo{
			PingOneContext:             ctx,
			PingOneApiClient:           apiClient,
			PingOneApiClientId:         *apiClientId,
			PingOneExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingOneSSOConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne SSO Resources...")

	exportableResources := []connector.ExportableResource{
		resources.Application(&c.clientInfo),
		resources.ApplicationAttributeMapping(&c.clientInfo),
		resources.ApplicationFlowPolicyAssignment(&c.clientInfo),
		resources.ApplicationResourceGrant(&c.clientInfo),
		resources.ApplicationRoleAssignment(&c.clientInfo),
		resources.ApplicationSecret(&c.clientInfo),
		resources.ApplicationSignOnPolicyAssignment(&c.clientInfo),
		resources.Group(&c.clientInfo),
		resources.GroupNesting(&c.clientInfo),
		resources.GroupRoleAssignment(&c.clientInfo),
		resources.IdentityProvider(&c.clientInfo),
		resources.IdentityProviderAttribute(&c.clientInfo),
		resources.PasswordPolicy(&c.clientInfo),
		resources.Population(&c.clientInfo),
		resources.PopulationDefault(&c.clientInfo),
		resources.PopulationDefaultIdp(&c.clientInfo),
		resources.Resource(&c.clientInfo),
		resources.ResourceAttribute(&c.clientInfo),
		resources.ResourceSecret(&c.clientInfo),
		resources.ResourceScope(&c.clientInfo),
		resources.ResourceScopeOpenId(&c.clientInfo),
		resources.ResourceScopePingOneApi(&c.clientInfo),
		resources.SchemaAttribute(&c.clientInfo),
		resources.SignOnPolicy(&c.clientInfo),
		resources.SignOnPolicyAction(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingOneSSOConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingOneSSOConnector) Login() error {
	return nil
}

func (c *PingOneSSOConnector) Logout() error {
	return nil
}
