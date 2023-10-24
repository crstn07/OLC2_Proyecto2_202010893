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
	Size         int
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
	return Scope{Anterior: &ant, Variables: make(map[string]Variable), Nombre: nombre}
}

func (s *Scope) totalSize() int {
	actual := s
	total := 0
	if actual != nil {
		total += actual.Size
		actual = actual.Anterior
	}
	return total
}

func (s *Scope) agregarVariable(variable Variable) bool {
	// Verificar si una variable ya existe en el ambito actual
	if _, existe := s.Variables[variable.ID]; existe {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("La variable '%v' ya existe\n", variable.ID),
		})
		//return Valor{Valor: fmt.Sprintf("Error: La variable '%v' ya existe\n", variable.ID), Tipo: Error}
		return false
	}
	variable.Posicion = s.totalSize()
	s.Variables[variable.ID] = variable // Agrega la variable al scope actual
	s.Size++                            // Aumenta el tamaño del scope actual
	return true
}

func (s *Scope) encontrarVariable(id string) interface{} {
	// Busca la variable en el ambito actual
	variable, existe := s.Variables[id]
	// Verificar si la variables existe en el ambito actual
	if existe {
		//return Valor{Valor: variable.Valor, Tipo: variable.Tipo}
		return variable
	} else {
		if s.Anterior != nil {
			return s.Anterior.encontrarVariable(id)
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("La variable '%s' no existe\n", id),
		})
		//return Valor{fmt.Sprintf("Error: La variable '%s' no existe\n", id), Error, false, false, false, nil, 0}
		return nil
	}
}

func (s *Scope) modificarVariable(id string, expr Valor) bool {
	// Busca la variable en el ambito actual
	variable, existe := s.Variables[id]
	// Verificar si la variables existe en el ambito actual
	if existe {
		if !variable.Constante {
			if variable.Tipo == expr.Tipo {
				variable.Valor = expr.Valor
				s.Variables[id] = variable // Edita la variable en el scope
				return true                //Se edito la variable
			}
			variable.Valor = nil
			s.Variables[id] = variable // Edita la variable con valor nil
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   "",
				Columna: "",
				Mensaje: fmt.Sprintf("El valor a asignar a la variable '%s' no es del mismo tipo, se le asignó 'nil'\n", id),
			})
			return false
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("'%s' es una constante, no puede cambiar su valor\n", id),
		})
		return false
	} else {
		if s.Anterior != nil {
			return s.Anterior.modificarVariable(id, expr)
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("La variable '%s' no existe\n", id),
		})
		return false
	}
}
