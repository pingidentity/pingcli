package resources

import (
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

func (r *PingFederateAuthenticationPoliciesFragmentResource) ResourceType() string {
	return "pingfederate_authentication_policies_fragment"
}

func (r *PingFederateAuthenticationPoliciesFragmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	fragmentData, err := r.getFragmentData()
	if err != nil {
		return nil, err
	}

	for fragmentId, fragmentName := range fragmentData {
		commentData := map[string]string{
			"Authentication Policies Fragment ID":   fragmentId,
			"Authentication Policies Fragment Name": fragmentName,
			"Resource Type":                         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fragmentName,
			ResourceID:         fragmentId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationPoliciesFragmentResource) getFragmentData() (map[string]string, error) {
	fragmentData := make(map[string]string)

	authnPoliciesFragments, response, err := r.clientInfo.ApiClient.AuthenticationPoliciesAPI.GetFragments(r.clientInfo.Context).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetFragments", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if authnPoliciesFragments == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	authnPoliciesFragmentsItems, authnPoliciesFragmentsItemsOk := authnPoliciesFragments.GetItemsOk()
	if !authnPoliciesFragmentsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, authnPoliciesFragment := range authnPoliciesFragmentsItems {
		authnPoliciesFragmentId, authnPoliciesFragmentIdOk := authnPoliciesFragment.GetIdOk()
		authnPoliciesFragmentName, authnPoliciesFragmentNameOk := authnPoliciesFragment.GetNameOk()

		if authnPoliciesFragmentIdOk && authnPoliciesFragmentNameOk {
			fragmentData[*authnPoliciesFragmentId] = *authnPoliciesFragmentName
		}
	}

	return fragmentData, nil
}
