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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOnePopulationDefaultResource
func PopulationDefault(clientInfo *connector.PingOneClientInfo) *PingOnePopulationDefaultResource {
	return &PingOnePopulationDefaultResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOnePopulationDefaultResource) ResourceType() string {
	return "pingone_population_default"
}

func (r *PingOnePopulationDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportPopulationDefault()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOnePopulationDefaultResource) exportPopulationDefault() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllPopulations", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, population := range embedded.GetPopulations() {
			populationDefault, populationDefaultOk := population.GetDefaultOk()

			if populationDefaultOk && *populationDefault {
				populationName, populationNameOk := population.GetNameOk()

				if populationNameOk {
					r.addImportBlock(*populationName)
				}
			}
		}
	}

	return nil
}

func (r *PingOnePopulationDefaultResource) addImportBlock(populationName string) {
	commentData := map[string]string{
		"Default Population Name": populationName,
		"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
		"Resource Type":           r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_population_default", populationName),
		ResourceID:         r.clientInfo.ExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
