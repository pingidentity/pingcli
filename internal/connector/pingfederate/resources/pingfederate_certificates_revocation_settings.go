package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateCertificatesRevocationSettingsResource{}
)

type PingFederateCertificatesRevocationSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateCertificatesRevocationSettingsResource
func CertificatesRevocationSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateCertificatesRevocationSettingsResource {
	return &PingFederateCertificatesRevocationSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateCertificatesRevocationSettingsResource) ResourceType() string {
	return "pingfederate_certificates_revocation_settings"
}

func (r *PingFederateCertificatesRevocationSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	certificatesRevocationSettingsId := "certificates_revocation_settings_singleton_id"
	certificatesRevocationSettingsName := "Certificates Revocation Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       certificatesRevocationSettingsName,
		ResourceID:         certificatesRevocationSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
