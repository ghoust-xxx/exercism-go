package isogram

import "strings"

func IsIsogram(word string) bool {
	letters := map[rune]int{}
	str := []rune(strings.ToLower(word))
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '-' {
			continue
		}
		letters[str[i]]++
		if letters[str[i]] > 1 {
			return false
		}
	}
	return true
}
