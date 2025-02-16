package email

import (
	"strconv"
	"time"

	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/templates"
)

type EmailFactory struct{}

func NewEmailFactory() *EmailFactory {
	return &EmailFactory{}
}

func (f *EmailFactory) CreatePickUpSendEmail(to, subject, recipientName, trackingCode string) models.SendEmailPayload {
	return models.SendEmailPayload{
		To:           to,
		Subject:      subject,
		TemplateName: templates.PickUpTemplate,
		Params: map[string]string{
			"recipient_name": recipientName,
			"tracking_code":  trackingCode,
			"current_year":   strconv.Itoa(time.Now().Year()),
		},
	}
}
