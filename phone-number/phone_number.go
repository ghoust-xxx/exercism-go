package phonenumber

import (
	"fmt"
	"errors"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var res []rune
	for _, val := range phoneNumber {
		if !unicode.IsDigit(val) {
			continue
		}
		res = append(res, val)
	}

	if len(res) == 11 {
		if res[0] == '1' {
			res = res[1:]
		} else {
			return "", errors.New("wrong")
		}
	}

	if res[0] == '1' || res[0] == '0' || res[3] == '1' || res[3] == '0' ||
		len(res) != 10 {
		return "", errors.New("wrong")
	}

	return string(res), nil
}

func AreaCode(phoneNumber string) (string, error) {
	clean, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return clean[:3], nil
}

func Format(phoneNumber string) (string, error) {
	clean, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", clean[0:3], clean[3:6], clean[6:]), nil
}
