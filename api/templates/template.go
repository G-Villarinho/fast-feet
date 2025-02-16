package templates

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/G-Villarinho/fast-feet-api/di"
)

type TemplateName string

const (
	PickUpTemplate TemplateName = "pick-up-template"
)

//go:generate mockery --name=TemplateService --output=../mocks --outpkg=mocks
type Template interface {
	RenderTemplate(templateName TemplateName, params map[string]string) (string, error)
}

type templateService struct {
	i    *di.Injector
	path string
}

func NewTemplate(i *di.Injector) (Template, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return &templateService{
		i:    i,
		path: filepath.Join(dir, "templates"),
	}, nil
}

func (t *templateService) RenderTemplate(templateName TemplateName, params map[string]string) (string, error) {
	content, err := os.ReadFile(filepath.Join(t.path, string(templateName)+".html"))
	if err != nil {
		return "", errors.New("read email template: " + err.Error())
	}

	template := string(content)
	for key, value := range params {
		placeholder := fmt.Sprintf("#%s#", key)
		template = strings.ReplaceAll(template, placeholder, value)
	}

	return template, nil
}
