package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]string{
		{"", ""},
		{"abc", "abc"},
		{"foo", "boo"},
		{"abc", "def"},
		{"+265", "265"},
		{"eveRyone", ""},
		{"HellO", "yoAll"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"hello woRld", "fam"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
	}
	for _, arg := range table {
		challenge.Function("Weareunique", student.Weareunique, solutions.Weareunique, arg)
	}
}
