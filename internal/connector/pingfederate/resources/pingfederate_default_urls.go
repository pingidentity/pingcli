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

type PingFederateDefaultURLsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateDefaultURLsResource
func DefaultURLs(clientInfo *connector.PingFederateClientInfo) *PingFederateDefaultURLsResource {
	return &PingFederateDefaultURLsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateDefaultURLsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	defaultURLsId := "default_urls_singleton_id"
	defaultURLsName := "Default URLs"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       defaultURLsName,
		ResourceID:         defaultURLsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateDefaultURLsResource) ResourceType() string {
	return "pingfederate_default_urls"
}
