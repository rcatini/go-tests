package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	years := []string{
		//CamelCase
		"CamelCase",
		"camelCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"ThisIsATest",
	}
	for _, arg := range years {
		challenge.Function("CamelToSnakeCase", student.CamelToSnakeCase, solutions.CamelToSnakeCase, arg)
	}
}
