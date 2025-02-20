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

func (r *PingOneCertificateResource) ResourceType() string {
	return "pingone_certificate"
}

func (r *PingOneCertificateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	certificateData, err := getCertificateData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for certificateId, certificateName := range certificateData {
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

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func getCertificateData(clientInfo *connector.PingOneClientInfo, resourceType string) (map[string]string, error) {
	certificateData := make(map[string]string)

	// TODO: Implement pagination once supported in the PingOne Go Client SDK
	entityArray, response, err := clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetCertificates(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetCertificates", resourceType)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if entityArray == nil {
		return nil, fmt.Errorf("failed to export resource '%s'.\n"+
			"PingOne API request for resource '%s' was not successful. response data is nil.\n"+
			"response code: %s\n"+
			"response body: %s",
			resourceType, resourceType, response.Status, response.Body)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		return nil, fmt.Errorf("failed to export resource '%s'.\n"+
			"PingOne API request for resource '%s' was not successful. response data is nil.\n"+
			"response code: %s\n"+
			"response body: %s",
			resourceType, resourceType, response.Status, response.Body)
	}

	for _, certificate := range embedded.GetCertificates() {
		certificateId, certificateIdOk := certificate.GetIdOk()
		certificateName, certificateNameOk := certificate.GetNameOk()

		if certificateIdOk && certificateNameOk {
			certificateData[*certificateId] = *certificateName
		}
	}

	return certificateData, nil
}
