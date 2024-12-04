package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFAFido2PolicyResource{}
)

type PingOneMFAFido2PolicyResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneMFAFido2PolicyResource
func MFAFido2Policy(clientInfo *connector.PingOneClientInfo) *PingOneMFAFido2PolicyResource {
	return &PingOneMFAFido2PolicyResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneMFAFido2PolicyResource) ResourceType() string {
	return "pingone_mfa_fido2_policy"
}

func (r *PingOneMFAFido2PolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportFido2Policies()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneMFAFido2PolicyResource) exportFido2Policies() error {
	iter := r.clientInfo.ApiClient.MFAAPIClient.FIDO2PolicyApi.ReadFIDO2Policies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadFIDO2Policies", r.ResourceType())
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

		for _, fido2Policy := range embedded.GetFido2Policies() {
			fido2PolicyId, fido2PolicyIdOk := fido2Policy.GetIdOk()
			fido2PolicyName, fido2PolicyNameOk := fido2Policy.GetNameOk()

			if fido2PolicyIdOk && fido2PolicyNameOk {
				r.addImportBlock(*fido2PolicyId, *fido2PolicyName)
			}
		}
	}

	return nil
}

func (r *PingOneMFAFido2PolicyResource) addImportBlock(fido2PolicyId, fido2PolicyName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"FIDO2 Policy ID":       fido2PolicyId,
		"FIDO2 Policy Name":     fido2PolicyName,
		"Resource Type":         r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fido2PolicyName,
		ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, fido2PolicyId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
