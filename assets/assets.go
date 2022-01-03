package assets

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type Assets struct {
	FileName string
	OutPath  string
	FuncMap  *template.FuncMap
	Data     interface{}
}

const TemplatePath = "/templates"

func JoinTemplatePath(fileName string) string {
	return filepath.Join(TemplatePath, fileName)
}

func (t *Assets) Generate() error {
	file, err := AssetFiles.Open(filepath.Join(TemplatePath, t.FileName+".tmpl"))
	if err != nil {
		return errors.WithStack(err)
	}
	inFile, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.WithStack(err)
	}
	outFile, err := t.GenerateTemplate(inFile)
	if err != nil {
		return errors.WithStack(err)
	}
	if err = os.WriteFile(filepath.Join(t.OutPath, t.FileName), outFile, os.ModePerm); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (t *Assets) GenerateTemplate(templData []byte) ([]byte, error) {
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
