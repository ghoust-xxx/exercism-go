package isbn

import (
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	i := 1
	sum := 0
	rest := ""
	for idx, v := range isbn {
		if i > 10 {
			rest = isbn[idx:]
			break
		}
		if unicode.IsNumber(v) || unicode.IsLetter(v) {
			var num int
			var err error
			if v == 'X' &&  i == 10 {
				num = 10
			} else {
				num, err = strconv.Atoi(string(v))
				if err != nil {
					return false
				}
			}
			sum += num * (11 - i)
			i++
		}
	}
	if i != 11 || rest != "" {
		return false
	}

	return (sum % 11) == 0
}
