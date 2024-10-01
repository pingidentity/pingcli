package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneFormResource{}
)

type PingoneFormResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneFormResource
func Form(clientInfo *connector.PingOneClientInfo) *PingoneFormResource {
	return &PingoneFormResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneFormResource) ExportAll() (*[]connector.ImportBlock, error) {
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

func (r *PingoneFormResource) ResourceType() string {
	return "pingone_form"
}
