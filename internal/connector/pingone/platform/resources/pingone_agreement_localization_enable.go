package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneAgreementLocalizationEnableResource{}
)

type PingOneAgreementLocalizationEnableResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneAgreementLocalizationEnableResource
func AgreementLocalizationEnable(clientInfo *connector.PingOneClientInfo) *PingOneAgreementLocalizationEnableResource {
	return &PingOneAgreementLocalizationEnableResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneAgreementLocalizationEnableResource) ResourceType() string {
	return "pingone_agreement_localization_enable"
}

func (r *PingOneAgreementLocalizationEnableResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	agreementLocalizationImportBlocks, err := AgreementLocalization(r.clientInfo).ExportAll()
	if err != nil {
		return nil, err
	}

	for _, importBlock := range *agreementLocalizationImportBlocks {
		importBlock = connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_enable", importBlock.ResourceName),
			ResourceID:         importBlock.ResourceID,
			CommentInformation: importBlock.CommentInformation,
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}
