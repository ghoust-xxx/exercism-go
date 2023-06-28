package allergies

var alls = map[string]uint{
	"eggs": 1,
	"peanuts": 2,
	"shellfish": 4,
	"strawberries": 8,
	"tomatoes": 16,
	"chocolate": 32,
	"pollen": 64,
	"cats": 128,
}

func Allergies(all uint) []string {
	var out []string
	for key, val := range alls {
		if val & all > 0 {
			out = append(out, key)
		}
	}
	return out
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies & alls[allergen] > 0
}
