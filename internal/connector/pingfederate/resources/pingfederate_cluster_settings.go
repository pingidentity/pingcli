package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateClusterSettingsResource{}
)

type PingFederateClusterSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateClusterSettingsResource
func ClusterSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateClusterSettingsResource {
	return &PingFederateClusterSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateClusterSettingsResource) ResourceType() string {
	return "pingfederate_cluster_settings"
}

func (r *PingFederateClusterSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	clusterSettingsId := "cluster_settings_singleton_id"
	clusterSettingsName := "Cluster Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       clusterSettingsName,
		ResourceID:         clusterSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
