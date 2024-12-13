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
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateLocalIdentityProfileResource
func LocalIdentityProfile(clientInfo *connector.PingFederateClientInfo) *PingFederateLocalIdentityProfileResource {
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

	identityProfileData, err := r.getIdentityProfileData()
	if err != nil {
		return nil, err
	}

	for identityProfileId, identityProfileName := range *identityProfileData {
		commentData := map[string]string{
			"Local Identity Profile ID":   identityProfileId,
			"Local Identity Profile Name": identityProfileName,
			"Resource Type":               r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       identityProfileName,
			ResourceID:         identityProfileId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateLocalIdentityProfileResource) getIdentityProfileData() (*map[string]string, error) {
	identityProfileData := make(map[string]string)

	identityProfiles, response, err := r.clientInfo.ApiClient.LocalIdentityIdentityProfilesAPI.GetIdentityProfiles(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetIdentityProfiles", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if identityProfiles == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	identityProfilesItems, identityProfilesItemsOk := identityProfiles.GetItemsOk()
	if !identityProfilesItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, identityProfile := range identityProfilesItems {
		identityProfileId, identityProfileIdOk := identityProfile.GetIdOk()
		identityProfileName, identityProfileNameOk := identityProfile.GetNameOk()

		if identityProfileIdOk && identityProfileNameOk {
			identityProfileData[*identityProfileId] = *identityProfileName
		}
	}

	return &identityProfileData, nil
}
