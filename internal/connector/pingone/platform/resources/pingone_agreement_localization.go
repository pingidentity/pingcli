// Copyright © 2025 Ping Identity Corporation

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
	_ connector.ExportableResource = &PingOneAgreementLocalizationResource{}
)

type PingOneAgreementLocalizationResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneAgreementLocalizationResource
func AgreementLocalization(clientInfo *connector.ClientInfo) *PingOneAgreementLocalizationResource {
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

	for agreementId, agreementName := range agreementData {
		agreementLocalizationData, err := r.getAgreementLocalizationData(agreementId)
		if err != nil {
			return nil, err
		}

		for agreementLocalizationId, agreementLocalizationLocale := range agreementLocalizationData {
			commentData := map[string]string{
				"Agreement ID":                  agreementId,
				"Agreement Name":                agreementName,
				"Agreement Localization ID":     agreementLocalizationId,
				"Agreement Localization Locale": agreementLocalizationLocale,
				"Export Environment ID":         r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                 r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", agreementName, agreementLocalizationLocale),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, agreementId, agreementLocalizationId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneAgreementLocalizationResource) getAgreementData() (map[string]string, error) {
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

func (r *PingOneAgreementLocalizationResource) getAgreementLocalizationData(agreementId string) (map[string]string, error) {
	agreementLocalizationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, agreementId).Execute()
	languageInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedLanguagesInner](iter, "ReadAllAgreementLanguages", "GetLanguages", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, languageInner := range languageInners {
		if languageInner.AgreementLanguage != nil {
			agreementLocalizationId, agreementLocalizationIdOk := languageInner.AgreementLanguage.GetIdOk()
			agreementLocalizationLocale, agreementLocalizationLocaleOk := languageInner.AgreementLanguage.GetLocaleOk()

			if agreementLocalizationIdOk && agreementLocalizationLocaleOk {
				agreementLocalizationData[*agreementLocalizationId] = *agreementLocalizationLocale
			}
		}
	}

	return agreementLocalizationData, nil
}
