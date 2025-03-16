package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneAgreementLocalizationRevisionResource{}
)

type PingOneAgreementLocalizationRevisionResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneAgreementLocalizationRevisionResource
func AgreementLocalizationRevision(clientInfo *connector.ClientInfo) *PingOneAgreementLocalizationRevisionResource {
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

	agreementData, err := r.getAgreementData()
	if err != nil {
		return nil, err
	}

	for agreementId, agreementName := range agreementData {
		agreementLocalizationData, err := r.getAgreementLanguageData(agreementId)
		if err != nil {
			return nil, err
		}

		for agreementLocalizationId, agreementLocalizationLocale := range agreementLocalizationData {
			agreementLocalizationRevisionData, err := r.getAgreementLocalizationRevisionData(agreementId, agreementLocalizationId)
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
					"Export Environment ID":              r.clientInfo.PingOneExportEnvironmentID,
					"Resource Type":                      r.ResourceType(),
				}

				importBlock := connector.ImportBlock{
					ResourceType:       r.ResourceType(),
					ResourceName:       fmt.Sprintf("%s_%s_%s", agreementName, agreementLocalizationLocale, agreementLocalizationRevisionId),
					ResourceID:         fmt.Sprintf("%s/%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, agreementId, agreementLocalizationId, agreementLocalizationRevisionId),
					CommentInformation: common.GenerateCommentInformation(commentData),
				}

				importBlocks = append(importBlocks, importBlock)
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneAgreementLocalizationRevisionResource) getAgreementData() (map[string]string, error) {
	agreementData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	agreements, err := pingone.GetManagementAPIObjectsFromIterator[management.Agreement](iter, "ReadAllAgreements", "GetAgreements", r.ResourceType())
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

func (r *PingOneAgreementLocalizationRevisionResource) getAgreementLanguageData(agreementId string) (map[string]string, error) {
	agreementLanguageData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, agreementId).Execute()
	languageInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedLanguagesInner](iter, "ReadAllAgreementLanguages", "GetLanguages", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, languageInner := range languageInners {
		if languageInner.AgreementLanguage != nil {
			agreementLanguageLocale, agreementLanguageLocaleOk := languageInner.AgreementLanguage.GetLocaleOk()
			agreementLanguageId, agreementLanguageIdOk := languageInner.AgreementLanguage.GetIdOk()

			if agreementLanguageLocaleOk && agreementLanguageIdOk {
				agreementLanguageData[*agreementLanguageId] = *agreementLanguageLocale
			}
		}
	}

	return agreementLanguageData, nil
}

func (r *PingOneAgreementLocalizationRevisionResource) getAgreementLocalizationRevisionData(agreementId, agreementLocalizationId string) ([]string, error) {
	agreementLocalizationRevisionData := []string{}

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, agreementId, agreementLocalizationId).Execute()
	agreementLocalizationRevisions, err := pingone.GetManagementAPIObjectsFromIterator[management.AgreementLanguageRevision](iter, "ReadAllAgreementLanguageRevisions", "GetRevisions", r.ResourceType())
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
