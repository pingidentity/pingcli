package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateCertificatesRevocationOCSPCertificateResource{}
)

type PingFederateCertificatesRevocationOCSPCertificateResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateCertificatesRevocationOCSPCertificateResource
func CertificatesRevocationOCSPCertificate(clientInfo *connector.PingFederateClientInfo) *PingFederateCertificatesRevocationOCSPCertificateResource {
	return &PingFederateCertificatesRevocationOCSPCertificateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateCertificatesRevocationOCSPCertificateResource) ResourceType() string {
	return "pingfederate_certificates_revocation_ocsp_certificate"
}

func (r *PingFederateCertificatesRevocationOCSPCertificateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ocspCertificateData, err := r.getOcspCertificateData()
	if err != nil {
		return nil, err
	}

	for ocspCertificateId, ocspCertificateInfo := range ocspCertificateData {
		ocspCertificateIssuerDN := ocspCertificateInfo[0]
		ocspCertificateSerialNumber := ocspCertificateInfo[1]

		commentData := map[string]string{
			"Certificate Revocation OCSP Certificate ID":            ocspCertificateId,
			"Certificate Revocation OCSP Certificate Issuer DN":     ocspCertificateIssuerDN,
			"Certificate Revocation OCSP Certificate Serial Number": ocspCertificateSerialNumber,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", ocspCertificateIssuerDN, ocspCertificateSerialNumber),
			ResourceID:         ocspCertificateId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateCertificatesRevocationOCSPCertificateResource) getOcspCertificateData() (map[string][]string, error) {
	ocspCertificateData := make(map[string][]string)

	ocspCertificates, response, err := r.clientInfo.ApiClient.CertificatesRevocationAPI.GetOcspCertificates(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetOcspCertificates", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if ocspCertificates == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	ocspCertificatesItems, ocspCertificatesItemsOk := ocspCertificates.GetItemsOk()
	if !ocspCertificatesItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, ocspCertificate := range ocspCertificatesItems {
		ocspCertificateId, ocspCertificateIdOk := ocspCertificate.GetIdOk()
		ocspCertificateIssuerDN, ocspCertificateIssuerDNOk := ocspCertificate.GetIssuerDNOk()
		ocspCertificateSerialNumber, ocspCertificateSerialNumberOk := ocspCertificate.GetSerialNumberOk()

		if ocspCertificateIdOk && ocspCertificateIssuerDNOk && ocspCertificateSerialNumberOk {
			ocspCertificateData[*ocspCertificateId] = []string{*ocspCertificateIssuerDN, *ocspCertificateSerialNumber}
		}
	}

	return ocspCertificateData, nil
}
