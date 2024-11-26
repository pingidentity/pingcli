package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneCustomDomainResource{}
)

type PingOneEnvironmentResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneEnvironmentResource
func Environment(clientInfo *connector.PingOneClientInfo) *PingOneEnvironmentResource {
	return &PingOneEnvironmentResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneEnvironmentResource) ResourceType() string {
	return "pingone_environment"
}

func (r *PingOneEnvironmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportEnvironments()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneEnvironmentResource) exportEnvironments() error {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	err = common.HandleClientResponse(response, err, "ReadOneEnvironment", r.ResourceType())
	if err != nil {
		return err
	}

	r.addImportBlock()

	return nil
}

func (r *PingOneEnvironmentResource) addImportBlock() {
	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
