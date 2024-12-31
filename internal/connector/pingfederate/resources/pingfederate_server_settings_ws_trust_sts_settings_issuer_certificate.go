package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource{}
)

type PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource
func ServerSettingsWsTrustStsSettingsIssuerCertificate(clientInfo *connector.PingFederateClientInfo) *PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource {
	return &PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource) ResourceType() string {
	return "pingfederate_server_settings_ws_trust_sts_settings_issuer_certificate"
}

func (r *PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	issuerCertsData, err := r.getIssuerCertsData()
	if err != nil {
		return nil, err
	}

	for issuerCertId, issuerCertInfo := range *issuerCertsData {
		issuerCertDN := issuerCertInfo[0]
		issuerCertSerialNumber := issuerCertInfo[1]

		commentData := map[string]string{
			"Issuer Certificate ID":            issuerCertId,
			"Issuer Certificate Issuer DN":     issuerCertDN,
			"Issuer Certificate Serial Number": issuerCertSerialNumber,
			"Resource Type":                    r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", issuerCertDN, issuerCertSerialNumber),
			ResourceID:         issuerCertId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource) getIssuerCertsData() (*map[string][]string, error) {
	issuerCertsData := make(map[string][]string)

	issuerCerts, response, err := r.clientInfo.ApiClient.ServerSettingsAPI.GetCerts(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetCerts", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if issuerCerts == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	issuerCertsItems, issuerCertsItemsOk := issuerCerts.GetItemsOk()
	if !issuerCertsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, issuerCert := range issuerCertsItems {
		issuerCertView, issuerCertViewOk := issuerCert.GetCertViewOk()

		if issuerCertViewOk {
			issuerCertId, issuerCertIdOk := issuerCertView.GetIdOk()
			issuerCertDN, issuerCertDNOk := issuerCertView.GetIssuerDNOk()
			issuerCertSerialNumber, issuerCertSerialNumberOk := issuerCertView.GetSerialNumberOk()

			if issuerCertIdOk && issuerCertDNOk && issuerCertSerialNumberOk {
				issuerCertsData[*issuerCertId] = []string{*issuerCertDN, *issuerCertSerialNumber}
			}
		}
	}

	return &issuerCertsData, nil
}
