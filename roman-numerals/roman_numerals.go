package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("out of range")
	}

	out := ""
	out += strings.Repeat("M", input/1000)
	input %= 1000

	i := input / 100
	switch {
	case i < 4:
		out += strings.Repeat("C", i)
	case i == 4:
		out += "CD"
	case i == 5:
		out += "D"
	case i < 9:
		out += "D" + strings.Repeat("C", i-5)
	case i == 9:
		out += "CM"
	}
	input %= 100

	i = input / 10
	switch {
	case i < 4:
		out += strings.Repeat("X", i)
	case i == 4:
		out += "XL"
	case i == 5:
		out += "L"
	case i < 9:
		out += "L" + strings.Repeat("X", i-5)
	case i == 9:
		out += "XC"
	}
	input %= 10

	i = input
	switch {
	case i < 4:
		out += strings.Repeat("I", i)
	case i == 4:
		out += "IV"
	case i == 5:
		out += "V"
	case i < 9:
		out += "V" + strings.Repeat("I", i-5)
	case i == 9:
		out += "IX"
	}

	return out, nil
}
