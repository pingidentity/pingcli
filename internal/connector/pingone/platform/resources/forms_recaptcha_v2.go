// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneFormsRecaptchaV2Resource{}
)

type PingOneFormsRecaptchaV2Resource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneFormsRecaptchaV2Resource
func FormsRecaptchaV2(clientInfo *connector.ClientInfo) *PingOneFormsRecaptchaV2Resource {
	return &PingOneFormsRecaptchaV2Resource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneFormsRecaptchaV2Resource) ResourceType() string {
	return "pingone_forms_recaptcha_v2"
}

func (r *PingOneFormsRecaptchaV2Resource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ok, err := r.checkFormsRecaptchaV2Data()
	if err != nil {
		return nil, err
	}
	if !ok {
		return &importBlocks, nil
	}

	commentData := map[string]string{
		"Resource Type":         r.ResourceType(),
		"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       r.ResourceType(),
		ResourceID:         r.clientInfo.PingOneExportEnvironmentID,
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	importBlocks = append(importBlocks, importBlock)

	return &importBlocks, nil
}

func (r *PingOneFormsRecaptchaV2Resource) checkFormsRecaptchaV2Data() (bool, error) {
	_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.RecaptchaConfigurationApi.ReadRecaptchaConfiguration(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	return common.CheckSingletonResource(response, err, "ReadRecaptchaConfiguration", r.ResourceType())
}
