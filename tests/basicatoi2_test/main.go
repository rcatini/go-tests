package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(random.IntBetween(0, random.MaxInt))
	}
	table = append(table,
		strconv.Itoa(random.MaxInt),
		"",
		"0",
		"Invalid123",
		"123Invalid",
		"Invalid",
		"1Invalid23",
		"123",
		"123.",
		"123.0",
	)
	for _, arg := range table {
		challenge.Function("BasicAtoi2", student.BasicAtoi2, solutions.BasicAtoi2, arg)
	}
}
