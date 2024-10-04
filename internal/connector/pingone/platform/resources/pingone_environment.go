package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneCustomDomainResource{}
)

type PingoneEnvironmentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneEnvironmentResource
func Environment(clientInfo *connector.PingOneClientInfo) *PingoneEnvironmentResource {
	return &PingoneEnvironmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneEnvironmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       "export_environment",
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingoneEnvironmentResource) ResourceType() string {
	return "pingone_environment"
}
