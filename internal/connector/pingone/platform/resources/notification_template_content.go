// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources

import (
	"errors"
	"fmt"
	"io"
	"net/http"
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
		templateContentData, err := r.getNotificationTemplateContentData(templateName)
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
					"Export Environment ID":            r.clientInfo.PingOneExportEnvironmentID,
					"Template Content ID":              templateContentId,
				}

				if templateContentVariant != "" {
					commentData["Template Content Variant"] = templateContentVariant
					templateContentVariant = fmt.Sprintf("_%s", templateContentVariant)
				}

				importBlock := connector.ImportBlock{
					ResourceType:       r.ResourceType(),
					ResourceName:       fmt.Sprintf("%s_%s_%s%s_%s", string(templateName), templateContentDeliveryMethod, templateContentLocale, templateContentVariant, templateContentId),
					ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, string(templateName), templateContentId),
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

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedLanguagesInner](iter, "ReadLanguages", "GetLanguages", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		if innerObj.Language != nil {
			languageLocale, languageLocaleOk := innerObj.Language.GetLocaleOk()
			languageEnabled, languageEnabledOk := innerObj.Language.GetEnabledOk()

			if languageLocaleOk && languageEnabledOk && *languageEnabled {
				enabledLocales[*languageLocale] = true
			}
		}
	}

	return enabledLocales, nil
}

func (r *PingOneNotificationTemplateContentResource) getTemplateNames() (arr []management.EnumTemplateName, err error) {
	templateNames := []management.EnumTemplateName{}

	for _, templateName := range management.AllowedEnumTemplateNameEnumValues {
		_, response, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadOneTemplate(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, templateName).Execute()
		// When PingOne services are not enabled in an environment,
		// the response code for the templates related to that service is
		// 400 Bad Request - "CONSTRAINT_VIOLATION"
		if err != nil && response.StatusCode == http.StatusBadRequest && response.Status == "400 Bad Request" {
			defer func() {
				cErr := response.Body.Close()
				if cErr != nil {
					err = errors.Join(err, cErr)
				}
			}()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}

			if strings.Contains(string(body), "CONSTRAINT_VIOLATION") {
				continue
			} // else fall through to handle other errors
		}

		// Handle all other errors or bad responses
		ok, err := common.HandleClientResponse(response, err, "ReadOneTemplate", r.ResourceType())
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		templateNames = append(templateNames, templateName)
	}

	return templateNames, nil
}

func (r *PingOneNotificationTemplateContentResource) getNotificationTemplateContentData(templateName management.EnumTemplateName) ([]NotificationTemplateContentData, error) {
	notificationTemplateContentData := []NotificationTemplateContentData{}

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadAllTemplateContents(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, templateName).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.TemplateContent](iter, "ReadAllTemplateContents", "GetContents", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, notificationTemplateContent := range apiObjs {
		var (
			notificationTemplateContentId               *string
			notificationTemplateContentIdOk             bool
			notificationTemplateContentDeliveryMethod   *management.EnumTemplateContentDeliveryMethod
			notificationTemplateContentDeliveryMethodOk bool
			notificationTemplateContentLocale           *string
			notificationTemplateContentLocaleOk         bool
			notificationTemplateContentVariant          string
		)

		switch {
		case notificationTemplateContent.TemplateContentPush != nil:
			notificationTemplateContentId, notificationTemplateContentIdOk = notificationTemplateContent.TemplateContentPush.GetIdOk()
			notificationTemplateContentDeliveryMethod, notificationTemplateContentDeliveryMethodOk = notificationTemplateContent.TemplateContentPush.GetDeliveryMethodOk()
			notificationTemplateContentLocale, notificationTemplateContentLocaleOk = notificationTemplateContent.TemplateContentPush.GetLocaleOk()
			notificationTemplateContentVariant = notificationTemplateContent.TemplateContentPush.GetVariant()
		case notificationTemplateContent.TemplateContentSMS != nil:
			notificationTemplateContentId, notificationTemplateContentIdOk = notificationTemplateContent.TemplateContentSMS.GetIdOk()
			notificationTemplateContentDeliveryMethod, notificationTemplateContentDeliveryMethodOk = notificationTemplateContent.TemplateContentSMS.GetDeliveryMethodOk()
			notificationTemplateContentLocale, notificationTemplateContentLocaleOk = notificationTemplateContent.TemplateContentSMS.GetLocaleOk()
			notificationTemplateContentVariant = notificationTemplateContent.TemplateContentSMS.GetVariant()
		case notificationTemplateContent.TemplateContentEmail != nil:
			notificationTemplateContentId, notificationTemplateContentIdOk = notificationTemplateContent.TemplateContentEmail.GetIdOk()
			notificationTemplateContentDeliveryMethod, notificationTemplateContentDeliveryMethodOk = notificationTemplateContent.TemplateContentEmail.GetDeliveryMethodOk()
			notificationTemplateContentLocale, notificationTemplateContentLocaleOk = notificationTemplateContent.TemplateContentEmail.GetLocaleOk()
			notificationTemplateContentVariant = notificationTemplateContent.TemplateContentEmail.GetVariant()
		case notificationTemplateContent.TemplateContentVoice != nil:
			notificationTemplateContentId, notificationTemplateContentIdOk = notificationTemplateContent.TemplateContentVoice.GetIdOk()
			notificationTemplateContentDeliveryMethod, notificationTemplateContentDeliveryMethodOk = notificationTemplateContent.TemplateContentVoice.GetDeliveryMethodOk()
			notificationTemplateContentLocale, notificationTemplateContentLocaleOk = notificationTemplateContent.TemplateContentVoice.GetLocaleOk()
			notificationTemplateContentVariant = notificationTemplateContent.TemplateContentVoice.GetVariant()
		default:
			output.Warn(fmt.Sprintf("Template content '%v' for template '%s' is not one of: Push, SMS, Email, or Voice. Skipping export.", notificationTemplateContent, templateName), nil)

			continue
		}

		if notificationTemplateContentIdOk && notificationTemplateContentDeliveryMethodOk && notificationTemplateContentLocaleOk {
			notificationTemplateContentData = append(notificationTemplateContentData, NotificationTemplateContentData{
				TemplateContentId:             *notificationTemplateContentId,
				TemplateContentDeliveryMethod: string(*notificationTemplateContentDeliveryMethod),
				TemplateContentLocale:         *notificationTemplateContentLocale,
				TemplateContentVariant:        notificationTemplateContentVariant,
			})
		}
	}

	return notificationTemplateContentData, nil
}
