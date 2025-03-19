// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateCertificatesRevocationOcspCertificateResource{}
)

type PingFederateCertificatesRevocationOcspCertificateResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateCertificatesRevocationOcspCertificateResource
func CertificatesRevocationOcspCertificate(clientInfo *connector.ClientInfo) *PingFederateCertificatesRevocationOcspCertificateResource {
	return &PingFederateCertificatesRevocationOcspCertificateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateCertificatesRevocationOcspCertificateResource) ResourceType() string {
	return "pingfederate_certificates_revocation_ocsp_certificate"
}

func (r *PingFederateCertificatesRevocationOcspCertificateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	certificatesRevocationOcspCertificateData, err := r.getCertificatesRevocationOcspCertificateData()
	if err != nil {
		return nil, err
	}

	for certificatesRevocationOcspCertificateId, certificatesRevocationOcspCertificateInfo := range certificatesRevocationOcspCertificateData {
		certificatesRevocationOcspCertificateIssuerDn := certificatesRevocationOcspCertificateInfo[0]
		certificatesRevocationOcspCertificateSerialNumber := certificatesRevocationOcspCertificateInfo[1]

		commentData := map[string]string{
			"Certificates Revocation Ocsp Certificate ID":            certificatesRevocationOcspCertificateId,
			"Certificates Revocation Ocsp Certificate Issuer DN":     certificatesRevocationOcspCertificateIssuerDn,
			"Certificates Revocation Ocsp Certificate Serial Number": certificatesRevocationOcspCertificateSerialNumber,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", certificatesRevocationOcspCertificateIssuerDn, certificatesRevocationOcspCertificateSerialNumber),
			ResourceID:         certificatesRevocationOcspCertificateId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateCertificatesRevocationOcspCertificateResource) getCertificatesRevocationOcspCertificateData() (map[string][]string, error) {
	certificatesRevocationOcspCertificateData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.CertificatesRevocationAPI.GetOcspCertificates(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetOcspCertificates", r.ResourceType())
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

	for _, certificatesRevocationOcspCertificate := range items {
		certificatesRevocationOcspCertificateId, certificatesRevocationOcspCertificateIdOk := certificatesRevocationOcspCertificate.GetIdOk()
		certificatesRevocationOcspCertificateIssuerDn, certificatesRevocationOcspCertificateIssuerDnOk := certificatesRevocationOcspCertificate.GetIssuerDNOk()
		certificatesRevocationOcspCertificateSerialNumber, certificatesRevocationOcspCertificateSerialNumberOk := certificatesRevocationOcspCertificate.GetSerialNumberOk()

		if certificatesRevocationOcspCertificateIdOk && certificatesRevocationOcspCertificateIssuerDnOk && certificatesRevocationOcspCertificateSerialNumberOk {
			certificatesRevocationOcspCertificateData[*certificatesRevocationOcspCertificateId] = []string{*certificatesRevocationOcspCertificateIssuerDn, *certificatesRevocationOcspCertificateSerialNumber}
		}
	}

	return certificatesRevocationOcspCertificateData, nil
}
