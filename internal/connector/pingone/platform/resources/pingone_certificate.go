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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneCertificateResource
func Certificate(clientInfo *connector.PingOneClientInfo) *PingOneCertificateResource {
	return &PingOneCertificateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneCertificateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetCertificates(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "GetCertificates"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, certificate := range embedded.GetCertificates() {
		certificateName, certificateNameOk := certificate.GetNameOk()
		certificateId, certificateIdOk := certificate.GetIdOk()

		if certificateNameOk && certificateIdOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Certificate Name":      *certificateName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Certificate ID":        *certificateId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *certificateName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *certificateId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneCertificateResource) ResourceType() string {
	return "pingone_certificate"
}
