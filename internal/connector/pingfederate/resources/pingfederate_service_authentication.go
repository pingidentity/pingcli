package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateServiceAuthenticationResource{}
)

type PingFederateServiceAuthenticationResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateServiceAuthenticationResource
func ServiceAuthentication(clientInfo *connector.PingFederateClientInfo) *PingFederateServiceAuthenticationResource {
	return &PingFederateServiceAuthenticationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateServiceAuthenticationResource) ResourceType() string {
	return "pingfederate_service_authentication"
}

func (r *PingFederateServiceAuthenticationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	serviceAuthenticationId := "service_authentication_singleton_id"
	serviceAuthenticationName := "Service Authentication"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       serviceAuthenticationName,
		ResourceID:         serviceAuthenticationId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}