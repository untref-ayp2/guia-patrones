package ejercicios

type PosicionesParesIterator[T comparable] struct {
	// Implementar
}

func (i *PosicionesParesIterator[T]) HasNext() bool {
	// Implementar
	return false
}

func (i *PosicionesParesIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}

func (a *Array[T]) PosicionesParesIterator() *PosicionesParesIterator[T] {
	// Implementar
	return nil
}
