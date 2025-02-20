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
	_ connector.ExportableResource = &PingOneTrustedEmailDomainResource{}
)

type PingOneTrustedEmailDomainResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOne Trusted Email Domain Resource
func TrustedEmailDomain(clientInfo *connector.PingOneClientInfo) *PingOneTrustedEmailDomainResource {
	return &PingOneTrustedEmailDomainResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneTrustedEmailDomainResource) ResourceType() string {
	return "pingone_trusted_email_domain"
}

func (r *PingOneTrustedEmailDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	trustedEmailDomainData, err := getTrustedEmailDomainData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for trustedEmailDomainId, trustedEmailDomainName := range trustedEmailDomainData {
		commentData := map[string]string{
			"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
			"Resource Type":             r.ResourceType(),
			"Trusted Email Domain ID":   trustedEmailDomainId,
			"Trusted Email Domain Name": trustedEmailDomainName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       trustedEmailDomainName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, trustedEmailDomainId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func getTrustedEmailDomainData(clientInfo *connector.PingOneClientInfo, resourceType string) (map[string]string, error) {
	trustedEmailDomainData := make(map[string]string)

	iter := clientInfo.ApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(clientInfo.Context, clientInfo.ExportEnvironmentID).Execute()
	trustedEmailDomains, err := common.GetManagementAPIObjectsFromIterator[management.EmailDomain](iter, "ReadAllTrustedEmailDomains", "GetEmailDomains", resourceType)
	if err != nil {
		return nil, err
	}

	for _, trustedEmailDomain := range trustedEmailDomains {
		trustedEmailDomainId, trustedEmailDomainIdOk := trustedEmailDomain.GetIdOk()
		trustedEmailDomainName, trustedEmailDomainNameOk := trustedEmailDomain.GetDomainNameOk()

		if trustedEmailDomainIdOk && trustedEmailDomainNameOk {
			trustedEmailDomainData[*trustedEmailDomainId] = *trustedEmailDomainName
		}
	}

	return trustedEmailDomainData, nil
}
