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

func Test_PingFederateSessionAuthenticationPolicy_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SessionAuthenticationPolicy(PingFederateClientInfo)

	testPCVId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)

	testIdpAdapterId, _ := createIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)
	defer deleteIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testIdpAdapterId)

	sessionAuthenticationPolicyId, sessionAuthenticationPolicyAuthSourceType, sessionAuthenticationPolicyAuthSourceRefId := createSessionAuthenticationPolicy(t, PingFederateClientInfo, resource.ResourceType(), testIdpAdapterId)
	defer deleteSessionAuthenticationPolicy(t, PingFederateClientInfo, resource.ResourceType(), sessionAuthenticationPolicyId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s_%s", sessionAuthenticationPolicyId, sessionAuthenticationPolicyAuthSourceType, sessionAuthenticationPolicyAuthSourceRefId),
			ResourceID:   sessionAuthenticationPolicyId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testIdpAdapterId string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.SessionAPI.CreateSourcePolicy(clientInfo.Context)
	result := client.AuthenticationSessionPolicy{
		AuthenticationSource: client.AuthenticationSource{
			SourceRef: client.ResourceLink{
				Id: testIdpAdapterId,
			},
			Type: "IDP_ADAPTER",
		},
		Id: utils.Pointer("TestSessionAuthenticationPolicyId"),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateSourcePolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.AuthenticationSource.Type, resource.AuthenticationSource.SourceRef.Id
}

func deleteSessionAuthenticationPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.SessionAPI.DeleteSourcePolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSourcePolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
