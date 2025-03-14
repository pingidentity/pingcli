package resources

import (
	"fmt"
	"io"
	"strings"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
)

type NotificationTemplateContentData struct {
	TemplateContentId             string
	TemplateContentDeliveryMethod string
	TemplateContentLocale         string
	TemplateContentVariant        string
}

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneNotificationTemplateContentResource{}
)

type PingOneNotificationTemplateContentResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneNotificationTemplateContentResource
func NotificationTemplateContent(clientInfo *connector.ClientInfo) *PingOneNotificationTemplateContentResource {
	return &PingOneNotificationTemplateContentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneNotificationTemplateContentResource) ResourceType() string {
	return "pingone_notification_template_content"
}

func (r *PingOneNotificationTemplateContentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	enabledLocales, err := r.getEnabledLocales()
	if err != nil {
		return nil, err
	}

	templateNames, err := r.getTemplateNames()
	if err != nil {
		return nil, err
	}

	for _, templateName := range templateNames {
		templateContentData, err := r.getTemplateContentData(templateName)
		if err != nil {
			return nil, err
		}

		for _, templateContentInfo := range templateContentData {
			templateContentId := templateContentInfo.TemplateContentId
			templateContentDeliveryMethod := templateContentInfo.TemplateContentDeliveryMethod
			templateContentLocale := templateContentInfo.TemplateContentLocale
			templateContentVariant := templateContentInfo.TemplateContentVariant

			// Only export template content if the locale is enabled
			if enabledLocales[templateContentLocale] {
				commentData := map[string]string{
					"Resource Type":                    r.ResourceType(),
					"Template Name":                    string(templateName),
					"Template Content Delivery Method": templateContentDeliveryMethod,
					"Template Content Locale":          templateContentLocale,
					"Export Environment ID":            r.clientInfo.ExportEnvironmentID,
					"Template Content ID":              templateContentId,
				}

				if templateContentVariant != "" {
					commentData["Template Content Variant"] = templateContentVariant
					templateContentVariant = fmt.Sprintf("_%s", templateContentVariant)
				}

				importBlock := connector.ImportBlock{
					ResourceType:       r.ResourceType(),
					ResourceName:       fmt.Sprintf("%s_%s_%s%s", string(templateName), templateContentDeliveryMethod, templateContentLocale, templateContentVariant),
					ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, string(templateName), templateContentId),
					CommentInformation: common.GenerateCommentInformation(commentData),
				}

				importBlocks = append(importBlocks, importBlock)
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneNotificationTemplateContentResource) getEnabledLocales() (map[string]bool, error) {
	enabledLocales := make(map[string]bool)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	languageInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedLanguagesInner](iter, "ReadLanguages", "GetLanguages", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, languageInner := range languageInners {
		if languageInner.Language != nil {
			languageLocale, languageLocaleOk := languageInner.Language.GetLocaleOk()
			languageEnabled, languageEnabledOk := languageInner.Language.GetEnabledOk()

			if languageLocaleOk && languageEnabledOk && *languageEnabled {
				enabledLocales[*languageLocale] = true
			}
		}
	}

	return enabledLocales, nil
}

func (r *PingOneNotificationTemplateContentResource) getTemplateNames() ([]management.EnumTemplateName, error) {
	templateNames := []management.EnumTemplateName{}

	for _, templateName := range management.AllowedEnumTemplateNameEnumValues {
		_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadOneTemplate(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, templateName).Execute()
		// When PingOne services are not enabled in an environment,
		// the response code for the templates related to that service is
		// 400 Bad Request - "CONSTRAINT_VIOLATION"
		if err != nil && response.StatusCode == 400 && response.Status == "400 Bad Request" {
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}

			if strings.Contains(string(body), "CONSTRAINT_VIOLATION") {
				continue
			}
		}

		// Handle all other errors or bad responses
		ok, err := common.HandleClientResponse(response, err, "ReadOneTemplate", r.ResourceType())
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}

		templateNames = append(templateNames, templateName)
	}

	return templateNames, nil
}

func (r *PingOneNotificationTemplateContentResource) getTemplateContentData(templateName management.EnumTemplateName) ([]NotificationTemplateContentData, error) {
	templateContentData := []NotificationTemplateContentData{}

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadAllTemplateContents(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, templateName).Execute()
	templateContents, err := pingone.GetManagementAPIObjectsFromIterator[management.TemplateContent](iter, "ReadAllTemplateContents", "GetContents", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, templateContent := range templateContents {
		var (
			templateContentId               *string
			templateContentIdOk             bool
			templateContentDeliveryMethod   *management.EnumTemplateContentDeliveryMethod
			templateContentDeliveryMethodOk bool
			templateContentLocale           *string
			templateContentLocaleOk         bool
			templateContentVariant          string
		)

		switch {
		case templateContent.TemplateContentPush != nil:
			templateContentId, templateContentIdOk = templateContent.TemplateContentPush.GetIdOk()
			templateContentDeliveryMethod, templateContentDeliveryMethodOk = templateContent.TemplateContentPush.GetDeliveryMethodOk()
			templateContentLocale, templateContentLocaleOk = templateContent.TemplateContentPush.GetLocaleOk()
			templateContentVariant = templateContent.TemplateContentPush.GetVariant()
		case templateContent.TemplateContentSMS != nil:
			templateContentId, templateContentIdOk = templateContent.TemplateContentSMS.GetIdOk()
			templateContentDeliveryMethod, templateContentDeliveryMethodOk = templateContent.TemplateContentSMS.GetDeliveryMethodOk()
			templateContentLocale, templateContentLocaleOk = templateContent.TemplateContentSMS.GetLocaleOk()
			templateContentVariant = templateContent.TemplateContentSMS.GetVariant()
		case templateContent.TemplateContentEmail != nil:
			templateContentId, templateContentIdOk = templateContent.TemplateContentEmail.GetIdOk()
			templateContentDeliveryMethod, templateContentDeliveryMethodOk = templateContent.TemplateContentEmail.GetDeliveryMethodOk()
			templateContentLocale, templateContentLocaleOk = templateContent.TemplateContentEmail.GetLocaleOk()
			templateContentVariant = templateContent.TemplateContentEmail.GetVariant()
		case templateContent.TemplateContentVoice != nil:
			templateContentId, templateContentIdOk = templateContent.TemplateContentVoice.GetIdOk()
			templateContentDeliveryMethod, templateContentDeliveryMethodOk = templateContent.TemplateContentVoice.GetDeliveryMethodOk()
			templateContentLocale, templateContentLocaleOk = templateContent.TemplateContentVoice.GetLocaleOk()
			templateContentVariant = templateContent.TemplateContentVoice.GetVariant()
		default:
			output.Warn(fmt.Sprintf("Template content '%v' for template '%s' is not one of: Push, SMS, Email, or Voice. Skipping export.", templateContent, templateName), nil)
			continue
		}

		if templateContentIdOk && templateContentDeliveryMethodOk && templateContentLocaleOk {
			templateContentData = append(templateContentData, NotificationTemplateContentData{
				TemplateContentId:             *templateContentId,
				TemplateContentDeliveryMethod: string(*templateContentDeliveryMethod),
				TemplateContentLocale:         *templateContentLocale,
				TemplateContentVariant:        templateContentVariant,
			})
		}
	}

	return templateContentData, nil
}
