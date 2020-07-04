package main

import (
	"log"
)

func main() {
	worker(hello, 10)
	conf := NewConfig(
		withNameOption("testconfig"),
		withAgeOption(10),
		withAddrOption("test config addr"),
	)
	log.Print(conf)
}

func hello(i interface{}) {
	log.Print("hello", i.(int))
}

func worker(f func(a interface{}), i interface{}) {
	f(i)
}

type option func(a *config)

type config struct {
	name string
	age  int
	addr string
}

func withNameOption(name string) option {
	return func(c *config) {
		c.name = name
	}
}
func withAgeOption(age int) option {
	return func(c *config) {
		c.age = age
	}
}
func withAddrOption(addr string) option {
	return func(c *config) {
		c.addr = addr
	}
}
func NewConfig(opts ...option) *config {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}
	return conf
}
