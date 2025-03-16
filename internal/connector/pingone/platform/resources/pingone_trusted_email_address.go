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
	_ connector.ExportableResource = &PingOneTrustedEmailAddressResource{}
)

type PingOneTrustedEmailAddressResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneTrustedEmailAddressResource
func TrustedEmailAddress(clientInfo *connector.ClientInfo) *PingOneTrustedEmailAddressResource {
	return &PingOneTrustedEmailAddressResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneTrustedEmailAddressResource) ResourceType() string {
	return "pingone_trusted_email_address"
}

func (r *PingOneTrustedEmailAddressResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	trustedEmailDomainData, err := r.getTrustedEmailDomainData()
	if err != nil {
		return nil, err
	}

	for trustedEmailDomainId, trustedEmailDomainName := range trustedEmailDomainData {
		trustedEmailAddressData, err := r.getTrustedEmailAddressData(trustedEmailDomainId)
		if err != nil {
			return nil, err
		}

		for trustedEmailId, trustedEmailAddress := range trustedEmailAddressData {
			commentData := map[string]string{
				"Export Environment ID":     r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":             r.ResourceType(),
				"Trusted Email Address":     trustedEmailAddress,
				"Trusted Email Address ID":  trustedEmailId,
				"Trusted Email Domain ID":   trustedEmailDomainId,
				"Trusted Email Domain Name": trustedEmailDomainName,
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", trustedEmailDomainName, trustedEmailAddress),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, trustedEmailDomainId, trustedEmailId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneTrustedEmailAddressResource) getTrustedEmailDomainData() (map[string]string, error) {
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

func (r *PingOneTrustedEmailAddressResource) getTrustedEmailAddressData(trustedEmailDomainId string) (map[string]string, error) {
	trustedEmailAddressData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, trustedEmailDomainId).Execute()
	trustedEmailAddresses, err := pingone.GetManagementAPIObjectsFromIterator[management.EmailDomainTrustedEmail](iter, "ReadAllTrustedEmailAddresses", "GetTrustedEmails", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, trustedEmail := range trustedEmailAddresses {
		trustedEmailAddress, trustedEmailAddressOk := trustedEmail.GetEmailAddressOk()
		trustedEmailId, trustedEmailIdOk := trustedEmail.GetIdOk()

		if trustedEmailAddressOk && trustedEmailIdOk {
			trustedEmailAddressData[*trustedEmailId] = *trustedEmailAddress
		}
	}

	return trustedEmailAddressData, nil
}
