package main

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/twpayne/go-jsonstruct"
	"gopkg.in/yaml.v3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	basePath = "example"
	config   = &AppConfig{}
)

func main() {
	configD, err := os.ReadFile(basePath + "/app.yaml")
	check(err)
	err = yaml.Unmarshal(configD, config)
	check(err)

	prepare(basePath, config)
	genTemplate(config)
}

func prepare(path string, config *AppConfig) {
	err := os.MkdirAll(path+"/.tmp", os.ModePerm)
	check(err)
	d, err := os.ReadFile("templates/go.mod.tmpl")
	check(err)
	s := generateTemplate(string(d), &template.FuncMap{}, map[string]string{
		"PackagePath": config.PackagePath,
	})
	err = os.WriteFile(path+"/.tmp/go.mod", s, os.ModePerm)
	check(err)
}

func genTemplate(config *AppConfig) {
	outY := createYaml(config.Schemas[1].Payload)
	reader1 := bytes.NewReader(outY)
	createStructFromYaml(reader1, "downstream")

	inY := createYaml(config.Schemas[0].Payload)
	reader2 := bytes.NewReader(inY)
	createStructFromYaml(reader2, "upstream")

	f, err := os.ReadFile("templates/mappingTemplate.go")
	check(err)
	var fm = &template.FuncMap{
		"parseFieldName": parseFieldName,
		"getFieldType":   getFieldType,
		"parseJSONPath":  parseJSONPath,
	}
	m := make(map[string]string)
	aaa, err := yaml.Marshal(config.Flows[0].Mapping)
	err = yaml.Unmarshal(aaa, &m)

	t := generateTemplate(string(f), fm, m)
	err = os.WriteFile(basePath+"/.tmp/main.go", t, os.ModePerm)
	check(err)
}

func createYaml(i interface{}) []byte {
	outY, err := yaml.Marshal(i)
	check(err)
	return replaceType(outY)
}

func generateTemplate(s string, fm *template.FuncMap, data map[string]string) []byte {
	t, err := template.New("").Funcs(*fm).Parse(s)
	check(err)
	b := new(strings.Builder)
	err = t.Execute(b, &data)
	check(err)
	println(b.String())
	return []byte(b.String())
}

func replaceType(data []byte) []byte {
	s := string(data)
	s = strings.ReplaceAll(s, "int", "111111")
	s = strings.ReplaceAll(s, "boolean", "true")
	return []byte(s)
}

func createStructFromYaml(reader io.Reader, typename string) []byte {
	observedValue, err := jsonstruct.ObserveYAML(reader)
	check(err)
	options := []jsonstruct.GeneratorOption{
		jsonstruct.WithTypeName(capFirstLetter(typename)),
		jsonstruct.WithPackageName("main"),
		jsonstruct.WithOmitEmpty(jsonstruct.OmitEmptyAuto),
		jsonstruct.WithSkipUnparseableProperties(false),
		jsonstruct.WithUseJSONNumber(false),
		jsonstruct.WithGoFormat(true),
	}
	goCode, err := jsonstruct.NewGenerator(options...).GoCode(observedValue)
	check(err)
	goCode = replaceInt2Int64(goCode)
	err = os.WriteFile(basePath+"/.tmp/"+typename+".go", goCode, 0644)
	check(err)
	return goCode
}

func replaceInt2Int64(data []byte) []byte {
	d := string(data)
	d = strings.ReplaceAll(d, "int", "int64")
	return []byte(d)
}

func capFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func parseFieldName(raw string) string {
	a := &jsonstruct.AbbreviationHandlingFieldNamer{}
	return a.FieldName(raw)
}
func parseJSONPath(s string) []string {
	return strings.Split(s, ".")
}

func getFieldType1(i interface{}, s string) string {
	b := reflect.ValueOf(i).Elem().FieldByName(parseFieldName(s))
	return b.Type().Name()
}

func getFieldType(s string) string {
	m := make(map[string]string)
	y, err := yaml.Marshal(config.Schemas[1].Payload)
	check(err)
	err = yaml.Unmarshal(y, m)
	return capFirstLetter(m[s])

}
