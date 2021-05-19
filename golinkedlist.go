package golinkedlist

import (
	"errors"
	"reflect"
)

type Node struct {
	value      interface{}
	next, prev *Node
}

type LinkedList struct {
	first, last *Node
	Size        int
}

func CreateList() LinkedList {
	var l LinkedList
	l.EmptyingList()
	return l
}

func (l *LinkedList) EmptyingList() {
	l.first = nil
	l.last = nil
	l.Size = 0
}

func AddNewNode(v interface{}) *Node {
	N := new(Node)
	N.value = v
	N.next = nil
	N.prev = nil
	return N
}

func (l *LinkedList) Add(v interface{}) {
	n := AddNewNode(v)
	l.Size++
	if l.first == nil {
		l.first = n
		l.last = n
		return
	}
	l.AddLast(n)
}

func (l *LinkedList) Get(i int) (interface{}, error) {
	if l.Size == 0 {
		return nil, errors.New("list empty")
	}
	if i >= l.Size {
		return nil, errors.New("out of bound")
	}
	n := l.first
	for j := 0; j < i; j++ {
		n = n.next
	}
	return n.value, nil
}

func (l *LinkedList) Remove(i int) error {
	if l.Size == 0 {
		return errors.New("list empty")
	}
	if i >= l.Size {
		return errors.New("out of bound")
	}
	l.Size--
	if l.first == l.last {
		l.first, l.last = nil, nil
		return nil
	}
	if i == 0 {
		l.removeFirst()
		return nil
	}

	if i == l.Size {
		l.removeLast()
		return nil
	}

	n := l.first
	for j := 0; j < i; j++ {
		n = n.next
	}

	l.removeMiddle(n)

	return nil
}

func (l *LinkedList) Sort(ascending bool) error {
	if l.first == nil {
		return errors.New("list empty")
	}
	if l.Size == 1 {
		return nil
	}

	for n := l.first.next; n.next != nil; n = n.next {
		if reflect.TypeOf(n.value) != reflect.TypeOf(n.prev.value) {
			return errors.New("unsortable, different data type")
		}
	}

	if ascending {
		return l.sortAscendingLinear()
	} else {
		return l.sortDescendingLinear()
	}

}
