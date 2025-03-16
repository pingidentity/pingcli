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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource
func ServerSettingsWsTrustStsSettingsIssuerCertificate(clientInfo *connector.ClientInfo) *PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource {
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

	serverSettingsWsTrustStsSettingsIssuerCertificateData, err := r.getServerSettingsWsTrustStsSettingsIssuerCertificateData()
	if err != nil {
		return nil, err
	}

	for serverSettingsWsTrustStsSettingsIssuerCertificateId, serverSettingsWsTrustStsSettingsIssuerCertificateInfo := range serverSettingsWsTrustStsSettingsIssuerCertificateData {
		serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDn := serverSettingsWsTrustStsSettingsIssuerCertificateInfo[0]
		serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumber := serverSettingsWsTrustStsSettingsIssuerCertificateInfo[1]

		commentData := map[string]string{
			"Server Settings Ws Trust Sts Settings Issuer Certificate ID":            serverSettingsWsTrustStsSettingsIssuerCertificateId,
			"Server Settings Ws Trust Sts Settings Issuer Certificate Issuer DN":     serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDn,
			"Server Settings Ws Trust Sts Settings Issuer Certificate Serial Number": serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumber,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDn, serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumber),
			ResourceID:         serverSettingsWsTrustStsSettingsIssuerCertificateId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateServerSettingsWsTrustStsSettingsIssuerCertificateResource) getServerSettingsWsTrustStsSettingsIssuerCertificateData() (map[string][]string, error) {
	serverSettingsWsTrustStsSettingsIssuerCertificateData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.ServerSettingsAPI.GetCerts(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetCerts", r.ResourceType())
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

	for _, serverSettingsWsTrustStsSettingsIssuerCertificate := range items {
		serverSettingsWsTrustStsSettingsIssuerCertificateCertView, serverSettingsWsTrustStsSettingsIssuerCertificateCertViewOk := serverSettingsWsTrustStsSettingsIssuerCertificate.GetCertViewOk()

		if serverSettingsWsTrustStsSettingsIssuerCertificateCertViewOk {
			serverSettingsWsTrustStsSettingsIssuerCertificateId, serverSettingsWsTrustStsSettingsIssuerCertificateIdOk := serverSettingsWsTrustStsSettingsIssuerCertificateCertView.GetIdOk()
			serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDn, serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDnOk := serverSettingsWsTrustStsSettingsIssuerCertificateCertView.GetIssuerDNOk()
			serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumber, serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumberOk := serverSettingsWsTrustStsSettingsIssuerCertificateCertView.GetSerialNumberOk()

			if serverSettingsWsTrustStsSettingsIssuerCertificateIdOk && serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDnOk && serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumberOk {
				serverSettingsWsTrustStsSettingsIssuerCertificateData[*serverSettingsWsTrustStsSettingsIssuerCertificateId] = []string{*serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDn, *serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumber}
			}
		}
	}

	return serverSettingsWsTrustStsSettingsIssuerCertificateData, nil
}
