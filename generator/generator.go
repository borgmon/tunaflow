package generator

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/borgmon/tunaflow/funcmap"
	"github.com/borgmon/tunaflow/schema"
	"github.com/pkg/errors"
)

type Generator struct {
	Config   *AppConfig
	BasePath string
}

const BuildPath = "build"

func (g *Generator) GetBuildPath() string { return filepath.Join(g.BasePath, BuildPath) }

func (g *Generator) generateTemplate(fileName string) error {
	t := &Template{
		InPath:  filepath.Join(templatePath, fileName+".tmpl"),
		OutPath: g.GetBuildPath(),
		Data:    g.Config,
	}
	inFile, err := os.ReadFile(t.InPath)
	if err != nil {
		return errors.WithStack(err)
	}
	outFile, err := t.GenerateTemplate(inFile)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.WriteFile(filepath.Join(t.OutPath, fileName), outFile, os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (g *Generator) GenerateBase() error {
	if err := os.MkdirAll(g.BasePath+"/build/schema", os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	if err := g.generateTemplate("go.mod"); err != nil {
		return errors.WithStack(err)
	}
	if err := g.generateTemplate("router.go"); err != nil {
		return errors.WithStack(err)
	}
	if err := g.generateTemplate("main.go"); err != nil {
		return errors.WithStack(err)
	}
	if err := CopyFile("./templates/go.sum", g.BasePath+"/build/go.sum"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g *Generator) generateSchema(schemaConfig *Schema) error {
	yamlData, err := schema.CreateYamlFromObj(schemaConfig.Payload)
	if err != nil {
		return errors.WithStack(err)
	}

	structData, err := schema.CreateStructFromYaml(yamlData, schemaConfig.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.WriteFile(filepath.Join(g.GetBuildPath(), "schema", schemaConfig.Name+".go"), structData, os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (g *Generator) GenerateSchema() error {
	for _, s := range g.Config.Schemas {
		if err := g.generateSchema(s); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := schema.GenEnDecoder(g.GetBuildPath()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (g *Generator) GenerateTransformer() error {
	for _, flow := range g.Config.Flows {
		if err := g.generateTransformer(flow); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (g *Generator) generateTransformer(flow *Flow) error {
	fm := &template.FuncMap{
		"ParseFieldName":  funcmap.ParseFieldName,
		"ParseFieldValue": funcmap.ParseFieldValue,
	}
	t := &Template{
		InPath:  filepath.Join(templatePath, "transform.go.tmpl"),
		OutPath: g.GetBuildPath(),
		Data:    flow,
		FuncMap: fm,
	}
	inFile, err := os.ReadFile(t.InPath)
	if err != nil {
		return errors.WithStack(err)
	}
	outFile, err := t.GenerateTemplate(inFile)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.WriteFile(filepath.Join(t.OutPath, flow.Name+".go"), outFile, os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
