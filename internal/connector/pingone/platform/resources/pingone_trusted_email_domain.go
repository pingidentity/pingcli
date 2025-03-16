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
	_ connector.ExportableResource = &PingOneTrustedEmailDomainResource{}
)

type PingOneTrustedEmailDomainResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOne Trusted Email Domain Resource
func TrustedEmailDomain(clientInfo *connector.ClientInfo) *PingOneTrustedEmailDomainResource {
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

	trustedEmailDomainData, err := r.getTrustedEmailDomainData()
	if err != nil {
		return nil, err
	}

	for trustedEmailDomainId, trustedEmailDomainName := range trustedEmailDomainData {
		commentData := map[string]string{
			"Export Environment ID":     r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":             r.ResourceType(),
			"Trusted Email Domain ID":   trustedEmailDomainId,
			"Trusted Email Domain Name": trustedEmailDomainName,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       trustedEmailDomainName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, trustedEmailDomainId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneTrustedEmailDomainResource) getTrustedEmailDomainData() (map[string]string, error) {
	trustedEmailDomainData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	trustedEmailDomains, err := pingone.GetManagementAPIObjectsFromIterator[management.EmailDomain](iter, "ReadAllTrustedEmailDomains", "GetEmailDomains", r.ResourceType())
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
