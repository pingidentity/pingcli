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

	trustedEmailDomainData, err := getTrustedEmailDomainData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	for trustedEmailDomainId, trustedEmailDomainName := range trustedEmailDomainData {
		trustedEmailAddressData, err := getTrustedEmailAddressData(r.clientInfo, r.ResourceType(), trustedEmailDomainId)
		if err != nil {
			return nil, err
		}

		for trustedEmailId, trustedEmailAddress := range trustedEmailAddressData {
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

func getTrustedEmailAddressData(clientInfo *connector.PingOneClientInfo, resourceType, trustedEmailDomainId string) (map[string]string, error) {
	trustedEmailAddressData := make(map[string]string)

	iter := clientInfo.ApiClient.ManagementAPIClient.TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses(clientInfo.Context, clientInfo.ExportEnvironmentID, trustedEmailDomainId).Execute()
	trustedEmailAddresses, err := common.GetManagementAPIObjectsFromIterator[management.EmailDomainTrustedEmail](iter, "ReadAllTrustedEmailAddresses", "GetTrustedEmails", resourceType)
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
