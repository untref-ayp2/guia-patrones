package ejercicios

type Iterator[T comparable] interface {
	// HasNext verifica si hay un siguiente elemento.
	HasNext() bool
	// Next devuelve el siguiente elemento.
	// Si no hay más elementos, devuelve un error.
	// Primero avanza el iterador y luego devuelve el elemento.
	Next() (T, error)
}

type DoubleIterator[T comparable] interface {
	// HasNext verifica si hay un siguiente elemento.
	// Si no hay más elementos, devuelve false.
	HasNext() bool
	// Next devuelve el siguiente elemento.
	// Si no hay más elementos, devuelve un error.
	Next() (T, error)
	//
	// HasPrevious verifica si hay un elemento anterior.
	// Si no hay más elementos, devuelve false.
	HasPrevious() bool
	// Previous devuelve el elemento anterior.
	// Si no hay más elementos, devuelve un error.
	Previous() (T, error)
}
