package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateCertificateCAResource{}
)

type PingFederateCertificateCAResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateCertificateCAResource
func CertificateCA(clientInfo *connector.PingFederateClientInfo) *PingFederateCertificateCAResource {
	return &PingFederateCertificateCAResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateCertificateCAResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.CertificatesCaAPI.GetTrustedCAs(r.clientInfo.Context).Execute
	apiFunctionName := "GetTrustedCAs"

	certViews, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if certViews == nil {
		l.Error().Msgf("Returned %s() certViews is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	certViewsItems, ok := certViews.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() certViews items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, certView := range certViewsItems {
		certViewId, certViewIdOk := certView.GetIdOk()
		certViewIssuerDN, certViewIssuerDNOk := certView.GetIssuerDNOk()
		certViewSerialNumber, certViewSerialNumberOk := certView.GetSerialNumberOk()

		if certViewIdOk && certViewIssuerDNOk && certViewSerialNumberOk {
			commentData := map[string]string{
				"Resource Type":                r.ResourceType(),
				"Certificate CA Resource ID":   *certViewId,
				"Certificate CA Issuer DN":     *certViewIssuerDN,
				"Certificate CA Serial Number": *certViewSerialNumber,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", *certViewIssuerDN, *certViewSerialNumber),
				ResourceID:         *certViewId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateCertificateCAResource) ResourceType() string {
	return "pingfederate_certificate_ca"
}
