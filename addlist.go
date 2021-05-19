package golinkedlist

func (l *LinkedList) AddFirst(n *Node) {
	n.next = l.first
	l.first.prev = n
	l.first = n
	if l.last == nil {
		l.last = n
	}
}

func (l *LinkedList) AddLast(n *Node) {
	n.prev = l.last
	l.last.next = n
	l.last = n
	if l.first == nil {
		l.first = n
	}
}
