package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGroupResource{}
)

type PingOneGroupResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneGroupResource
func Group(clientInfo *connector.PingOneClientInfo) *PingOneGroupResource {
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

	for groupId, groupName := range *groupData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Group ID":              groupId,
			"Group Name":            groupName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       groupName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, groupId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneGroupResource) getGroupData() (*map[string]string, error) {
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

		for _, group := range embedded.GetGroups() {
			groupId, groupIdOk := group.GetIdOk()
			groupName, groupNameOk := group.GetNameOk()

			if groupIdOk && groupNameOk {
				groupData[*groupId] = *groupName
			}
		}
	}

	return &groupData, nil
}
