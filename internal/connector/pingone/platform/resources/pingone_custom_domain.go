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

func (r *PingOneCustomDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.CustomDomainsApi.ReadAllDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllDomains"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, customDomain := range embedded.GetCustomDomains() {
		customDomainName, customDomainNameOk := customDomain.GetDomainNameOk()
		customDomainId, customDomainIdOk := customDomain.GetIdOk()

		if customDomainIdOk && customDomainNameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Custom Domain Name":    *customDomainName,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Custom Domain ID":      *customDomainId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *customDomainName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *customDomainId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingOneCustomDomainResource) ResourceType() string {
	return "pingone_custom_domain"
}
