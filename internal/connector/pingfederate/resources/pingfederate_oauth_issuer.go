package resources

import (
	"fmt"

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

func (r *PingFederateOAuthIssuerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.OauthIssuersAPI.GetOauthIssuers(r.clientInfo.Context).Execute
	apiFunctionName := "GetOauthIssuers"

	issuers, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if issuers == nil {
		l.Error().Msgf("Returned %s() issuers is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	issuersItems, ok := issuers.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() issuers items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, issuer := range issuersItems {
		issuerId, issuerIdOk := issuer.GetIdOk()
		issuerName, issuerNameOk := issuer.GetNameOk()

		if issuerIdOk && issuerNameOk {
			commentData := map[string]string{
				"Resource Type":              r.ResourceType(),
				"OAuth Issuer Resource ID":   *issuerId,
				"OAuth Issuer Resource Name": *issuerName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *issuerName,
				ResourceID:         *issuerId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateOAuthIssuerResource) ResourceType() string {
	return "pingfederate_oauth_issuer"
}
