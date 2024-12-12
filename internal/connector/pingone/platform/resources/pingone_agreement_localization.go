package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneAgreementLocalizationResource{}
)

type PingOneAgreementLocalizationResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneAgreementLocalizationResource
func AgreementLocalization(clientInfo *connector.PingOneClientInfo) *PingOneAgreementLocalizationResource {
	return &PingOneAgreementLocalizationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneAgreementLocalizationResource) ResourceType() string {
	return "pingone_agreement_localization"
}

func (r *PingOneAgreementLocalizationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	agreementData, err := r.getAgreementData()
	if err != nil {
		return nil, err
	}

	for agreementId, agreementName := range *agreementData {
		agreementLanguageData, err := r.getAgreementLanguageData(agreementId)
		if err != nil {
			return nil, err
		}

		for agreementLanguageId, agreementLanguageLocale := range *agreementLanguageData {
			commentData := map[string]string{
				"Agreement ID":              agreementId,
				"Agreement Language ID":     agreementLanguageId,
				"Agreement Language Locale": agreementLanguageLocale,
				"Agreement Name":            agreementName,
				"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
				"Resource Type":             r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", agreementName, agreementLanguageLocale),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, agreementId, agreementLanguageId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneAgreementLocalizationResource) getAgreementData() (*map[string]string, error) {
	agreementData := make(map[string]string)
	iter := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllAgreements", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, agreement := range embedded.GetAgreements() {
			agreementId, agreementIdOk := agreement.GetIdOk()
			agreementName, agreementNameOk := agreement.GetNameOk()

			if agreementIdOk && agreementNameOk {
				agreementData[*agreementId] = *agreementName
			}
		}
	}

	return &agreementData, nil
}

func (r *PingOneAgreementLocalizationResource) getAgreementLanguageData(agreementId string) (*map[string]string, error) {
	agreementLanguageData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, agreementId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllAgreementLanguages", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, languageInner := range embedded.GetLanguages() {
			if languageInner.AgreementLanguage != nil {
				agreementLanguageId, agreementLanguageIdOk := languageInner.AgreementLanguage.GetIdOk()
				agreementLanguageLocale, agreementLanguageLocaleOk := languageInner.AgreementLanguage.GetLocaleOk()

				if agreementLanguageIdOk && agreementLanguageLocaleOk {
					agreementLanguageData[*agreementLanguageId] = *agreementLanguageLocale
				}
			}
		}
	}

	return &agreementLanguageData, nil
}
