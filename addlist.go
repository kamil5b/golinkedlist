package golinkedlist

func (l *DoublyLinkedList) AddFirst(n *Doubly_Node) {
	n.next = l.first
	l.first.prev = n
	l.first = n
	if l.last == nil {
		l.last = n
	}
}

func (l *DoublyLinkedList) AddLast(n *Doubly_Node) {
	n.prev = l.last
	l.last.next = n
	l.last = n
	if l.first == nil {
		l.first = n
	}
}

func (l *SinglyLinkedList) AddFirst(n *Singly_Node) {
	n.next = l.first
	l.first = n
}

func (l *SinglyLinkedList) AddLast(n *Singly_Node) {

	m := l.first
	for m.next != nil {
		m = m.next
	}

	m.next = n

}
