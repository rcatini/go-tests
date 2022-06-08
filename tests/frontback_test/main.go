package main

import (
	"strings"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := []string{
		"A", // A ---> AB
		"B", // B ---> ABC
		"C", // C --- BCD
		"D", // D ---> CDE
		"AB", // AB ---> \n
		"a", // a ---> ab
		"z", // z ---> yz
		"HELLO WORLD", // \n
		" ",// \n
	}

	for _, s := range table {
		challenge.Program("frontback", strings.Fields(s)...)
	}
	challenge.Program("frontback", "1")
}
