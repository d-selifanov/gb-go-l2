package main

import (
	"embed"
	"log"
	"os"
	"text/template"
)

type TplData struct {
	Package  string
	TypeName string
	Fields   []TplField
}

type TplField struct {
	Name     string
	JsonName string
	Type     string
}

type MyStruct struct {
	MyField string `json:"myJsonField"`
}

func (s *MyStruct) FromMap(m map[string]interface{}) {
	s.MyField = m["myJsonField"].(string)
}

//go:embed template.tpl
var f embed.FS

func main() {
	tpl, _ := f.ReadFile("template.tpl")

	templateData := TplData{
		Package:  "main",
		TypeName: "Product",
		Fields: []TplField{
			{
				Name:     "Code",
				JsonName: "code",
				Type:     "uuid.UUID",
			},
			{
				Name:     "Name",
				JsonName: "name",
				Type:     "string",
			},
			{
				Name:     "Price",
				JsonName: "price",
				Type:     "float64",
			},
			{
				Name:     "Count",
				JsonName: "count",
				Type:     "int64",
			},
		},
	}

	genFile, err := os.Create("generated_file.go")
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("generated_file").Parse(string(tpl)))
	if err := t.Execute(genFile, templateData); err != nil {
		log.Fatal(err)
	}
}
