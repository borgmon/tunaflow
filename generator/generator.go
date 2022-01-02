package generator

import (
	"bytes"
	"html/template"
	"os"

	"gopkg.in/yaml.v2"
)

type Generator struct {
	Config   *AppConfig
	BasePath string
}

func (g *Generator) Prepare() error {
	err := os.MkdirAll(g.BasePath+"/.tmp", os.ModePerm)
	if err != nil {
		return err
	}
	d, err := os.ReadFile("templates/go.mod.tmpl")
	if err != nil {
		return err
	}
	s, err := g.GenerateTemplate(string(d), &template.FuncMap{}, map[string]string{
		"PackagePath": g.Config.PackagePath,
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(g.BasePath+"/.tmp/go.mod", s, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) GenTemplate() error {
	outY, err := g.CreateYaml(g.Config.Schemas[1].Payload)
	if err != nil {
		return err
	}
	reader1 := bytes.NewReader(outY)
	g.CreateStructFromYaml(reader1, "downstream")

	inY, err := g.CreateYaml(g.Config.Schemas[0].Payload)
	if err != nil {
		return err
	}
	reader2 := bytes.NewReader(inY)
	g.CreateStructFromYaml(reader2, "upstream")

	f, err := os.ReadFile("templates/mappingTemplate.go.tmpl")
	if err != nil {
		return err
	}
	var fm = &template.FuncMap{
		"parseFieldName": g.ParseFieldName,
		"getFieldType":   g.GetFieldType,
		"parseJSONPath":  g.ParseJSONPath,
	}
	m := make(map[string]string)
	aaa, err := yaml.Marshal(g.Config.Flows[0].Mapping)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(aaa, &m)
	if err != nil {
		return err
	}
	t, err := g.GenerateTemplate(string(f), fm, m)
	if err != nil {
		return err
	}
	err = os.WriteFile(g.BasePath+"/.tmp/main.go", t, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
