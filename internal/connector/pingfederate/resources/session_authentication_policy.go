package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSessionAuthenticationPolicyResource{}
)

type PingFederateSessionAuthenticationPolicyResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateSessionAuthenticationPolicyResource
func SessionAuthenticationPolicy(clientInfo *connector.PingFederateClientInfo) *PingFederateSessionAuthenticationPolicyResource {
	return &PingFederateSessionAuthenticationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSessionAuthenticationPolicyResource) ResourceType() string {
	return "pingfederate_session_authentication_policy"
}

func (r *PingFederateSessionAuthenticationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	sessionAuthenticationPolicyData, err := r.getSessionAuthenticationPolicyData()
	if err != nil {
		return nil, err
	}

	for sessionAuthenticationPolicyId, sessionAuthenticationPolicyInfo := range *sessionAuthenticationPolicyData {
		sessionAuthenticationPolicyAuthenticationSourceType := sessionAuthenticationPolicyInfo[0]
		sessionAuthenticationPolicyAuthenticationSourceSourceRefId := sessionAuthenticationPolicyInfo[1]
		commentData := map[string]string{
			"Session Authentication Policy ID":                           sessionAuthenticationPolicyId,
			"Session Authentication Policy Authentication Source Type":   sessionAuthenticationPolicyAuthenticationSourceType,
			"Session Authentication Policy Authentication Source Ref ID": sessionAuthenticationPolicyAuthenticationSourceSourceRefId,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s_%s", sessionAuthenticationPolicyId, sessionAuthenticationPolicyAuthenticationSourceType, sessionAuthenticationPolicyAuthenticationSourceSourceRefId),
			ResourceID:         sessionAuthenticationPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSessionAuthenticationPolicyResource) getSessionAuthenticationPolicyData() (*map[string][]string, error) {
	sessionAuthenticationPolicyData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.ApiClient.SessionAPI.GetSourcePolicies(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetSourcePolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, sessionAuthenticationPolicy := range items {
		sessionAuthenticationPolicyId, sessionAuthenticationPolicyIdOk := sessionAuthenticationPolicy.GetIdOk()
		sessionAuthenticationPolicyAuthenticationSource, sessionAuthenticationPolicyAuthenticationSourceOk := sessionAuthenticationPolicy.GetAuthenticationSourceOk()

		if sessionAuthenticationPolicyIdOk && sessionAuthenticationPolicyAuthenticationSourceOk {
			sessionAuthenticationPolicyAuthenticationSourceType, sessionAuthenticationPolicyAuthenticationSourceTypeOk := sessionAuthenticationPolicyAuthenticationSource.GetTypeOk()
			sessionAuthenticationPolicyAuthenticationSourceSourceRef, sessionAuthenticationPolicyAuthenticationSourceSourceRefOk := sessionAuthenticationPolicyAuthenticationSource.GetSourceRefOk()

			if sessionAuthenticationPolicyAuthenticationSourceTypeOk && sessionAuthenticationPolicyAuthenticationSourceSourceRefOk {
				sessionAuthenticationPolicyAuthenticationSourceSourceRefId, sessionAuthenticationPolicyAuthenticationSourceSourceRefIdOk := sessionAuthenticationPolicyAuthenticationSourceSourceRef.GetIdOk()

				if sessionAuthenticationPolicyAuthenticationSourceSourceRefIdOk {
					sessionAuthenticationPolicyData[*sessionAuthenticationPolicyId] = []string{*sessionAuthenticationPolicyAuthenticationSourceType, *sessionAuthenticationPolicyAuthenticationSourceSourceRefId}
				}
			}
		}
	}

	return &sessionAuthenticationPolicyData, nil
}
