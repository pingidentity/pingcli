package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneLanguageResource{}
)

type PingoneLanguageResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneLanguageResource
func Language(clientInfo *connector.PingOneClientInfo) *PingoneLanguageResource {
	return &PingoneLanguageResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneLanguageResource) ExportAll() (*[]connector.ImportBlock, error) {
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

			// If language is not customer added, skip it
			languageCustomerAdded, languageCustomerAddedOk := language.GetCustomerAddedOk()
			if languageCustomerAddedOk && !*languageCustomerAdded {
				continue
			}

			languageId, languageIdOk := language.GetIdOk()
			languageName, languageNameOk := language.GetNameOk()

			if languageIdOk && languageNameOk && languageCustomerAddedOk {
				commentData := map[string]string{
					"Resource Type":         r.ResourceType(),
					"Language Name":         *languageName,
					"Export Environment ID": r.clientInfo.ExportEnvironmentID,
					"Language ID":           *languageId,
				}

				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType:       r.ResourceType(),
					ResourceName:       *languageName,
					ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *languageId),
					CommentInformation: common.GenerateCommentInformation(commentData),
				})
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneLanguageResource) ResourceType() string {
	return "pingone_language"
}
