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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneLanguageUpdateResource
func LanguageUpdate(clientInfo *connector.PingOneClientInfo) *PingOneLanguageUpdateResource {
	return &PingOneLanguageUpdateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneLanguageUpdateResource) ResourceType() string {
	return "pingone_language_update"
}

func (r *PingOneLanguageUpdateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	languageData, err := r.getLanguageData()
	if err != nil {
		return nil, err
	}

	for languageId, languageName := range *languageData {
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

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneLanguageUpdateResource) getLanguageData() (*map[string]string, error) {
	languageData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadLanguages", r.ResourceType())
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
							languageData[*languageId] = *languageName
						}
					}
				}
			}
		}
	}

	return &languageData, nil
}
