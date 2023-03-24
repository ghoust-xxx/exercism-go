package cipher

import (
	"strings"
	"unicode"
)

type shift struct {
	offset int
}

type vigenere struct {
	key string
}

var abc string = "abcdefghijklmnopqrstuvwxyz"

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return shift{offset: distance}
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	res := ""
	for _, val := range input {
		if !unicode.IsLetter(val) {
			continue
		}
		idx := (int(val) - 'a' + c.offset + len(abc)) % len(abc)
		res += string(abc[idx])
	}

	return res
}

func (c shift) Decode(input string) string {
	res := ""
	for _, val := range input {
		idx := (int(val) - 'a' - c.offset + len(abc)) % len(abc)
		res += string(abc[idx])
	}

	return res
}

func NewVigenere(keyInput string) Cipher {
	notA := false
	for _, val := range keyInput {
		if !unicode.IsLetter(val) || unicode.IsUpper(val) {
			return nil
		}
		if val != 'a' {
			notA = true
		}
	}
	if !notA {
		return nil
	}
	return vigenere{key: keyInput}
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	res := ""
	cnt := 0
	for _, val := range input {
		if !unicode.IsLetter(val) {
			continue
		}
		idx := (int(val) + (int(v.key[cnt % len(v.key)]) - 'a') - 'a' + len(abc)) % len(abc)
		res += string(abc[idx])
		cnt++
	}

	return res
}

func (v vigenere) Decode(input string) string {
	res := ""
	for cnt, val := range input {
		idx := (int(val) - int(v.key[cnt % len(v.key)]) + len(abc)) % len(abc)
		res += string(abc[idx])
	}

	return res
}
