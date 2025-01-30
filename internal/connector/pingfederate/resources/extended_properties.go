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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateExtendedPropertiesResource
func ExtendedProperties(clientInfo *connector.ClientInfo) *PingFederateExtendedPropertiesResource {
	return &PingFederateExtendedPropertiesResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateExtendedPropertiesResource) ResourceType() string {
	return "pingfederate_extended_properties"
}

func (r *PingFederateExtendedPropertiesResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	extendedPropertiesId := "extended_properties_singleton_id"
	extendedPropertiesName := "Extended Properties"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       extendedPropertiesName,
		ResourceID:         extendedPropertiesId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
