package ejercicios

import "fmt"

type List[T comparable] struct {
	sentinelHead *Node[T]
	sentinelTail *Node[T]
	size         int
}

// NewList crea una nueva lista enlazada
func NewList[T comparable]() *List[T] {
	var def T
	sentinelHead, sentinelTail := NewNode(def), NewNode(def)
	sentinelHead.SetNext(sentinelTail)
	sentinelTail.SetPrev(sentinelHead)

	return &List[T]{sentinelHead: sentinelHead, sentinelTail: sentinelTail}
}

// Append agrega un nuevo nodo al final de la lista
func (l *List[T]) Append(dato T) {
	l.sentinelTail.InsertBefore(dato)
	l.size++
}

// Prepend agrega un nuevo nodo al principio de la lista
func (l *List[T]) Prepend(dato T) {
	l.sentinelHead.InsertAfter(dato)
	l.size++
}

// RemoveFirst elimina el primer nodo de la lista
func (l *List[T]) RemoveFirst() {
	if !l.IsEmpty() {
		l.Head().Remove()
		l.size--
	}
}

// RemoveLast elimina el último nodo de la lista
func (l *List[T]) RemoveLast() {
	if !l.IsEmpty() {
		l.Tail().Remove()
		l.size--
	}
}

// Remove elimina el nodo con el dato pasado por parámetro
func (l *List[T]) Remove(dato T) {
	if node := l.Find(dato); node != nil {
		node.Remove()
		l.size--
	}
}

// IsEmpty devuelve true si la lista está vacía
func (l *List[T]) IsEmpty() bool {
	return l.Size() == 0
}

// Size devuelve la cantidad de nodos en la lista
func (l *List[T]) Size() int {
	return l.size
}

// Head devuelve el dato del primer nodo de la lista
func (l *List[T]) Head() *Node[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.sentinelHead.Next()
}

// Tail devuelve el dato del último nodo de la lista
func (l *List[T]) Tail() *Node[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.sentinelTail.Prev()
}

// Find busca un nodo con el dato pasado por parámetro y lo devuelve
func (l *List[T]) Find(dato T) *Node[T] {
	for node := l.sentinelHead.next; node != l.sentinelTail; node = node.Next() {
		if node.Data() == dato {
			return node
		}
	}
	return nil
}

// Clear elimina todos los nodos de la lista
func (l *List[T]) Clear() {
	l.sentinelHead.SetNext(l.sentinelTail)
	l.sentinelTail.SetPrev(l.sentinelHead)
	l.size = 0
}

// String devuelve una representación en cadena de la lista
func (l *List[T]) String() string {
	result := "List: ["
	for node := l.sentinelHead.Next(); node != l.sentinelTail; node = node.Next() {
		result += fmt.Sprintf("%v", node.Data())
		if node.Next() != l.sentinelTail {
			result += ", "
		}
	}
	result += "]"
	return result
}
