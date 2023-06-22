package robotname

import (
	"errors"
	"fmt"
)

// Define the Robot type here.
type Robot struct {
	name string;
}

var names = make(map[string]bool)
var curr_id = 0

var letters = []rune("FJUQVIDNPRTHKMBCYXZGLAESOW")
var digits = []rune("3016987254")

func (r *Robot) Name() (string, error) {
	if curr_id >= 676000 {
		return "", errors.New("No more unique name")
	}
	if r.name == "" {
		r.Reset()
	}
	return fmt.Sprintf("%s", r.name), nil
}

func (r *Robot) Reset() {
	word := make([]rune, 5)
	id := curr_id

	word[4] = digits[id % 10]
	id /= 10

	word[3] = digits[id % 10]
	id /= 10

	word[2] = digits[id % 10]
	id /= 10

	word[1] = letters[id % 26]
	id /= 26

	word[0] = letters[id]

	r.name = string(word)
	curr_id++
}
