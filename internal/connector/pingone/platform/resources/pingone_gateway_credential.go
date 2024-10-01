package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneGatewayCredentialResource{}
)

type PingoneGatewayCredentialResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneGatewayCredentialResource
func GatewayCredential(clientInfo *connector.PingOneClientInfo) *PingoneGatewayCredentialResource {
	return &PingoneGatewayCredentialResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneGatewayCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllGateways"

	gatewaysEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, gatewayInner := range gatewaysEmbedded.GetGateways() {
		var (
			gatewayId     *string
			gatewayName   *string
			gatewayIdOk   bool
			gatewayNameOk bool
		)

		switch {
		case gatewayInner.Gateway != nil:
			gatewayId, gatewayIdOk = gatewayInner.Gateway.GetIdOk()
			gatewayName, gatewayNameOk = gatewayInner.Gateway.GetNameOk()
		case gatewayInner.GatewayTypeLDAP != nil:
			gatewayId, gatewayIdOk = gatewayInner.GatewayTypeLDAP.GetIdOk()
			gatewayName, gatewayNameOk = gatewayInner.GatewayTypeLDAP.GetNameOk()
		case gatewayInner.GatewayTypeRADIUS != nil:
			gatewayId, gatewayIdOk = gatewayInner.GatewayTypeRADIUS.GetIdOk()
			gatewayName, gatewayNameOk = gatewayInner.GatewayTypeRADIUS.GetNameOk()
		default:
			continue
		}

		if gatewayIdOk && gatewayNameOk {
			apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GatewayCredentialsApi.ReadAllGatewayCredentials(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *gatewayId).Execute
			apiFunctionName := "ReadAllGatewayCredentials"

			gatewayCredentialsEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for gatewayCredentialIndex, gatewayCredential := range gatewayCredentialsEmbedded.GetCredentials() {
				gatewayCredentialId, gatewayCredentialIdOk := gatewayCredential.GetIdOk()

				if gatewayCredentialIdOk {
					commentData := map[string]string{
						"Resource Type":         r.ResourceType(),
						"Gateway Name":          *gatewayName,
						"Export Environment ID": r.clientInfo.ExportEnvironmentID,
						"Gateway ID":            *gatewayId,
						"Gateway Credential ID": *gatewayCredentialId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_credential_%d", *gatewayName, (gatewayCredentialIndex + 1)),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *gatewayId, *gatewayCredentialId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneGatewayCredentialResource) ResourceType() string {
	return "pingone_gateway_credential"
}
