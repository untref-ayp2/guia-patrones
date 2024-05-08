package ejercicios

type Parte interface {
	ObtenerPrecio() float32
	ObtenerDescripcion() string
}

type ParteSimple struct {
	// Implementar
}

func NewParteSimple(descripcion string, precio float32) *ParteSimple {
	// Implementar
	return nil
}

func (p *ParteSimple) ObtenerPrecio() float32 {
	// Implementar
	return 0.0
}

func (p *ParteSimple) ObtenerDescripcion() string {
	// Implementar
	return ""
}

type Ensamble struct {
	// Implementar
}

func NewEnsamble() *Ensamble {
	// Implementar
	return nil
}

func (e *Ensamble) AgregarParte(p Parte) {
	// Implementar
}

func (e *Ensamble) ObtenerPrecio() float32 {
	// Implementar
	return 0.0
}

func (e *Ensamble) ObtenerDescripcion() string {
	// Implementar
	return ""
}
