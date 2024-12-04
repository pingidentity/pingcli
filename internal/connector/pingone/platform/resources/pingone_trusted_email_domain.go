package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneTrustedEmailDomainResource{}
)

type PingOneTrustedEmailDomainResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOne Trusted Email Domain Resource
func TrustedEmailDomain(clientInfo *connector.PingOneClientInfo) *PingOneTrustedEmailDomainResource {
	return &PingOneTrustedEmailDomainResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneTrustedEmailDomainResource) ResourceType() string {
	return "pingone_trusted_email_domain"
}

func (r *PingOneTrustedEmailDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportTrustedEmailDomains()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneTrustedEmailDomainResource) exportTrustedEmailDomains() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllTrustedEmailDomains", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, emailDomain := range embedded.GetEmailDomains() {
			emailDomainId, emailDomainIdOk := emailDomain.GetIdOk()
			emailDomainName, emailDomainNameOk := emailDomain.GetDomainNameOk()

			if emailDomainIdOk && emailDomainNameOk {
				r.addImportBlock(*emailDomainId, *emailDomainName)
			}
		}
	}

	return nil
}

func (r *PingOneTrustedEmailDomainResource) addImportBlock(emailDomainId, emailDomainName string) {
	commentData := map[string]string{
		"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
		"Resource Type":             r.ResourceType(),
		"Trusted Email Domain ID":   emailDomainId,
		"Trusted Email Domain Name": emailDomainName,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       emailDomainName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, emailDomainId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
