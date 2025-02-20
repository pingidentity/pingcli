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

func (r *PingFederateCertificateCAResource) ResourceType() string {
	return "pingfederate_certificate_ca"
}

func (r *PingFederateCertificateCAResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	trustedCAData, err := r.getTrustedCAData()
	if err != nil {
		return nil, err
	}

	for certViewId, certViewInfo := range trustedCAData {
		certViewIssuerDN := certViewInfo[0]
		certViewSerialNumber := certViewInfo[1]

		commentData := map[string]string{
			"Certificate CA Issuer DN":     certViewIssuerDN,
			"Certificate CA ID":            certViewId,
			"Certificate CA Serial Number": certViewSerialNumber,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", certViewIssuerDN, certViewSerialNumber),
			ResourceID:         certViewId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateCertificateCAResource) getTrustedCAData() (map[string][]string, error) {
	trustedCAData := make(map[string][]string)

	certViews, response, err := r.clientInfo.ApiClient.CertificatesCaAPI.GetTrustedCAs(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetTrustedCAs", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if certViews == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	certViewsItems, certViewsItemsOk := certViews.GetItemsOk()
	if !certViewsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, certView := range certViewsItems {
		certViewId, certViewIdOk := certView.GetIdOk()
		certViewIssuerDN, certViewIssuerDNOk := certView.GetIssuerDNOk()
		certViewSerialNumber, certViewSerialNumberOk := certView.GetSerialNumberOk()

		if certViewIdOk && certViewIssuerDNOk && certViewSerialNumberOk {
			trustedCAData[*certViewId] = []string{*certViewIssuerDN, *certViewSerialNumber}
		}
	}

	return trustedCAData, nil
}
