package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value interface{}
	n *Node
	p *Node
}

type List struct {
	first *Node
	last *Node
}

func NewList(elements ...interface{}) *List {
	list := new(List)
	for _, val := range elements {
		list.Push(val)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.n
}

func (n *Node) Prev() *Node {
	return n.p
}

func (l *List) Unshift(v interface{}) {
	n := new(Node)
	n.Value = v
	if l.first == nil {
		l.first = n
		l.last = n
		return
	}

	n.n = l.first
	l.first.p = n
	l.first = n
}

func (l *List) Push(v interface{}) {
	n := new(Node)
	n.Value = v
	if l.first == nil {
		l.first = n
		l.last = n
		return
	}

	n.p = l.last
	l.last.n = n
	l.last = n
}

func (l *List) Shift() (interface{}, error) {
	if l.first == nil {
		return nil, errors.New("Empty list")
	}

	if l.first == l.last {
		n := l.first.Value
		l.first = nil
		l.last = nil
		return n, nil
	}

	n := l.first
	l.first = l.first.n
	l.first.p = nil
	return n.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.first == nil {
		return nil, errors.New("Empty list")
	}

	if l.first == l.last {
		n := l.first.Value
		l.first = nil
		l.last = nil
		return n, nil
	}

	n := l.last
	l.last = l.last.p
	l.last.n = nil
	return n.Value, nil
}

func (l *List) Reverse() {
	n := l.first
	if n == nil {
		return
	}

	for n.n != nil {
		x := n.n
		n.n = n.n.n
		if n.n != nil {
			n.n.p = n
		} else {
			l.last = n
		}
		l.Unshift(x.Value)
	}
	n = l.first
	for n != nil {
		n = n.n
	}
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
