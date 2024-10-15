package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationPoliciesFragmentResource{}
)

type PingFederateAuthenticationPoliciesFragmentResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationPoliciesFragmentResource
func AuthenticationPoliciesFragment(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationPoliciesFragmentResource {
	return &PingFederateAuthenticationPoliciesFragmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationPoliciesFragmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthenticationPoliciesAPI.GetFragments(r.clientInfo.Context).Execute
	apiFunctionName := "GetFragments"

	authnPoliciesFragments, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if authnPoliciesFragments == nil {
		l.Error().Msgf("Returned %s() authnPoliciesFragments is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	authnPoliciesFragmentsItems, ok := authnPoliciesFragments.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() authnPoliciesFragment items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authnPoliciesFragment := range authnPoliciesFragmentsItems {
		authnPoliciesFragmentId, authnPoliciesFragmentIdOk := authnPoliciesFragment.GetIdOk()
		authnPoliciesFragmentName, authnPoliciesFragmentNameOk := authnPoliciesFragment.GetNameOk()

		if authnPoliciesFragmentIdOk && authnPoliciesFragmentNameOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authentication Policies Fragment Resource ID":   *authnPoliciesFragmentId,
				"Authentication Policies Fragment Resource Name": *authnPoliciesFragmentName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authnPoliciesFragmentName,
				ResourceID:         *authnPoliciesFragmentId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationPoliciesFragmentResource) ResourceType() string {
	return "pingfederate_authentication_policies_fragment"
}
