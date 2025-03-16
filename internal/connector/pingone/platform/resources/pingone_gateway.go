package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGatewayResource{}
)

type PingOneGatewayResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneGatewayResource
func Gateway(clientInfo *connector.ClientInfo) *PingOneGatewayResource {
	return &PingOneGatewayResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneGatewayResource) ResourceType() string {
	return "pingone_gateway"
}

func (r *PingOneGatewayResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	gatewayData, err := r.getGatewayData()
	if err != nil {
		return nil, err
	}

	for gatewayId, gatewayName := range gatewayData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Gateway ID":            gatewayId,
			"Gateway Name":          gatewayName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       gatewayName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, gatewayId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneGatewayResource) getGatewayData() (map[string]string, error) {
	gatewayData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	gateways, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedGatewaysInner](iter, "ReadAllGateways", "GetGateways", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, gatewayInner := range gateways {
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

	return gatewayData, nil
}
