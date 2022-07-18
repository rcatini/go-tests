package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]string{
		{"foo", "boo"},
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"", ""},
		{"abc", "abc"},
	}
	for _, arg := range table {
		challenge.Function("Weareunique", student.Weareunique, solutions.Weareunique, arg)
	}
}
