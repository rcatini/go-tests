package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		s         string
		toCompare string
	}

	table := []node{}

	// the first 15 values are returning 0 for this test
	for i := 0; i < 15; i++ {
		wordToTest := random.RandASCII()

		val := node{
			s:         wordToTest,
			toCompare: wordToTest,
		}
		table = append(table, val)
	}

	// the next 15 values are supposed to return 1 or -1 for this test
	for i := 0; i < 15; i++ {
		wordToTest := random.RandASCII()
		wrongMatch := random.RandASCII()

		val := node{
			s:         wordToTest,
			toCompare: wrongMatch,
		}
		table = append(table, val)
	}
	// those are the test values from the README examples
	table = append(table,
		node{s: "Hello!", toCompare: "Hello!"},
		node{s: "Salut!", toCompare: "lut!"},
		node{s: "Ola!", toCompare: "Ol"},
	)

	for _, arg := range table {
		challenge.Function("Compare", student.Compare, solutions.Compare, arg.s, arg.toCompare)
	}
}
