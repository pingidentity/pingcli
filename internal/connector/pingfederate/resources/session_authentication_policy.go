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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSessionAuthenticationPolicyResource
func SessionAuthenticationPolicy(clientInfo *connector.ClientInfo) *PingFederateSessionAuthenticationPolicyResource {
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

	for sessionAuthenticationPolicyId, sessionAuthenticationPolicyInfo := range sessionAuthenticationPolicyData {
		sessionAuthenticationPolicyAuthenticationSourceType := sessionAuthenticationPolicyInfo[0]
		sessionAuthenticationPolicyAuthenticationSourceRefId := sessionAuthenticationPolicyInfo[1]

		commentData := map[string]string{
			"Session Authentication Policy ID":                           sessionAuthenticationPolicyId,
			"Session Authentication Policy Authentication Source Type":   sessionAuthenticationPolicyAuthenticationSourceType,
			"Session Authentication Policy Authentication Source Ref ID": sessionAuthenticationPolicyAuthenticationSourceRefId,
			"Resource Type": r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s_%s", sessionAuthenticationPolicyId, sessionAuthenticationPolicyAuthenticationSourceType, sessionAuthenticationPolicyAuthenticationSourceRefId),
			ResourceID:         sessionAuthenticationPolicyId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateSessionAuthenticationPolicyResource) getSessionAuthenticationPolicyData() (map[string][]string, error) {
	sessionAuthenticationPolicyData := make(map[string][]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.SessionAPI.GetSourcePolicies(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetSourcePolicies", r.ResourceType())
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

	for _, sessionAuthenticationPolicy := range items {
		sessionAuthenticationPolicyId, sessionAuthenticationPolicyIdOk := sessionAuthenticationPolicy.GetIdOk()
		sessionAuthenticationPolicyAuthenticationSource, sessionAuthenticationPolicyAuthenticationSourceOk := sessionAuthenticationPolicy.GetAuthenticationSourceOk()

		if sessionAuthenticationPolicyIdOk && sessionAuthenticationPolicyAuthenticationSourceOk {
			sessionAuthenticationPolicyAuthenticationSourceType, sessionAuthenticationPolicyAuthenticationSourceTypeOk := sessionAuthenticationPolicyAuthenticationSource.GetTypeOk()
			sessionAuthenticationPolicyAuthenticationSourceRef, sessionAuthenticationPolicyAuthenticationSourceRefOk := sessionAuthenticationPolicyAuthenticationSource.GetSourceRefOk()

			if sessionAuthenticationPolicyAuthenticationSourceTypeOk && sessionAuthenticationPolicyAuthenticationSourceRefOk {
				sessionAuthenticationPolicyAuthenticationSourceRefId, sessionAuthenticationPolicyAuthenticationSourceRefIdOk := sessionAuthenticationPolicyAuthenticationSourceRef.GetIdOk()

				if sessionAuthenticationPolicyAuthenticationSourceRefIdOk {
					sessionAuthenticationPolicyData[*sessionAuthenticationPolicyId] = []string{*sessionAuthenticationPolicyAuthenticationSourceType, *sessionAuthenticationPolicyAuthenticationSourceRefId}
				}
			}
		}
	}

	return sessionAuthenticationPolicyData, nil
}
