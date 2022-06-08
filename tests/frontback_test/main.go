package main

import (
	"strings"
	// student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	// "github.com/01-edu/go-tests/lib/random"
	// "github.com/01-edu/go-tests/solutions"
)

// func ourSolution(value string) string {
// 	expectedResult := map[string]string{
// 		"A":           "AB",
// 		"B":           "ABC",
// 		"C":           "BCD",
// 		"D":           "CDE",
// 		"E":           "DEF",
// 		"F":           "EFG",
// 		"G":           "FGH",
// 		"H":           "GHI",
// 		"I":           "HIJ",
// 		"J":           "IJK",
// 		"K":           "JKL",
// 		"L":           "KLM",
// 		"M":           "LMN",
// 		"N":           "MNO",
// 		"O":           "NOP",
// 		"P":           "OPQ",
// 		"Q":           "PQR",
// 		"R":           "QRS",
// 		"S":           "RST",
// 		"T":           "STU",
// 		"U":           "TUV",
// 		"V":           "UVW",
// 		"W":           "VXY",
// 		"X":           "WXY",
// 		"Y":           "XYZ",
// 		"Z":           "YZ",
// 		"a":           "ab",
// 		"b":           "abc",
// 		"c":           "bcd",
// 		"d":           "cde",
// 		"e":           "def",
// 		"f":           "efg",
// 		"g":           "fgh",
// 		"h":           "ghi",
// 		"i":           "hij",
// 		"j":           "ijk",
// 		"k":           "jkl",
// 		"l":           "klm",
// 		"m":           "lmn",
// 		"n":           "mno",
// 		"o":           "nop",
// 		"p":           "opq",
// 		"q":           "pqr",
// 		"r":           "qrs",
// 		"s":           "rst",
// 		"t":           "stu",
// 		"u":           "tuv",
// 		"v":           "uvw",
// 		"w":           "vxy",
// 		"x":           "wxy",
// 		"y":           "xyz",
// 		"z":           "yz",
// 		"HELLO WORLD": "\n",
// 		"ab":          "\n",
// 	}
// 	return expectedResult[value]
// }

func main() {
	table := []string{
		"A",
		"B",
		"C",
		"D",
		"AB",
		"HELLO WORLD",
	}

	for _, s := range table {
		challenge.Program("frontback", strings.Fields(s)...)
	}
	challenge.Program("frontback", "1", "a")
}
