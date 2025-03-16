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
	_ connector.ExportableResource = &PingOneLanguageResource{}
)

type PingOneLanguageResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneLanguageResource
func Language(clientInfo *connector.ClientInfo) *PingOneLanguageResource {
	return &PingOneLanguageResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneLanguageResource) ResourceType() string {
	return "pingone_language"
}

func (r *PingOneLanguageResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	languageData, err := r.getLanguageData()
	if err != nil {
		return nil, err
	}

	for languageId, languageName := range languageData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Language ID":           languageId,
			"Language Name":         languageName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       languageName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, languageId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneLanguageResource) getLanguageData() (map[string]string, error) {
	languageData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	languageInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedLanguagesInner](iter, "ReadLanguages", "GetLanguages", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, languageInner := range languageInners {
		if languageInner.Language != nil {
			// If language is not customer added, skip it
			languageCustomerAdded, languageCustomerAddedOk := languageInner.Language.GetCustomerAddedOk()
			if !languageCustomerAddedOk || !*languageCustomerAdded {
				continue
			}

			languageId, languageIdOk := languageInner.Language.GetIdOk()
			languageName, languageNameOk := languageInner.Language.GetNameOk()

			if languageIdOk && languageNameOk {
				languageData[*languageId] = *languageName
			}
		}
	}

	return languageData, nil
}
