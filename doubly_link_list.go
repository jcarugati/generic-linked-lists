package linked_list

type DoublyLinkedList[T any] struct {
	Link[T]
}

func (dll *DoublyLinkedList[T]) Prev() *node[T] {
	if dll.current == nil {
		dll.current = dll.head
		return nil
	}

	if dll.current == dll.head {
		return nil
	}

	current := dll.current.prev()
	dll.current = current
	return current
}

func NewDoublyLink[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}
