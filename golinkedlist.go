package golinkedlist

type List interface {
	EmptyingList()
	Add(interface{})
	Get(int) (interface{}, error)
	Remove(int) error
	Sort(bool) error
	sortAscending() error
	sortDescending() error
}

type Doubly_Node struct {
	value      interface{}
	next, prev *Doubly_Node
}

type DoublyLinkedList struct {
	first, last *Doubly_Node
	Size        int
}

func CreateList() DoublyLinkedList {
	var l DoublyLinkedList
	l.EmptyingList()
	return l
}

func (l *DoublyLinkedList) EmptyingList() {
	l.first = nil
	l.last = nil
	l.Size = 0
}

func AddNewNode(v interface{}) *Doubly_Node {
	N := new(Doubly_Node)
	N.value = v
	N.next = nil
	N.prev = nil
	return N
}
