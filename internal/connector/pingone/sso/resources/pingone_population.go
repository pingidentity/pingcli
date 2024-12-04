package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOnePopulationResource{}
)

type PingOnePopulationResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOnePopulationResource
func Population(clientInfo *connector.PingOneClientInfo) *PingOnePopulationResource {
	return &PingOnePopulationResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOnePopulationResource) ResourceType() string {
	return "pingone_population"
}

func (r *PingOnePopulationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportPopulations()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOnePopulationResource) exportPopulations() error {
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
			populationId, populationIdOk := population.GetIdOk()
			populationName, populationNameOk := population.GetNameOk()

			if populationIdOk && populationNameOk {
				r.addImportBlock(*populationId, *populationName)
			}
		}
	}

	return nil
}

func (r *PingOnePopulationResource) addImportBlock(populationId, populationName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Population ID":         populationId,
		"Population Name":       populationName,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       populationName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, populationId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
