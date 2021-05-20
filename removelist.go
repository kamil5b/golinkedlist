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
