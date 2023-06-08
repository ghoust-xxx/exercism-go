package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	var res *Node

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	for _, rec := range records {
		if rec.ID >= len(records) {
			return nil, errors.New("Wrong number")
		}
		var err error
		res, err = insert(res, rec)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func insert(tree *Node, rec Record) (*Node, error) {
	if tree == nil {
		if rec.ID != rec.Parent {
			return nil, errors.New("Wrong root parent")
		}
		return &Node{ID: rec.ID, Children: []*Node{}}, nil
	}

	if rec.ID == rec.Parent {
		return nil, errors.New("Dublicate root")
	}
	node := search(tree, rec)
	if node == nil {
		return nil, errors.New("Can't fine node")
	}
	for _, val := range node.Children {
		if val.ID == rec.ID {
			return nil, errors.New("Dublicate note")
		}
	}
	node.Children = append(node.Children, &Node{ID: rec.ID, Children: []*Node{}})

	return tree, nil
}

func search(tree *Node, rec Record) *Node {
	if tree.ID == rec.Parent {
		return tree
	}
	for _, val := range tree.Children {
		node := search(val, rec)
		if node != nil {
			return node
		}
	}
	return nil
}
