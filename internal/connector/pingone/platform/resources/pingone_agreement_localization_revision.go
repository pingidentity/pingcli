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
	_ connector.ExportableResource = &PingOneAgreementLocalizationRevisionResource{}
)

type PingOneAgreementLocalizationRevisionResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneAgreementLocalizationRevisionResource
func AgreementLocalizationRevision(clientInfo *connector.PingOneClientInfo) *PingOneAgreementLocalizationRevisionResource {
	return &PingOneAgreementLocalizationRevisionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneAgreementLocalizationRevisionResource) ResourceType() string {
	return "pingone_agreement_localization_revision"
}

func (r *PingOneAgreementLocalizationRevisionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	agreementData, err := getAgreementData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for agreementId, agreementName := range agreementData {
		agreementLocalizationData, err := getAgreementLocalizationData(r.clientInfo, r.ResourceType(), agreementId)
		if err != nil {
			return nil, err
		}

		for agreementLocalizationId, agreementLocalizationLocale := range agreementLocalizationData {
			agreementLocalizationRevisionData, err := getAgreementLocalizationRevisionData(r.clientInfo, r.ResourceType(), agreementId, agreementLocalizationId)
			if err != nil {
				return nil, err
			}

			for _, agreementLocalizationRevisionId := range agreementLocalizationRevisionData {
				commentData := map[string]string{
					"Agreement ID":                       agreementId,
					"Agreement Name":                     agreementName,
					"Agreement Localization ID":          agreementLocalizationId,
					"Agreement Localization Locale":      agreementLocalizationLocale,
					"Agreement Localization Revision ID": agreementLocalizationRevisionId,
					"Export Environment ID":              r.clientInfo.ExportEnvironmentID,
					"Resource Type":                      r.ResourceType(),
				}

				importBlock := connector.ImportBlock{
					ResourceType:       r.ResourceType(),
					ResourceName:       fmt.Sprintf("%s_%s_%s", agreementName, agreementLocalizationLocale, agreementLocalizationRevisionId),
					ResourceID:         fmt.Sprintf("%s/%s/%s/%s", r.clientInfo.ExportEnvironmentID, agreementId, agreementLocalizationId, agreementLocalizationRevisionId),
					CommentInformation: common.GenerateCommentInformation(commentData),
				}

				importBlocks = append(importBlocks, importBlock)
			}
		}
	}

	return &importBlocks, nil
}

func getAgreementLocalizationRevisionData(clientInfo *connector.PingOneClientInfo, resourceType, agreementId, agreementLocalizationId string) ([]string, error) {
	agreementLocalizationRevisionData := []string{}

	iter := clientInfo.ApiClient.ManagementAPIClient.AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions(clientInfo.Context, clientInfo.ExportEnvironmentID, agreementId, agreementLocalizationId).Execute()
	agreementLocalizationRevisions, err := common.GetManagementAPIObjectsFromIterator[management.AgreementLanguageRevision](iter, "ReadAllAgreementLanguageRevisions", "GetRevisions", resourceType)
	if err != nil {
		return nil, err
	}

	for _, agreementLocalizationRevision := range agreementLocalizationRevisions {
		agreementLocalizationRevisionId, agreementLocalizationRevisionIdOk := agreementLocalizationRevision.GetIdOk()

		if agreementLocalizationRevisionIdOk {
			agreementLocalizationRevisionData = append(agreementLocalizationRevisionData, *agreementLocalizationRevisionId)
		}
	}

	return agreementLocalizationRevisionData, nil
}
