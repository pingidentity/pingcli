package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOnePopulationDefaultResource{}
)

type PingOnePopulationDefaultResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOnePopulationDefaultResource
func PopulationDefault(clientInfo *connector.PingOneClientInfo) *PingOnePopulationDefaultResource {
	return &PingOnePopulationDefaultResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOnePopulationDefaultResource) ResourceType() string {
	return "pingone_population_default"
}

func (r *PingOnePopulationDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	defaultPopulationName, err := r.getDefaultPopulationName()
	if err != nil {
		return nil, err
	}

	commentData := map[string]string{
		"Default Population Name": *defaultPopulationName,
		"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
		"Resource Type":           r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_population_default", *defaultPopulationName),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOnePopulationDefaultResource) getDefaultPopulationName() (*string, error) {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllPopulations", r.ResourceType())
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

		for _, population := range embedded.GetPopulations() {
			populationDefault, populationDefaultOk := population.GetDefaultOk()

			if populationDefaultOk && *populationDefault {
				populationName, populationNameOk := population.GetNameOk()

				if populationNameOk {
					return populationName, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("Unable to find the name of the default population")
}
