package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneIdentityProviderAttributeResource{}
)

type PingoneIdentityProviderAttributeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneIdentityProviderAttributeResource
func IdentityProviderAttribute(clientInfo *connector.PingOneClientInfo) *PingoneIdentityProviderAttributeResource {
	return &PingoneIdentityProviderAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneIdentityProviderAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	idpsApiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.IdentityProvidersApi.ReadAllIdentityProviders(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	idpsApiFunctionName := "ReadAllIdentityProviders"

	embedded, err := common.GetManagementEmbedded(idpsApiExecuteFunc, idpsApiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, idp := range embedded.GetIdentityProviders() {
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
			apiExecuteIdpAttributesFunc := r.clientInfo.ApiClient.ManagementAPIClient.IdentityProviderAttributesApi.ReadAllIdentityProviderAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *idpId).Execute
			apiIdpAttributesFunctionName := "ReadAllIdentityProviderAttributes"

			idpAttributesEmbedded, err := common.GetManagementEmbedded(apiExecuteIdpAttributesFunc, apiIdpAttributesFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, idpAttribute := range idpAttributesEmbedded.GetAttributes() {
				idpAttributeId, idpAttributeIdOk := idpAttribute.IdentityProviderAttribute.GetIdOk()
				idpAttributeName, idpAttributeNameOk := idpAttribute.IdentityProviderAttribute.GetNameOk()
				if idpAttributeIdOk && idpAttributeNameOk {
					commentData := map[string]string{
						"Resource Type":                    r.ResourceType(),
						"Identity Provider Name":           *idpName,
						"Identity Provider Attribute Name": *idpAttributeName,
						"Export Environment ID":            r.clientInfo.ExportEnvironmentID,
						"Identity Provider ID":             *idpId,
						"Identity Provider Attribute ID":   *idpAttributeId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *idpName, *idpAttributeName),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *idpId, *idpAttributeId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneIdentityProviderAttributeResource) ResourceType() string {
	return "pingone_identity_provider_attribute"
}
