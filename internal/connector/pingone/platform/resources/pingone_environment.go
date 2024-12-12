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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneEnvironmentResource
func Environment(clientInfo *connector.PingOneClientInfo) *PingOneEnvironmentResource {
	return &PingOneEnvironmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneEnvironmentResource) ResourceType() string {
	return "pingone_environment"
}

func (r *PingOneEnvironmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	err := r.checkEnvironmentData()
	if err != nil {
		return nil, err
	}

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

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneEnvironmentResource) checkEnvironmentData() error {
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	err = common.HandleClientResponse(response, err, "ReadOneEnvironment", r.ResourceType())
	if err != nil {
		return err
	}

	return nil
}
