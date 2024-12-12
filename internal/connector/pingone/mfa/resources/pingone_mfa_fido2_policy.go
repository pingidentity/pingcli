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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneMFAFido2PolicyResource
func MFAFido2Policy(clientInfo *connector.PingOneClientInfo) *PingOneMFAFido2PolicyResource {
	return &PingOneMFAFido2PolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneMFAFido2PolicyResource) ResourceType() string {
	return "pingone_mfa_fido2_policy"
}

func (r *PingOneMFAFido2PolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	fido2PolicyData, err := r.getFido2PolicyData()
	if err != nil {
		return nil, err
	}

	for fido2PolicyId, fido2PolicyName := range *fido2PolicyData {
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

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneMFAFido2PolicyResource) getFido2PolicyData() (*map[string]string, error) {
	fido2PolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.MFAAPIClient.FIDO2PolicyApi.ReadFIDO2Policies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadFIDO2Policies", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, fido2Policy := range embedded.GetFido2Policies() {
			fido2PolicyId, fido2PolicyIdOk := fido2Policy.GetIdOk()
			fido2PolicyName, fido2PolicyNameOk := fido2Policy.GetNameOk()

			if fido2PolicyIdOk && fido2PolicyNameOk {
				fido2PolicyData[*fido2PolicyId] = *fido2PolicyName
			}
		}
	}

	return &fido2PolicyData, nil
}
