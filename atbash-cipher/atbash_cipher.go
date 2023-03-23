package atbash

import (
	"strings"
	"unicode"
)


func Atbash(s string) string {
	var res string
	var c []rune = []rune("zyxwvutsrqponmlkjihgfedcba")
	s = strings.ToLower(s)
	cnt := 0
	for _, val := range s {
		if !(unicode.IsDigit(val) || unicode.IsLetter(val)) {
			continue
		}

		if cnt != 0 && cnt % 5 == 0 {
			res += " "
		}
		cnt++

		if unicode.IsDigit(val) {
			res += string(val)
		} else {
			res += string(c[val - 'a'])
		}
	}

	return res
}
