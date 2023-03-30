package acronym

import "unicode"

func Abbreviate(s string) string {
	nWord := true
	res := ""
	for _, val := range s {
		if !(unicode.IsLetter(val) || val == '-' || val == ' ') {
			continue
		}
		if val == '-' || val == ' ' {
			nWord = true
			continue
		}
		if nWord {
			res += string(unicode.ToUpper(val))
			nWord = false
		}
	}
	return res
}
