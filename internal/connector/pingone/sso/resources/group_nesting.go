// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGroupNestingResource{}
)

type PingOneGroupNestingResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneGroupNestingResource
func GroupNesting(clientInfo *connector.ClientInfo) *PingOneGroupNestingResource {
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

	for groupId, groupName := range groupData {
		groupNestingData, err := r.getGroupNestingData(groupId)
		if err != nil {
			return nil, err
		}

		for nestedGroupId, nestedGroupName := range groupNestingData {
			commentData := map[string]string{
				"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
				"Nested Group ID":       nestedGroupId,
				"Nested Group Name":     nestedGroupName,
				"Parent Group ID":       groupId,
				"Parent Group Name":     groupName,
				"Resource Type":         r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", groupName, nestedGroupName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, groupId, nestedGroupId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneGroupNestingResource) getGroupData() (map[string]string, error) {
	groupData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.Group](iter, "ReadAllGroups", "GetGroups", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, group := range apiObjs {
		groupId, groupIdOk := group.GetIdOk()
		groupName, groupNameOk := group.GetNameOk()

		if groupIdOk && groupNameOk {
			groupData[*groupId] = *groupName
		}
	}

	return groupData, nil
}

func (r *PingOneGroupNestingResource) getGroupNestingData(groupId string) (map[string]string, error) {
	groupNestingData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GroupsApi.ReadGroupNesting(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, groupId).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.GroupMembership](iter, "ReadGroupNesting", "GetGroupMemberships", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, nestedGroup := range apiObjs {
		nestedGroupId, nestedGroupIdOk := nestedGroup.GetIdOk()
		nestedGroupName, nestedGroupNameOk := nestedGroup.GetNameOk()

		if nestedGroupIdOk && nestedGroupNameOk {
			groupNestingData[*nestedGroupId] = *nestedGroupName
		}
	}

	return groupNestingData, nil
}
