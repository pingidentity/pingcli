package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOauthIssuerResource{}
)

type PingFederateOauthIssuerResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOauthIssuerResource
func OauthIssuer(clientInfo *connector.ClientInfo) *PingFederateOauthIssuerResource {
	return &PingFederateOauthIssuerResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOauthIssuerResource) ResourceType() string {
	return "pingfederate_oauth_issuer"
}

func (r *PingFederateOauthIssuerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oauthIssuerData, err := r.getOauthIssuerData()
	if err != nil {
		return nil, err
	}

	for oauthIssuerId, oauthIssuerName := range oauthIssuerData {
		commentData := map[string]string{
			"Oauth Issuer ID":   oauthIssuerId,
			"Oauth Issuer Name": oauthIssuerName,
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

func (r *PingFederateOauthIssuerResource) getOauthIssuerData() (map[string]string, error) {
	oauthIssuerData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthIssuersAPI.GetOauthIssuers(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetOauthIssuers", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, oauthIssuer := range items {
		oauthIssuerId, oauthIssuerIdOk := oauthIssuer.GetIdOk()
		oauthIssuerName, oauthIssuerNameOk := oauthIssuer.GetNameOk()

		if oauthIssuerIdOk && oauthIssuerNameOk {
			oauthIssuerData[*oauthIssuerId] = *oauthIssuerName
		}
	}

	return oauthIssuerData, nil
}
