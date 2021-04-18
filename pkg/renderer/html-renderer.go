package renderer

import (
	"bytes"
	"github.com/chrootlogin/stato/pkg/models"
	"html/template"
	"strings"
)

type HtmlRenderer struct {}

func init() {
	appendRenderer(&HtmlRenderer{})
}

func (r *HtmlRenderer) IsSupported(path string) bool {
	if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".htm") {
		return true
	}

	return false
}

func (r *HtmlRenderer) Render(data *models.ViewData) error {
	// parse template
	tmpl, err := template.New("tpl").Parse(data.Content);
	if err != nil {
		return ErrParsingTemplateFile
	}

	// create new buffer
	buf := new(bytes.Buffer)

	// execute template
	if err := tmpl.Execute(buf, data); err != nil {
		return ErrRenderingTemplateFile
	}

	// put data in buf
	data.Content = buf.String()

	return nil
}

