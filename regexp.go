package main

import (
	"regexp"
)

func match() {
	//countRegexp := regexp.MustCompile(`count(\((.*?)\))`)
	countRegexp := regexp.MustCompile(`count\(.*?\)`)
	matches := countRegexp.FindAllString("count(regex_match(paoding_uNJvEqvc,\".+goodjob.+.ok\"))")
	fmt.Println(matches)
}
