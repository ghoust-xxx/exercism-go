package proverb

import "fmt"

func Proverb(rhyme []string) (out []string) {
	if len(rhyme) == 0 {
		return nil
	}
	for i := 0; i < len(rhyme) - 1; i++ {
		out = append(out, fmt.Sprintf("For want of a %s the %s was lost.",
			rhyme[i], rhyme[i+1]))
	}
	out = append(out, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return
}
