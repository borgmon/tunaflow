package generator

import (
	"io"
	"os"
	"strings"

	"github.com/twpayne/go-jsonstruct"
	"gopkg.in/yaml.v2"
)

func (g *Generator) CreateYaml(i interface{}) ([]byte, error) {
	outY, err := yaml.Marshal(i)
	if err != nil {
		return nil, err
	}
	return g.ReplaceType(outY), nil
}

func (g *Generator) ReplaceType(data []byte) []byte {
	s := string(data)
	s = strings.ReplaceAll(s, "int", "111111")
	s = strings.ReplaceAll(s, "boolean", "true")
	return []byte(s)
}
func (g *Generator) CreateStructFromYaml(reader io.Reader, typename string) ([]byte, error) {
	observedValue, err := jsonstruct.ObserveYAML(reader)
	if err != nil {
		return nil, err
	}
	options := []jsonstruct.GeneratorOption{
		jsonstruct.WithTypeName(g.CapFirstLetter(typename)),
		jsonstruct.WithPackageName("main"),
		jsonstruct.WithOmitEmpty(jsonstruct.OmitEmptyAuto),
		jsonstruct.WithSkipUnparseableProperties(false),
		jsonstruct.WithUseJSONNumber(false),
		jsonstruct.WithGoFormat(true),
	}
	goCode, err := jsonstruct.NewGenerator(options...).GoCode(observedValue)
	if err != nil {
		return nil, err
	}
	goCode = g.ReplaceInt2Int64(goCode)
	err = os.WriteFile(g.BasePath+"/.tmp/"+typename+".go", goCode, 0644)
	if err != nil {
		return nil, err
	}
	return goCode, nil
}
