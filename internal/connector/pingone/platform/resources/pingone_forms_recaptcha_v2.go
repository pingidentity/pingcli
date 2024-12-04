package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneFormRecaptchaV2Resource{}
)

type PingOneFormRecaptchaV2Resource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneFormRecaptchaV2Resource
func FormRecaptchaV2(clientInfo *connector.PingOneClientInfo) *PingOneFormRecaptchaV2Resource {
	return &PingOneFormRecaptchaV2Resource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneFormRecaptchaV2Resource) ResourceType() string {
	return "pingone_forms_recaptcha_v2"
}

func (r *PingOneFormRecaptchaV2Resource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportFormRecaptchaV2()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneFormRecaptchaV2Resource) exportFormRecaptchaV2() error {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.RecaptchaConfigurationApi.ReadRecaptchaConfiguration(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	err = common.HandleClientResponse(response, err, "ReadRecaptchaConfiguration", r.ResourceType())
	if err != nil {
		return err
	}

	if response.StatusCode == 204 {
		return common.DataNilError(r.ResourceType(), response)
	}

	r.addImportBlock()

	return nil
}

func (r *PingOneFormRecaptchaV2Resource) addImportBlock() {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
