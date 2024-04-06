package swaggeneration

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

type property struct {
	DataType   string              `yaml:"type,omitempty"`
	Format     string              `yaml:"format,omitempty"`
	Example    string              `yaml:"example,omitempty"`
	Properties map[string]property `yaml:"properties,omitempty"`
}

func generateBodySchemaFromAllExample(example map[interface{}]interface{}) string {
	uniqueKeyBody := map[string]property{}
	for _, e := range example {
		entry := e.(map[interface{}]interface{})
		bdy := entry["requestBody"].(string)

		var jsonBody map[string]interface{}
		err := json.Unmarshal([]byte(bdy), &jsonBody)
		if err != nil {
			panic(err)
		}

		uniqueKeyBody = parseObject(jsonBody, uniqueKeyBody)
	}

	yamlData, err := yaml.Marshal(uniqueKeyBody)
	if err != nil {
		panic(err)
	}
	return string(yamlData)
}

func generateResponseBodySchemaFromExample(example interface{}) string {
	uniqueKeyBody := map[string]property{}
	entry := example.(map[interface{}]interface{})
	bdy := entry["responseBody"].(string)

	var jsonBody map[string]interface{}
	err := json.Unmarshal([]byte(bdy), &jsonBody)
	if err != nil {
		panic(err)
	}

	uniqueKeyBody = parseObject(jsonBody, uniqueKeyBody)

	yamlData, err := yaml.Marshal(uniqueKeyBody)
	if err != nil {
		panic(err)
	}
	return string(yamlData)
}

func parseObject(jsonBody map[string]interface{}, parsedData map[string]property) map[string]property {
	for key, b := range jsonBody {
		var t property
		switch val := b.(type) {
		case int:
			t = property{
				DataType: "integer",
				Example:  strconv.Itoa(val),
			}
		case float64:
			t = property{
				DataType: "number",
				Format:   "float",
				Example:  fmt.Sprintf("%f", val),
			}
		case string:
			t = property{
				DataType: "string",
				Example:  val,
			}
		case bool:
			t = property{
				DataType: "boolean",
				Example:  strconv.FormatBool(val),
			}
		default:
			childJson := val.(map[string]interface{})
			parsedChild := parseObject(childJson, map[string]property{})
			t = property{
				DataType:   "object",
				Properties: parsedChild,
			}
		}

		parsedData[key] = t
	}

	return parsedData
}
