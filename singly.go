package golinkedlist

import (
	"errors"
	"reflect"
)

func (l *SinglyLinkedList) Add(v interface{}) {
	n := CreateSinglyNode(v)
	l.Size++
	if l.first == nil {
		l.first = n
		return
	}
	l.AddLast(n)
}

func (l *SinglyLinkedList) Get(i int) (interface{}, error) {
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

func (l *SinglyLinkedList) Remove(i int) error {
	if l.Size == 0 {
		return errors.New("list empty")
	}
	if i >= l.Size {
		return errors.New("out of bound")
	}
	if l.Size == 1 {
		l.EmptyingList()
		return nil
	}

	l.Size--
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

func (l *SinglyLinkedList) Sort(ascending bool) error {
	if l.first == nil {
		return errors.New("list empty")
	}
	if l.Size == 1 {
		return nil
	}
	m := l.first
	for n := l.first.next; n.next != nil; n = n.next {
		if reflect.TypeOf(n.value) != reflect.TypeOf(m.value) {
			return errors.New("unsortable, different data type")
		}
		m = m.next
	}

	if ascending {
		return l.sortAscending()
	} else {
		return l.sortDescending()
	}
}
