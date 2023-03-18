package rotationalcipher

import (
	"strings"
	"unicode"
)

var baseL string = "abcdefghijklmnopqrstuvwxyz"
var baseU string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RotationalCipher(plain string, shiftKey int) string {
	res := make([]rune, len(plain))
	for i, val := range plain {
		idx := -1
		if unicode.IsLower(val) {
			idx = strings.IndexRune(baseL, val)
			if idx == -1 {
				res[i] = val
			} else {
				res[i] = rune(baseL[(idx + shiftKey) % 26])
			}
			continue
		}
		if unicode.IsUpper(val) {
			idx = strings.IndexRune(baseU, val)
			if idx == -1 {
				res[i] = val
			} else {
				res[i] = rune(baseU[(idx + shiftKey) % 26])
			}
			continue
		}
		res[i] = val
	}

	return string(res)
}
