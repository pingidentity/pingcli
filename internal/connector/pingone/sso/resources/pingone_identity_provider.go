package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneIdentityProviderResource{}
)

type PingOneIdentityProviderResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneIdentityProviderResource
func IdentityProvider(clientInfo *connector.PingOneClientInfo) *PingOneIdentityProviderResource {
	return &PingOneIdentityProviderResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneIdentityProviderResource) ResourceType() string {
	return "pingone_identity_provider"
}

func (r *PingOneIdentityProviderResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportIdentityProviders()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneIdentityProviderResource) exportIdentityProviders() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.IdentityProvidersApi.ReadAllIdentityProviders(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllIdentityProviders", r.ResourceType())
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
				r.addImportBlock(*idpId, *idpName)
			}
		}
	}

	return nil
}

func (r *PingOneIdentityProviderResource) addImportBlock(idpId, idpName string) {
	commentData := map[string]string{
		"Export Environment ID":  r.clientInfo.ExportEnvironmentID,
		"Identity Provider ID":   idpId,
		"Identity Provider Name": idpName,
		"Resource Type":          r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       idpName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, idpId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
