package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, i := range layers {
		if i == "noodles" {
			noodles++
		} else if i == "sauce" {
			sauce++
		}
	}
	return 50 * noodles, 0.2 * float64(sauce)
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(src []float64, cnt int) []float64 {
	var out []float64
	for i := 0; i < len(src); i++ {
		out = append(out, src[i]/float64(2) * float64(cnt))
	}
	return out
}
