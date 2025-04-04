// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func ServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createServerSettingsWsTrustStsSettingsIssuerCertificate,
		DeleteFunc:         deleteServerSettingsWsTrustStsSettingsIssuerCertificate,
		Dependencies:       nil,
		ExportableResource: resources.ServerSettingsWsTrustStsSettingsIssuerCertificate(clientInfo),
	}
}

func createServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Fatalf("Unexpected number of arguments provided to createServerSettingsWsTrustStsSettingsIssuerCertificate(): %v", strArgs)
	}

	fileData, err := testutils.CreateX509Certificate()
	if err != nil {
		t.Errorf("Failed to create test %s: %v", resourceType, err)

		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingFederateApiClient.ServerSettingsAPI.ImportCertificate(clientInfo.PingFederateContext)
	clientStruct := client.X509File{
		FileData: fileData,
		Id:       utils.Pointer("testx509fileid"),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "ImportCertificate", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			*resource.CertView.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:            *resource.CertView.Id,
			testutils_resource.ENUM_ISSUER_DN:     *resource.CertView.IssuerDN,
			testutils_resource.ENUM_SERIAL_NUMBER: *resource.CertView.SerialNumber,
		},
	}
}

func deleteServerSettingsWsTrustStsSettingsIssuerCertificate(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteServerSettingsWsTrustStsSettingsIssuerCertificate(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.ServerSettingsAPI.DeleteCertificate(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteCertificate", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
