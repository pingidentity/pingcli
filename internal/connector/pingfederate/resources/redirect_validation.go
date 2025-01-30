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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateRedirectValidationResource
func RedirectValidation(clientInfo *connector.ClientInfo) *PingFederateRedirectValidationResource {
	return &PingFederateRedirectValidationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateRedirectValidationResource) ResourceType() string {
	return "pingfederate_redirect_validation"
}

func (r *PingFederateRedirectValidationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	redirectValidationId := "redirect_validation_singleton_id"
	redirectValidationName := "Redirect Validation"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       redirectValidationName,
		ResourceID:         redirectValidationId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
