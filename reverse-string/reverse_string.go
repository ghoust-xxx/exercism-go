package reverse

func Reverse(input string) string {
	in := []rune(input)
	for i := 0; i < len(in)/2; i++ {
		tmp := in[i]
		in[i] = in[len(in)-i-1]
		in[len(in)-i-1] = tmp
	}
	return string(in)
}
