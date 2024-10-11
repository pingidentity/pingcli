package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneFormResource{}
)

type PingOneFormResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneFormResource
func Form(clientInfo *connector.PingOneClientInfo) *PingOneFormResource {
	return &PingOneFormResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneFormResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.FormManagementApi.ReadAllForms(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllForms"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, form := range embedded.GetForms() {
		formId, formIdOk := form.GetIdOk()
		formName, formNameOk := form.GetNameOk()

		if formIdOk && formNameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Form Name":             *formName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Form ID":               *formId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *formName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *formId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneFormResource) ResourceType() string {
	return "pingone_form"
}
