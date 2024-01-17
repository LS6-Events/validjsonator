# ValidJSONator

ValidJSONator is a basic Go library that takes a [Go Validator V10](https://github.com/go-playground/validator) validation string and converts it into a Schema that can be used in JSONSchema or OpenAPI.

This library is used to share this logic between our other libraries [Astra](https://github.com/LS6-Events/astra) and [Confuse](https://github.com/LS6-Events/confuse), and hopefully will be useful to others.


## Installation
```bash
go get github.com/ls6-events/validjsonator
```

## Usage

```go
package main

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/ls6-events/validjsonator"
)

type Example struct {
	ExampleString string  `validate:"required,base64"`
	ExampleInt    int     `validate:"required,min=1,max=10"`
	ExampleFloat  float64 `validate:"required,min=1,max=10"`
}

func main() {
	Example := Example{}
	exampleType := reflect.TypeOf(Example)

	structFieldsSchema := make(map[string]validjsonator.Schema)
	for i := 0; i < exampleType.NumField(); i++ {
		fieldTag := exampleType.Field(i).Tag
		fieldName := exampleType.Field(i).Name

		validationTag := fieldTag.Get("validate")

		validationSchema, _ := validjsonator.ValidationTagsToSchema(validationTag)

		structFieldsSchema[fieldName] = validationSchema
	}

	spew.Dump(structFieldsSchema)
}

```
