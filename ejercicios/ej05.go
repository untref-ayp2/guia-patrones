package ejercicios

// Implementar un iterador que permita recorrer la matriz por fila.
type MatrizFilaIterator[T comparable] struct {
	matriz [][]T
	fila   int
	col    int
}

func (i *MatrizFilaIterator[T]) HasNext() bool {
	// Implementar
	return false
}
func (i *MatrizFilaIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}

func (m *Matriz) FilaIterator() *MatrizFilaIterator[int] {
	// Implementar
	return nil
}

// Implementar un iterador que permita recorrer la matriz por columna.
type MatrizColumnaIterator[T comparable] struct {
	matriz [][]T
	fila   int
	col    int
}

func (i *MatrizColumnaIterator[T]) HasNext() bool {
	// Implementar
	return false
}

func (i *MatrizColumnaIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}
