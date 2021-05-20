package golinkedlist

func (l *DoublyLinkedList) removeFirst() {
	n := l.first
	l.first = n.next
	l.first.prev = nil

}

func (l *DoublyLinkedList) removeLast() {
	n := l.last
	l.last = n.prev
	l.last.next = nil

}

func (l *DoublyLinkedList) removeMiddle(n *Doubly_Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *SinglyLinkedList) removeFirst() {
	n := l.first
	l.first = n.next

}

func (l *SinglyLinkedList) removeLast() {

	m := l.first
	for m.next.next != nil {
		m = m.next
	}
	m.next = nil

}

func (l *SinglyLinkedList) removeMiddle(n *Singly_Node) {

	m := l.first
	for m.next != n && m != nil {
		m = m.next
	}

	if m != nil {
		m.next = n.next
	}

}
