package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneAgreementResource{}
)

type PingOneAgreementResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneAgreementResource
func Agreement(clientInfo *connector.PingOneClientInfo) *PingOneAgreementResource {
	return &PingOneAgreementResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneAgreementResource) ResourceType() string {
	return "pingone_agreement"
}

func (r *PingOneAgreementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportAgreements()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneAgreementResource) exportAgreements() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllAgreements", r.ResourceType())
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

		for _, agreement := range embedded.GetAgreements() {
			agreementId, agreementIdOk := agreement.GetIdOk()
			agreementName, agreementNameOk := agreement.GetNameOk()

			if agreementIdOk && agreementNameOk {
				r.addImportBlock(*agreementId, *agreementName)
			}
		}
	}

	return nil
}

func (r *PingOneAgreementResource) addImportBlock(agreementId, agreementName string) {
	commentData := map[string]string{
		"Agreement ID":          agreementId,
		"Agreement Name":        agreementName,
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       agreementName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, agreementId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
