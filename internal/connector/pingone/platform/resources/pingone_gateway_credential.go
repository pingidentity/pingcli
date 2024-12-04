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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneGatewayCredentialResource
func GatewayCredential(clientInfo *connector.PingOneClientInfo) *PingOneGatewayCredentialResource {
	return &PingOneGatewayCredentialResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneGatewayCredentialResource) ResourceType() string {
	return "pingone_gateway_credential"
}

func (r *PingOneGatewayCredentialResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportGatewayCredentials()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneGatewayCredentialResource) exportGatewayCredentials() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGateways", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
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
				err := r.exportGatewayCredentialsByGateway(*gatewayId, *gatewayName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneGatewayCredentialResource) exportGatewayCredentialsByGateway(gatewayId, gatewayName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewayCredentialsApi.ReadAllGatewayCredentials(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, gatewayId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGatewayCredentials", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, gatewayCredential := range embedded.GetCredentials() {
			gatewayCredentialId, gatewayCredentialIdOk := gatewayCredential.GetIdOk()

			if gatewayCredentialIdOk {
				r.addImportBlock(gatewayId, gatewayName, *gatewayCredentialId)
			}
		}
	}

	return nil
}

func (r *PingOneGatewayCredentialResource) addImportBlock(gatewayId, gatewayName, gatewayCredentialId string) {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
