package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
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

	domainData, err := getCustomDomainData(r.clientInfo, r.ResourceType())
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

func getCustomDomainData(clientInfo *connector.PingOneClientInfo, resourceType string) (*map[string]string, error) {
	domainData := make(map[string]string)

	iter := clientInfo.ApiClient.ManagementAPIClient.CustomDomainsApi.ReadAllDomains(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()
	customDomains, err := common.GetManagementAPIObjectsFromIterator[management.CustomDomain](iter, "ReadAllDomains", "GetCustomDomains", resourceType)
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

	return &domainData, nil
}
