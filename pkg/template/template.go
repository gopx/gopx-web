package template

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"gopx.io/gopx-web/pkg/config"
	"gopx.io/gopx-web/pkg/log"
)

var pageTemplate = template.New("page-template")

func init() {
	p, err := templateGlobPattern()
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	_, err = pageTemplate.ParseGlob(p)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}

func templateGlobPattern() (s string, err error) {
	s, err = filepath.Abs(config.Web.PageTemplatePath)
	if err != nil {
		return
	}
	s = filepath.Join(s, "*"+config.Web.PageTemplateExtension)
	return
}

// Render writes HTML representation of the template to a Writer with specified data.
func Render(name string, w io.Writer, data interface{}) error {
	if !strings.HasSuffix(name, config.Web.PageTemplateExtension) {
		name += config.Web.PageTemplateExtension
	}
	return pageTemplate.ExecuteTemplate(w, name, data)
}
