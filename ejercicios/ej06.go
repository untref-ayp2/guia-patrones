package ejercicios

type ListIterator[T comparable] struct {
	// Implementar
}

// NewListIterator crea un nuevo iterador para la lista especificada.
func NewListIterator[T comparable](list *List[T]) *ListIterator[T] {
	// Implementar
	return nil
}

// Next devuelve el siguiente elemento en la lista y avanza la posición del iterador.
func (i *ListIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}

// HasNext verifica si hay un siguiente elemento en la lista.
func (i *ListIterator[T]) HasNext() bool {
	// Implementar
	return false
}

// Previous devuelve el elemento anterior en la lista y retrocede la posición del iterador.
func (i *ListIterator[T]) Previous() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}

// HasPrevious verifica si hay un elemento anterior en la lista.
func (i *ListIterator[T]) HasPrevious() bool {
	// Implementar
	return false
}

// Iterator crea un nuevo iterador para la lista.
func (l *List[T]) Iterator() *ListIterator[T] {
	// Implementar
	return nil
}
