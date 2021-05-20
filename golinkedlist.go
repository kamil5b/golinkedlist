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

//SINGLY
type Singly_Node struct {
	value interface{}
	next  *Singly_Node
}

type SinglyLinkedList struct {
	first *Singly_Node
	Size  int
}

func (l *SinglyLinkedList) EmptyingList() {
	l.first = nil
	l.Size = 0
}

func CreateSinglyNode(v interface{}) *Singly_Node {
	N := new(Singly_Node)
	N.value = v
	N.next = nil
	return N
}

//DOUBLY
type Doubly_Node struct {
	value      interface{}
	next, prev *Doubly_Node
}

type DoublyLinkedList struct {
	first, last *Doubly_Node
	Size        int
}

func (l *DoublyLinkedList) EmptyingList() {
	l.first = nil
	l.last = nil
	l.Size = 0
}

func CreateDoublyNode(v interface{}) *Doubly_Node {
	N := new(Doubly_Node)
	N.value = v
	N.next = nil
	N.prev = nil
	return N
}
