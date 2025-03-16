package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateCertificateCaResource{}
)

type PingFederateCertificateCaResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateCertificateCaResource
func CertificateCa(clientInfo *connector.ClientInfo) *PingFederateCertificateCaResource {
	return &PingFederateCertificateCaResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateCertificateCaResource) ResourceType() string {
	return "pingfederate_certificate_ca"
}

func (r *PingFederateCertificateCaResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	certificateCaData, err := r.getCertificateCaData()
	if err != nil {
		return nil, err
	}

	for certificateCaId, certificateCaInfo := range certificateCaData {
		certificateCaIssuerDn := certificateCaInfo[0]
		certificateCaSerialNumber := certificateCaInfo[1]

		commentData := map[string]string{
			"Certificate Ca ID":            certificateCaId,
			"Certificate Ca Issuer DN":     certificateCaIssuerDn,
			"Certificate Ca Serial Number": certificateCaSerialNumber,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", certificateCaIssuerDn, certificateCaSerialNumber),
			ResourceID:         certificateCaId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateCertificateCaResource) getCertificateCaData() (map[string][]string, error) {
	certificateCaData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.CertificatesCaAPI.GetTrustedCAs(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetTrustedCAs", r.ResourceType())
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

	for _, certificateCa := range items {
		certificateCaId, certificateCaIdOk := certificateCa.GetIdOk()
		certificateCaIssuerDn, certificateCaIssuerDnOk := certificateCa.GetIssuerDNOk()
		certificateCaSerialNumber, certificateCaSerialNumberOk := certificateCa.GetSerialNumberOk()

		if certificateCaIdOk && certificateCaIssuerDnOk && certificateCaSerialNumberOk {
			certificateCaData[*certificateCaId] = []string{*certificateCaIssuerDn, *certificateCaSerialNumber}
		}
	}

	return certificateCaData, nil
}
