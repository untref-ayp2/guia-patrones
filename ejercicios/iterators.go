package ejercicios

type Iterator[T comparable] interface {
	// HasNext verifica si hay un siguiente elemento.
	HasNext() bool
	// Next devuelve el siguiente elemento.
	// Si no hay más elementos, devuelve un error.
	// Devuelve el elemento actual y luego avanza el iterador.
	Next() (T, error)
}

type DoubleIterator[T comparable] interface {
	// HasNext verifica si hay un siguiente elemento.
	// Si no hay más elementos, devuelve false.
	HasNext() bool
	// Next devuelve el elemento actual y luego avanza el iterador
	// Si no hay más elementos, devuelve un error.
	Next() (T, error)
	//
	// HasPrevious verifica si hay un elemento anterior.
	// Si no hay más elementos, devuelve false.
	HasPrevious() bool
	// Previous devuelve el elemento actual y luego retrocede el iterador.
	// Si no hay más elementos, devuelve un error.
	Previous() (T, error)
}
