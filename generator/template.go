package generator

import (
	"html/template"
	"strings"

	"github.com/pkg/errors"
)

type Template struct {
	InPath  string
	OutPath string
	FuncMap *template.FuncMap
	Data    interface{}
}

const templatePath = "./templates"

func (t *Template) GenerateTemplate(templData []byte) ([]byte, error) {
	templ := template.New("")

	if t.FuncMap != nil {
		templ = templ.Funcs(*t.FuncMap)
	}
	templ, err := templ.Parse(string(templData))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	b := new(strings.Builder)
	err = templ.Execute(b, t.Data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return []byte(b.String()), nil
}
