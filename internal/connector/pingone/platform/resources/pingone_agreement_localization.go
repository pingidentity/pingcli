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
			commentData := map[string]string{
				"Agreement ID":                  agreementId,
				"Agreement Name":                agreementName,
				"Agreement Localization ID":     agreementLocalizationId,
				"Agreement Localization Locale": agreementLocalizationLocale,
				"Export Environment ID":         r.clientInfo.ExportEnvironmentID,
				"Resource Type":                 r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", agreementName, agreementLocalizationLocale),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, agreementId, agreementLocalizationId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func getAgreementLocalizationData(clientInfo *connector.PingOneClientInfo, resourceType, agreementId string) (map[string]string, error) {
	agreementLocalizationData := make(map[string]string)

	iter := clientInfo.ApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(clientInfo.Context, clientInfo.ExportEnvironmentID, agreementId).Execute()
	agreementLocalizations, err := common.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedLanguagesInner](iter, "ReadAllAgreementLanguages", "GetLanguages", resourceType)
	if err != nil {
		return nil, err
	}

	for _, agreementLocalization := range agreementLocalizations {
		if agreementLocalization.AgreementLanguage != nil {
			agreementLocalizationId, agreementLocalizationIdOk := agreementLocalization.AgreementLanguage.GetIdOk()
			agreementLocalizationLocale, agreementLocalizationLocaleOk := agreementLocalization.AgreementLanguage.GetLocaleOk()

			if agreementLocalizationIdOk && agreementLocalizationLocaleOk {
				agreementLocalizationData[*agreementLocalizationId] = *agreementLocalizationLocale
			}
		}
	}

	return agreementLocalizationData, nil
}
