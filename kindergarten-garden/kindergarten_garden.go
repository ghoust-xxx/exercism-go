package kindergarten

import (
	"errors"
	"sort"
)

type Garden struct {
	plants string;
	children []string;
}

var list = map[string]string{
	"G": "grass",
	"C": "clover",
	"R": "radishes",
	"V": "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	rLen := len(diagram)/2
	if diagram[0] != '\n' || diagram[rLen] != '\n' || rLen % 2 == 0 {
		return nil, errors.New("error")
	}
	for i, val := range diagram {
		if i == 0 || i == rLen {
			continue
		}
		if !(val == 'G' || val == 'C' || val == 'R' || val == 'V') {
			return nil, errors.New("error")
		}
	}
	var empty []string
	list := append(empty, children...)
	sort.Strings(list)
	lastChild := ""
	for _, val := range list {
		if val == lastChild {
			return nil, errors.New("error")
		}
		lastChild = val
	}
	return &Garden{diagram, list}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	var res []string
	idx := -1
	for i, val := range g.children {
		if val == child {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, false
	}
	rLen := len(g.plants)/2
	idx *= 2
	res = append(res, list[string(g.plants[1+idx])])
	res = append(res, list[string(g.plants[1+idx+1])])
	res = append(res, list[string(g.plants[1+idx+rLen])])
	res = append(res, list[string(g.plants[1+idx+rLen+1])])

	return res, true
}
