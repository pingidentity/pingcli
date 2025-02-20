package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
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

	for parentGroupId, parentGroupName := range groupData {
		groupNestingData, err := r.getGroupNestingData(parentGroupId)
		if err != nil {
			return nil, err
		}

		for nestedGroupId, nestedGroupName := range groupNestingData {
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

func (r *PingOneGroupNestingResource) getGroupData() (map[string]string, error) {
	groupData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	groups, err := common.GetManagementAPIObjectsFromIterator[management.Group](iter, "ReadAllGroups", "GetGroups", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, parentGroup := range groups {
		parentGroupId, parentGroupIdOk := parentGroup.GetIdOk()
		parentGroupName, parentGroupNameOk := parentGroup.GetNameOk()

		if parentGroupIdOk && parentGroupNameOk {
			groupData[*parentGroupId] = *parentGroupName
		}
	}

	return groupData, nil
}

func (r *PingOneGroupNestingResource) getGroupNestingData(parentGroupId string) (map[string]string, error) {
	groupNestingData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadGroupNesting(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, parentGroupId).Execute()
	groupNestings, err := common.GetManagementAPIObjectsFromIterator[management.GroupMembership](iter, "ReadGroupNesting", "GetGroupMemberships", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, nestedGroup := range groupNestings {
		nestedGroupId, nestedGroupIdOk := nestedGroup.GetIdOk()
		nestedGroupName, nestedGroupNameOk := nestedGroup.GetNameOk()

		if nestedGroupIdOk && nestedGroupNameOk {
			groupNestingData[*nestedGroupId] = *nestedGroupName
		}
	}

	return groupNestingData, nil
}
