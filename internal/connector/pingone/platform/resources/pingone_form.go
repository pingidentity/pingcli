package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
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

func (r *PingOneFormResource) ResourceType() string {
	return "pingone_form"
}

func (r *PingOneFormResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	formData, err := r.getFormData()
	if err != nil {
		return nil, err
	}

	for formId, formName := range formData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Form ID":               formId,
			"Form Name":             formName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       formName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, formId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneFormResource) getFormData() (map[string]string, error) {
	formData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.FormManagementApi.ReadAllForms(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	forms, err := common.GetManagementAPIObjectsFromIterator[management.Form](iter, "ReadAllForms", "GetForms", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, form := range forms {
		formId, formIdOk := form.GetIdOk()
		formName, formNameOk := form.GetNameOk()

		if formIdOk && formNameOk {
			formData[*formId] = *formName
		}
	}

	return formData, nil
}
