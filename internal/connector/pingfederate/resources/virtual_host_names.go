package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateVirtualHostNamesResource{}
)

type PingFederateVirtualHostNamesResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateVirtualHostNamesResource
func VirtualHostNames(clientInfo *connector.ClientInfo) *PingFederateVirtualHostNamesResource {
	return &PingFederateVirtualHostNamesResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateVirtualHostNamesResource) ResourceType() string {
	return "pingfederate_virtual_host_names"
}

func (r *PingFederateVirtualHostNamesResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	virtualHostNamesId := "virtual_host_names_singleton_id"
	virtualHostNamesName := "Virtual Host Names"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       virtualHostNamesName,
		ResourceID:         virtualHostNamesId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
