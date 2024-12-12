package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGroupNestingResource{}
)

type PingOneGroupNestingResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneGroupNestingResource
func GroupNesting(clientInfo *connector.PingOneClientInfo) *PingOneGroupNestingResource {
	return &PingOneGroupNestingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneGroupNestingResource) ResourceType() string {
	return "pingone_group_nesting"
}

func (r *PingOneGroupNestingResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	groupData, err := r.getGroupData()
	if err != nil {
		return nil, err
	}

	for parentGroupId, parentGroupName := range *groupData {
		groupNestingData, err := r.getGroupNestingData(parentGroupId)
		if err != nil {
			return nil, err
		}

		for nestedGroupId, nestedGroupName := range *groupNestingData {
			commentData := map[string]string{
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Nested Group ID":       nestedGroupId,
				"Nested Group Name":     nestedGroupName,
				"Parent Group ID":       parentGroupId,
				"Parent Group Name":     parentGroupName,
				"Resource Type":         r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", parentGroupName, nestedGroupName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, parentGroupId, nestedGroupId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneGroupNestingResource) getGroupData() (*map[string]string, error) {
	groupData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGroups", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, parentGroup := range embedded.GetGroups() {
			parentGroupId, parentGroupIdOk := parentGroup.GetIdOk()
			parentGroupName, parentGroupNameOk := parentGroup.GetNameOk()

			if parentGroupIdOk && parentGroupNameOk {
				groupData[*parentGroupId] = *parentGroupName
			}
		}
	}

	return &groupData, nil
}

func (r *PingOneGroupNestingResource) getGroupNestingData(parentGroupId string) (*map[string]string, error) {
	groupNestingData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadGroupNesting(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, parentGroupId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadGroupNesting", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, nestedGroup := range embedded.GetGroupMemberships() {
			nestedGroupId, nestedGroupIdOk := nestedGroup.GetIdOk()
			nestedGroupName, nestedGroupNameOk := nestedGroup.GetNameOk()

			if nestedGroupIdOk && nestedGroupNameOk {
				groupNestingData[*nestedGroupId] = *nestedGroupName
			}
		}
	}

	return &groupNestingData, nil
}
