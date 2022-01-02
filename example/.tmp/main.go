package main

import (
	"encoding/json"

	"github.com/buger/jsonparser"
)

type Data []byte
 
func (data Data) mapping() (downstream Downstream, err error) {
	
	downstream.Name, err = jsonparser.GetString(data, []string{
		"msg",
	}...)
	
	
	downstream.Time, err = jsonparser.GetString(data, []string{
		"msg",
	}...)
	
	
	downstream.Value, err = jsonparser.GetInt(data, []string{
		"nested","value",
	}...)
	
	

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