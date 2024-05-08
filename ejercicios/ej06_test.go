package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListIterator_Next(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	iterator := list.Iterator()

	// Test Next
	element, err := iterator.Next()
	assert.NoError(t, err)
	assert.Equal(t, 1, element)

	element, err = iterator.Next()
	assert.NoError(t, err)
	assert.Equal(t, 2, element)

	element, err = iterator.Next()
	assert.NoError(t, err)
	assert.Equal(t, 3, element)

	// Test Next when no more elements
	_, err = iterator.Next()
	assert.Error(t, err)
}

func TestListIterator_HasNext(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)

	iterator := list.Iterator()

	// Test HasNext
	assert.True(t, iterator.HasNext())
	iterator.Next()
	assert.True(t, iterator.HasNext())
	iterator.Next()
	assert.False(t, iterator.HasNext())
}

func TestListIterator_Previous(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	iterator := list.Iterator()

	// Move to the end of the list
	iterator.Next()
	iterator.Next()
	iterator.Next()

	// Test Previous
	element, err := iterator.Previous()
	assert.NoError(t, err)
	assert.Equal(t, 3, element)

	element, err = iterator.Previous()
	assert.NoError(t, err)
	assert.Equal(t, 2, element)

	element, err = iterator.Previous()
	assert.NoError(t, err)
	assert.Equal(t, 1, element)

	// Test Previous when no more elements
	_, err = iterator.Previous()
	assert.Error(t, err)
}

func TestListIterator_HasPrevious(t *testing.T) {
	list := NewList[int]()
	list.Append(1)
	list.Append(2)

	iterator := list.Iterator()

	// Move to the end of the list
	iterator.Next()
	iterator.Next()

	// Test HasPrevious
	assert.True(t, iterator.HasPrevious())
	iterator.Previous()
	assert.True(t, iterator.HasPrevious())
	iterator.Previous()
	assert.False(t, iterator.HasPrevious())
}
