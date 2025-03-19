// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOnePopulationDefaultResource{}
)

type PingOnePopulationDefaultResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOnePopulationDefaultResource
func PopulationDefault(clientInfo *connector.ClientInfo) *PingOnePopulationDefaultResource {
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
		"Export Environment ID":   r.clientInfo.PingOneExportEnvironmentID,
		"Resource Type":           r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_population_default", *defaultPopulationName),
		ResourceID:         r.clientInfo.PingOneExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOnePopulationDefaultResource) getDefaultPopulationName() (*string, error) {
	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	populations, err := pingone.GetManagementAPIObjectsFromIterator[management.Population](iter, "ReadAllPopulations", "GetPopulations", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, population := range populations {
		populationDefault, populationDefaultOk := population.GetDefaultOk()

		if populationDefaultOk && *populationDefault {
			populationName, populationNameOk := population.GetNameOk()

			if populationNameOk {
				return populationName, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to find the name of the default population")
}
