package resistorcolorduo

var cols = []string {
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var cmap map[string]int
	cmap = make(map[string]int)
	for i, v := range cols {
		cmap[v] = i
	}

	var out int
	if len(colors) > 0 {
		out = 10 * cmap[colors[0]]
	}
	if len(colors) > 1 {
		out += cmap[colors[1]]
	}
	return out
}
