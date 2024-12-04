package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGatewayResource{}
)

type PingOneGatewayResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneGatewayResource
func Gateway(clientInfo *connector.PingOneClientInfo) *PingOneGatewayResource {
	return &PingOneGatewayResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneGatewayResource) ResourceType() string {
	return "pingone_gateway"
}

func (r *PingOneGatewayResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportGateways()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneGatewayResource) exportGateways() error {
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
				r.addImportBlock(*gatewayId, *gatewayName)
			}
		}
	}

	return nil
}

func (r *PingOneGatewayResource) addImportBlock(gatewayId, gatewayName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Gateway ID":            gatewayId,
		"Gateway Name":          gatewayName,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       gatewayName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, gatewayId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
