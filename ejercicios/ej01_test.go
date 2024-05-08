package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArchivosMusica(t *testing.T) {
	original := NewFileMp3("el gato volador.mp3")

	var adaptador ArchivoDeMusica = NewArchivoDeMusicaAdapter(original)

	require.Implements(t, (*ArchivoDeMusica)(nil), adaptador, "El adaptador debe implementar la interfaz ArchivoDeMusica")
	assert.Equal(t, "Reproduciendo archivo MP3. Nombre: el gato volador.mp3", adaptador.Reproducir())
}
