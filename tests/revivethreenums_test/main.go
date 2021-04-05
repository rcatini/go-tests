package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][4]int{
		{3, 6, 5, 4},
		{40, 40, 40, 60},
		{201, 101, 101, 200},
	}

	for i := 0; i < 25; i++ {
		first := random.IntBetween(0, 877)
		second := random.IntBetween(0, 877)
		third := random.IntBetween(0, 877)
		table = append(table, [4]int{
			first + second,
			second + third,
			first + third,
			first + second + third,
		})
	}
	for _, arg := range table {
		challenge.Function("Revivethreenums", student.ReviveThreeNums, solutions.ReviveThreeNums, arg[0], arg[1], arg[2], arg[3])
	}
}
