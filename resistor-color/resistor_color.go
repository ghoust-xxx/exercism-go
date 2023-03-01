package resistorcolor

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

var cmap map[string]int

// Colors should return the list of all colors.
func Colors() []string {
	return cols
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	if cmap == nil {
		cmap = make(map[string]int)
		for i, v := range cols {
			cmap[v] = i
		}
	}
	return cmap[color]
}
