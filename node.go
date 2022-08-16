package linked_list

type node[T any] struct {
	item     T
	follows  *node[T]
	previous *node[T]
}

func (n *node[T]) next() *node[T] {
	return n.follows
}

func (n *node[T]) GetValue() T {
	return n.item
}

func (n *node[T]) prev() *node[T] {
	return n.previous
}
