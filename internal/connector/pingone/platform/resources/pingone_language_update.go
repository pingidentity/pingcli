package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneLanguageUpdateResource{}
)

type PingoneLanguageUpdateResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneLanguageUpdateResource
func LanguageUpdate(clientInfo *connector.PingOneClientInfo) *PingoneLanguageUpdateResource {
	return &PingoneLanguageUpdateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneLanguageUpdateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadLanguages"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, languageInner := range embedded.GetLanguages() {
		if languageInner.Language != nil {
			language := languageInner.Language

			languageCreatedAt, languageCreatedAtOk := language.GetCreatedAtOk()
			languageUpdatedAt, languageUpdatedAtOk := language.GetUpdatedAtOk()

			// if language update time is equal to creation time, skip it as it has not been updated
			if languageCreatedAtOk && languageUpdatedAtOk && (*languageCreatedAt).Equal(*languageUpdatedAt) {
				continue
			}

			languageId, languageIdOk := language.GetIdOk()
			languageName, languageNameOk := language.GetNameOk()

			if languageIdOk && languageNameOk && languageCreatedAtOk && languageUpdatedAtOk {
				commentData := map[string]string{
					"Resource Type":         r.ResourceType(),
					"Language Name":         *languageName,
					"Export Environment ID": r.clientInfo.ExportEnvironmentID,
					"Language ID":           *languageId,
				}

				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType:       r.ResourceType(),
					ResourceName:       fmt.Sprintf("%s_update", *languageName),
					ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *languageId),
					CommentInformation: common.GenerateCommentInformation(commentData),
				})
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneLanguageUpdateResource) ResourceType() string {
	return "pingone_language_update"
}
