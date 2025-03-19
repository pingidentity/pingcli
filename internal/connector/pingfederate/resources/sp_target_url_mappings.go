// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateSpTargetUrlMappingsResource{}
)

type PingFederateSpTargetUrlMappingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateSpTargetUrlMappingsResource
func SpTargetUrlMappings(clientInfo *connector.ClientInfo) *PingFederateSpTargetUrlMappingsResource {
	return &PingFederateSpTargetUrlMappingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateSpTargetUrlMappingsResource) ResourceType() string {
	return "pingfederate_sp_target_url_mappings"
}

func (r *PingFederateSpTargetUrlMappingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	spTargetUrlMappingsId := "sp_target_url_mappings_singleton_id"
	spTargetUrlMappingsName := "Sp Target Url Mappings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       spTargetUrlMappingsName,
		ResourceID:         spTargetUrlMappingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
