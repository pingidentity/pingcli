package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneNotificationTemplateContentResource{}
)

type PingOneNotificationTemplateContentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneNotificationTemplateContentResource
func NotificationTemplateContent(clientInfo *connector.PingOneClientInfo) *PingOneNotificationTemplateContentResource {
	return &PingOneNotificationTemplateContentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneNotificationTemplateContentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	// Get all enabled language locales
	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadLanguages"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	enabledLocales := map[string]bool{}
	for _, languageInner := range embedded.GetLanguages() {
		if languageInner.Language != nil {
			language := languageInner.Language

			languageEnabled, languageEnabledOk := language.GetEnabledOk()
			languageLocale, languageLocaleOk := language.GetLocaleOk()

			if languageEnabledOk && languageLocaleOk && *languageEnabled {
				enabledLocales[*languageLocale] = true
			}
		}
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	// This is weird... the provider mentions many possible template types,
	// but pingone console and the API only support the following types:
	validTemplateNames := []management.EnumTemplateName{
		management.ENUMTEMPLATENAME_DEVICE_PAIRING,
		management.ENUMTEMPLATENAME_EMAIL_VERIFICATION_ADMIN,
		management.ENUMTEMPLATENAME_EMAIL_VERIFICATION_USER,
		management.ENUMTEMPLATENAME_GENERAL,
		management.ENUMTEMPLATENAME_NEW_DEVICE_PAIRED,
		management.ENUMTEMPLATENAME_STRONG_AUTHENTICATION,
		management.ENUMTEMPLATENAME_TRANSACTION,
		management.ENUMTEMPLATENAME_VERIFICATION_CODE_TEMPLATE,
	}

	// TODO: When the above hard-coded values are fixed, use the following
	// for _, templateNameEnum := range management.AllowedEnumTemplateNameEnumValues {
	for _, templateNameEnum := range validTemplateNames {
		apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadAllTemplateContents(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, templateNameEnum).Execute
		apiFunctionName := "ReadAllTemplateContents"

		embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
		if err != nil {
			return nil, err
		}

		for _, templateContents := range embedded.GetContents() {
			var (
				templateContentsId       *string
				templateContentsIdOk     bool
				templateDeliveryMethod   *management.EnumTemplateContentDeliveryMethod
				templateDeliveryMethodOk bool
				templateLocale           *string
				templateLocaleOk         bool
				templateVariant          *string
				templateVariantOk        bool
			)

			switch {
			case templateContents.TemplateContentEmail != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentEmail.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentEmail.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentEmail.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentEmail.GetVariantOk()
			case templateContents.TemplateContentPush != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentPush.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentPush.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentPush.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentPush.GetVariantOk()
			case templateContents.TemplateContentSMS != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentSMS.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentSMS.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentSMS.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentSMS.GetVariantOk()
			case templateContents.TemplateContentVoice != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentVoice.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentVoice.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentVoice.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentVoice.GetVariantOk()
			default:
				continue
			}

			// If the template content locale is not enabled, skip it
			if templateLocaleOk && !enabledLocales[*templateLocale] {
				continue
			}

			// This variable handles the case where template type, locale,
			// and delivery method are the same across two content instances
			// Append it to the ResourceName if present from SDK
			if templateVariantOk {
				*templateVariant = "_" + *templateVariant
			} else {
				emptyString := ""
				templateVariant = &emptyString
			}

			if templateContentsIdOk && templateDeliveryMethodOk && templateLocaleOk {
				commentData := map[string]string{
					"Resource Type":            r.ResourceType(),
					"Template Name":            string(templateNameEnum),
					"Template Delivery Method": string(*templateDeliveryMethod),
					"Template Locale":          string(*templateLocale),
					"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
					"Template Contents ID":     string(*templateContentsId),
				}

				if templateVariantOk {
					commentData["Template Variant"] = *templateVariant
				}

				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType:       r.ResourceType(),
					ResourceName:       fmt.Sprintf("%s_%s_%s%s", templateNameEnum, *templateDeliveryMethod, *templateLocale, *templateVariant),
					ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, templateNameEnum, *templateContentsId),
					CommentInformation: common.GenerateCommentInformation(commentData),
				})
			}
		}

	}

	return &importBlocks, nil
}

func (r *PingOneNotificationTemplateContentResource) ResourceType() string {
	return "pingone_notification_template_content"
}
