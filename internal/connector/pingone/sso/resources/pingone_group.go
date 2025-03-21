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
	_ connector.ExportableResource = &PingOneGroupResource{}
)

type PingOneGroupResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneGroupResource
func Group(clientInfo *connector.ClientInfo) *PingOneGroupResource {
	return &PingOneGroupResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneGroupResource) ResourceType() string {
	return "pingone_group"
}

func (r *PingOneGroupResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	groupData, err := r.getGroupData()
	if err != nil {
		return nil, err
	}

	for groupId, groupName := range groupData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Group ID":              groupId,
			"Group Name":            groupName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       groupName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, groupId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneGroupResource) getGroupData() (map[string]string, error) {
	groupData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	groups, err := pingone.GetManagementAPIObjectsFromIterator[management.Group](iter, "ReadAllGroups", "GetGroups", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		groupId, groupIdOk := group.GetIdOk()
		groupName, groupNameOk := group.GetNameOk()

		if groupIdOk && groupNameOk {
			groupData[*groupId] = *groupName
		}
	}

	return groupData, nil
}
