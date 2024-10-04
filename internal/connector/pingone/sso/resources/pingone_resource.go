package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceResource{}
)

type PingoneResourceResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneResourceResource
func Resource(clientInfo *connector.PingOneClientInfo) *PingoneResourceResource {
	return &PingoneResourceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllResources"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, resourceInner := range embedded.GetResources() {
		resource := resourceInner.Resource
		resourceId, resourceIdOk := resource.GetIdOk()
		resourceName, resourceNameOk := resource.GetNameOk()

		if resourceIdOk && resourceNameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Resource Name":         *resourceName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Resource ID":           *resourceId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *resourceName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *resourceId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceResource) ResourceType() string {
	return "pingone_resource"
}
