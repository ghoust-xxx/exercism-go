package cryptosquare

import (
	"strings"
	"unicode"
)

func Encode(pt string) string {
	if pt == "" {
		return ""
	}

	var in string
	for _, val := range pt {
		if unicode.IsLetter(val) || unicode.IsDigit(val) {
			in += string(val)
		}
	}
	in = strings.ToLower(in)

	c, r := 1, 1
	for c * r < len(in) {
		if c == r {
			c++
		} else {
			r++
		}
	}

	var res string
	for i := 0; i < c*r; i++ {
		idx := i%r * c + i/r
		if idx >= len(in) {
			res += " "
		} else {
			res += string(in[idx])
		}
		if (i + 1) % r == 0 && i != c*r - 1 {
			res += " "
		}
	}

	return res
}
