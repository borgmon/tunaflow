package generator

import (
	"log"
	"reflect"
	"strings"

	"github.com/twpayne/go-jsonstruct"
	"gopkg.in/yaml.v2"
)

func (g *Generator) CapFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func (g *Generator) ParseFieldName(raw string) string {
	a := &jsonstruct.AbbreviationHandlingFieldNamer{}
	return a.FieldName(raw)
}
func (g *Generator) ParseJSONPath(s string) []string {
	return strings.Split(s, ".")
}

func (g *Generator) GetFieldType1(i interface{}, s string) string {
	b := reflect.ValueOf(i).Elem().FieldByName(g.ParseFieldName(s))
	return b.Type().Name()
}

func (g *Generator) GetFieldType(s string) string {
	m := make(map[string]string)
	y, err := yaml.Marshal(g.Config.Schemas[1].Payload)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(y, m)
	if err != nil {
		log.Fatalln(err)
	}
	return g.CapFirstLetter(m[s])
}

func (g *Generator) ReplaceInt2Int64(data []byte) []byte {
	d := string(data)
	d = strings.ReplaceAll(d, "int", "int64")
	return []byte(d)
}
