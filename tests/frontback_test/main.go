package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"

	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		random.IntSliceBetween(0, 127),
		'A',
		'B',
		'C',
		'D',
	)

	for _, arg := range table {
		challenge.Function("frontback", student.frontback, solutions.frontback, arg)
	}
}
