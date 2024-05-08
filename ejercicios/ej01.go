package ejercicios

type FileMp3 struct {
	name string
}

func NewFileMp3(name string) *FileMp3 {
	return &FileMp3{name: name}
}

func (f *FileMp3) PlayMp3() string {
	return "Reproduciendo archivo MP3. Nombre: " + f.name
}

type ArchivoDeMusica interface {
	Reproducir() string
}

type ArchivoDeMusicaAdapter struct {
	// Implementar
}

func NewArchivoDeMusicaAdapter(original *FileMp3) *ArchivoDeMusicaAdapter {
	// Implementar
	return nil
}

func (a *ArchivoDeMusicaAdapter) Reproducir() string {
	// Implementar
	return ""
}
