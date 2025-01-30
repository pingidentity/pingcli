package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateKerberosRealmSettingsResource{}
)

type PingFederateKerberosRealmSettingsResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateKerberosRealmSettingsResource
func KerberosRealmSettings(clientInfo *connector.ClientInfo) *PingFederateKerberosRealmSettingsResource {
	return &PingFederateKerberosRealmSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateKerberosRealmSettingsResource) ResourceType() string {
	return "pingfederate_kerberos_realm_settings"
}

func (r *PingFederateKerberosRealmSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	kerberosRealmSettingsId := "kerberos_realm_settings_singleton_id"
	kerberosRealmSettingsName := "Kerberos Realm Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       kerberosRealmSettingsName,
		ResourceID:         kerberosRealmSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}
