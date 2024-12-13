package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOAuthIssuerResource{}
)

type PingFederateOAuthIssuerResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOAuthIssuerResource
func OAuthIssuer(clientInfo *connector.PingFederateClientInfo) *PingFederateOAuthIssuerResource {
	return &PingFederateOAuthIssuerResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOAuthIssuerResource) ResourceType() string {
	return "pingfederate_oauth_issuer"
}

func (r *PingFederateOAuthIssuerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthIssuerData, err := r.getOAuthIssuerData()
	if err != nil {
		return nil, err
	}

	for oauthIssuerId, oauthIssuerName := range *oauthIssuerData {
		commentData := map[string]string{
			"OAuth Issuer ID":   oauthIssuerId,
			"OAuth Issuer Name": oauthIssuerName,
			"Resource Type":     r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       oauthIssuerName,
			ResourceID:         oauthIssuerId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthIssuerResource) getOAuthIssuerData() (*map[string]string, error) {
	issuerData := make(map[string]string)

	issuers, response, err := r.clientInfo.ApiClient.OauthIssuersAPI.GetOauthIssuers(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetOauthIssuers", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if issuers == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	issuersItems, issuersItemsOk := issuers.GetItemsOk()
	if !issuersItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, issuer := range issuersItems {
		issuerId, issuerIdOk := issuer.GetIdOk()
		issuerName, issuerNameOk := issuer.GetNameOk()

		if issuerIdOk && issuerNameOk {
			issuerData[*issuerId] = *issuerName
		}
	}

	return &issuerData, nil
}
