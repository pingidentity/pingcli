package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneAgreementResource{}
)

type PingOneAgreementResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneAgreementResource
func Agreement(clientInfo *connector.PingOneClientInfo) *PingOneAgreementResource {
	return &PingOneAgreementResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneAgreementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllAgreements"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())
	for _, agreement := range embedded.GetAgreements() {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()

		commentData := map[string]string{
			"Resource Type":           r.ResourceType(),
			"Agreement Resource Name": *agreementName,
			"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
			"Agreement Resource ID":   *agreementId,
		}

		if agreementIdOk && agreementNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *agreementName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *agreementId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneAgreementResource) ResourceType() string {
	return "pingone_agreement"
}
