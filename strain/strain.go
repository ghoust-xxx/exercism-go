package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var j Ints
	for _, item := range i {
		if filter(item) {
			j = append(j, item)
		}
	}
	return j
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var j Ints
	for _, item := range i {
		if !filter(item) {
			j = append(j, item)
		}
	}
	return j
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var j Lists
	for _, item := range l {
		if filter(item) {
			j = append(j, item)
		}
	}
	return j
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var j Strings
	for _, item := range s {
		if filter(item) {
			j = append(j, item)
		}
	}
	return j
}
