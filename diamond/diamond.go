package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	out := strings.Builder{}

	if char < 'A' || char > 'Z' {
		return "", errors.New("Out of range")
	}

	if char == 'A' {
		return "A", nil
	}

	for sp := byte('A'); sp <= char; sp++ {
		out.WriteString(Line(sp, int(char - 'A') * 2 + 1))
		out.WriteString("\n")
	}
	for sp := char - 1; sp >= 'A'; sp-- {
		out.WriteString(Line(sp, int(char - 'A') * 2 + 1))
		if sp != 'A' {
			out.WriteString("\n")
		}
	}


	return out.String(), nil
}

func Line(char byte, l int) string {
	out := strings.Builder{}

	for i := 0; i < l; i++ {
		if i == l / 2 - int(char - 'A') || i == l / 2 + int(char - 'A') {
			out.WriteString(string(char))
		} else {
			out.WriteString(string(" "))
		}
	}

	return out.String()
}
