package main

import (
	"regexp"
	"strings"
	"fmt"
)

func main() {
	expr := regexp.MustCompile("\\b\\w")

	transformadora := func(s string) string {
		return strings.ToUpper(s)
	}

	texto := "antonio carlos jobim"
	fmt.Println(transformadora(texto))
	fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}

