package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementLocalizationResource{}
)

type PingoneAgreementLocalizationResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAgreementLocalizationResource
func AgreementLocalization(clientInfo *connector.PingOneClientInfo) *PingoneAgreementLocalizationResource {
	return &PingoneAgreementLocalizationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAgreementLocalizationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllAgreements"

	agreementEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, agreement := range agreementEmbedded.GetAgreements() {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()

		if agreementIdOk && agreementNameOk {
			apiExecuteFunc = r.clientInfo.ApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *agreement.Id).Execute
			apiFunctionName = "ReadAllAgreementLanguages"

			agreementLanguageEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, languageWrapper := range agreementLanguageEmbedded.GetLanguages() {
				if languageWrapper.AgreementLanguage != nil {
					agreementLanguage := languageWrapper.AgreementLanguage

					agreementLanguageLocale, agreementLanguageLocaleOk := agreementLanguage.GetLocaleOk()
					agreementLanguageId, agreementLanguageIdOk := agreementLanguage.GetIdOk()

					if agreementLanguageLocaleOk && agreementLanguageIdOk {
						commentData := map[string]string{
							"Resource Type":             r.ResourceType(),
							"Agreement Name":            *agreementName,
							"Agreement Language Locale": *agreementLanguageLocale,
							"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
							"Agreement ID":              *agreementId,
							"Agreement Language ID":     *agreementLanguageId,
						}

						importBlocks = append(importBlocks, connector.ImportBlock{
							ResourceType:       r.ResourceType(),
							ResourceName:       fmt.Sprintf("%s_%s", *agreementName, *agreementLanguageLocale),
							ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *agreementId, *agreementLanguageId),
							CommentInformation: common.GenerateCommentInformation(commentData),
						})
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementLocalizationResource) ResourceType() string {
	return "pingone_agreement_localization"
}
