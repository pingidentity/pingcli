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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneAgreementEnableResource
func AgreementEnable(clientInfo *connector.PingOneClientInfo) *PingOneAgreementEnableResource {
	return &PingOneAgreementEnableResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneAgreementEnableResource) ResourceType() string {
	return "pingone_agreement_enable"
}

func (r *PingOneAgreementEnableResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportAgreementEnables()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneAgreementEnableResource) exportAgreementEnables() error {
	agreementImportBlocks, err := Agreement(r.clientInfo).ExportAll()
	if err != nil {
		return err
	}

	for _, importBlock := range *agreementImportBlocks {
		importBlock = connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_enable", importBlock.ResourceName),
			ResourceID:         importBlock.ResourceID,
			CommentInformation: importBlock.CommentInformation,
		}

		*r.importBlocks = append(*r.importBlocks, importBlock)
	}

	return nil
}
