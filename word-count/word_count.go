package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	res := make(Frequency)
	phrase = strings.ToLower(phrase)
	word := ""
	for _, c := range phrase {
		if !(unicode.IsLetter(c) || unicode.IsDigit(c) || c == '\'') {
			if len(word) != 0 {
				add_word(res, word)
				word = ""
			}
			continue
		}
		word += string(c)
	}
	add_word(res, word)

	return res
}

func add_word(res Frequency, word string) {
	if (len(word) > 0 && word[0] == byte('\'')) {
		word = word[1:]
	}
	if (len(word) > 0 && word[len(word)-1] == byte('\'')) {
		word = word[:len(word)-1]
	}
	if len(word) > 0 {
		res[word]++
	}
}
