// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateLocalIdentityProfileResource{}
)

type PingFederateLocalIdentityProfileResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateLocalIdentityProfileResource
func LocalIdentityProfile(clientInfo *connector.ClientInfo) *PingFederateLocalIdentityProfileResource {
	return &PingFederateLocalIdentityProfileResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateLocalIdentityProfileResource) ResourceType() string {
	return "pingfederate_local_identity_profile"
}

func (r *PingFederateLocalIdentityProfileResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	localIdentityProfileData, err := r.getLocalIdentityProfileData()
	if err != nil {
		return nil, err
	}

	for localIdentityProfileId, localIdentityProfileName := range localIdentityProfileData {
		commentData := map[string]string{
			"Local Identity Profile ID":   localIdentityProfileId,
			"Local Identity Profile Name": localIdentityProfileName,
			"Resource Type":               r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       localIdentityProfileName,
			ResourceID:         localIdentityProfileId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateLocalIdentityProfileResource) getLocalIdentityProfileData() (map[string]string, error) {
	localIdentityProfileData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.LocalIdentityIdentityProfilesAPI.GetIdentityProfiles(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetIdentityProfiles", r.ResourceType())
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

	for _, localIdentityProfile := range items {
		localIdentityProfileId, localIdentityProfileIdOk := localIdentityProfile.GetIdOk()
		localIdentityProfileName, localIdentityProfileNameOk := localIdentityProfile.GetNameOk()

		if localIdentityProfileIdOk && localIdentityProfileNameOk {
			localIdentityProfileData[*localIdentityProfileId] = *localIdentityProfileName
		}
	}

	return localIdentityProfileData, nil
}
