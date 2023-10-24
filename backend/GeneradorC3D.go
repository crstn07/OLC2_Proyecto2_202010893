package main

import "fmt"

type Generador struct {
	Temporal      int
	Etiqueta      int
	Codigo        string
	CodigoFuncion string
	CodigoFinal   string
	printStrFlag  bool
	esFuncion     bool
}

// Generar de Temporales
func (g *Generador) nuevoTemporal() string {
	g.Temporal++
	return fmt.Sprint("t", g.Temporal)
}

// Generador de Etiquetas
func (g *Generador) nuevaEtiqueta() string {
	g.Etiqueta++
	return fmt.Sprint("L", g.Etiqueta)
}

// Generador del Encabezado
func (g *Generador) crearEncabezado() string {
	encabezado := `#include <stdio.h>

float stack[100000];	//Stack
float heap[100000];	//Heap
float P;		//Puntero del Stack
float H;		//Puntero del Heap
`
	if g.Temporal > 0 {
		encabezado += "float "
		for i := 1; i <= g.Temporal; i++ {
			encabezado += fmt.Sprint("t", i)
			if i < g.Temporal {
				encabezado += ", "
			}
		}
		encabezado += ";	//Temporales\n\n"
	}

	return encabezado
}

// Generador del Codigo Final
func (g *Generador) crearCodigoFinal() string {
	codigo := g.crearEncabezado()
	codigo += g.CodigoFuncion
	codigo += "\nint main(){\n"
	codigo += g.Codigo
	codigo += "\n    return 0;\n"
	codigo += "}\n"
	return codigo
}

func (g *Generador) agregarCodigo(codigo string) {
	if g.esFuncion {
		g.CodigoFuncion += "    " + codigo
	} else {
		g.Codigo += "    " + codigo
	}
}

func (g *Generador) Goto(etiqueta string) {
	g.agregarCodigo(fmt.Sprintf("goto %v;\n", etiqueta))
}

func (g *Generador) imprimirEtiqueta(etiqueta string) {
	g.agregarCodigo(fmt.Sprintf("%v:\n", etiqueta))
}

func (g *Generador) inicioFuncion(nombre string) {
	g.esFuncion = true
	g.CodigoFuncion += fmt.Sprint("void ", nombre, "() {\n")
}

func (g *Generador) finFuncion() {
	g.CodigoFuncion += "    return;\n}\n"
	g.esFuncion = false
}

// res = left op right
func (g *Generador) Expresion(res string, left string, op string, right string) {
	g.agregarCodigo(fmt.Sprintf("%v = %v%v%v;\n", res, left, op, right))
}

// temporal = stack[posicion]
func (g *Generador) getStack(temporal string, posicion string) {
	g.agregarCodigo(fmt.Sprintf("%v = stack[(int)%v];\n", temporal, posicion))
}

// stack[posicion] = valor
func (g *Generador) setStack(posicion string, valor string) {
	g.agregarCodigo(fmt.Sprintf("stack[(int)%v] = %v;\n", posicion, valor))
}

// temporal = heap[posicion]
func (g *Generador) getHeap(temporal string, posicion string) {
	g.agregarCodigo(fmt.Sprintf("%v = heap[(int)%v];\n", temporal, posicion))
}

// heap[posicion] = valor
func (g *Generador) setHeap(posicion string, valor string) {
	g.agregarCodigo(fmt.Sprintf("heap[(int)%v] = %v;\n", posicion, valor))
}

// H = H + 1
func (g *Generador) nextHeap() {
	g.agregarCodigo("H = H + 1;\n")
}

// if (left op right) goto etiqueta;
func (g *Generador) If(left string, op string, right string, etiqueta string) {
	g.agregarCodigo(fmt.Sprintf("if (%v %v %v) goto %v;\n", left, op, right, etiqueta))
}

// P = P + size;
func (g *Generador) nuevoAmbito(size string) {
	g.agregarCodigo(fmt.Sprintf("P = P + %v;\n", size))
}

// P = P - size;
func (g *Generador) getAmbito(size string) {
	g.agregarCodigo(fmt.Sprintf("P = P - %v;\n", size))
}

// nombre();
func (g *Generador) getFuncion(nombre string) {
	g.agregarCodigo(fmt.Sprintf("%v();\n", nombre))
}

// printf("%tipo", casteo valor);
func (g *Generador) Printf(tipo string, cast string, valor interface{}) {
	g.agregarCodigo(fmt.Sprintf("printf(\"%%%v\", %v%v);\n", tipo, cast, valor.(string)))
}

func (g *Generador) Comentario(comentario string) {
	g.agregarCodigo(fmt.Sprintf("//%v\n", comentario))
}

func (g *Generador) imprimirString() {
	if g.printStrFlag {
		return
	}

	g.printStrFlag = true
	//this.esFuncion = true;

	g.inicioFuncion("imprimirString")
	salida := g.nuevaEtiqueta()
	ciclo := g.nuevaEtiqueta()

	tmpP := g.nuevoTemporal()
	tmpH := g.nuevoTemporal()

	g.Expresion(tmpP, "P", "+", "1")
	g.getStack(tmpH, tmpP)

	tmp := g.nuevoTemporal()
	g.imprimirEtiqueta(ciclo)
	g.getHeap(tmp, tmpH)
	g.If(tmp, "==", "-1", salida)
	g.Printf("c", "(int)", tmp)
	g.Expresion(tmpH, tmpH, "+", "1")
	g.Goto(ciclo)
	g.imprimirEtiqueta(salida)
	g.finFuncion()
	//g.isFuncion = false;
}
