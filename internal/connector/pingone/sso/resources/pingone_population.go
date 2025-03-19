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
	_ connector.ExportableResource = &PingOnePopulationResource{}
)

type PingOnePopulationResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOnePopulationResource
func Population(clientInfo *connector.ClientInfo) *PingOnePopulationResource {
	return &PingOnePopulationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOnePopulationResource) ResourceType() string {
	return "pingone_population"
}

func (r *PingOnePopulationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	populationData, err := r.getPopulationData()
	if err != nil {
		return nil, err
	}

	for populationId, populationName := range populationData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Population ID":         populationId,
			"Population Name":       populationName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       populationName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, populationId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOnePopulationResource) getPopulationData() (map[string]string, error) {
	populationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	populations, err := pingone.GetManagementAPIObjectsFromIterator[management.Population](iter, "ReadAllPopulations", "GetPopulations", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, population := range populations {
		populationId, populationIdOk := population.GetIdOk()
		populationName, populationNameOk := population.GetNameOk()

		if populationIdOk && populationNameOk {
			populationData[*populationId] = *populationName
		}
	}

	return populationData, nil
}
