package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
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

type mfaFido2PolicyImportBlockData struct {
	ID   string
	Name string
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

func (r *PingOneMFAFido2PolicyResource) fido2Policies() (fido2Policies []mfa.FIDO2Policy, err error) {
	fido2Policies = []mfa.FIDO2Policy{}

	// Fetch all pingone_mfa_application_push_credentials resources for the given pingone_application
	entityArrayPagedIterator := r.clientInfo.ApiClient.MFAAPIClient.FIDO2PolicyApi.ReadFIDO2Policies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	allEmbedded, err := common.GetAllMFAEmbedded(entityArrayPagedIterator, "ReadFIDO2Policies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, embedded := range allEmbedded {
		for _, fido2Policy := range embedded.GetFido2Policies() {
			fido2Policies = append(fido2Policies, fido2Policy)
		}
	}

	return fido2Policies, nil
}

func (r *PingOneMFAFido2PolicyResource) importBlockData() (importBlockData []mfaFido2PolicyImportBlockData, err error) {
	importBlockData = []mfaFido2PolicyImportBlockData{}

	fido2Policies, err := r.fido2Policies()
	if err != nil {
		return nil, err
	}

	for _, fido2Policy := range fido2Policies {
		fido2PolicyName, fido2PolicyNameOk := fido2Policy.GetNameOk()
		fido2PolicyId, fido2PolicyIdOk := fido2Policy.GetIdOk()

		if fido2PolicyNameOk && fido2PolicyIdOk {
			importBlockData = append(importBlockData, mfaFido2PolicyImportBlockData{
				ID:   *fido2PolicyId,
				Name: *fido2PolicyName,
			})
		}
	}

	return importBlockData, nil
}

func (r *PingOneMFAFido2PolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlockData, err := r.importBlockData()
	if err != nil {
		return nil, err
	}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, data := range importBlockData {
		commentData := map[string]string{
			"Resource Type":         r.ResourceType(),
			"FIDO2 Policy Name":     data.Name,
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"FIDO2 Policy ID":       data.ID,
		}

		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       data.Name,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, data.ID),
			CommentInformation: common.GenerateCommentInformation(commentData),
		})
	}

	return &importBlocks, nil
}
