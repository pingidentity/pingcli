package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneTrustedEmailAddressResource{}
)

type PingOneTrustedEmailAddressResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneTrustedEmailAddressResource
func TrustedEmailAddress(clientInfo *connector.PingOneClientInfo) *PingOneTrustedEmailAddressResource {
	return &PingOneTrustedEmailAddressResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneTrustedEmailAddressResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllTrustedEmailDomains"

	emailDomainEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, trustedEmailDomain := range emailDomainEmbedded.GetEmailDomains() {
		trustedEmailDomainId, trustedEmailDomainIdOk := trustedEmailDomain.GetIdOk()
		trustedEmailDomainName, trustedEmailDomainNameOk := trustedEmailDomain.GetDomainNameOk()

		if trustedEmailDomainIdOk && trustedEmailDomainNameOk {
			apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *trustedEmailDomainId).Execute
			apiFunctionName := "ReadAllTrustedEmailAddresses"

			trustedEmailAddressEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, trustedEmail := range trustedEmailAddressEmbedded.GetTrustedEmails() {
				trustedEmailAddress, trustedEmailAddressOk := trustedEmail.GetEmailAddressOk()
				trustedEmailId, trustedEmailIdOk := trustedEmail.GetIdOk()

				if trustedEmailAddressOk && trustedEmailIdOk {
					commentData := map[string]string{
						"Resource Type":             r.ResourceType(),
						"Trusted Email Domain Name": *trustedEmailDomainName,
						"Trusted Email Address":     *trustedEmailAddress,
						"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
						"Trusted Email Domain ID":   *trustedEmailDomainId,
						"Trusted Email Address ID":  *trustedEmailId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *trustedEmailDomainName, *trustedEmailAddress),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *trustedEmailDomainId, *trustedEmailId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneTrustedEmailAddressResource) ResourceType() string {
	return "pingone_trusted_email_address"
}
