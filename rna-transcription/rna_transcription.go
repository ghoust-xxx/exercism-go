package strand

func ToRNA(dna string) string {
	out := ""
	for _, c := range dna {
		switch c {
		case 'G':
			out += "C"
		case 'C':
			out += "G"
		case 'T':
			out += "A"
		case 'A':
			out += "U"
		default:
			return ""
		}
	}
	return out
}
