// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	_ connector.ExportableResource = &PingOneGatewayCredentialResource{}
)

type PingOneGatewayCredentialResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneGatewayCredentialResource
func GatewayCredential(clientInfo *connector.ClientInfo) *PingOneGatewayCredentialResource {
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

	gatewayCredentialData, err := r.getGatewayCredentialData()
	if err != nil {
		return nil, err
	}

	for gatewayCredentialId, gatewayCredentialName := range gatewayCredentialData {
		commentData := map[string]string{
			"Gateway Credential ID":   gatewayCredentialId,
			"Gateway Credential Name": gatewayCredentialName,
			"Export Environment ID":   r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":           r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       gatewayCredentialName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, gatewayCredentialId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneGatewayCredentialResource) getGatewayData() (map[string]string, error) {
	//TODO
}

func (r *PingOneGatewayCredentialResource) getGatewayCredentialData() (map[string]string, error) {
	gatewayCredentialData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GatewayCredentialsApi.ReadAllGatewayCredentials(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.GatewayCredential](iter, "ReadAllGatewayCredentials", "GetCredentials", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, gatewayCredential := range apiObjs {
		gatewayCredentialId, gatewayCredentialIdOk := gatewayCredential.GetIdOk()
		gatewayCredentialName, gatewayCredentialNameOk := gatewayCredential.GetNameOk()

		if gatewayCredentialIdOk && gatewayCredentialNameOk {
			gatewayCredentialData[*gatewayCredentialId] = *gatewayCredentialName
		}
	}

	return gatewayCredentialData, nil
}
