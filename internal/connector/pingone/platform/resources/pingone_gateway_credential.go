package resources

import (
	"fmt"

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

	gatewayData, err := r.getGatewayData()
	if err != nil {
		return nil, err
	}

	for gatewayId, gatewayName := range *gatewayData {
		gatewayCredentialData, err := r.getGatewayCredentialData(gatewayId)
		if err != nil {
			return nil, err
		}

		for _, gatewayCredentialId := range *gatewayCredentialData {
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

func (r *PingOneGatewayCredentialResource) getGatewayData() (*map[string]string, error) {
	gatewayData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGateways", r.ResourceType())
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, gatewayInner := range embedded.GetGateways() {
			var (
				gatewayId     *string
				gatewayIdOk   bool
				gatewayName   *string
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
				gatewayData[*gatewayId] = *gatewayName
			}
		}
	}

	return &gatewayData, nil
}

func (r *PingOneGatewayCredentialResource) getGatewayCredentialData(gatewayId string) (*[]string, error) {
	gatewayCredentialData := []string{}

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewayCredentialsApi.ReadAllGatewayCredentials(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, gatewayId).Execute()

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGatewayCredentials", r.ResourceType())
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, gatewayCredential := range embedded.GetCredentials() {
			gatewayCredentialId, gatewayCredentialIdOk := gatewayCredential.GetIdOk()

			if gatewayCredentialIdOk {
				gatewayCredentialData = append(gatewayCredentialData, *gatewayCredentialId)
			}
		}
	}

	return &gatewayCredentialData, nil
}
