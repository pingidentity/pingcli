package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneResourceResource{}
)

type PingOneResourceResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneResourceResource
func Resource(clientInfo *connector.PingOneClientInfo) *PingOneResourceResource {
	return &PingOneResourceResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneResourceResource) ResourceType() string {
	return "pingone_resource"
}

func (r *PingOneResourceResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportResources()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneResourceResource) exportResources() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllResources", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, resourceInner := range embedded.GetResources() {
			if resourceInner.Resource != nil {
				resourceId, resourceIdOk := resourceInner.Resource.GetIdOk()
				resourceName, resourceNameOk := resourceInner.Resource.GetNameOk()

				if resourceIdOk && resourceNameOk {
					r.addImportBlock(*resourceId, *resourceName)
				}
			}
		}
	}

	return nil
}

func (r *PingOneResourceResource) addImportBlock(resourceId, resourceName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource ID":           resourceId,
		"Resource Name":         resourceName,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       resourceName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, resourceId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
