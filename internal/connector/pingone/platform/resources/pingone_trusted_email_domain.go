package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneTrustedEmailDomainResource{}
)

type PingoneTrustedEmailDomainResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a Pingone Trusted Email Domain Resource
func TrustedEmailDomain(clientInfo *connector.PingOneClientInfo) *PingoneTrustedEmailDomainResource {
	return &PingoneTrustedEmailDomainResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneTrustedEmailDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllTrustedEmailDomains"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, emailDomain := range embedded.GetEmailDomains() {
		emailDomainId, emailDomainIdOk := emailDomain.GetIdOk()
		emailDomainName, emailDomainNameOk := emailDomain.GetDomainNameOk()

		if emailDomainIdOk && emailDomainNameOk {
			commentData := map[string]string{
				"Resource Type":             r.ResourceType(),
				"Trusted Email Domain Name": *emailDomainName,
				"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
				"Trusted Email Domain ID":   *emailDomainId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *emailDomainName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *emailDomainId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneTrustedEmailDomainResource) ResourceType() string {
	return "pingone_trusted_email_domain"
}
