package resistorcolortrio

import (
	"fmt"
)

var cmap map[string]int = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var mul map[int]string = map[int]string{
	0: "",
	1: "kilo",
	2: "mega",
	3: "giga",
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	res := 0
	res = cmap[colors[0]]*10 + cmap[colors[1]]
	for i := 0; i < cmap[colors[2]]; i++ {
		res *= 10
	}

	if res == 0 {
		return "0 ohms"
	}

	add := 0
	for res%1000 == 0 {
		res /= 1000
		add++
	}

	return fmt.Sprintf("%d %sohms", res, mul[add])
}
