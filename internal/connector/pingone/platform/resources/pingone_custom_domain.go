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
	_ connector.ExportableResource = &PingOneCustomDomainResource{}
)

type PingOneCustomDomainResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneCustomDomainResource
func CustomDomain(clientInfo *connector.ClientInfo) *PingOneCustomDomainResource {
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

	for domainId, domainName := range domainData {
		commentData := map[string]string{
			"Custom Domain ID":      domainId,
			"Custom Domain Name":    domainName,
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       domainName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, domainId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneCustomDomainResource) getCustomDomainData() (map[string]string, error) {
	domainData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.CustomDomainsApi.ReadAllDomains(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	customDomains, err := pingone.GetManagementAPIObjectsFromIterator[management.CustomDomain](iter, "ReadAllDomains", "GetCustomDomains", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, customDomain := range customDomains {
		customDomainName, customDomainNameOk := customDomain.GetDomainNameOk()
		customDomainId, customDomainIdOk := customDomain.GetIdOk()

		if customDomainIdOk && customDomainNameOk {
			domainData[*customDomainId] = *customDomainName
		}
	}

	return domainData, nil
}
