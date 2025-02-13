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
	jsonStr := `{"name": "John Doe","age": 25}`

	// Load schema and JSON data as loaders
	schemaLoader := gojsonschema.NewStringLoader(schemaStr)
	jsonLoader := gojsonschema.NewStringLoader(jsonStr)

	// Validate JSON against schema
	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		fmt.Println(err)
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

	// check for regex expression
	schema2 := `{
		"type": "string",
		"pattern": "^(?:[01]\\d|2[0-3]):[0-5]\\d$"
	}`

	time := "11:11"

	// Load schema and JSON data as loaders
	s := gojsonschema.NewStringLoader(schema2)
	// for string we need to have string of for `"value"`
	// for schema to understand that this is of type string
	j := gojsonschema.NewStringLoader(fmt.Sprintf(`"%s"`, time)) 

	// Validate JSON against schema
	r, err := gojsonschema.Validate(s, j)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)

	// Check validation result
	if r.Valid() {
		fmt.Println("JSON is valid!")
	} else {
		fmt.Println("JSON is invalid. Errors:")
		for _, desc := range r.Errors() {
			fmt.Println("-", desc)
		}
	}

	// solution 1
	// - create different attribute type and validate it accordingly
	// solution 2
	// - using type decide whether to use string or json validator
	// solution 3
	// - use only json to store data , remove strings

}
