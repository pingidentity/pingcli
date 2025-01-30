package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateConfigurationEncryptionKeysRotateResource{}
)

type PingFederateConfigurationEncryptionKeysRotateResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateConfigurationEncryptionKeysRotateResource
func ConfigurationEncryptionKeysRotate(clientInfo *connector.ClientInfo) *PingFederateConfigurationEncryptionKeysRotateResource {
	return &PingFederateConfigurationEncryptionKeysRotateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateConfigurationEncryptionKeysRotateResource) ResourceType() string {
	return "pingfederate_configuration_encryption_keys_rotate"
}

func (r *PingFederateConfigurationEncryptionKeysRotateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	configurationEncryptionKeysRotateId := "configuration_encryption_keys_rotate_singleton_id"
	configurationEncryptionKeysRotateName := "Configuration Encryption Keys Rotate"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       configurationEncryptionKeysRotateName,
		ResourceID:         configurationEncryptionKeysRotateId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
