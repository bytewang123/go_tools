package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"regexp"
	"strconv"
	"strings"
)

var data = []byte(`
		{
		"zion":{
			"detection":{
                "ave":{
                    "malware_name":"Sorter.AutoAdd.123.5001",
                    "version":"1.1",
					"level":50,
					"detail":{
						"sublevel":1
					}
                },
                "bd":{
                    "malware_name":"bd-virus1",
                    "version":"1.1",
					"level":60,
					"detail":{
						"sublevel":2
					}
                },
                "owl":{
                    "malware_name":"owl-virus1",
                    "version":"1.1",
					"level":70,
					"detail":{
						"sublevel":3
					}
                }
		}
		},
		"vt":{
			"detection":{
                "ave":{
                    "malware_name":"vt-Sorter.AutoAdd.123.5001",
                    "version":"1.1",
					"level":70,
					"detail":{
						"sublevel":1
					}
                },
                "bd":{
                    "malware_name":"vt-bd-virus1",
                    "version":"1.1",
					"level":60,
					"detail":{
						"sublevel":2
					}
                },
                "owl":{
                    "malware_name":"vt-owl-virus1",
                    "version":"1.1",
					"level":50,
					"detail":{
						"sublevel":3
					}
                }
		}
		}
		}`)

var report1 = `
[{"report_uri":1111},{"report_uri":1112},{"report_uri":1113}]
`

type rule struct {
	FieldsOrder []string
}

type ParseRule struct {
	Path        []string
	Type        string
	TargetValue string
	Expr        string
}

func main() {
	/*
		getByPath()
		arrayEachStruct()
		arrayEachSimple()
		objectEachField()
		setValueInArray()
		setValueByPath()
	*/
	//getValueBySimplePath()
	//match()
	//queryWithMultiWildCard()
	//regexpJson()
	//getReport()
	recursiveQueryWithWildcards()
}

func getReport() {
	jsonparser.ArrayEach([]byte(report1), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		//v, _, _, err := jsonparser.Get(value, "report_uri")
		fmt.Println("report_uri", string(value), err)
	})
}

func queryWithMultiWildCard() {
	//*.detection.*.level
	path := []string{"*", "detection", "*", "level"}
	getValue := data
	getEachValues := [][]byte{}
	result := map[string]string{}

	for n, k := range path {
		if k == "*" {
			if len(getEachValues) != 0 {
				fmt.Println(k)
				newEachValues := [][]byte{}
				for _, v := range getEachValues {
					jsonparser.ObjectEach(v, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
						newEachValues = append(newEachValues, value)
						return nil
					})
				}
				getEachValues = newEachValues
			} else {
				jsonparser.ObjectEach(getValue, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
					getEachValues = append(getEachValues, value)
					return nil
				})
				fmt.Println(len(getEachValues))
			}
		} else {
			var err error
			if len(getEachValues) != 0 {
				newEachValues := [][]byte{}
				for _, v := range getEachValues {
					getValue, _, _, err = jsonparser.Get(v, k)
					if err != nil {
						panic(err)
					}
					fmt.Println(string(getValue))
					newEachValues = append(newEachValues, getValue)
				}
				getEachValues = newEachValues
			} else {
				getValue, _, _, err = jsonparser.Get(getValue, k)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(getValue))
			}
		}
		fmt.Println(k, "==========================")

		if n == len(path)-1 {
			for iter, iterVal := range getEachValues {
				result[fmt.Sprintf("eachValue_%+v", iter)] = string(iterVal)
			}
		}
	}

	fmt.Println(result)
}

func recursiveQueryWithWildcards() {
	result := []string{}
	recursiveGetData([]string{"*", "*", "*", "malware_name"}, data, &result, 0)
	fmt.Println(result)
}

func recursiveGetData(keys []string, data []byte, result *[]string, layer int) {

	if layer == len(keys) {
		*result = append(*result, string(data))
	} else {
		switch keys[layer] {
		case "*":
			err := jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
				recursiveGetData(keys, value, result, layer+1)
				return nil
			})
			if err != nil {
				panic(err)
			}
		default:
			getted, _, _, err := jsonparser.Get(data, keys[layer])
			if err != nil {
				panic(err)
			}
			recursiveGetData(keys, getted, result, layer+1)
		}
	}
}

func match() {
	//countRegexp := regexp.MustCompile(`count(\((.*?)\))`)
	countRegexp := regexp.MustCompile(`count\(.*?\)`)
	matches := countRegexp.FindAllString("count(levels == 70 && v1 == \"bd-virus1\")  > 1 || count(levels2 == 60) && v2 == \"ave-virus1\"", -1)
	fmt.Println(matches)
}

func flat() {
	cond := `levels == 70 && 
    b  not match \"(heur)|(\bgen\b)|(\bgeneric.*?)|(^not-a-virus:)\"  && 
    c  not match \"(heur)|(\bgen\b)|(\bgeneric.*?)|(confidence)|(^pua)\" && 
    d  not match \"(heur)|(\bgen\b)|(\bgeneric.*?)|(pup-)|(artemis)\" && 
    e  not match \"(heur)|(\bgen\b)|(\bgeneric.*?)|(_generic)|(^pua)|(^Possible_)\"`
	aa := []int{50, 60, 70}
	for _, sa := range aa {
		res := strings.Replace(cond, "levels", strconv.Itoa(sa), -1)
		fmt.Println(res)
	}

}

func getValueBySimplePath() {
	v, _, _, err := jsonparser.Get(data, "zion", "detection2")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))
	err = jsonparser.ObjectEach(v, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		level, _, _, err := jsonparser.Get(value, "detail", "sublevel")
		if err != nil {
			panic(err)
		}
		fmt.Printf("level:%+v\n", string(level))
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func setValueByPath() {
	v, err := jsonparser.Set(data, []byte("100"), "zion", "detection2")
	if err != nil {
		fmt.Printf("set subelevel value error:%+v\n", err)
		panic(err)
	}
	fmt.Printf("set value:%+v\n", string(v))
}

func setValueInArray() {
	i := 0
	_, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Printf("idx:%+v\n", i)
		i++
		eachValue, err := jsonparser.Set(value, []byte("100"), "sublevel")
		if err != nil {
			fmt.Printf("set each value error:%+v\n", err)
		}

		fmt.Printf("each value:%+v\n", string(eachValue))
	}, "zion", "detection")
	if err != nil {
		panic(err)
	}
	fmt.Printf("all:%+v\n", string(data))
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
		eachValue, _, _, err := jsonparser.Get(value, "level")
		if err != nil {
			fmt.Printf("get each value error:%+v\n", err)
		}
		fmt.Printf("each value:%+v\n", string(eachValue))
	}, "zion", "detection")
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
