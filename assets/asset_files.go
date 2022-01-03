package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _AssetFiles66179fd3f41ff75583c039352f246b25cbf76ae2 = "version: 1\nname: {{.AppName}}\napp-version: v1.0\npackage-path: {{.AppName}}\n\nschemas:\n  - name: input\n    payload:\n      stringField: string\n      boolField: boolean\n      nested:\n        value: int\n\n  - name: output\n    payload:\n      name: string\n      isFish: boolean\n      value: int\n\nflows:\n  - name: testflow\n    upstream: input\n    downstream: output\n    mapping:\n      name: stringField\n      isOpen: boolField\n      value: nested.value\n"
var _AssetFiles77d92646711ea780d561b03eca99f41bd8377c35 = "github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=\ngithub.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=\ngithub.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=\ngithub.com/gin-contrib/sse v0.1.0 h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=\ngithub.com/gin-contrib/sse v0.1.0/go.mod h1:RHrZQHXnP2xjPF+u1gW/2HnVO7nvIa9PG3Gm+fLHvGI=\ngithub.com/gin-gonic/gin v1.7.7 h1:3DoBmSbJbZAWqXJC3SLjAPfutPJJRN1U5pALB7EeTTs=\ngithub.com/gin-gonic/gin v1.7.7/go.mod h1:axIBovoeJpVj8S3BwE0uPMTeReE4+AfFtqpqaZ1qq1U=\ngithub.com/go-playground/assert/v2 v2.0.1 h1:MsBgLAaY856+nPRTKrp3/OZK38U/wa0CcBYNjji3q3A=\ngithub.com/go-playground/assert/v2 v2.0.1/go.mod h1:VDjEfimB/XKnb+ZQfWdccd7VUvScMdVu0Titje2rxJ4=\ngithub.com/go-playground/locales v0.13.0 h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=\ngithub.com/go-playground/locales v0.13.0/go.mod h1:taPMhCMXrRLJO55olJkUXHZBHCxTMfnGwq/HNwmWNS8=\ngithub.com/go-playground/universal-translator v0.17.0 h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=\ngithub.com/go-playground/universal-translator v0.17.0/go.mod h1:UkSxE5sNxxRwHyU+Scu5vgOQjsIJAF8j9muTVoKLVtA=\ngithub.com/go-playground/validator/v10 v10.4.1 h1:pH2c5ADXtd66mxoE0Zm9SUhxE20r7aM3F26W0hOn+GE=\ngithub.com/go-playground/validator/v10 v10.4.1/go.mod h1:nlOn6nFhuKACm19sB/8EGNn9GlaMV7XkbRSipzJ0Ii4=\ngithub.com/golang/protobuf v1.3.3 h1:gyjaxf+svBWX08ZjK86iN9geUJF0H6gp2IRKX6Nf6/I=\ngithub.com/golang/protobuf v1.3.3/go.mod h1:vzj43D7+SQXF/4pzW/hwtAqwc6iTitCiVSaWz5lYuqw=\ngithub.com/google/gofuzz v1.0.0/go.mod h1:dBl0BpW6vV/+mYPU4Po3pmUjxk6FQPldtuIdl/M65Eg=\ngithub.com/josharian/intern v1.0.0 h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=\ngithub.com/josharian/intern v1.0.0/go.mod h1:5DoeVV0s6jJacbCEi61lwdGj/aVlrQvzHFFd8Hwg//Y=\ngithub.com/json-iterator/go v1.1.9 h1:9yzud/Ht36ygwatGx56VwCZtlI/2AD15T1X2sjSuGns=\ngithub.com/json-iterator/go v1.1.9/go.mod h1:KdQUCv79m/52Kvf8AW2vK1V8akMuk1QjK/uOdHXbAo4=\ngithub.com/leodido/go-urn v1.2.0 h1:hpXL4XnriNwQ/ABnpepYM/1vCLWNDfUNts8dX3xTG6Y=\ngithub.com/leodido/go-urn v1.2.0/go.mod h1:+8+nEpDfqqsY+g338gtMEUOtuK+4dEMhiQEgxpxOKII=\ngithub.com/mailru/easyjson v0.7.7 h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=\ngithub.com/mailru/easyjson v0.7.7/go.mod h1:xzfreul335JAWq5oZzymOObrkdz5UnU4kGfJJLY9Nlc=\ngithub.com/mattn/go-isatty v0.0.12 h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=\ngithub.com/mattn/go-isatty v0.0.12/go.mod h1:cbi8OIDigv2wuxKPP5vlRcQ1OAZbq2CE4Kysco4FUpU=\ngithub.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 h1:ZqeYNhU3OHLH3mGKHDcjJRFFRrJa6eAM5H+CtDdOsPc=\ngithub.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=\ngithub.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 h1:Esafd1046DLDQ0W1YjYsBW+p8U2u7vzgW2SQVmlNazg=\ngithub.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742/go.mod h1:bx2lNnkwVCuqBIxFjflWJWanXIb3RllmbCylyMrvgv0=\ngithub.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=\ngithub.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=\ngithub.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=\ngithub.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=\ngithub.com/stretchr/testify v1.4.0 h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=\ngithub.com/stretchr/testify v1.4.0/go.mod h1:j7eGeouHqKxXV5pUuKE4zz7dFj8WfuZ+81PSLYec5m4=\ngithub.com/ugorji/go v1.1.7 h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=\ngithub.com/ugorji/go v1.1.7/go.mod h1:kZn38zHttfInRq0xu/PH0az30d+z6vm202qpg1oXVMw=\ngithub.com/ugorji/go/codec v1.1.7 h1:2SvQaVZ1ouYrrKKwoSk2pzd4A9evlKJb9oTL+OaLUSs=\ngithub.com/ugorji/go/codec v1.1.7/go.mod h1:Ax+UKWsSmolVDwsd+7N3ZtXu+yMGCf907BLYF3GoBXY=\ngolang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=\ngolang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 h1:psW17arqaxU48Z5kZ0CQnkZWQJsqcURM6tKiBApRjXI=\ngolang.org/x/crypto v0.0.0-20200622213623-75b288015ac9/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=\ngolang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=\ngolang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=\ngolang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=\ngolang.org/x/sys v0.0.0-20200116001909-b77594299b42 h1:vEOn+mP2zCOVzKckCZy6YsCtDblrpj/w7B9nxGNELpg=\ngolang.org/x/sys v0.0.0-20200116001909-b77594299b42/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=\ngolang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=\ngolang.org/x/text v0.3.2/go.mod h1:bEr9sfX3Q8Zfm5fL9x+3itogRgK3+ptLWKqgva+5dAk=\ngolang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=\ngopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=\ngopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=\ngopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=\ngopkg.in/yaml.v2 v2.2.8 h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=\ngopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=\n"
var _AssetFiles9bf78cc4c5207582e4f00d8dee38ef5bd775baae = "package main\n\nimport (\n\t\"example/schema\"\n\n\t\"github.com/mailru/easyjson\"\n)\nfunc {{.Name}}Mapping(in *schema.{{ParseFieldName .Upstream}}) (*schema.{{ParseFieldName .Downstream}}, error) {\n\tresult := &schema.{{ParseFieldName .Downstream}}{\n\t\t{{range $key, $val := .Mapping}}{{ParseFieldName $key}}: in.{{ParseFieldValue $val}},\n\t\t{{end}}\n\t}\n\treturn result, nil\n}\n\nfunc {{.Name}}Transform(data []byte) ([]byte, error) {\n\tinput := &schema.{{ParseFieldName .Upstream}}{}\n\tif err := easyjson.Unmarshal(data, input); err != nil {\n\t\tprintln(err)\n\t\treturn nil, err\n\t}\n\tds, err := {{.Name}}Mapping(input)\n\tif err != nil {\n\t\tprintln(err)\n\t\treturn nil, err\n\t}\n\tresult, err := easyjson.Marshal(ds)\n\tif err != nil {\n\t\tprintln(err)\n\t\treturn nil, err\n\t}\n\treturn result, nil\n}\n"
var _AssetFilesfcf621875c5f7ca89a3ea0157e685819ec5b7a45 = "package main\n\nimport (\n\t\"net/http\"\n\n\t\"github.com/gin-gonic/gin\"\n)\n\ntype Router struct {\n\trouter *gin.Engine\n}\n\nfunc InitRouter() *Router {\n\treturn &Router{gin.Default()}\n}\n\nfunc (r *Router) Start() {\n\t{{range $key, $val := .Flows}}\n\tr.router.POST(\"/{{$val.Name}}\", r.{{$val.Name}}Handler)\n\t{{end}}\n\tr.router.Run(\"localhost:8080\")\n}\n\n{{range $key, $val := .Flows}}\nfunc (r *Router) {{$val.Name}}Handler(c *gin.Context) {\n\tdata, err := c.GetRawData()\n\tif err != nil {\n\t\tc.JSON(http.StatusBadRequest, err)\n\t}\n\tresult, err := {{$val.Name}}Transform(data)\n\tif err != nil {\n\t\tc.JSON(http.StatusInternalServerError, err)\n\t}\n\tc.Data(http.StatusOK, \"application/json\", result)\n}\n{{end}}\n"
var _AssetFiles75fd09ca72e7afa573a17ae46afb2edda6310d4d = "package main\n\nfunc main() {\n\tr := InitRouter()\n\tr.Start()\n}\n"
var _AssetFilesa63462477b679f62a1b5996015cb9741bd4a126b = "module {{.PackagePath}}\n\ngo 1.17\n\nrequire (\n\tgithub.com/gin-gonic/gin v1.7.7\n\tgithub.com/mailru/easyjson v0.7.7\n)\n\nrequire (\n\tgithub.com/gin-contrib/sse v0.1.0 // indirect\n\tgithub.com/go-playground/locales v0.13.0 // indirect\n\tgithub.com/go-playground/universal-translator v0.17.0 // indirect\n\tgithub.com/go-playground/validator/v10 v10.4.1 // indirect\n\tgithub.com/golang/protobuf v1.3.3 // indirect\n\tgithub.com/josharian/intern v1.0.0 // indirect\n\tgithub.com/json-iterator/go v1.1.9 // indirect\n\tgithub.com/leodido/go-urn v1.2.0 // indirect\n\tgithub.com/mattn/go-isatty v0.0.12 // indirect\n\tgithub.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect\n\tgithub.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect\n\tgithub.com/ugorji/go/codec v1.1.7 // indirect\n\tgolang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect\n\tgolang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect\n\tgopkg.in/yaml.v2 v2.2.8 // indirect\n)\n"

// AssetFiles returns go-assets FileSystem
var AssetFiles = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"transform.go.tmpl", "router.go.tmpl", "main.go.tmpl", "go.mod.tmpl", "app.yaml.tmpl", "go.sum"}}, map[string]*assets.File{
	"/templates/go.mod.tmpl": &assets.File{
		Path:     "/templates/go.mod.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1641164600, 1641164600688427300),
		Data:     []byte(_AssetFilesa63462477b679f62a1b5996015cb9741bd4a126b),
	}, "/templates/app.yaml.tmpl": &assets.File{
		Path:     "/templates/app.yaml.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1641169044, 1641169044617000000),
		Data:     []byte(_AssetFiles66179fd3f41ff75583c039352f246b25cbf76ae2),
	}, "/templates/go.sum": &assets.File{
		Path:     "/templates/go.sum",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1641165314, 1641165314952000000),
		Data:     []byte(_AssetFiles77d92646711ea780d561b03eca99f41bd8377c35),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1641194132, 1641194132008427300),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1641169044, 1641169044608427300),
		Data:     nil,
	}, "/templates/transform.go.tmpl": &assets.File{
		Path:     "/templates/transform.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1641165082, 1641165082038427300),
		Data:     []byte(_AssetFiles9bf78cc4c5207582e4f00d8dee38ef5bd775baae),
	}, "/templates/router.go.tmpl": &assets.File{
		Path:     "/templates/router.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1641165065, 1641165065568427300),
		Data:     []byte(_AssetFilesfcf621875c5f7ca89a3ea0157e685819ec5b7a45),
	}, "/templates/main.go.tmpl": &assets.File{
		Path:     "/templates/main.go.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1641155433, 1641155433174000000),
		Data:     []byte(_AssetFiles75fd09ca72e7afa573a17ae46afb2edda6310d4d),
	}}, "")
