package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateVirtualHostNamesResource{}
)

type PingFederateVirtualHostNamesResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateVirtualHostNamesResource
func VirtualHostNames(clientInfo *connector.PingFederateClientInfo) *PingFederateVirtualHostNamesResource {
	return &PingFederateVirtualHostNamesResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateVirtualHostNamesResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	virtualHostNamesId := "virtual_host_names_singleton_id"
	virtualHostNamesName := "Virtual Host Names"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       virtualHostNamesName,
		ResourceID:         virtualHostNamesId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateVirtualHostNamesResource) ResourceType() string {
	return "pingfederate_virtual_host_names"
}
