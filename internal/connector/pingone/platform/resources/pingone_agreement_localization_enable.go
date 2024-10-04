package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementLocalizationEnableResource{}
)

type PingoneAgreementLocalizationEnableResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAgreementLocalizationEnableResource
func AgreementLocalizationEnable(clientInfo *connector.PingOneClientInfo) *PingoneAgreementLocalizationEnableResource {
	return &PingoneAgreementLocalizationEnableResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAgreementLocalizationEnableResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement_localization_enable resources...")

	localizationImportBlocks, err := AgreementLocalization(r.clientInfo).ExportAll()
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_localization_enable resources...")

	for _, importBlock := range *localizationImportBlocks {
		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_enable", importBlock.ResourceName),
			ResourceID:         importBlock.ResourceID,
			CommentInformation: importBlock.CommentInformation,
		})
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementLocalizationEnableResource) ResourceType() string {
	return "pingone_agreement_localization_enable"
}
