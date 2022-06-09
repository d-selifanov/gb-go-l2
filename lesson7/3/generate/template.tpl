package {{.Package}}

func (s *{{.TypeName}}) FromMap(m map[string]interface{}) {
    {{range .Fields}}s.{{.Name}} = m["{{.JsonName}}"].({{.Type}})
    {{end}}
}