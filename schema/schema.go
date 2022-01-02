package schema

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/mailru/easyjson/bootstrap"
	"github.com/mailru/easyjson/parser"
	"github.com/pkg/errors"
	"github.com/twpayne/go-jsonstruct"
	"gopkg.in/yaml.v2"
)

const schemaPath = "schema"

func CreateYamlFromObj(i interface{}) ([]byte, error) {
	outY, err := yaml.Marshal(i)
	if err != nil {
		return nil, err
	}
	return replaceType(outY), nil
}

func CreateStructFromYaml(yamlData []byte, typename string) ([]byte, error) {
	reader := bytes.NewReader(yamlData)
	observedValue, err := jsonstruct.ObserveYAML(reader)
	if err != nil {
		return nil, err
	}
	options := []jsonstruct.GeneratorOption{
		jsonstruct.WithTypeName(capFirstLetter(typename)),
		jsonstruct.WithPackageName(schemaPath),
		jsonstruct.WithOmitEmpty(jsonstruct.OmitEmptyAuto),
		jsonstruct.WithSkipUnparseableProperties(false),
		jsonstruct.WithUseJSONNumber(false),
		jsonstruct.WithGoFormat(true),
	}
	goCode, err := jsonstruct.NewGenerator(options...).GoCode(observedValue)
	if err != nil {
		return nil, err
	}
	goCode = replaceInt2Int64(goCode)
	return goCode, nil
}

//FIXME
func replaceType(data []byte) []byte {
	s := string(data)
	s = strings.ReplaceAll(s, "int", "111111")
	s = strings.ReplaceAll(s, "boolean", "true")
	return []byte(s)
}

//FIXME
func replaceInt2Int64(data []byte) []byte {
	d := string(data)
	d = strings.ReplaceAll(d, "int", "int64")
	return []byte(d)
}

func capFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func GenEnDecoder(buildPath string) error {
	p := parser.Parser{AllStructs: true}
	pwd, err := os.Getwd()
	if err != nil {
		return errors.WithStack(err)
	}
	fullPath := filepath.Join(pwd, filepath.Join(buildPath, schemaPath))

	if err := p.Parse(fullPath, true); err != nil {
		return errors.WithStack(err)
	}
	jsonGenerator := &bootstrap.Generator{
		PkgPath: p.PkgPath,
		PkgName: p.PkgName,
		Types:   p.StructNames,
		OutName: fullPath + "/easyjson.go",
	}

	if err := jsonGenerator.Run(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
