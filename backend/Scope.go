package main

import (
	"Proyecto1/parser"
	"fmt"
)

const (
	String    = iota // 0
	Bool             // 1
	Character        // 2
	Int              // 3
	Float            // 4
	Error            // 5
	Void             // 6
)

type Valor struct {
	Valor interface{} // Valor
	Tipo  int         // Tipo de dato
}

type Variable struct {
	ID        string
	Tipo      int
	Valor     interface{}
	Constante bool
	Posicion  int
	enHeap    bool
	etqTrue   string
	etqFalse  string
	Linea     int
	Columna   int
}

type Scope struct {
	Nombre       string
	Anterior     *Scope
	Variables    map[string]Variable
	Funciones    map[string]Funcion
	Size         map[string]int
	etqSalida    string
	etqCiclo     string
	Tipo         int
	ValorRetorno string
}

// Definiciones de tipos
type Funcion struct {
	ID         string
	Parametros parser.IParametrosContext
	Bloque     parser.IBlockContext //interface{}
	TipoReturn int
	Linea      int
	Columna    int
}

type Parametro struct {
	ID_Externo string
	ID_Interno string
	Tipo       int
	Referencia bool
}

type Error_ struct {
	Tipo    string
	Linea   string
	Columna string
	Mensaje string
}

// Lista de errores
var listaErrores []Error_

func NuevoScope(ant Scope, nombre string) Scope {
	tam := make(map[string]int)
	tam["Size"] = 0
	return Scope{Anterior: &ant, Variables: make(map[string]Variable), Nombre: nombre, Size: tam, etqSalida: ant.etqSalida, etqCiclo: ant.etqCiclo}
}

func (s *Scope) totalSize() int {
	actual := s
	total := 0
	for actual != nil {
		total += actual.Size["Size"]
		actual = actual.Anterior
	}
	return total
}

func (s *Scope) agregarVariable(variable Variable) int {
	// Verificar si una variable ya existe en el ambito actual
	if _, existe := s.Variables[variable.ID]; existe {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(variable.Linea),
			Columna: fmt.Sprint(variable.Columna),
			Mensaje: fmt.Sprintf("La variable '%v' ya existe\n", variable.ID),
		})
		//fmt.Println("YA EXISE LA VARIABLE")
		return -1
	}
	variable.Posicion = s.totalSize()
	/* 	fmt.Println("POSICION: ", variable.Posicion)
	fmt.Println("VALOR: ", variable.Valor) */

	s.Variables[variable.ID] = variable // Agrega la variable al scope actual
	s.Size["Size"] = s.Size["Size"] + 1 // Aumenta el tamaño del scope actual
	return variable.Posicion
}

func (s *Scope) encontrarVariable(id string, linea int, columna int) interface{} {
	// Busca la variable en el ambito actual
	variable, existe := s.Variables[id]
	// Verificar si la variables existe en el ambito actual
	if existe {
		return variable
	} else {
		if s.Anterior != nil {
			return s.Anterior.encontrarVariable(id, linea, columna)
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(linea),
			Columna: fmt.Sprint(columna),
			Mensaje: fmt.Sprintf("La variable '%s' no existe\n", id),
		})
		return nil
	}
}

func (s *Scope) modificarVariable(id string, expr Valor, linea int, columna int) int {
	// Busca la variable en el ambito actual
	variable, existe := s.Variables[id]
	// Verificar si la variables existe en el ambito actual
	if existe {
		if !variable.Constante {
			if variable.Tipo == expr.Tipo {
				variable.Valor = expr.Valor
				s.Variables[id] = variable // Edita la variable en el scope
				return variable.Posicion   //Se edito la variable
			}
			variable.Valor = nil
			s.Variables[id] = variable // Edita la variable con valor nil
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(linea),
				Columna: fmt.Sprint(columna),
				Mensaje: fmt.Sprintf("El valor a asignar a la variable '%s' no es del mismo tipo, se le asignó 'nil'\n", id),
			})
			return -1
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(linea),
			Columna: fmt.Sprint(columna),
			Mensaje: fmt.Sprintf("'%s' es una constante, no puede cambiar su valor\n", id),
		})
		return -1
	} else {
		if s.Anterior != nil {
			return s.Anterior.modificarVariable(id, expr, linea, columna)
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(linea),
			Columna: fmt.Sprint(columna),
			Mensaje: fmt.Sprintf("La variable '%s' no existe\n", id),
		})
		return -1
	}
}
