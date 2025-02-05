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

	for trustedEmailDomainId, trustedEmailDomainName := range *trustedEmailDomainData {
		trustedEmailAddressData, err := r.getTrustedEmailAddressData(trustedEmailDomainId)
		if err != nil {
			return nil, err
		}

		for trustedEmailId, trustedEmailAddress := range *trustedEmailAddressData {
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

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneTrustedEmailAddressResource) getTrustedEmailDomainData() (*map[string]string, error) {
	trustedEmailDomainData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllTrustedEmailDomains", r.ResourceType())
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, trustedEmailDomain := range embedded.GetEmailDomains() {
			trustedEmailDomainId, trustedEmailDomainIdOk := trustedEmailDomain.GetIdOk()
			trustedEmailDomainName, trustedEmailDomainNameOk := trustedEmailDomain.GetDomainNameOk()

			if trustedEmailDomainIdOk && trustedEmailDomainNameOk {
				trustedEmailDomainData[*trustedEmailDomainId] = *trustedEmailDomainName
			}
		}
	}

	return &trustedEmailDomainData, nil
}

func (r *PingOneTrustedEmailAddressResource) getTrustedEmailAddressData(trustedEmailDomainId string) (*map[string]string, error) {
	trustedEmailAddressData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, trustedEmailDomainId).Execute()

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllTrustedEmailAddresses", r.ResourceType())
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, trustedEmail := range embedded.GetTrustedEmails() {
			trustedEmailAddress, trustedEmailAddressOk := trustedEmail.GetEmailAddressOk()
			trustedEmailId, trustedEmailIdOk := trustedEmail.GetIdOk()

			if trustedEmailAddressOk && trustedEmailIdOk {
				trustedEmailAddressData[*trustedEmailId] = *trustedEmailAddress
			}
		}
	}

	return &trustedEmailAddressData, nil
}
