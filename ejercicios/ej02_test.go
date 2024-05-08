package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjemplo(t *testing.T) {
	parte1 := NewParteSimple("tornillo", 5)
	parte2 := NewParteSimple("tuerca", 3)
	// parte3 := NewParteSimple("ruleman", 50)
	parte4 := NewParteSimple("varilla", 5)

	assert.Equal(t, float32(5.0), parte1.ObtenerPrecio())
	assert.Equal(t, "tuerca", parte2.ObtenerDescripcion())

	// tiene un tornillo, dos tuercas y una varilla
	ensamble := NewEnsamble()
	ensamble.AgregarParte(parte1)
	ensamble.AgregarParte(parte2)
	ensamble.AgregarParte(parte2)
	ensamble.AgregarParte(parte4)

	assert.Equal(t, float32(16.0), ensamble.ObtenerPrecio())
	assert.Equal(t, "tornillo, tuerca, tuerca, varilla", ensamble.ObtenerDescripcion())
}
