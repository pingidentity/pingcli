package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOpenIDConnectPolicyResource{}
)

type PingFederateOpenIDConnectPolicyResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateOpenIDConnectPolicyResource
func OpenIDConnectPolicy(clientInfo *connector.PingFederateClientInfo) *PingFederateOpenIDConnectPolicyResource {
	return &PingFederateOpenIDConnectPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOpenIDConnectPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.OauthOpenIdConnectAPI.GetOIDCPolicies(r.clientInfo.Context).Execute
	apiFunctionName := "GetOIDCPolicies"

	openIDConnectPolicies, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if openIDConnectPolicies == nil {
		l.Error().Msgf("Returned %s() openIDConnectPolicies is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	openIDConnectPoliciesItems, ok := openIDConnectPolicies.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() openIDConnectPolicies items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, openIDConnectPolicy := range openIDConnectPoliciesItems {
		openIDConnectPolicyId, openIDConnectPolicyIdOk := openIDConnectPolicy.GetIdOk()
		openIDConnectPolicyName, openIDConnectPolicyNameOk := openIDConnectPolicy.GetNameOk()

		if openIDConnectPolicyIdOk && openIDConnectPolicyNameOk {
			commentData := map[string]string{
				"Resource Type":                       r.ResourceType(),
				"OpenID Connect Policy Resource ID":   *openIDConnectPolicyId,
				"OpenID Connect Policy Resource Name": *openIDConnectPolicyName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *openIDConnectPolicyName,
				ResourceID:         *openIDConnectPolicyId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateOpenIDConnectPolicyResource) ResourceType() string {
	return "pingfederate_openid_connect_policy"
}
