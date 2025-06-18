package service

import (
	"bytes"
	"text/template"

	"scaffold/domain/model"
)

type TemplateService interface {
	Render(templateText string, project model.Project) (string, error)
}

type templateServiceImpl struct{}

func NewTemplateService() TemplateService {
	return &templateServiceImpl{}
}

// Render テンプレート文字列にプロジェクト情報を埋め込む
func (s *templateServiceImpl) Render(templateText string, project model.Project) (string, error) {
	tmpl, err := template.New("file").Parse(templateText)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, project); err != nil {
		return "", err
	}

	return buf.String(), nil
}
