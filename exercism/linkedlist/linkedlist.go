package linkedlist

import (
	"errors"
)

type Node struct {
	Value      interface{}
	prev, next *Node
}

type List struct {
	first, last *Node
	NumNodes    int
}

func NewList(args ...interface{}) *List {
	nodes := make([]Node, len(args))
	for i, arg := range args {
		nodes[i] = Node{Value: arg}
		if i > 0 {
			nodes[i].prev = &nodes[i-1]
			nodes[i-1].next = &nodes[i]
		}
	}
	return &List{first: &nodes[0], last: &nodes[len(args)-1], NumNodes: len(args)}
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) Push(v interface{}) {
	first := l.First()
	new := Node{Value: v}

	first.prev = &new
	new.next = first
	l.first = &new
	l.NumNodes++
}

func (l *List) Unshift(v interface{}) {
	last := l.Last()
	new := Node{Value: v}

	last.next = &new
	new.prev = last
	l.last = &new
	l.NumNodes++
}

func (l *List) Pop() (interface{}, error) {
	if l.NumNodes == 1 {
		only := l.First()
		l.first = nil
		l.last = nil
		l.NumNodes--
		return only.Value, nil
	}
	first := l.First()
	if first == nil {
		return nil, errors.New("Empty list or lacking first node")
	}
	newFirst := first.Next()
	l.first = newFirst
	newFirst.prev = nil
	first.next = nil

	l.NumNodes--
	return first.Value, nil

}

func (l *List) Shift() (interface{}, error) {
	if l.NumNodes == 1 {
		only := l.Last()
		l.first = nil
		l.last = nil
		l.NumNodes--
		return only.Value, nil
	}
	last := l.Last()
	if last == nil {
		return nil, errors.New("Empty list or lacking last node")
	}
	newLast := last.Prev()
	l.last = newLast
	newLast.next = nil
	last.prev = nil

	l.NumNodes--
	return last.Value, nil

}

func (l *List) Reverse() {
	if l.NumNodes == 1 {
		return
	}
	first := l.First()
	last := l.Last()
	curr := l.First()
	for i := 0; i < l.NumNodes; i++ {
		nextCurr := curr.Next()

		// if curr != nil {
		// 	fmt.Printf("%d -- ", curr.Value)
		// }
		// if nextCurr != nil {
		// 	fmt.Printf("next: %d \n", nextCurr.Value)
		// }

		curr.prev, curr.next = curr.next, curr.prev
		curr = nextCurr

	}
	l.first = last
	l.last = first

}
