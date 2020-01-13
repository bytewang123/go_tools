package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)

var data = []byte(`
		{
		"person": {
		"name": {
		"first": "Leonid",
		"last": "Bugaev",
		"fullName": "Leonid Bugaev"
		},
		"github": {
		"handle": "buger",
		"followers": 109
		},
		"avatars": [
		{"index":100},
		{ "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
		{ "url1": "https://avatars2.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
		]
		},
		"company": {
		"name": "Acme"
		}
		}
		`)

func main() {
	r := rule{
		FieldsOrder: []string{"person", "name", "first"},
	}
	v, dataType, err := parseFields(data, &r)
	fmt.Printf("value:%+v\tdataType:%+v\terr:%+v\n", string(v), dataType, err)

	r = rule{
		FieldsOrder: []string{"person", "avatars", "[0]", "index"},
	}
	v, dataType, err = parseFields(data, &r)
	fmt.Printf("value:%+v\tdataType:%+v\terr:%+v\n", string(v), dataType, err)
}

type rule struct {
	FieldsOrder []string
}

func parseFields(data []byte, r *rule) ([]byte, jsonparser.ValueType, error) {
	var value []byte
	var dataType jsonparser.ValueType
	//var offset int
	var err error

	value = data
	for _, field := range r.FieldsOrder {
		value, dataType, _, err = jsonparser.Get(value, field)
		if err != nil {
			return nil, jsonparser.Unknown, err
		}
	}
	return value, dataType, nil
}
