package resources

import (
	"fmt"

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

func (r *PingFederateLocalIdentityProfileResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.LocalIdentityIdentityProfilesAPI.GetIdentityProfiles(r.clientInfo.Context).Execute
	apiFunctionName := "GetIdentityProfiles"

	localIdentityProfiles, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if localIdentityProfiles == nil {
		l.Error().Msgf("Returned %s() localIdentityProfiles is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	localIdentityProfilesItems, ok := localIdentityProfiles.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() localIdentityProfiles items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, localIdentityProfile := range localIdentityProfilesItems {
		localIdentityProfileId, localIdentityProfileIdOk := localIdentityProfile.GetIdOk()
		localIdentityProfileName, localIdentityProfileNameOk := localIdentityProfile.GetNameOk()

		if localIdentityProfileIdOk && localIdentityProfileNameOk {
			commentData := map[string]string{
				"Resource Type":                        r.ResourceType(),
				"Local Identity Profile Resource ID":   *localIdentityProfileId,
				"Local Identity Profile Resource Name": *localIdentityProfileName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *localIdentityProfileName,
				ResourceID:         *localIdentityProfileId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateLocalIdentityProfileResource) ResourceType() string {
	return "pingfederate_local_identity_profile"
}
