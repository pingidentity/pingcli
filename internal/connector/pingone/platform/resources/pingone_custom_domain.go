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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneCustomDomainResource
func CustomDomain(clientInfo *connector.PingOneClientInfo) *PingOneCustomDomainResource {
	return &PingOneCustomDomainResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneCustomDomainResource) ResourceType() string {
	return "pingone_custom_domain"
}

func (r *PingOneCustomDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportCustomDomains()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneCustomDomainResource) exportCustomDomains() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.CustomDomainsApi.ReadAllDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllDomains", r.ResourceType())
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

		for _, customDomain := range embedded.GetCustomDomains() {
			customDomainName, customDomainNameOk := customDomain.GetDomainNameOk()
			customDomainId, customDomainIdOk := customDomain.GetIdOk()

			if customDomainIdOk && customDomainNameOk {
				r.addImportBlock(*customDomainId, *customDomainName)
			}
		}
	}

	return nil
}

func (r *PingOneCustomDomainResource) addImportBlock(customDomainId, customDomainName string) {
	commentData := map[string]string{
		"Custom Domain ID":      customDomainId,
		"Custom Domain Name":    customDomainName,
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       customDomainName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, customDomainId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
