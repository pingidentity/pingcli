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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneNotificationTemplateContentResource
func NotificationTemplateContent(clientInfo *connector.PingOneClientInfo) *PingOneNotificationTemplateContentResource {
	return &PingOneNotificationTemplateContentResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneNotificationTemplateContentResource) ResourceType() string {
	return "pingone_notification_template_content"
}

func (r *PingOneNotificationTemplateContentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportNotificationTemplateContents()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneNotificationTemplateContentResource) exportNotificationTemplateContents() error {
	for _, templateName := range management.AllowedEnumTemplateNameEnumValues {
		err := r.exportNotificationTemplateContentsByTemplate(templateName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PingOneNotificationTemplateContentResource) exportNotificationTemplateContentsByTemplate(templateName management.EnumTemplateName) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadAllTemplateContents(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, templateName).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllTemplateContents", r.ResourceType())
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

		for _, templateContent := range embedded.GetContents() {
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
				continue
			}

			if templateContentIdOk && templateContentDeliveryMethodOk && templateContentLocaleOk {
				r.addImportBlock(string(templateName), *templateContentId, string(*templateContentDeliveryMethod), *templateContentLocale, templateContentVariant)
			}
		}
	}

	return nil
}

func (r *PingOneNotificationTemplateContentResource) addImportBlock(templateName, templateContentId, templateContentDeliveryMethod, templateContentLocale, templateContentVariant string) {
	commentData := map[string]string{
		"Resource Type":                    r.ResourceType(),
		"Template Name":                    templateName,
		"Template Content Delivery Method": templateContentDeliveryMethod,
		"Template Content Locale":          templateContentLocale,
		"Export Environment ID":            r.clientInfo.ExportEnvironmentID,
		"Template Content ID":              templateContentId,
	}

	if templateContentVariant != "" {
		commentData["Template Content Variant"] = templateContentVariant
		templateContentVariant = "_" + templateContentVariant
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s_%s%s", templateName, templateContentDeliveryMethod, templateContentLocale, templateContentVariant),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, templateName, templateContentId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
