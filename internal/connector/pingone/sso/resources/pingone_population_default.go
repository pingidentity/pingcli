package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingonePopulationDefaultDefaultResource{}
)

type PingonePopulationDefaultDefaultResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingonePopulationDefaultDefaultResource
func PopulationDefault(clientInfo *connector.PingOneClientInfo) *PingonePopulationDefaultDefaultResource {
	return &PingonePopulationDefaultDefaultResource{
		clientInfo: clientInfo,
	}
}

func (r *PingonePopulationDefaultDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllPopulations"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	foundDefault := false
	var defaultPopulation management.Population
	for _, population := range embedded.GetPopulations() {
		if population.GetDefault() {
			foundDefault = true
			defaultPopulation = population
			break
		}
	}

	if !foundDefault {
		l.Debug().Msgf("No exportable %s resource found", r.ResourceType())
		return &[]connector.ImportBlock{}, nil
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	defaultPopulationName, defaultPopulationNameOk := defaultPopulation.GetNameOk()

	if defaultPopulationNameOk {
		commentData := map[string]string{
			"Resource Type":           r.ResourceType(),
			"Default Population Name": *defaultPopulationName,
			"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
		}

		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_population_default", *defaultPopulationName),
			ResourceID:         r.clientInfo.ExportEnvironmentID,
			CommentInformation: common.GenerateCommentInformation(commentData),
		})
	}

	return &importBlocks, nil
}

func (r *PingonePopulationDefaultDefaultResource) ResourceType() string {
	return "pingone_population_default"
}
