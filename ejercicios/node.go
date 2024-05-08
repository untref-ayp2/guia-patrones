package ejercicios

type Node[T comparable] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

// NewNode crea un nuevo nodo con el dato proporcionado.
func NewNode[T comparable](dato T) *Node[T] {
	return &Node[T]{data: dato, next: nil, prev: nil}
}

// Data devuelve el dato almacenado en el nodo.
func (n *Node[T]) Data() T {
	return n.data
}

// Next devuelve el siguiente nodo en la lista.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Prev devuelve el nodo anterior en la lista.
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// SetData establece un nuevo dato en el nodo.
func (n *Node[T]) SetData(dato T) {
	n.data = dato
}

// SetNext establece el siguiente nodo en la lista.
func (n *Node[T]) SetNext(sig *Node[T]) {
	n.next = sig
}

// SetPrev establece el nodo anterior en la lista.
func (n *Node[T]) SetPrev(prev *Node[T]) {
	n.prev = prev
}

// InsertAfter inserta un nuevo nodo con el dato proporcionado después del nodo actual.
func (n *Node[T]) InsertAfter(dato T) {
	newNode := NewNode(dato)
	newNode.prev = n
	newNode.next = n.next
	if n.next != nil {
		n.next.prev = newNode
	}
	n.next = newNode
}

// InsertBefore inserta un nuevo nodo con el dato proporcionado antes del nodo actual.
func (n *Node[T]) InsertBefore(dato T) {
	newNode := NewNode(dato)
	newNode.next = n
	newNode.prev = n.prev
	if n.prev != nil {
		n.prev.next = newNode
	}
	n.prev = newNode
}

// Remove elimina el nodo actual de la lista.
func (n *Node[T]) Remove() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	n.next = nil
	n.prev = nil
}

// RemoveNext elimina el nodo siguiente al nodo actual.
func (n *Node[T]) RemoveNext() {
	if n.next != nil {
		nextNode := n.next
		n.next = nextNode.next
		if nextNode.next != nil {
			nextNode.next.prev = n
		}
		nextNode.next = nil
		nextNode.prev = nil
	}
}

// RemovePrev elimina el nodo anterior al nodo actual.
func (n *Node[T]) RemovePrev() {
	if n.prev != nil {
		prevNode := n.prev
		n.prev = prevNode.prev
		if prevNode.prev != nil {
			prevNode.prev.next = n
		}
		prevNode.next = nil
		prevNode.prev = nil
	}
}
