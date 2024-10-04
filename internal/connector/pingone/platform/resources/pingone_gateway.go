package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneGatewayResource{}
)

type PingoneGatewayResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneGatewayResource
func Gateway(clientInfo *connector.PingOneClientInfo) *PingoneGatewayResource {
	return &PingoneGatewayResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneGatewayResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllGateways"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

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
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Gateway Name":          *gatewayName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Gateway ID":            *gatewayId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *gatewayName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *gatewayId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneGatewayResource) ResourceType() string {
	return "pingone_gateway"
}
