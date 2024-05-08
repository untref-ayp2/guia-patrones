package ejercicios

type Componente interface {
	Resolver() float32
	String() string
}

type Operador struct {
	// Implementar
}

func NewOperador(operando string, izq, der Componente) *Operador {
	// Implementar
	return nil
}

func (o *Operador) String() string {
	// Implementar
	return ""
}

func (o *Operador) Resolver() float32 {
	// Implementar
	return 0.0
}

type Operando struct {
	// Implementar
}

func NewOperando(valor float32) *Operando {
	// Implementar
	return nil
}

func (o *Operando) String() string {
	// Implementar
	return ""
}

func (o *Operando) Resolver() float32 {
	// Implementar
	return 0.0
}
