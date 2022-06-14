package solutions

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result += "_"
			result += string(s[i])
		} else {
			result += string(s[i])
		}
	}
	return result
}
