package ejercicios

type Array[T comparable] struct {
	data []T
}

func NewArray[T comparable](data []T) *Array[T] {
	return &Array[T]{data: data}
}
