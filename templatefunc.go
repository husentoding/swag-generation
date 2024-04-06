package swaggeneration

import (
	"html/template"
	"strings"
)

func templateFunc() template.FuncMap {
	return template.FuncMap{
		"bodyFromExample":   generateBodySchemaFromAllExample,
		"headerFromExample": generateHeaderSchemaFromAllExample,
		"responseBody":      generateResponseBodySchemaFromExample,
		"indent":            indent,
	}
}

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}
