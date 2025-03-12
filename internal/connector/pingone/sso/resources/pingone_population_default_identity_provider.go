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
	_ connector.ExportableResource = &PingOnePopulationDefaultIdpResource{}
)

type PingOnePopulationDefaultIdpResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOnePopulationDefaultIdpResource
func PopulationDefaultIdp(clientInfo *connector.PingOneClientInfo) *PingOnePopulationDefaultIdpResource {
	return &PingOnePopulationDefaultIdpResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOnePopulationDefaultIdpResource) ResourceType() string {
	return "pingone_population_default_identity_provider"
}

func (r *PingOnePopulationDefaultIdpResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	populationData, err := r.getPopulationData()
	if err != nil {
		return nil, err
	}

	for populationId, populationName := range populationData {
		ok, err := r.checkPopulationDefaultIdp(populationId)
		if err != nil {
			return nil, err
		}
		if !ok {
			return &importBlocks, nil
		}

		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Population ID":         populationId,
			"Population Name":       populationName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_default_identity_provider", populationName),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, populationId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOnePopulationDefaultIdpResource) getPopulationData() (map[string]string, error) {
	populationData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
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

func (r *PingOnePopulationDefaultIdpResource) checkPopulationDefaultIdp(populationId string) (bool, error) {
	_, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.PopulationsApi.ReadOnePopulationDefaultIdp(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, populationId).Execute()
	return pingone.CheckSingletonResource(resp, err, "ReadOnePopulationDefaultIdp", r.ResourceType())
}
