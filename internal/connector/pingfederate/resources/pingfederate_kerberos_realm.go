package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateKerberosRealmResource{}
)

type PingFederateKerberosRealmResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateKerberosRealmResource
func KerberosRealm(clientInfo *connector.PingFederateClientInfo) *PingFederateKerberosRealmResource {
	return &PingFederateKerberosRealmResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateKerberosRealmResource) ResourceType() string {
	return "pingfederate_kerberos_realm"
}

func (r *PingFederateKerberosRealmResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	kerberosRealmData, err := r.getKerberosRealmData()
	if err != nil {
		return nil, err
	}

	for kerberosRealmId, kerberosRealmName := range *kerberosRealmData {
		commentData := map[string]string{
			"Kerberos Realm ID":   kerberosRealmId,
			"Kerberos Realm Name": kerberosRealmName,
			"Resource Type":       r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       kerberosRealmName,
			ResourceID:         kerberosRealmId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateKerberosRealmResource) getKerberosRealmData() (*map[string]string, error) {
	kerberosRealmData := make(map[string]string)

	kerberosRealms, response, err := r.clientInfo.ApiClient.KerberosRealmsAPI.GetKerberosRealms(r.clientInfo.Context).Execute()
	err = common.HandleClientResponse(response, err, "GetKerberosRealms", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if kerberosRealms == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	kerberosRealmsItems, kerberosRealmsItemsOk := kerberosRealms.GetItemsOk()
	if !kerberosRealmsItemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, kerberosRealm := range kerberosRealmsItems {
		kerberosRealmId, kerberosRealmIdOk := kerberosRealm.GetIdOk()
		kerberosRealmName, kerberosRealmNameOk := kerberosRealm.GetKerberosRealmNameOk()

		if kerberosRealmIdOk && kerberosRealmNameOk {
			kerberosRealmData[*kerberosRealmId] = *kerberosRealmName
		}
	}

	return &kerberosRealmData, nil
}
