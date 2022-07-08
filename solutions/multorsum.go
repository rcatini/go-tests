package solutions

func MultOrSum(ints []int, init int) int {
	result := init
	for _, num := range ints {
		if num%2 == 0 {
			result += num
			continue
		}
		result *= num
	}
	return result
}
