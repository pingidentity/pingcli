package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneIdentityProviderResource{}
)

type PingoneIdentityProviderResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneIdentityProviderResource
func IdentityProvider(clientInfo *connector.PingOneClientInfo) *PingoneIdentityProviderResource {
	return &PingoneIdentityProviderResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneIdentityProviderResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteIdentityProvidersFunc := r.clientInfo.ApiClient.ManagementAPIClient.IdentityProvidersApi.ReadAllIdentityProviders(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiIdentityProviderFunctionName := "ReadAllIdentityProviders"

	embedded, err := common.GetManagementEmbedded(apiExecuteIdentityProvidersFunc, apiIdentityProviderFunctionName, r.ResourceType())
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
			commentData := map[string]string{
				"Resource Type":          r.ResourceType(),
				"Identity Provider Name": *idpName,
				"Export Environment ID":  r.clientInfo.ExportEnvironmentID,
				"Identity Provider ID":   *idpId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *idpName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *idpId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneIdentityProviderResource) ResourceType() string {
	return "pingone_identity_provider"
}
