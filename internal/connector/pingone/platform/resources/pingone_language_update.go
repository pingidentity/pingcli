package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneLanguageUpdateResource{}
)

type PingOneLanguageUpdateResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneLanguageUpdateResource
func LanguageUpdate(clientInfo *connector.PingOneClientInfo) *PingOneLanguageUpdateResource {
	return &PingOneLanguageUpdateResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneLanguageUpdateResource) ResourceType() string {
	return "pingone_language_update"
}

func (r *PingOneLanguageUpdateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportLanguageUpdates()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneLanguageUpdateResource) exportLanguageUpdates() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadLanguages", r.ResourceType())
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
			if languageInner.Language != nil {
				languageEnabled, languageEnabledOk := languageInner.Language.GetEnabledOk()
				languageLocale, languageLocaleOk := languageInner.Language.GetLocaleOk()
				languageDefault, languageDefaultOk := languageInner.Language.GetDefaultOk()

				if languageEnabledOk && languageLocaleOk && languageDefaultOk {
					// Export the language if it meets any of the criteria of the following 3 conditions:
					// 1) Any language enabled
					// 2) The 'en' language disabled
					// 3) If any language other than 'en' is the default

					if *languageEnabled || (*languageLocale == "en" && !*languageEnabled) || (*languageLocale != "en" && *languageDefault) {
						languageId, languageIdOk := languageInner.Language.GetIdOk()
						languageName, languageNameOk := languageInner.Language.GetNameOk()

						if languageIdOk && languageNameOk {
							r.addImportBlock(*languageId, *languageName)
						}
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneLanguageUpdateResource) addImportBlock(languageId, languageName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Language ID":           languageId,
		"Language Name":         languageName,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_update", languageName),
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, languageId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
