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

type PingFederateIDPDefaultURLsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateIDPDefaultURLsResource
func IDPDefaultURLs(clientInfo *connector.PingFederateClientInfo) *PingFederateIDPDefaultURLsResource {
	return &PingFederateIDPDefaultURLsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIDPDefaultURLsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	idpDefaultURLsId := "idp_default_urls_singleton_id"
	idpDefaultURLsName := "IDP Default URLs"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       idpDefaultURLsName,
		ResourceID:         idpDefaultURLsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateIDPDefaultURLsResource) ResourceType() string {
	return "pingfederate_idp_default_urls"
}
