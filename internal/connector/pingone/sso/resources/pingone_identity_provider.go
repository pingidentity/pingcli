// Copyright © 2025 Ping Identity Corporation

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

	for idpId, idpName := range identityProviderData {
		commentData := map[string]string{
			"Export Environment ID":  r.clientInfo.PingOneExportEnvironmentID,
			"Identity Provider ID":   idpId,
			"Identity Provider Name": idpName,
			"Resource Type":          r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       idpName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, idpId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneIdentityProviderResource) getIdentityProviderData() (map[string]string, error) {
	identityProviderData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.IdentityProvidersApi.ReadAllIdentityProviders(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	identityProviders, err := pingone.GetManagementAPIObjectsFromIterator[management.IdentityProvider](iter, "ReadAllIdentityProviders", "GetIdentityProviders", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, idp := range identityProviders {
		var (
			idpId     *string
			idpIdOk   bool
			idpName   *string
			idpNameOk bool
		)

		switch {
		case idp.IdentityProviderApple != nil:
			idpId, idpIdOk = idp.IdentityProviderApple.GetIdOk()
			idpName, idpNameOk = idp.IdentityProviderApple.GetNameOk()
		case idp.IdentityProviderClientIDClientSecret != nil:
			idpId, idpIdOk = idp.IdentityProviderClientIDClientSecret.GetIdOk()
			idpName, idpNameOk = idp.IdentityProviderClientIDClientSecret.GetNameOk()
		case idp.IdentityProviderFacebook != nil:
			idpId, idpIdOk = idp.IdentityProviderFacebook.GetIdOk()
			idpName, idpNameOk = idp.IdentityProviderFacebook.GetNameOk()
		case idp.IdentityProviderOIDC != nil:
			idpId, idpIdOk = idp.IdentityProviderOIDC.GetIdOk()
			idpName, idpNameOk = idp.IdentityProviderOIDC.GetNameOk()
		case idp.IdentityProviderPaypal != nil:
			idpId, idpIdOk = idp.IdentityProviderPaypal.GetIdOk()
			idpName, idpNameOk = idp.IdentityProviderPaypal.GetNameOk()
		case idp.IdentityProviderSAML != nil:
			idpId, idpIdOk = idp.IdentityProviderSAML.GetIdOk()
			idpName, idpNameOk = idp.IdentityProviderSAML.GetNameOk()
		default:
			continue
		}

		if idpIdOk && idpNameOk {
			identityProviderData[*idpId] = *idpName
		}
	}

	return identityProviderData, nil
}
