package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SchoolList struct {
	School []struct {
		Name   string `yaml:"name"`
		Level  string `yaml:"level"`
		Grades []struct {
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
			Classes     []struct {
				Name      string `yaml:"name"`
				Students  string `yaml:"students,omitempty"`
				Studentes string `yaml:"studentes,omitempty"`
			} `yaml:"classes"`
		} `yaml:"grades"`
	} `yaml:"school"`
}

func ParseYAMLFile(path string) (*SchoolList, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	schoolList := &SchoolList{}
	err = yaml.Unmarshal(content, schoolList)
	if err != nil {
		return nil, err
	}
	fmt.Printf("schoolList = %+v\n", schoolList)
	return schoolList, nil
}

func main() {
	_, err := ParseYAMLFile("./test.yaml")
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}
}
