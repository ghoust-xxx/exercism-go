package piglatin

import (
	"strings"
)

var vowels = []string{ "a", "e", "i", "o", "u", "y", "xr", "yt"}

func Sentence(s string) string {
	words := strings.Split(s, " ")
	var out []string
	for _, word := range words {
		out = append(out, Word(word))
	}
	return strings.Join(out, " ")
}


func Word (s string) string {
	v_map := make(map[string]bool)
	for _, val := range vowels {
		v_map[val] = true
	}

	// If vowel
	if string(s[0]) == "y" && string(s[1]) != "t" {
		return string(s[1:]) + "yay"
	}
	for _, val := range vowels {
		if strings.HasPrefix(s, val) {
			return s + "ay"
		}
	}

	// 3rd rule
	begin := ""

	if len(s) >= 2 {
		if string(s[1]) == "q" && string(s[2]) == "u" {
			return string(s[3:]) + string(s[0]) + "quay"
		}

		if string(s[0]) == "q" && string(s[1]) == "u" {
			return string(s[2:]) + "quay"
		}
	}

	// If consonant
	end := ""
	begin = ""
	pref := true
	for _, c := range s {
		if pref {
			_, ok := v_map[string(c)]
			if ok {
				pref = false
				begin = string(c)
			} else {
				end += string(c)
			}
		} else {
			begin += string(c)
		}
	}
	return begin + end + "ay"
}
