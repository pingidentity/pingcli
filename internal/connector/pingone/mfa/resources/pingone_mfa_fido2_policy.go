package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneMFAFido2PolicyResource{}
)

type PingoneMFAFido2PolicyResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneMFAFido2PolicyResource
func MFAFido2Policy(clientInfo *connector.PingOneClientInfo) *PingoneMFAFido2PolicyResource {
	return &PingoneMFAFido2PolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneMFAFido2PolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.MFAAPIClient.FIDO2PolicyApi.ReadFIDO2Policies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadFIDO2Policies"

	embedded, err := common.GetMFAEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, fido2Policy := range embedded.GetFido2Policies() {
		fido2PolicyName, fido2PolicyNameOk := fido2Policy.GetNameOk()
		fido2PolicyId, fido2PolicyIdOk := fido2Policy.GetIdOk()

		if fido2PolicyNameOk && fido2PolicyIdOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"FIDO2 Policy Name":     *fido2PolicyName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"FIDO2 Policy ID":       *fido2PolicyId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *fido2PolicyName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *fido2PolicyId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneMFAFido2PolicyResource) ResourceType() string {
	return "pingone_mfa_fido2_policy"
}
