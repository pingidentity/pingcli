package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneCertificateResource{}
)

type PingOneCertificateResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneCertificateResource
func Certificate(clientInfo *connector.PingOneClientInfo) *PingOneCertificateResource {
	return &PingOneCertificateResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneCertificateResource) ResourceType() string {
	return "pingone_certificate"
}

func (r *PingOneCertificateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportCertificates()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneCertificateResource) exportCertificates() error {
	// TODO: Implement pagination once supported in the PingOne Go Client SDK
	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetCertificates(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "GetCertificates", r.ResourceType())
	if err != nil {
		return err
	}

	if entityArray == nil {
		return common.DataNilError(r.ResourceType(), response)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		return common.DataNilError(r.ResourceType(), response)
	}

	for _, certificate := range embedded.GetCertificates() {
		certificateId, certificateIdOk := certificate.GetIdOk()
		certificateName, certificateNameOk := certificate.GetNameOk()

		if certificateIdOk && certificateNameOk {
			r.addImportBlock(*certificateId, *certificateName)
		}
	}

	return nil
}

func (r *PingOneCertificateResource) addImportBlock(certificateId string, certificateName string) {
	commentData := map[string]string{
		"Certificate ID":        certificateId,
		"Certificate Name":      certificateName,
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       certificateName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, certificateId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
