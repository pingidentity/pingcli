package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneAgreementEnableResource{}
)

type PingOneAgreementEnableResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneAgreementEnableResource
func AgreementEnable(clientInfo *connector.PingOneClientInfo) *PingOneAgreementEnableResource {
	return &PingOneAgreementEnableResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneAgreementEnableResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement_enable resources...")

	agreementImportBlocks, err := Agreement(r.clientInfo).ExportAll()
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_enable resources...")

	for _, importBlock := range *agreementImportBlocks {
		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_enable", importBlock.ResourceName),
			ResourceID:         importBlock.ResourceID,
			CommentInformation: importBlock.CommentInformation,
		})
	}

	return &importBlocks, nil
}

func (r *PingOneAgreementEnableResource) ResourceType() string {
	return "pingone_agreement_enable"
}
