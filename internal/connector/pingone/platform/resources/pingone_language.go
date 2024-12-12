package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneLanguageResource{}
)

type PingOneLanguageResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneLanguageResource
func Language(clientInfo *connector.PingOneClientInfo) *PingOneLanguageResource {
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

	for languageId, languageName := range *languageData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Language ID":           languageId,
			"Language Name":         languageName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       languageName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, languageId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneLanguageResource) getLanguageData() (*map[string]string, error) {
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
	}

	return &languageData, nil
}
