package generator

import (
	"html/template"
	"strings"
)

func (g *Generator) GenerateTemplate(s string, fm *template.FuncMap, data map[string]string) ([]byte, error) {
	t, err := template.New("").Funcs(*fm).Parse(s)
	if err != nil {
		return nil, err
	}
	b := new(strings.Builder)
	err = t.Execute(b, &data)
	if err != nil {
		return nil, err
	}
	return []byte(b.String()), nil
}
