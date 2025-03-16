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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneEnvironmentResource
func Environment(clientInfo *connector.ClientInfo) *PingOneEnvironmentResource {
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

	ok, err := r.checkEnvironmentData()
	if err != nil {
		return nil, err
	}
	if !ok {
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.PingOneExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneEnvironmentResource) checkEnvironmentData() (bool, error) {
	_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	return common.CheckSingletonResource(response, err, "ReadOneEnvironment", r.ResourceType())
}
