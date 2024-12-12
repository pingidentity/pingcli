package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneIdentityProviderAttributeResource{}
)

type PingOneIdentityProviderAttributeResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneIdentityProviderAttributeResource
func IdentityProviderAttribute(clientInfo *connector.PingOneClientInfo) *PingOneIdentityProviderAttributeResource {
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

	for idpId, idpName := range *identityProviderData {
		identityProviderAttributeData, err := r.getIdentityProviderAttributeData(idpId)
		if err != nil {
			return nil, err
		}

		for idpAttributeId, idpAttributeName := range *identityProviderAttributeData {
			commentData := map[string]string{
				"Export Environment ID":            r.clientInfo.ExportEnvironmentID,
				"Identity Provider Attribute ID":   idpAttributeId,
				"Identity Provider Attribute Name": idpAttributeName,
				"Identity Provider ID":             idpId,
				"Identity Provider Name":           idpName,
				"Resource Type":                    r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", idpName, idpAttributeName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, idpId, idpAttributeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneIdentityProviderAttributeResource) getIdentityProviderData() (*map[string]string, error) {
	identityProviderData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.IdentityProvidersApi.ReadAllIdentityProviders(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllIdentityProviders", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

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
				identityProviderData[*idpId] = *idpName
			}
		}
	}

	return &identityProviderData, nil
}

func (r *PingOneIdentityProviderAttributeResource) getIdentityProviderAttributeData(idpId string) (*map[string]string, error) {
	identityProviderAttributeData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.IdentityProviderAttributesApi.ReadAllIdentityProviderAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, idpId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllIdentityProviderAttributes", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, idpAttribute := range embedded.GetAttributes() {
			idpAttributeId, idpAttributeIdOk := idpAttribute.IdentityProviderAttribute.GetIdOk()
			idpAttributeName, idpAttributeNameOk := idpAttribute.IdentityProviderAttribute.GetNameOk()

			if idpAttributeIdOk && idpAttributeNameOk {
				identityProviderAttributeData[*idpAttributeId] = *idpAttributeName
			}
		}
	}

	return &identityProviderAttributeData, nil
}
