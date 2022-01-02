package generator

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/mailru/easyjson/bootstrap"
	"github.com/mailru/easyjson/parser"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Generator struct {
	Config   *AppConfig
	BasePath string
}

func (g *Generator) Prepare() error {
	if err := os.MkdirAll(g.BasePath+"/build/schema", os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	d, err := os.ReadFile("templates/go.mod.tmpl")
	if err != nil {
		return errors.WithStack(err)
	}
	s, err := g.GenerateTemplate(string(d), &template.FuncMap{}, map[string]string{
		"PackagePath": g.Config.PackagePath,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.WriteFile(g.BasePath+"/build/go.mod", s, os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	if err = g.CopyFile("template/go.sum", g.BasePath+"/build"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
func (g *Generator) CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return errors.WithStack(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		return errors.WithStack(err)
	}

	source, err := os.Open(src)
	if err != nil {
		return errors.WithStack(err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return errors.WithStack(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return errors.WithStack(err)
}
func (g *Generator) GenTemplate() error {
	inY, err := g.CreateYaml(g.Config.Schemas[0].Payload)
	if err != nil {
		return errors.WithStack(err)
	}
	reader2 := bytes.NewReader(inY)
	if g.CreateStructFromYaml(reader2, g.Config.Schemas[0].Name); err != nil {
		return errors.WithStack(err)
	}
	outY, err := g.CreateYaml(g.Config.Schemas[1].Payload)
	if err != nil {
		return errors.WithStack(err)
	}
	reader1 := bytes.NewReader(outY)
	if g.CreateStructFromYaml(reader1, g.Config.Schemas[1].Name); err != nil {
		return errors.WithStack(err)
	}

	// if err = g.GenEnDecoder(); err != nil {
	// 	return errors.WithStack(err)
	// }

	f, err := os.ReadFile("templates/mappingTemplate.go.tmpl")
	if err != nil {
		return errors.WithStack(err)
	}
	var fm = &template.FuncMap{
		"parseFieldName":    g.ParseFieldName,
		"getFieldType":      g.GetFieldType,
		"parseJSONPath":     g.ParseJSONPath,
		"getDownstreamName": g.GetDownstreamName,
	}
	m := make(map[string]string)
	flowMapping, err := yaml.Marshal(g.Config.Flows[0].Mapping)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = yaml.Unmarshal(flowMapping, &m); err != nil {
		return errors.WithStack(err)
	}
	t, err := g.GenerateTemplate(string(f), fm, m)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.WriteFile(g.BasePath+"/build/main.go", t, os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (g *Generator) GenEnDecoder() error {
	p := parser.Parser{AllStructs: true}
	pwd, err := os.Getwd()
	if err != nil {
		return errors.WithStack(err)
	}
	fullPath := filepath.Join(pwd, g.BasePath+"/build/schema")

	if err := p.Parse(fullPath, true); err != nil {
		return errors.WithStack(err)
	}
	jsonGenerator := &bootstrap.Generator{
		PkgPath: p.PkgPath,
		PkgName: p.PkgName,
		Types:   p.StructNames,
		OutName: fullPath + "/coder.go",
	}

	if err := jsonGenerator.Run(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
