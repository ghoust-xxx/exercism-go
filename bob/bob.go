package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.Trim(remark, " ")

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	isSpace := true
	isUpper := false
	isLower := false
	for _, val := range remark {
		if unicode.IsSpace(val) {
			continue
		}

		isSpace = false
		if unicode.IsUpper(val) {
			isUpper = true
		}
		if unicode.IsLower(val) {
			isLower = true
		}
	}

	if isSpace {
		return "Fine. Be that way!"
	}

	if isUpper && !isLower {
		if remark[len(remark)-1] == '?' {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	}

	if remark[len(remark)-1] == '?' {
		return "Sure."
	}

	return "Whatever."
}
