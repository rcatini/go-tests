package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"strings"
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
		"1",
		"2",
		" ",
		"",
	}
	for _, s := range table {
		challenge.Program("frontback", strings.Fields(s)...)
	}
}
