// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	_ connector.ExportableResource = &PingOnePopulationDefaultIdentityProviderResource{}
)

type PingOnePopulationDefaultIdentityProviderResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOnePopulationDefaultIdentityProviderResource
func PopulationDefaultIdentityProvider(clientInfo *connector.ClientInfo) *PingOnePopulationDefaultIdentityProviderResource {
	return &PingOnePopulationDefaultIdentityProviderResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOnePopulationDefaultIdentityProviderResource) ResourceType() string {
	return "pingone_population_default_identity_provider"
}

func (r *PingOnePopulationDefaultIdentityProviderResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	populationData, err := r.getPopulationData()
	if err != nil {
		return nil, err
	}

	for populationId, populationName := range populationData {
		ok, err := r.checkPopulationDefaultIdentityProviderData(populationId)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		commentData := map[string]string{
			"Population ID":         populationId,
			"Population Name":       populationName,
			"Resource Type":         r.ResourceType(),
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_default_identity_provider", populationName),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, populationId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOnePopulationDefaultIdentityProviderResource) getPopulationData() (map[string]string, error) {
	populationData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.Population](iter, "ReadAllPopulations", "GetPopulations", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, population := range apiObjs {
		populationId, populationIdOk := population.GetIdOk()
		populationName, populationNameOk := population.GetNameOk()

		if populationIdOk && populationNameOk {
			populationData[*populationId] = *populationName
		}
	}

	return populationData, nil
}

func (r *PingOnePopulationDefaultIdentityProviderResource) checkPopulationDefaultIdentityProviderData(populationId string) (bool, error) {
	_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.PopulationsApi.ReadOnePopulationDefaultIdp(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, populationId).Execute()

	return common.CheckSingletonResource(response, err, "ReadOnePopulationDefaultIdp", r.ResourceType())
}
