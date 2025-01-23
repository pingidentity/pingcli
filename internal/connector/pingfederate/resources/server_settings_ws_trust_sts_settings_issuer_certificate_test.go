package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateServerSettingsWsTrustStsSettingsIssuerCertificate_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.ServerSettingsWsTrustStsSettingsIssuerCertificate(PingFederateClientInfo)

	serverSettingsWsTrustStsSettingsIssuerCertificateId, serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDn, serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumber := createServerSettingsWsTrustStsSettingsIssuerCertificate(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteServerSettingsWsTrustStsSettingsIssuerCertificate(t, PingFederateClientInfo, resource.ResourceType(), serverSettingsWsTrustStsSettingsIssuerCertificateId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", serverSettingsWsTrustStsSettingsIssuerCertificateIssuerDn, serverSettingsWsTrustStsSettingsIssuerCertificateSerialNumber),
			ResourceID:   serverSettingsWsTrustStsSettingsIssuerCertificateId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string, string) {
	t.Helper()

	fileData, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	request := clientInfo.ApiClient.ServerSettingsAPI.ImportCertificate(clientInfo.Context)
	result := client.X509File{
		FileData: fileData,
		Id:       utils.Pointer("testx509fileid"),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "ImportCertificate", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.CertView.Id, *resource.CertView.IssuerDN, *resource.CertView.SerialNumber
}

func deleteServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.ServerSettingsAPI.DeleteCertificate(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteCertificate", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
