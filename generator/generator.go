package generator

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/borgmon/tunaflow/assets"
	"github.com/borgmon/tunaflow/funcmap"
	"github.com/borgmon/tunaflow/schema"
	"github.com/pkg/errors"
)

type Generator struct {
	Config *AppConfig
}

func GetPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return "./"
	}
	return pwd
}

func (g *Generator) generateTemplate(fileName string) error {
	t := &assets.Assets{
		FileName: fileName,
		OutPath:  GetPwd(),
		Data:     g.Config,
	}
	return t.Generate()
}

func (g *Generator) GenerateBase() error {
	if err := os.MkdirAll(filepath.Join(GetPwd(), schema.SchemaPath), os.ModePerm); err != nil {
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
	if err := CopyFile(assets.JoinTemplatePath("go.sum"), filepath.Join(GetPwd()+"/go.sum")); err != nil {
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
	if err = os.WriteFile(filepath.Join(GetPwd(), "schema", schemaConfig.Name+".go"), structData, os.ModePerm); err != nil {
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

	if err := schema.GenEnDecoder(GetPwd()); err != nil {
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
	t := &assets.Assets{
		FileName: filepath.Join(assets.TemplatePath, "transform.go.tmpl"),
		OutPath:  GetPwd(),
		Data:     flow,
		FuncMap:  fm,
	}
	file, err := assets.AssetFiles.Open(t.FileName)
	if err != nil {
		return errors.WithStack(err)
	}
	inFile, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.WithStack(err)
	}
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
