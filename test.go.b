package main

import (
	"html/template"
	"os"
	"strings"

	"github.com/borgmon/tunaflow/mapping"
	"github.com/twpayne/go-jsonstruct"
	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// createStructFromYaml("example/testflow.yaml", "mapping/MappingConfig.go", "MappingConfig", "mapping")
	createStructFromYaml("example/input.yaml", "example/.tmp/schemas/Upstream.go", "Upstream", "schemas")
	createStructFromYaml("example/output.yaml", "example/.tmp/schemas/Downstream.go", "Downstream", "schemas")
	templateStuff()
}

func createStructFromYaml(filepath string, out string, typename string, packagename string) {
	reader, err := os.Open(filepath)
	check(err)
	observedValue, err := jsonstruct.ObserveYAML(reader)
	check(err)
	options := []jsonstruct.GeneratorOption{
		jsonstruct.WithTypeName(typename),
		jsonstruct.WithPackageName(packagename),
		jsonstruct.WithOmitEmpty(jsonstruct.OmitEmptyAuto),
		jsonstruct.WithSkipUnparseableProperties(false),
		jsonstruct.WithUseJSONNumber(false),
		jsonstruct.WithGoFormat(true),
	}
	goCode, err := jsonstruct.NewGenerator(options...).GoCode(observedValue)
	check(err)
	err = os.WriteFile(out, goCode, 0644)
	check(err)
}

func parseType(raw string) string {
	a := &jsonstruct.AbbreviationHandlingFieldNamer{}
	return a.FieldName(raw)
}

func parseMapping(raw string) string {
	a := &jsonstruct.AbbreviationHandlingFieldNamer{}
	return a.FieldName(raw)
}

func templateStuff() {
	var fm = &template.FuncMap{
		"parseMapping": parseMapping,
	}
	d, err := os.ReadFile("mapping/mappingTemplate.go.tmpl")
	check(err)

	yamld, err := os.ReadFile("example/testflow.yaml")
	check(err)
	mappingConfig := &mapping.MappingConfig{}
	yaml.Unmarshal(yamld, &mappingConfig)

	generateTemplate(string(d), fm, mappingConfig.Mapping)
}
func generateTemplate(s string, fm *template.FuncMap, data map[string]string) string {

	t, err := template.New("").Funcs(*fm).Parse(s)
	check(err)
	b := new(strings.Builder)
	err = t.Execute(b, &data)
	check(err)
	println(b.String())
	return b.String()
}
