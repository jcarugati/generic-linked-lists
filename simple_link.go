package linked_list

type Link[T any] struct {
	head, current, tail *node[T]
	length              int
}

func NewSimpleLink[T any]() *Link[T] {
	return &Link[T]{
		length: 0,
	}
}

func (l *Link[T]) Add(h T) {
	n := &node[T]{
		item: h,
	}

	l.length++

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	tail := l.tail
	tail.follows = n

	l.tail = tail.next()
}

func (l *Link[T]) Next() *node[T] {
	if l.current == nil {
		l.current = l.head
		return l.head
	}
	if l.current == l.tail {
		return nil
	}

	current := l.current.next()
	l.current = current

	return current
}
