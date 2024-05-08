package ejercicios

import "fmt"

type Matriz struct {
	Filas    int
	Columnas int
	Datos    [][]int
}

// NewMatriz crea una nueva matriz con el número de filas y columnas especificadas.
// Inicializa todos los elementos de la matriz a 0.
func NewMatriz(filas, columnas int) *Matriz {
	matriz := &Matriz{
		Filas:    filas,
		Columnas: columnas,
		Datos:    make([][]int, filas),
	}

	for i := range filas {
		matriz.Datos[i] = make([]int, columnas)
	}

	return matriz
}

// Set establece el valor en la posición (fila, columna) de la matriz.
func (m *Matriz) Set(fila, columna, valor int) {
	if fila < 0 || fila >= m.Filas || columna < 0 || columna >= m.Columnas {
		panic("índice fuera de rango")
	}
	m.Datos[fila][columna] = valor
}

// Get obtiene el valor en la posición (fila, columna) de la matriz.
func (m *Matriz) Get(fila, columna int) int {
	if fila < 0 || fila >= m.Filas || columna < 0 || columna >= m.Columnas {
		panic("índice fuera de rango")
	}
	return m.Datos[fila][columna]
}

func (m *Matriz) String() string {
	s := ""
	for i := range m.Filas {
		for j := range m.Columnas {
			s += fmt.Sprintf("%d ", m.Datos[i][j])
		}
		s += "\n"
	}
	return s
}
