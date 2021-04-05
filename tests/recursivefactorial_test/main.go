package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := random.Ints()
	for i := 0; i <= 20; i++ {
		table = append(table, i)
	}
	for _, arg := range table {
		challenge.Function("RecursiveFactorial", student.RecursiveFactorial, solutions.RecursiveFactorial, arg)
	}
}
