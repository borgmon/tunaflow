package main

import (
	"encoding/json"

	"github.com/buger/jsonparser"
)

type Data []byte
 
func (data Data) mapping() (downstream Downstream, err error) {
	{{range $key, $val := .}}
	downstream.{{parseFieldName $key}}, err = jsonparser.Get{{getFieldType $key}}(data, []string{
		{{range $seg := parseJSONPath $val}}"{{$seg}}",{{end}}
	}...)
	
	{{end}}

	return
}

func main(){
	test := `{"msg":"hello","enabled":true,"nested":{"value":123}}`
	data := Data([]byte(test))
	ds, err := data.mapping()
	if err != nil {
		println(err)
	}
	data, err = json.Marshal(ds)
	if err != nil {
		println(err)
	}
	println(string(data))
}