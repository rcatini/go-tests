package main

import (
	"strings"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := []string{
		"A",
		"B",
		"C",
		"D",
		"AB",
		"a",
		"z",
		"HELLO WORLD",
		" ",
		"",
	}
	for _, s := range table {
		challenge.Program("frontback", strings.Fields(s)...)
	}
}
