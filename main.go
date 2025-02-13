package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// JSON Schema as a string
	schemaStr := `{
		"type": "object",
		"properties": {
			"name": { "type": "string" },
			"age": { "type": "integer", "minimum": 18 }
		},
		"required": ["name", "age"]
	}`

	// JSON Data as a string
	jsonStr := `{
		"name": "John Doe",
		"age": 25
	}`

	// Load schema and JSON data as loaders
	schemaLoader := gojsonschema.NewStringLoader(schemaStr)
	jsonLoader := gojsonschema.NewStringLoader(jsonStr)

	// Validate JSON against schema
	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		return 
	}

	fmt.Println(result)

	// Check validation result
	if result.Valid() {
		fmt.Println("JSON is valid!")
	} else {
		fmt.Println("JSON is invalid. Errors:")
		for _, desc := range result.Errors() {
			fmt.Println("-", desc)
		}
	}
}
