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
				{ "url": "https://avatars2.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
				{ "url": "https://avatars3.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
			],
			"types":[
				1,2,3
			]
		},
		"company": {
			"name": "Acme"
		},
		"zion":{
			"detection":[ 
				"ave": {
					"level":50,
					"malware_names":"ave-virus1",
				},
				"bd": {
					"level":70,
					"malware_names":"bd-virus1",
				},
				"owl": {
					"level":70,
					"malware_names":"owl-virus1",
				}
			]
		}
		}
		`)

type rule struct {
	FieldsOrder []string
}

func main() {
	getByPath()
	arrayEachStruct()
	arrayEachSimple()
	objectEachField()
}

func objectEachField() {
	err := jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		fmt.Printf("Key: '%s'\nValue: '%s'\nType: %s\n", string(key), string(value), dataType)
		return nil
	}, "person", "name")
	if err != nil {
		panic(err)
	}
}

func arrayEachStruct() {
	_, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		eachValue, _, _, err := jsonparser.Get(value, "url")
		if err != nil {
			fmt.Printf("get each value error:%+v\n", err)
		}
		fmt.Printf("each value:%+v\n", string(eachValue))
	}, "person", "avatars")
	if err != nil {
		panic(err)
	}
}

func arrayEachSimple() {
	_, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Printf("each value:%+v\n", string(value))
	}, "person", "types")
	if err != nil {
		panic(err)
	}
}

func getByPath() {
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
