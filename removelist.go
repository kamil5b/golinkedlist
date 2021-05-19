package golinkedlist

func (l *LinkedList) removeFirst() {

	n := l.first
	l.first = n.next
	l.first.prev = nil

}

func (l *LinkedList) removeLast() {

	n := l.last
	l.last = n.prev
	l.last.next = nil

}

func (l *LinkedList) removeMiddle(n *Node) {

	n.prev.next = n.next
	n.next.prev = n.prev
}
