package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var out string
	x := rune(input[0])
	cnt := 1
	for _, val := range input[1:] {
		if x == val {
			cnt++
			continue
		}

		if cnt > 1 {
			out += fmt.Sprintf("%d%s", cnt, string(x))
		} else {
			out += string(x)
		}
		cnt = 1
		x = val
	}
	if cnt > 1 {
		out += fmt.Sprintf("%d%s", cnt, string(x))
	} else {
		out += string(x)
	}
	return out
}

func RunLengthDecode(input string) string {
	var out string
	tmp := ""
	for _, val := range input {
		if unicode.IsDigit(val) {
			tmp += string(val)
			continue
		}

		if len(tmp) == 0 {
			out += string(val)
			continue
		}

		cnt, _ := strconv.Atoi(tmp)
		tmp = ""
		out += strings.Repeat(string(val), cnt)
	}
	return out
}
