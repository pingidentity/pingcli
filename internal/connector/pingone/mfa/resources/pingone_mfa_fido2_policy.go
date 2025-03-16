package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneMFAFido2PolicyResource{}
)

type PingOneMFAFido2PolicyResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneMFAFido2PolicyResource
func MFAFido2Policy(clientInfo *connector.ClientInfo) *PingOneMFAFido2PolicyResource {
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

	for fido2PolicyId, fido2PolicyName := range fido2PolicyData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"FIDO2 Policy ID":       fido2PolicyId,
			"FIDO2 Policy Name":     fido2PolicyName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fido2PolicyName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, fido2PolicyId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneMFAFido2PolicyResource) getFido2PolicyData() (map[string]string, error) {
	fido2PolicyData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.MFAAPIClient.FIDO2PolicyApi.ReadFIDO2Policies(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	fido2Policies, err := pingone.GetMfaAPIObjectsFromIterator[mfa.FIDO2Policy](iter, "ReadFIDO2Policies", "GetFido2Policies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, fido2Policy := range fido2Policies {
		fido2PolicyId, fido2PolicyIdOk := fido2Policy.GetIdOk()
		fido2PolicyName, fido2PolicyNameOk := fido2Policy.GetNameOk()

		if fido2PolicyIdOk && fido2PolicyNameOk {
			fido2PolicyData[*fido2PolicyId] = *fido2PolicyName
		}
	}

	return fido2PolicyData, nil
}
