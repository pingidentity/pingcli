package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGatewayCredentialResource{}
)

type PingOneGatewayCredentialResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneGatewayCredentialResource
func GatewayCredential(clientInfo *connector.PingOneClientInfo) *PingOneGatewayCredentialResource {
	return &PingOneGatewayCredentialResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneGatewayCredentialResource) ResourceType() string {
	return "pingone_gateway_credential"
}

func (r *PingOneGatewayCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	gatewayData, err := getGatewayData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for gatewayId, gatewayName := range gatewayData {
		gatewayCredentialData, err := getGatewayCredentialData(r.clientInfo, r.ResourceType(), gatewayId)
		if err != nil {
			return nil, err
		}

		for _, gatewayCredentialId := range gatewayCredentialData {
			commentData := map[string]string{
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Gateway Credential ID": gatewayCredentialId,
				"Gateway ID":            gatewayId,
				"Gateway Name":          gatewayName,
				"Resource Type":         r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_credential_%s", gatewayName, gatewayCredentialId),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, gatewayId, gatewayCredentialId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func getGatewayCredentialData(clientInfo *connector.PingOneClientInfo, resourceType, gatewayId string) ([]string, error) {
	gatewayCredentialData := []string{}

	iter := clientInfo.ApiClient.ManagementAPIClient.GatewayCredentialsApi.ReadAllGatewayCredentials(clientInfo.Context, clientInfo.ExportEnvironmentID, gatewayId).Execute()
	gatewayCredentials, err := common.GetManagementAPIObjectsFromIterator[management.GatewayCredential](iter, "ReadAllGatewayCredentials", "GetGatewayCredentials", resourceType)
	if err != nil {
		return nil, err
	}

	for _, gatewayCredential := range gatewayCredentials {
		gatewayCredentialId, gatewayCredentialIdOk := gatewayCredential.GetIdOk()

		if gatewayCredentialIdOk {
			gatewayCredentialData = append(gatewayCredentialData, *gatewayCredentialId)
		}
	}

	return gatewayCredentialData, nil
}
