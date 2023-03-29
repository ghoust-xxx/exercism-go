package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if len(l1) < len(l2) {
		for i := 0; i + len(l1) <= len(l2); i++ {
			if isEqual(l1, l2[i:i+len(l1)]) {
				return RelationSublist
			}
		}
	} else if len(l1) > len(l2) {
		for i := 0; i + len(l2) <= len(l1); i++ {
			if isEqual(l2, l1[i:i+len(l2)]) {
				return RelationSuperlist
			}
		}
	} else {
		if isEqual(l1, l2) {
			return RelationEqual
		}
	}

	return RelationUnequal
}

func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for k, _ := range l1 {
		if l1[k] != l2[k] {
			return false
		}
	}

	return true
}
