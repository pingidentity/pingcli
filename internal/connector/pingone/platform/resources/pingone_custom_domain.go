package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneCustomDomainResource{}
)

type PingOneCustomDomainResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneCustomDomainResource
func CustomDomain(clientInfo *connector.PingOneClientInfo) *PingOneCustomDomainResource {
	return &PingOneCustomDomainResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneCustomDomainResource) ResourceType() string {
	return "pingone_custom_domain"
}

func (r *PingOneCustomDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	domainData, err := r.getCustomDomainData()
	if err != nil {
		return nil, err
	}

	for domainId, domainName := range *domainData {
		commentData := map[string]string{
			"Custom Domain ID":      domainId,
			"Custom Domain Name":    domainName,
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       domainName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, domainId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneCustomDomainResource) getCustomDomainData() (*map[string]string, error) {
	domainData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.CustomDomainsApi.ReadAllDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllDomains", r.ResourceType())
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

		for _, customDomain := range embedded.GetCustomDomains() {
			customDomainName, customDomainNameOk := customDomain.GetDomainNameOk()
			customDomainId, customDomainIdOk := customDomain.GetIdOk()

			if customDomainIdOk && customDomainNameOk {
				domainData[*customDomainId] = *customDomainName
			}
		}
	}

	return &domainData, nil
}
