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
	_ connector.ExportableResource = &PingOnePopulationDefaultIdp{}
)

type PingOnePopulationDefaultIdp struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOnePopulationDefaultIdp
func PopulationDefaultIdp(clientInfo *connector.PingOneClientInfo) *PingOnePopulationDefaultIdp {
	return &PingOnePopulationDefaultIdp{
		clientInfo: clientInfo,
	}
}

func (r *PingOnePopulationDefaultIdp) ResourceType() string {
	return "pingone_population_default_identity_provider"
}

func (r *PingOnePopulationDefaultIdp) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	populationData, err := r.getPopulationData()
	if err != nil {
		return nil, err
	}

	for populationId, populationName := range populationData {
		populationDefaultIdp, err := r.getPopulationDefaultIdp(populationId)
		if err != nil {
			return nil, err
		}
		if populationDefaultIdp == nil {
			continue
		}

		commentData := map[string]string{
			"Export Environment ID":                r.clientInfo.ExportEnvironmentID,
			"Population Default Identity Provider": *populationDefaultIdp,
			"Population ID":                        populationId,
			"Population Name":                      populationName,
			"Resource Type":                        r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       populationName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, populationId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOnePopulationDefaultIdp) getPopulationData() (map[string]string, error) {
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

func (r *PingOnePopulationDefaultIdp) getPopulationDefaultIdp(populationId string) (*string, error) {
	populationDefaultIdp, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.PopulationsApi.ReadOnePopulationDefaultIdp(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, populationId).Execute()
	ok, err := common.HandleClientResponse(resp, err, "ReadOnePopulationDefaultIdp", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if populationDefaultIdp != nil {
		populationDefaultIdpId, populationDefaultIdpIdOk := populationDefaultIdp.GetIdOk()
		if populationDefaultIdpIdOk {
			return populationDefaultIdpId, nil
		}
	}

	return nil, nil
}
