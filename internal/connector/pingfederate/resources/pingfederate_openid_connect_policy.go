package resources

import (
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

func (r *PingFederateOpenIDConnectPolicyResource) ResourceType() string {
	return "pingfederate_openid_connect_policy"
}

func (r *PingFederateOpenIDConnectPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	oidcPolicyData, err := r.getOIDCPolicyData()
	if err != nil {
		return nil, err
	}

	for oidcPolicyId, oidcPolicyName := range *oidcPolicyData {
		commentData := map[string]string{
			"OpenID Connect Policy Resource ID":   oidcPolicyId,
			"OpenID Connect Policy Resource Name": oidcPolicyName,
			"Resource Type":                       r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       oidcPolicyName,
			ResourceID:         oidcPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOpenIDConnectPolicyResource) getOIDCPolicyData() (*map[string]string, error) {
	oidcPolicyData := make(map[string]string)

	oidcPolicies, response, err := r.clientInfo.ApiClient.OauthOpenIdConnectAPI.GetOIDCPolicies(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetOIDCPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	oidcPoliciesItems, ok := oidcPolicies.GetItemsOk()
	if !ok {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, oidcPolicy := range oidcPoliciesItems {
		oidcPolicyId, oidcPolicyIdOk := oidcPolicy.GetIdOk()
		oidcPolicyName, oidcPolicyNameOk := oidcPolicy.GetNameOk()

		if oidcPolicyIdOk && oidcPolicyNameOk {
			oidcPolicyData[*oidcPolicyId] = *oidcPolicyName
		}
	}

	return &oidcPolicyData, nil
}
