// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateDefaultUrlsResource{}
)

type PingFederateDefaultUrlsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateDefaultUrlsResource
func DefaultUrls(clientInfo *connector.ClientInfo) *PingFederateDefaultUrlsResource {
	return &PingFederateDefaultUrlsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateDefaultUrlsResource) ResourceType() string {
	return "pingfederate_default_urls"
}

func (r *PingFederateDefaultUrlsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	defaultUrlsId := "default_urls_singleton_id"
	defaultUrlsName := "Default Urls"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       defaultUrlsName,
		ResourceID:         defaultUrlsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
