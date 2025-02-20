package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
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

func (r *PingOneAgreementResource) ResourceType() string {
	return "pingone_agreement"
}

func (r *PingOneAgreementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	agreementData, err := getAgreementData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for agreementId, agreementName := range agreementData {
		commentData := map[string]string{
			"Agreement ID":          agreementId,
			"Agreement Name":        agreementName,
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       agreementName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, agreementId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func getAgreementData(clientInfo *connector.PingOneClientInfo, resourceType string) (map[string]string, error) {
	agreementData := make(map[string]string)

	iter := clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()
	agreements, err := common.GetManagementAPIObjectsFromIterator[management.Agreement](iter, "ReadAllAgreements", "GetAgreements", resourceType)
	if err != nil {
		return nil, err
	}

	for _, agreement := range agreements {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()

		if agreementIdOk && agreementNameOk {
			agreementData[*agreementId] = *agreementName
		}
	}

	return agreementData, nil
}
