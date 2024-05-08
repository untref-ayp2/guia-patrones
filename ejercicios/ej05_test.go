package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrizFilaIterator(t *testing.T) {
	matriz := Matriz{
		Datos: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	iter := matriz.FilaIterator()

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 0

	for iter.HasNext() {
		val, err := iter.Next()
		assert.NoError(t, err)
		assert.Equal(t, expected[index], val)
		index++
	}

	assert.Equal(t, len(expected), index)
}

func TestMatrizColumnaIterator(t *testing.T) {
	matriz := Matriz{
		Datos: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	iter := &MatrizColumnaIterator[int]{matriz: matriz.Datos, fila: 0, col: 0}

	expected := []int{1, 4, 7, 2, 5, 8, 3, 6, 9}
	index := 0

	for iter.HasNext() {
		val, err := iter.Next()
		assert.NoError(t, err)
		assert.Equal(t, expected[index], val)
		index++
	}

	assert.Equal(t, len(expected), index)
}

func TestMatrizFilaIterator_EmptyMatrix(t *testing.T) {
	matriz := Matriz{
		Datos: [][]int{},
	}

	iter := matriz.FilaIterator()

	assert.False(t, iter.HasNext())
}

func TestMatrizColumnaIterator_EmptyMatrix(t *testing.T) {
	matriz := Matriz{
		Datos: [][]int{},
	}

	iter := &MatrizColumnaIterator[int]{matriz: matriz.Datos, fila: 0, col: 0}

	assert.False(t, iter.HasNext())
}
