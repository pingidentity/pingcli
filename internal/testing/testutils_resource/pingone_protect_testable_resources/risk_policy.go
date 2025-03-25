// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_protect_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
)

func RiskPolicy(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createRiskPolicy,
		DeleteFunc:         deleteRiskPolicy,
		Dependencies:       nil,
		ExportableResource: resources.RiskPolicy(clientInfo),
	}
}

func createRiskPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createRiskPolicy(): %v", strArgs)
		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingOneApiClient.RiskAPIClient.RiskPoliciesApi.CreateRiskPolicySet(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := risk.RiskPolicySet{
		Name:    "Score-based policy",
		Default: utils.Pointer(false),
		DefaultResult: &risk.RiskPolicySetDefaultResult{
			Level: risk.ENUMRISKPOLICYRESULTLEVEL_LOW,
		},
		RiskPolicies: []risk.RiskPolicy{
			{
				Name: "ANONYMOUS_NETWORK_DETECTION",
				Result: risk.RiskPolicyResult{
					Level: risk.ENUMRISKLEVEL_HIGH,
				},
				Condition: risk.RiskPolicyCondition{
					Value: utils.Pointer("${details.anonymousNetworkDetected}"),
					Equals: &risk.RiskPolicyConditionEquals{
						Bool: utils.Pointer(true),
					},
				},
			},
			{
				Name: "GEOVELOCITY_ANOMALY",
				Result: risk.RiskPolicyResult{
					Level: risk.ENUMRISKLEVEL_MEDIUM,
				},
				Condition: risk.RiskPolicyCondition{
					Value: utils.Pointer("${details.impossibleTravel}"),
					Equals: &risk.RiskPolicyConditionEquals{
						Bool: utils.Pointer(true),
					},
				},
			},
		},
	}

	request = request.RiskPolicySet(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateRiskPolicySet", resourceType)
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
		},
	}
}

func deleteRiskPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteRiskPolicy(): %v", ids)
		return
	}

	request := clientInfo.PingOneApiClient.RiskAPIClient.RiskPoliciesApi.DeleteRiskPolicySet(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteRiskPolicySet", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}
}
