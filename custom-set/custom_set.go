package stringset

type Set struct {
	s map[string]bool
}

func New() Set {
	res := Set{}
	res.s = make(map[string]bool)
	return res
}

func NewFromSlice(l []string) Set {
	res := New()
	for _, val := range l {
		res.Add(val)
	}
	return res
}

func (s Set) String() string {
	res := "{";
	isFirst := true
	for val := range s.s {
		if isFirst {
			isFirst = false
		} else {
			res += ", "
		}
		res += "\"" + val + "\""
	}
	res += "}"
	return res
}

func (s Set) IsEmpty() bool {
	return len(s.s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for val := range s1.s {
		if !s2.Has(val) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for val := range s1.s {
		if s2.Has(val) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.s) != len(s2.s) {
		return false
	}
	for val := range s1.s {
		if !s2.Has(val) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	res := New()
	for val := range s1.s {
		if s2.Has(val) {
			res.Add(val)
		}
	}
	return res
}

func Difference(s1, s2 Set) Set {
	res := New()
	for val := range s1.s {
		if !s2.Has(val) {
			res.Add(val)
		}
	}
	return res
}

func Union(s1, s2 Set) Set {
	res := New()
	for val := range s1.s {
		res.Add(val)
	}
	for val := range s2.s {
		res.Add(val)
	}
	return res
}
