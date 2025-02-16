package email

import (
	"context"
	"log/slog"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/templates"
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendEmail(ctx context.Context, payload models.SendEmailPayload)
}

type emailService struct {
	i *di.Injector
	t templates.Template
}

func NewEmailService(i *di.Injector) (EmailService, error) {
	t, err := di.Invoke[templates.Template](i)
	if err != nil {
		return nil, err
	}

	return &emailService{
		i: i,
		t: t,
	}, nil
}

func (e *emailService) SendEmail(ctx context.Context, payload models.SendEmailPayload) {
	log := slog.With(
		slog.String("service", "email"),
		slog.String("func", "SendEmail"),
	)

	content, err := e.t.RenderTemplate(payload.TemplateName, payload.Params)
	if err != nil {
		log.Error("render %s.html email template: %w", string(payload.TemplateName), err)
		return
	}

	dialer := gomail.NewDialer(config.Env.SMTP.Host, config.Env.SMTP.Port, config.Env.SMTP.User, config.Env.SMTP.Password)

	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Env.SMTP.User)
	msg.SetHeader("To", payload.To)
	msg.SetHeader("Subject", payload.Subject)
	msg.SetBody("text/html", content)

	if err := dialer.DialAndSend(msg); err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("email sent succefully")
}
