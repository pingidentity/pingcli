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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneFormResource
func Form(clientInfo *connector.PingOneClientInfo) *PingOneFormResource {
	return &PingOneFormResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneFormResource) ResourceType() string {
	return "pingone_form"
}

func (r *PingOneFormResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportForms()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneFormResource) exportForms() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.FormManagementApi.ReadAllForms(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllForms", r.ResourceType())
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

		for _, form := range embedded.GetForms() {
			formId, formIdOk := form.GetIdOk()
			formName, formNameOk := form.GetNameOk()

			if formIdOk && formNameOk {
				r.addImportBlock(*formId, *formName)
			}
		}
	}

	return nil
}

func (r *PingOneFormResource) addImportBlock(formId, formName string) {
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

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
