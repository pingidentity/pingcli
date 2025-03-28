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
	_ connector.ExportableResource = &PingOneIdentityProviderResource{}
)

type PingOneIdentityProviderResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneIdentityProviderResource
func IdentityProvider(clientInfo *connector.ClientInfo) *PingOneIdentityProviderResource {
	return &PingOneIdentityProviderResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneIdentityProviderResource) ResourceType() string {
	return "pingone_identity_provider"
}

func (r *PingOneIdentityProviderResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	identityProviderData, err := r.getIdentityProviderData()
	if err != nil {
		return nil, err
	}

	for identityProviderId, identityProviderName := range identityProviderData {
		commentData := map[string]string{
			"Identity Provider ID":   identityProviderId,
			"Identity Provider Name": identityProviderName,
			"Export Environment ID":  r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       identityProviderName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, identityProviderId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneIdentityProviderResource) getIdentityProviderData() (map[string]string, error) {
	identityProviderData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.IdentityProvidersApi.ReadAllIdentityProviders(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.IdentityProvider](iter, "ReadAllIdentityProviders", "GetIdentityProviders", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		var (
			identityProviderId     *string
			identityProviderIdOk   bool
			identityProviderName   *string
			identityProviderNameOk bool
		)

		switch {
		case innerObj.IdentityProviderApple != nil:
			identityProviderId, identityProviderIdOk = innerObj.IdentityProviderApple.GetIdOk()
			identityProviderName, identityProviderNameOk = innerObj.IdentityProviderApple.GetNameOk()
		case innerObj.IdentityProviderClientIDClientSecret != nil:
			identityProviderId, identityProviderIdOk = innerObj.IdentityProviderClientIDClientSecret.GetIdOk()
			identityProviderName, identityProviderNameOk = innerObj.IdentityProviderClientIDClientSecret.GetNameOk()
		case innerObj.IdentityProviderFacebook != nil:
			identityProviderId, identityProviderIdOk = innerObj.IdentityProviderFacebook.GetIdOk()
			identityProviderName, identityProviderNameOk = innerObj.IdentityProviderFacebook.GetNameOk()
		case innerObj.IdentityProviderOIDC != nil:
			identityProviderId, identityProviderIdOk = innerObj.IdentityProviderOIDC.GetIdOk()
			identityProviderName, identityProviderNameOk = innerObj.IdentityProviderOIDC.GetNameOk()
		case innerObj.IdentityProviderPaypal != nil:
			identityProviderId, identityProviderIdOk = innerObj.IdentityProviderPaypal.GetIdOk()
			identityProviderName, identityProviderNameOk = innerObj.IdentityProviderPaypal.GetNameOk()
		case innerObj.IdentityProviderSAML != nil:
			identityProviderId, identityProviderIdOk = innerObj.IdentityProviderSAML.GetIdOk()
			identityProviderName, identityProviderNameOk = innerObj.IdentityProviderSAML.GetNameOk()
		default:
			continue
		}

		if identityProviderIdOk && identityProviderNameOk {
			identityProviderData[*identityProviderId] = *identityProviderName
		}
	}

	return identityProviderData, nil
}
