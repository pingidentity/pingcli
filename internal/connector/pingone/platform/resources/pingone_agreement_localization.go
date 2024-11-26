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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneAgreementLocalizationResource
func AgreementLocalization(clientInfo *connector.PingOneClientInfo) *PingOneAgreementLocalizationResource {
	return &PingOneAgreementLocalizationResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneAgreementLocalizationResource) ResourceType() string {
	return "pingone_agreement_localization"
}

func (r *PingOneAgreementLocalizationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportAgreementLocalizations()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneAgreementLocalizationResource) exportAgreementLocalizations() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllAgreements", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, agreement := range embedded.GetAgreements() {
			agreementId, agreementIdOk := agreement.GetIdOk()
			agreementName, agreementNameOk := agreement.GetNameOk()

			if agreementIdOk && agreementNameOk {
				err := r.exportAgreementLocalizationsByAgreement(*agreementId, *agreementName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneAgreementLocalizationResource) exportAgreementLocalizationsByAgreement(agreementId, agreementName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, agreementId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllAgreementLanguages", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, languageInner := range embedded.GetLanguages() {
			if languageInner.AgreementLanguage == nil {
				agreementLanguageId, agreementLanguageIdOk := languageInner.AgreementLanguage.GetIdOk()
				agreementLanguageLocale, agreementLanguageLocaleOk := languageInner.AgreementLanguage.GetLocaleOk()

				if agreementLanguageIdOk && agreementLanguageLocaleOk {
					r.addImportBlock(agreementId, agreementName, *agreementLanguageId, *agreementLanguageLocale)
				}
			}
		}
	}

	return nil
}

func (r *PingOneAgreementLocalizationResource) addImportBlock(agreementId, agreementName, agreementLanguageId, agreementLanguageLocale string) {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
