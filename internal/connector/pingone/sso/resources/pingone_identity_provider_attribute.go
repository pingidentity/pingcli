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
	_ connector.ExportableResource = &PingOneIdentityProviderAttributeResource{}
)

type PingOneIdentityProviderAttributeResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneIdentityProviderAttributeResource
func IdentityProviderAttribute(clientInfo *connector.ClientInfo) *PingOneIdentityProviderAttributeResource {
	return &PingOneIdentityProviderAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneIdentityProviderAttributeResource) ResourceType() string {
	return "pingone_identity_provider_attribute"
}

func (r *PingOneIdentityProviderAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	identityProviderData, err := r.getIdentityProviderData()
	if err != nil {
		return nil, err
	}

	for idpId, idpName := range identityProviderData {
		identityProviderAttributeData, err := r.getIdentityProviderAttributeData(idpId)
		if err != nil {
			return nil, err
		}

		for idpAttributeId, idpAttributeName := range identityProviderAttributeData {
			commentData := map[string]string{
				"Export Environment ID":            r.clientInfo.PingOneExportEnvironmentID,
				"Identity Provider Attribute ID":   idpAttributeId,
				"Identity Provider Attribute Name": idpAttributeName,
				"Identity Provider ID":             idpId,
				"Identity Provider Name":           idpName,
				"Resource Type":                    r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", idpName, idpAttributeName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, idpId, idpAttributeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneIdentityProviderAttributeResource) getIdentityProviderData() (map[string]string, error) {
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

func (r *PingOneIdentityProviderAttributeResource) getIdentityProviderAttributeData(idpId string) (map[string]string, error) {
	identityProviderAttributeData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.IdentityProviderAttributesApi.ReadAllIdentityProviderAttributes(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, idpId).Execute()
	attributeInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedAttributesInner](iter, "ReadAllIdentityProviderAttributes", "GetAttributes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, attributeInners := range attributeInners {
		if attributeInners.IdentityProviderAttribute != nil {
			idpAttributeId, idpAttributeIdOk := attributeInners.IdentityProviderAttribute.GetIdOk()
			idpAttributeName, idpAttributeNameOk := attributeInners.IdentityProviderAttribute.GetNameOk()

			if idpAttributeIdOk && idpAttributeNameOk {
				identityProviderAttributeData[*idpAttributeId] = *idpAttributeName
			}
		}
	}

	return identityProviderAttributeData, nil
}
