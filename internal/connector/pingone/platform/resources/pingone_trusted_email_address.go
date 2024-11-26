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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneTrustedEmailAddressResource
func TrustedEmailAddress(clientInfo *connector.PingOneClientInfo) *PingOneTrustedEmailAddressResource {
	return &PingOneTrustedEmailAddressResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneTrustedEmailAddressResource) ResourceType() string {
	return "pingone_trusted_email_address"
}

func (r *PingOneTrustedEmailAddressResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportTrustedEmailAddresses()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneTrustedEmailAddressResource) exportTrustedEmailAddresses() error {
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

		for _, trustedEmailDomain := range embedded.GetEmailDomains() {
			trustedEmailDomainId, trustedEmailDomainIdOk := trustedEmailDomain.GetIdOk()
			trustedEmailDomainName, trustedEmailDomainNameOk := trustedEmailDomain.GetDomainNameOk()

			if trustedEmailDomainIdOk && trustedEmailDomainNameOk {
				err := r.exportTrustedEmailAddressesByDomain(*trustedEmailDomainId, *trustedEmailDomainName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneTrustedEmailAddressResource) exportTrustedEmailAddressesByDomain(trustedEmailDomainId, trustedEmailDomainName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, trustedEmailDomainId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllTrustedEmailAddresses", r.ResourceType())
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

		for _, trustedEmail := range embedded.GetTrustedEmails() {
			trustedEmailAddress, trustedEmailAddressOk := trustedEmail.GetEmailAddressOk()
			trustedEmailId, trustedEmailIdOk := trustedEmail.GetIdOk()

			if trustedEmailAddressOk && trustedEmailIdOk {
				r.addImportBlock(trustedEmailDomainId, trustedEmailDomainName, *trustedEmailId, *trustedEmailAddress)
			}
		}
	}

	return nil
}

func (r *PingOneTrustedEmailAddressResource) addImportBlock(trustedEmailDomainId, trustedEmailDomainName, trustedEmailId, trustedEmailAddress string) {
	commentData := map[string]string{
		"Export Environment ID":     r.clientInfo.ExportEnvironmentID,
		"Resource Type":             r.ResourceType(),
		"Trusted Email Address":     trustedEmailAddress,
		"Trusted Email Address ID":  trustedEmailId,
		"Trusted Email Domain ID":   trustedEmailDomainId,
		"Trusted Email Domain Name": trustedEmailDomainName,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s", trustedEmailDomainName, trustedEmailAddress),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, trustedEmailDomainId, trustedEmailId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
