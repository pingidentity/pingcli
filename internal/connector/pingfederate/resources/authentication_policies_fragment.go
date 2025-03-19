// Copyright © 2025 Ping Identity Corporation

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateAuthenticationPoliciesFragmentResource
func AuthenticationPoliciesFragment(clientInfo *connector.ClientInfo) *PingFederateAuthenticationPoliciesFragmentResource {
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

	authenticationPoliciesFragmentData, err := r.getAuthenticationPoliciesFragmentData()
	if err != nil {
		return nil, err
	}

	for authenticationPoliciesFragmentId, authenticationPoliciesFragmentName := range authenticationPoliciesFragmentData {
		commentData := map[string]string{
			"Authentication Policies Fragment ID":   authenticationPoliciesFragmentId,
			"Authentication Policies Fragment Name": authenticationPoliciesFragmentName,
			"Resource Type":                         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       authenticationPoliciesFragmentName,
			ResourceID:         authenticationPoliciesFragmentId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationPoliciesFragmentResource) getAuthenticationPoliciesFragmentData() (map[string]string, error) {
	authenticationPoliciesFragmentData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.AuthenticationPoliciesAPI.GetFragments(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetFragments", r.ResourceType())
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

	for _, authenticationPoliciesFragment := range items {
		authenticationPoliciesFragmentId, authenticationPoliciesFragmentIdOk := authenticationPoliciesFragment.GetIdOk()
		authenticationPoliciesFragmentName, authenticationPoliciesFragmentNameOk := authenticationPoliciesFragment.GetNameOk()

		if authenticationPoliciesFragmentIdOk && authenticationPoliciesFragmentNameOk {
			authenticationPoliciesFragmentData[*authenticationPoliciesFragmentId] = *authenticationPoliciesFragmentName
		}
	}

	return authenticationPoliciesFragmentData, nil
}
