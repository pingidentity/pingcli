package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateExtendedPropertiesResource{}
)

type PingFederateExtendedPropertiesResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateExtendedPropertiesResource
func ExtendedProperties(clientInfo *connector.PingFederateClientInfo) *PingFederateExtendedPropertiesResource {
	return &PingFederateExtendedPropertiesResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateExtendedPropertiesResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	extendedPropertiesId := "extended_properties_singleton_id"
	extendedPropertiesName := "Extended Properties"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       extendedPropertiesName,
		ResourceID:         extendedPropertiesId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateExtendedPropertiesResource) ResourceType() string {
	return "pingfederate_extended_properties"
}
