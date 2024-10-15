package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateRedirectValidationResource{}
)

type PingFederateRedirectValidationResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateRedirectValidationResource
func RedirectValidation(clientInfo *connector.PingFederateClientInfo) *PingFederateRedirectValidationResource {
	return &PingFederateRedirectValidationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateRedirectValidationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	redirectValidationId := "redirect_validation_singleton_id"
	redirectValidationName := "Redirect Validation"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       redirectValidationName,
		ResourceID:         redirectValidationId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateRedirectValidationResource) ResourceType() string {
	return "pingfederate_redirect_validation"
}
