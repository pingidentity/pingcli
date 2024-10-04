package resources

import (
	"fmt"

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

func (r *PingFederateKerberosRealmResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.KerberosRealmsAPI.GetKerberosRealms(r.clientInfo.Context).Execute
	apiFunctionName := "GetKerberosRealms"

	kerberosRealms, response, err := apiExecuteFunc()

	err = common.HandleClientResponse(response, err, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	if kerberosRealms == nil {
		l.Error().Msgf("Returned %s() kerberosRealms is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	kerberosRealmsItems, ok := kerberosRealms.GetItemsOk()
	if !ok {
		l.Error().Msgf("Failed to get %s() kerberosRealms items.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", r.ResourceType(), apiFunctionName)
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, kerberosRealm := range kerberosRealmsItems {
		kerberosRealmId, kerberosRealmIdOk := kerberosRealm.GetIdOk()
		kerberosRealmName, kerberosRealmNameOk := kerberosRealm.GetKerberosRealmNameOk()

		if kerberosRealmIdOk && kerberosRealmNameOk {
			commentData := map[string]string{
				"Resource Type":                r.ResourceType(),
				"Kerberos Realm Resource ID":   *kerberosRealmId,
				"Kerberos Realm Resource Name": *kerberosRealmName,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *kerberosRealmName,
				ResourceID:         *kerberosRealmId,
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingFederateKerberosRealmResource) ResourceType() string {
	return "pingfederate_kerberos_realm"
}
