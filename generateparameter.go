package swaggeneration

import (
	"gopkg.in/yaml.v2"
)

type parameter struct {
	InType     string    `yaml:"in,omitempty"`
	Name       string    `yaml:"name,omitempty"`
	Schema     *property `yaml:"schema,omitempty"`
	IsRequired bool      `yaml:"required,omitempty"`
}

func generateHeaderSchemaFromAllExample(example map[interface{}]interface{}) string {
	uniqueKeyBody := map[string]parameter{}
	for _, e := range example {
		entry := e.(map[interface{}]interface{})
		headerList := entry["requestHeader"].(map[string]string)

		for key, i := range headerList {
			uniqueKeyBody[key] = parameter{
				InType: "header",
				Name:   key,
				Schema: &property{
					DataType: "string",
					Example:  i,
				},
				IsRequired: false, // TODO
			}
		}
	}

	yamlData, err := yaml.Marshal(uniqueKeyBody)
	if err != nil {
		panic(err)
	}
	return string(yamlData)
}
