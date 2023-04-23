package flatten

func Flatten(nested interface{}) []interface{} {
	res := []interface{}{}
	deep(nested, &res)

	return res
}

func deep(item interface{}, res *[]interface{}) {
	if item == nil {
		return
	}
	switch item.(type) {
	case []interface{}:
		for _, val := range item.([]interface{}) {
			deep(val, res)
		}
	default:
		*res = append(*res, item)
	}
}
