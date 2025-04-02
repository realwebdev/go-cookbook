package main

import (
	"fmt"
	"strings"
)

// call back events

type CallbackFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(s string) string {
	return "Foo_" + s
}

func executeCallback(s string, callback CallbackFunc) {
	fmt.Println(callback(s))
}

func main() {
	executeCallback("hello", Uppercase)
	executeCallback("hello", Prefixer)
}
