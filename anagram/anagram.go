package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	let := make(map[rune]int)
	for _, val := range strings.ToLower(subject) {
		let[val]++
	}

	var out []string
	words:
	for _, word := range candidates {
		if len(word) != len(subject) ||
			strings.ToLower(word) == strings.ToLower(subject) {
			continue
		}
		c := make(map[rune]int)
		for _, val := range strings.ToLower(word) {
			c[val]++
		}
		if len(let) != len(c) {
			continue
		}
		for i, _ := range c {
			if c[i] != let[i] {
				continue words
			}
		}
		out = append(out, word)
	}

	return out
}
