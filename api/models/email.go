package models

import "github.com/G-Villarinho/fast-feet-api/templates"

type SendEmailPayload struct {
	To           string
	Subject      string
	TemplateName templates.TemplateName
	Params       map[string]string
}
