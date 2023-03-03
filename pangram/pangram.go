package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	var al [26]int
	for _, val := range input {
		if val > 'z' || val < 'a' {
			continue
		}
		al[val - 'a']++
	}
	for _, val := range al {
		if val == 0 {
			return false
		}
	}
	return true
}
