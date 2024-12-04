package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneAgreementLocalizationRevisionResource{}
)

type PingOneAgreementLocalizationRevisionResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneAgreementLocalizationRevisionResource
func AgreementLocalizationRevision(clientInfo *connector.PingOneClientInfo) *PingOneAgreementLocalizationRevisionResource {
	return &PingOneAgreementLocalizationRevisionResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneAgreementLocalizationRevisionResource) ResourceType() string {
	return "pingone_agreement_localization_revision"
}

func (r *PingOneAgreementLocalizationRevisionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportAgreementLocalizationRevisions()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneAgreementLocalizationRevisionResource) exportAgreementLocalizationRevisions() error {
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
				err := r.exportAgreementLocalizationRevisionsByAgreement(*agreementId, *agreementName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneAgreementLocalizationRevisionResource) exportAgreementLocalizationRevisionsByAgreement(agreementId, agreementName string) error {
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
			if languageInner.AgreementLanguage != nil {
				agreementLanguageLocale, agreementLanguageLocaleOk := languageInner.AgreementLanguage.GetLocaleOk()
				agreementLanguageId, agreementLanguageIdOk := languageInner.AgreementLanguage.GetIdOk()

				if agreementLanguageLocaleOk && agreementLanguageIdOk {
					err := r.exportAgreementLocalizationRevisionsByAgreementLanguage(agreementId, agreementName, *agreementLanguageId, *agreementLanguageLocale)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneAgreementLocalizationRevisionResource) exportAgreementLocalizationRevisionsByAgreementLanguage(agreementId, agreementName, agreementLanguageId, agreementLanguageLocale string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, agreementId, agreementLanguageId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllAgreementLanguageRevisions", r.ResourceType())
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

		for _, revision := range embedded.GetRevisions() {
			agreementLanguageRevisionId, agreementLanguageRevisionIdOk := revision.GetIdOk()

			if agreementLanguageRevisionIdOk {
				r.addImportBlock(agreementId, agreementName, agreementLanguageId, agreementLanguageLocale, *agreementLanguageRevisionId)
			}
		}
	}

	return nil
}

func (r *PingOneAgreementLocalizationRevisionResource) addImportBlock(agreementId, agreementName, agreementLanguageId, agreementLanguageLocale, agreementLanguageRevisionId string) {
	commentData := map[string]string{
		"Agreement ID":                       agreementId,
		"Agreement Language ID":              agreementLanguageId,
		"Agreement Language Locale":          agreementLanguageLocale,
		"Agreement Localization Revision ID": agreementLanguageRevisionId,
		"Agreement Name":                     agreementName,
		"Export Environment ID":              r.clientInfo.ExportEnvironmentID,
		"Resource Type":                      r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s_%s", agreementName, agreementLanguageLocale, agreementLanguageRevisionId),
		ResourceID:         fmt.Sprintf("%s/%s/%s/%s", r.clientInfo.ExportEnvironmentID, agreementId, agreementLanguageId, agreementLanguageRevisionId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
