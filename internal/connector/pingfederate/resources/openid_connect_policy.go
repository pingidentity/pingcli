package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateOpenidConnectPolicyResource{}
)

type PingFederateOpenidConnectPolicyResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateOpenidConnectPolicyResource
func OpenidConnectPolicy(clientInfo *connector.ClientInfo) *PingFederateOpenidConnectPolicyResource {
	return &PingFederateOpenidConnectPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateOpenidConnectPolicyResource) ResourceType() string {
	return "pingfederate_openid_connect_policy"
}

func (r *PingFederateOpenidConnectPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	openidConnectPolicyData, err := r.getOpenidConnectPolicyData()
	if err != nil {
		return nil, err
	}

	for openidConnectPolicyId, openidConnectPolicyName := range openidConnectPolicyData {
		commentData := map[string]string{
			"Openid Connect Policy ID":   openidConnectPolicyId,
			"Openid Connect Policy Name": openidConnectPolicyName,
			"Resource Type":              r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       openidConnectPolicyName,
			ResourceID:         openidConnectPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateOpenidConnectPolicyResource) getOpenidConnectPolicyData() (map[string]string, error) {
	openidConnectPolicyData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.OauthOpenIdConnectAPI.GetOIDCPolicies(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetOIDCPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, openidConnectPolicy := range items {
		openidConnectPolicyId, openidConnectPolicyIdOk := openidConnectPolicy.GetIdOk()
		openidConnectPolicyName, openidConnectPolicyNameOk := openidConnectPolicy.GetNameOk()

		if openidConnectPolicyIdOk && openidConnectPolicyNameOk {
			openidConnectPolicyData[*openidConnectPolicyId] = *openidConnectPolicyName
		}
	}

	return openidConnectPolicyData, nil
}
