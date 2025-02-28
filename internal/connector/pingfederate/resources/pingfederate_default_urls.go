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

func (r *PingFederateDefaultURLsResource) ResourceType() string {
	return "pingfederate_default_urls"
}

func (r *PingFederateDefaultURLsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	defaultURLsId := "default_urls_singleton_id"
	defaultURLsName := "Default URLs"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       defaultURLsName,
		ResourceID:         defaultURLsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
