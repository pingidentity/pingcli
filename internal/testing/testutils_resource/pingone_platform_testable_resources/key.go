// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_platform_testable_resources

import (
	"math/big"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
)

func Key(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createKey,
		DeleteFunc:         deleteKey,
		Dependencies:       nil,
		ExportableResource: resources.Key(clientInfo),
	}
}

func createKey(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createKey(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.CertificateManagementApi.CreateKey(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.Certificate{
		Name:               "Doc test cert",
		SerialNumber:       big.NewInt(1575483893597),
		SubjectDN:          "CN=Doc test cert, OU=Ping Identity, O=Ping Identity, L=, ST=, C=US",
		Algorithm:          "RSA",
		KeyLength:          2048,
		ValidityPeriod:     365,
		SignatureAlgorithm: "SHA256withRSA",
		UsageType:          "SIGNING",
		Status:             utils.Pointer(management.ENUMCERTIFICATEKEYSTATUS_VALID),
		Default:            utils.Pointer(false),
	}

	request = request.Certificate(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateKey", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
			testutils_resource.ENUM_TYPE: string(resource.UsageType),
		},
	}
}

func deleteKey(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteKey(): %v", ids)

		return
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.CertificateManagementApi.DeleteKey(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteKey", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
