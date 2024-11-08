package main

import (
	"Proyecto1/parser"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rs/cors"
)

type Visitor struct {
	parser.SwiftGrammarVisitor
}

var tablaSimbolos Scope
var ts_funciones map[string]Scope

func (v *Visitor) Visit(tree antlr.ParseTree, ts Scope, generador *Generador) interface{} {
	switch val := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(val, ts, generador)
	case *parser.BlockContext:
		return v.VisitBlock(val, ts, generador)
	case *parser.InstrContext:
		return v.VisitInstr(val, ts, generador)
	case *parser.DeclaracionContext:
		return v.VisitDeclaracion(val, ts, generador)
	case *parser.DeclaracionTipoValorContext:
		return v.VisitDeclaracionTipoValor(val, ts, generador)
	case *parser.DeclaracionTipoContext:
		return v.VisitDeclaracionTipo(val, ts, generador)
	case *parser.DeclaracionValorContext:
		return v.VisitDeclaracionValor(val, ts, generador)
	case *parser.AsignacionContext:
		return v.VisitAsignacion(val, ts, generador)
	case *parser.Print_instrContext:
		return v.VisitPrint_instr(val, ts, generador)
	case *parser.If_instrContext:
		return v.VisitIf_instr(val, ts, generador)
	case *parser.IfContext:
		return v.VisitIf(val, ts, generador)
	case *parser.IfElseContext:
		return v.VisitIfElse(val, ts, generador)
	case *parser.ElseIfContext:
		return v.VisitElseIf(val, ts, generador)
	case *parser.Switch_instrContext:
		return v.VisitSwitch(val, ts, generador)
	case *parser.While_instrContext:
		return v.VisitWhile_instr(val, ts, generador)
	case *parser.For_instrContext:
		return v.VisitFor(val, ts, generador)
	case *parser.GuardContext:
		return v.VisitGuard(val, ts, generador)
	case *parser.FuncionContext:
		return v.VisitFuncion(val, ts, generador)
	case *parser.Llamada_funcContext:
		return v.VisitLlamada_func(val, ts, generador)
	case *parser.LlamadaFuncExprContext:
		return v.VisitLlamadaFuncExpr(val, ts, generador)
	case *parser.Dec_vectorContext:
		return v.VisitDeclaracion_vector(val, ts, generador)
	case *parser.Dec_matrizContext:
		return v.VisitDeclaracion_matriz(val, ts, generador)
	case *parser.Modificacion_vectorContext:
		return v.VisitModificacion_vector(val, ts, generador)
	case *parser.AppendContext:
		return v.VisitAppend(val, ts, generador)
	case *parser.RemoveLastContext:
		return v.VisitRemoveLast(val, ts, generador)
	case *parser.RemoveAtContext:
		return v.VisitRemoveAt(val, ts, generador)
	case *parser.UmenosExprContext:
		return v.VisitUMenosExpr(val, ts, generador)
	case *parser.NotExprContext:
		return v.VisitNotExpr(val, ts, generador)
	case *parser.OpExprContext:
		return v.VisitOpExpr(val, ts, generador)
	case *parser.ParExprContext:
		return v.VisitParExpr(val, ts, generador)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val, ts, generador)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val, ts, generador)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val, ts, generador)
	case *parser.FloatExprContext:
		return v.VisitFloatExpr(val, ts, generador)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val, ts, generador)
	case *parser.CharExprContext:
		return v.VisitCharExpr(val, ts, generador)
	case *parser.NilExprContext:
		return v.VisitNilExpr(val, ts, generador)
	case *parser.IntCastExprContext:
		return v.VisitIntCastExpr(val, ts, generador)
	case *parser.FloatCastExprContext:
		return v.VisitFloatCastExpr(val, ts, generador)
	case *parser.StringCastExprContext:
		return v.VisitStringCastExpr(val, ts, generador)
	case *parser.AccesoVectorContext:
		return v.VisitAccesoVector(val, ts, generador)
	case *parser.IsEmptyContext:
		return v.VisitIsEmpty(val, ts, generador)
	case *parser.CountContext:
		return v.VisitCount(val, ts, generador)
	case *parser.AccesoMatrizContext:
		return v.VisitAccesoMatriz(val, ts, generador)
	default:
		return fmt.Sprintf("Contexto No reconocido: %v", val)
	}

}

func (v *Visitor) VisitProg(ctx *parser.ProgContext, ts Scope, generador *Generador) interface{} {
	return v.Visit(ctx.Block(), ts, generador)
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext, ts Scope, generador *Generador) interface{} {
	out := ""
	for i := 0; ctx.Instr(i) != nil; i++ {
		v.Visit(ctx.Instr(i), ts, generador)
	}
	return Valor{Valor: out}
}

func (v *Visitor) VisitInstr(ctx *parser.InstrContext, ts Scope, generador *Generador) interface{} {
	if ctx.Declaracion() != nil {
		return v.Visit(ctx.Declaracion(), ts, generador)
	}
	if ctx.Asignacion() != nil {
		return v.Visit(ctx.Asignacion(), ts, generador)
	}
	if ctx.If_instr() != nil {
		return v.Visit(ctx.If_instr(), ts, generador)
	}
	if ctx.Print_instr() != nil {
		return v.Visit(ctx.Print_instr(), ts, generador)
	}
	if ctx.Switch_instr() != nil {
		return v.Visit(ctx.Switch_instr(), ts, generador)
	}
	if ctx.While_instr() != nil {
		return v.Visit(ctx.While_instr(), ts, generador)
	}
	if ctx.For_instr() != nil {
		return v.Visit(ctx.For_instr(), ts, generador)
	}
	if ctx.Guard() != nil {
		return v.Visit(ctx.Guard(), ts, generador)
	}
	if ctx.CONTINUE() != nil {
		generador.Comentario("Instrucción Continue")
		generador.Goto(ts.etqCiclo)
		return Valor{}
		//return Valor{Valor: 999999999, Continue: true}
	}
	if ctx.BREAK() != nil {
		generador.Comentario("Instrucción Break")
		generador.Goto(ts.etqSalida)
		return Valor{}
		//return Valor{Valor: 999999999, Break: true}
	}
	if ctx.RETURN() != nil {
		if ctx.Expr() != nil {
			retorno := v.Visit(ctx.Expr(), ts, generador).(Valor)
			fmt.Println("Retorno: ", retorno.Valor, retorno.Tipo)
			generador.Comentario("Instrucción Return")
			generador.setStack("P", fmt.Sprint(retorno.Valor))
			//generador.Goto(ts.etqSalida)
			return Valor{Valor: retorno.Valor, Tipo: retorno.Tipo}
		}
		return Valor{}
	}
	if ctx.Funcion() != nil {
		return v.Visit(ctx.Funcion(), ts, generador)
	}
	if ctx.Llamada_func() != nil {
		return v.Visit(ctx.Llamada_func(), ts, generador)
	}
	if ctx.Dec_vector() != nil {
		return v.Visit(ctx.Dec_vector(), ts, generador)
	}
	if ctx.Dec_matriz() != nil {
		return v.Visit(ctx.Dec_matriz(), ts, generador)
	}
	if ctx.Copia_vector() != nil {
		return v.Visit(ctx.Copia_vector(), ts, generador)
	}
	if ctx.Modificacion_vector() != nil {
		return v.Visit(ctx.Modificacion_vector(), ts, generador)
	}
	if ctx.Append_() != nil {
		return v.Visit(ctx.Append_(), ts, generador)
	}
	if ctx.RemoveLast() != nil {
		return v.Visit(ctx.RemoveLast(), ts, generador)
	}
	if ctx.RemoveAt() != nil {
		return v.Visit(ctx.RemoveAt(), ts, generador)
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: "Instrucción no reconocida",
	})
	return Valor{Valor: "Error: Instrucción no reconocida\n", Tipo: Error}
}

func (v *Visitor) VisitDeclaracion(ctx *parser.DeclaracionContext, ts Scope, generador *Generador) interface{} {
	fmt.Println("Declaración no valida")
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: "Declaración no válida",
	})
	return Valor{Valor: "Error: declaración no válida\n", Tipo: Error}
}

func (v *Visitor) VisitDeclaracionTipoValor(ctx *parser.DeclaracionTipoValorContext, ts Scope, generador *Generador) interface{} {
	expr := v.Visit(ctx.Expr(), ts, generador).(Valor) // Obtiene el valor de la expresión y su tipo
	constante := true
	if ctx.LET() == nil {
		constante = false
	}
	var tipo int // Variable para almacenar el tipo de la variable
	switch ctx.Tipo().GetText() {
	case "String":
		tipo = String
	case "Bool":
		tipo = Bool
	case "Character":
		tipo = Character
	case "Int":
		tipo = Int
	case "Float":
		tipo = Float
	}

	/* 	if ctx.Tipo().Tipo() != nil && reflect.TypeOf(expr.Valor).Kind() == reflect.Slice {
		valores := make([]interface{}, len(expr.Valor.([]interface{})))
		copy(valores, expr.Valor.([]interface{}))
		fmt.Println("******************************TIPO DEL VECTOR DECLARADO COMO VARIABLE***************************\n", tipo)
		var tipo int // Variable para almacenar el tipo de la variable
		switch ctx.Tipo().Tipo().GetText() {
		case "String":
			tipo = String
		case "Bool":
			tipo = Bool
		case "Character":
			tipo = Character
		case "Int":
			tipo = Int
		case "Float":
			tipo = Float
		}
		//return ts.agregarVariable(Variable{ctx.ID().GetText(), tipo, valores, false, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false,1})
	} */
	var pos int
	if tipo == expr.Tipo {
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), tipo, expr.Valor, constante, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1}) // Agrega la variable al scope
	} else if tipo == Float && expr.Tipo == Int {
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), Float, expr.Valor, constante, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1}) // Agrega la variable al scope de tipo float
	} else if tipo == Character && expr.Tipo == String && len(expr.Valor.(string)) == 1 {
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), Character, expr.Valor, constante, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1}) // Agrega la variable al scope de tipo Character
	} else {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
			Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
			Mensaje: fmt.Sprintf("El valor de la expresión '%v' no coincide con el tipo de la variable '%v', no se puede crear.\n", expr.Valor, ctx.ID().GetText()),
		})
		return Valor{}
	}
	if pos == -1 {
		generador.Comentario("Error: La variable ya existe")
		return Valor{}
	}
	generador.Comentario("Declaración de variable: " + ctx.ID().GetText())
	generador.setStack(fmt.Sprint(pos), fmt.Sprint(expr.Valor))
	generador.agregarCodigo("\n")
	return Valor{}
}

func (v *Visitor) VisitDeclaracionTipo(ctx *parser.DeclaracionTipoContext, ts Scope, generador *Generador) interface{} {
	var pos int
	if ctx.LET() == nil {
		var tipo int // Variable para almacenar el tipo de la variable
		switch ctx.Tipo().GetText() {
		case "String":
			tipo = String
		case "Bool":
			tipo = Bool
		case "Character":
			tipo = Character
		case "Int":
			tipo = Int
		case "Float":
			tipo = Float
		}
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), tipo, "nil", false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1}) // Agrega la variable al scope con valor nil
		if pos == -1 {
			generador.Comentario("Error: La variable ya existe")
			return Valor{}
		}
		generador.Comentario("Declaración de variable: " + ctx.ID().GetText())
		generador.setStack(fmt.Sprint(pos), fmt.Sprint(999999999))
		generador.agregarCodigo("\n")
		return Valor{}
	} else {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
			Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
			Mensaje: fmt.Sprintf("No se puede declarar la constante '%v' sin valor\n", ctx.ID().GetText()),
		})
		return Valor{}
	}
}

func (v *Visitor) VisitDeclaracionValor(ctx *parser.DeclaracionValorContext, ts Scope, generador *Generador) interface{} {
	expr := v.Visit(ctx.Expr(), ts, generador).(Valor) // Obtiene el valor de la expresión y su tipo
	var pos int
	if expr.Tipo != Error {
		constante := true
		if ctx.LET() == nil {
			constante = false
		}
		/* 		if expr.Tipo == String {
			heap = true
		} */
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), expr.Tipo, expr.Valor, constante, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1})
		generador.Comentario("Declaración de variable: " + ctx.ID().GetText())
		generador.setStack(fmt.Sprint(pos), fmt.Sprint(expr.Valor))
		generador.agregarCodigo("\n")
		return Valor{}
	} else {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
			Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
			Mensaje: fmt.Sprintf("No se puede declarar la variable '%v'\n       > %v\n", ctx.ID().GetText(), expr.Valor.(string)),
		})
		return Valor{}
	}
}

func (v *Visitor) VisitAsignacion(ctx *parser.AsignacionContext, ts Scope, generador *Generador) interface{} {
	expr := v.Visit(ctx.Expr(), ts, generador).(Valor) // Obtiene el valor de la expresión y su tipo
	mas := ctx.MAS()
	//menos := ctx.MENOS()
	variable := ts.encontrarVariable(ctx.ID().GetText(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	if variable != nil {
		var res Valor
		temp := generador.nuevoTemporal()
		if mas != nil {
			if variable.(Variable).Tipo == String && expr.Tipo == String {
				generador.Comentario("Concatenación de Strings")
				generador.Expresion(temp, "H", "", "") //Se almacena el inicio de la nueva cadena

				//Primer Palabra
				generador.Comentario("Primera Palabra")
				ciclo := generador.nuevaEtiqueta()
				salida := generador.nuevaEtiqueta()
				caracter := generador.nuevoTemporal()

				generador.imprimirEtiqueta(ciclo)
				generador.getHeap(caracter, fmt.Sprint(variable.(Variable).Valor))
				generador.If(caracter, "==", "-1", salida)
				generador.setHeap("H", caracter)
				generador.nextHeap()
				generador.Expresion(fmt.Sprint(variable.(Variable).Valor), fmt.Sprint(variable.(Variable).Valor), "+", "1")
				generador.Goto(ciclo)
				generador.imprimirEtiqueta(salida)

				//Segunda Palabra
				generador.Comentario("Segunda Palabra")
				ciclo = generador.nuevaEtiqueta()
				salida = generador.nuevaEtiqueta()
				caracter = generador.nuevoTemporal()

				generador.imprimirEtiqueta(ciclo)
				generador.getHeap(caracter, fmt.Sprint(expr.Valor))
				generador.If(caracter, "==", "-1", salida)
				generador.setHeap("H", caracter)
				generador.nextHeap()
				generador.Expresion(fmt.Sprint(expr.Valor), fmt.Sprint(expr.Valor), "+", "1")
				generador.Goto(ciclo)
				generador.imprimirEtiqueta(salida)

				//Fin de la concatenacion
				generador.setHeap("H", "-1")
				generador.nextHeap()
				generador.agregarCodigo("\n")
				res = Valor{Valor: temp, Tipo: String}
			} else if variable.(Variable).Tipo == Int && expr.Tipo == Int {
				temp2 := generador.nuevoTemporal()
				temp3 := generador.nuevoTemporal()
				generador.Expresion(temp2, "P", "+", fmt.Sprint(variable.(Variable).Posicion))
				generador.getStack(temp, temp2)
				generador.Expresion(temp3, "(int)"+temp, "+", "(int)"+fmt.Sprint(expr.Valor))
				res = Valor{Valor: temp3, Tipo: Int}
			} else if variable.(Variable).Tipo == Float && expr.Tipo == Float || (variable.(Variable).Tipo == Float && expr.Tipo == Int) || variable.(Variable).Tipo == Int && expr.Tipo == Float {
				temp2 := generador.nuevoTemporal()
				generador.getStack(temp, fmt.Sprint(variable.(Variable).Posicion))
				generador.Expresion(temp2, temp, "+", fmt.Sprint(expr.Valor))
				res = Valor{Valor: temp2, Tipo: Float}
			} else {
				listaErrores = append(listaErrores, Error_{
					Tipo:    "SEMANTICO",
					Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
					Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
					Mensaje: fmt.Sprintf("No se puede realizar la suma entre %v y %v, combinación de tipos no válida", variable.(Variable).Valor, expr.Valor),
				})
				generador.Comentario(fmt.Sprintf("Error: No se puede realizar la suma entre %v y %v, combinación de tipos no válida", variable.(Variable).Valor, expr.Valor))
				return Valor{Valor: 999999999, Tipo: Error}
			}
			pos := ts.modificarVariable(ctx.ID().GetText(), res, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
			fmt.Println("POSICION: ", pos)
			if pos != -1 {
				tmp1 := generador.nuevoTemporal()
				generador.Comentario("Asignación de variable: " + ctx.ID().GetText())
				generador.Expresion(tmp1, "P", "+", fmt.Sprint(pos))
				generador.setStack(tmp1, fmt.Sprint(res.Valor))
				generador.agregarCodigo("\n")
			}
			return Valor{}
		}
		/* 		if menos != nil {
			if variable.Tipo != Error {
				if variable.Tipo == Int && expr.Tipo == Int {
					res = variable.Valor.(int) - expr.Valor.(int)
				} else if variable.Tipo == Float && expr.Tipo == Float {
					res = variable.Valor.(float64) - expr.Valor.(float64)
				} else if variable.Tipo == Float && expr.Tipo == Int {
					res = variable.Valor.(float64) - float64(expr.Valor.(int))
					expr.Tipo = Float
				} else if variable.Tipo == Int && expr.Tipo == Float {
					res = float64(variable.Valor.(int)) - expr.Valor.(float64)
					expr.Tipo = Float
				} else {
					listaErrores = append(listaErrores, Error_{
						Tipo:    "SEMANTICO",
						Linea:   "",
						Columna: "",
						Mensaje: fmt.Sprintf("No se puede realizar la resta entre %v y %v, combinación de tipos no válida\n", expr.Valor, variable.Valor),
					})
					return Valor{Valor: fmt.Sprintf("Error: No se puede realizar la resta entre %v y %v, combinación de tipos no válida\n", expr.Valor, variable.Valor), Tipo: Error}
				}
				return ts.modificarVariable(ctx.ID().GetText(), Valor{res, expr.Tipo}, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
			}
			return variable
		} */

		pos := ts.modificarVariable(ctx.ID().GetText(), expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
		if pos != -1 {
			tmp1 := generador.nuevoTemporal()
			generador.Comentario("Asignación de variable: " + ctx.ID().GetText())
			generador.Expresion(tmp1, "P", "+", fmt.Sprint(pos))
			generador.setStack(tmp1, fmt.Sprint(expr.Valor))
			generador.agregarCodigo("\n")
		}
		return Valor{}
	}
	return Valor{}
}

func (v *Visitor) VisitPrint_instr(ctx *parser.Print_instrContext, ts Scope, generador *Generador) interface{} {
	for i := 0; ctx.Expr(i) != nil; i++ {
		valor := v.Visit(ctx.Expr(i), ts, generador).(Valor)
		generador.Comentario("Instrucción Print")
		True := generador.nuevaEtiqueta()
		False := generador.nuevaEtiqueta()
		Salida := generador.nuevaEtiqueta()
		generador.If(fmt.Sprint(valor.Valor), "==", "999999999", True)
		generador.Goto(False)
		generador.imprimirEtiqueta(True)
		generador.Printf("c", "", "110") //n
		generador.Printf("c", "", "105") //i
		generador.Printf("c", "", "108") //l
		generador.Printf("c", "", "32")  //espacio
		generador.Goto(Salida)
		generador.imprimirEtiqueta(False)
		if valor.Tipo == Int {
			generador.Printf("d", "(int)", fmt.Sprintf("%v", valor.Valor))
		} else if valor.Tipo == Float {
			generador.Printf("f", "", fmt.Sprintf("%v", valor.Valor))
		} else if valor.Tipo == Character {
			generador.Printf("c", "(int)", fmt.Sprintf("%v", valor.Valor))
		} else if valor.Tipo == String {
			generador.Comentario("Imprimir String")
			generador.printStrFlag = true

			temp := generador.nuevoTemporal()

			generador.Expresion(temp, "P", "+", fmt.Sprint(ts.totalSize()))
			generador.Expresion(temp, temp, "+", "1")
			generador.setStack(temp, fmt.Sprint(valor.Valor))
			generador.nuevoAmbito(fmt.Sprint(ts.totalSize()))
			generador.getFuncion("imprimirString")
			generador.getAmbito(fmt.Sprint(ts.totalSize()))
		} else if valor.Tipo == Bool {
			generador.Comentario("Imprimir Bool")
			Salida := generador.nuevaEtiqueta()
			True := generador.nuevaEtiqueta()
			False := generador.nuevaEtiqueta()
			generador.If(fmt.Sprint(valor.Valor), "==", "1", True)
			generador.Goto(False)
			generador.imprimirEtiqueta(True)
			generador.Printf("c", "", "116")
			generador.Printf("c", "", "114")
			generador.Printf("c", "", "117")
			generador.Printf("c", "", "101")
			generador.Goto(Salida)
			generador.imprimirEtiqueta(False)
			generador.Printf("c", "", "102")
			generador.Printf("c", "", "97")
			generador.Printf("c", "", "108")
			generador.Printf("c", "", "115")
			generador.Printf("c", "", "101")
			generador.imprimirEtiqueta(Salida)

		}
		generador.Printf("c", "", "32")
		generador.agregarCodigo("\n")
		generador.imprimirEtiqueta(Salida)
	}
	generador.Printf("c", "", "10")
	generador.agregarCodigo("\n")

	return nil
}

func (v *Visitor) VisitIf_instr(ctx *parser.If_instrContext, ts Scope, generador *Generador) interface{} {
	fmt.Println("If no valido")
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: "If no válido",
	})
	return Valor{Valor: "Error: If no válido\n", Tipo: Error}
}

func (v *Visitor) VisitIf(ctx *parser.IfContext, ts Scope, generador *Generador) interface{} {
	condicion := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if condicion.Tipo == Bool {
		Salida := generador.nuevaEtiqueta()
		True := generador.nuevaEtiqueta()
		False := generador.nuevaEtiqueta()
		new_ts := NuevoScope(ts, "If")
		//new_ts.etqSalida = Salida

		generador.Comentario("Instrucción IF")
		generador.If(fmt.Sprint(condicion.Valor), "==", "1", True)
		generador.Goto(Salida)
		generador.imprimirEtiqueta(True)
		v.Visit(ctx.Block(), new_ts, generador)
		generador.Goto(Salida)
		generador.imprimirEtiqueta(False)
		generador.imprimirEtiqueta(Salida)
	} else if condicion.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("Expresión no válida en el IF\n       > %v\n", condicion.Valor),
		})
	} else {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("La expresión '%v' en el IF no es booleana\n", condicion.Valor),
		})
	}
	return Valor{}
}

func (v *Visitor) VisitIfElse(ctx *parser.IfElseContext, ts Scope, generador *Generador) interface{} {
	condicion := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if condicion.Tipo == Bool {
		Salida := generador.nuevaEtiqueta()
		True := generador.nuevaEtiqueta()
		False := generador.nuevaEtiqueta()
		new_ts := NuevoScope(ts, "If")
		//new_ts.etqSalida = Salida

		generador.Comentario("Instrucción IF-ELSE")
		generador.If(fmt.Sprint(condicion.Valor), "==", "1", True)
		generador.Goto(False)
		generador.imprimirEtiqueta(True)
		v.Visit(ctx.Block(0), new_ts, generador)
		generador.Goto(Salida)
		generador.imprimirEtiqueta(False)
		v.Visit(ctx.Block(1), new_ts, generador)
		generador.imprimirEtiqueta(Salida)
	} else if condicion.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("Expresión no válida en el IF\n       > %v\n", condicion.Valor),
		})
	} else {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("La expresión '%v' en el IF no es booleana\n", condicion.Valor),
		})
	}
	return Valor{}
}

func (v *Visitor) VisitElseIf(ctx *parser.ElseIfContext, ts Scope, generador *Generador) interface{} {
	condicion := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if condicion.Tipo == Bool {
		Salida := generador.nuevaEtiqueta()
		True := generador.nuevaEtiqueta()
		False := generador.nuevaEtiqueta()
		new_ts := NuevoScope(ts, "If")
		//new_ts.etqSalida = Salida

		generador.Comentario("Instrucción ELSE-IF")
		generador.If(fmt.Sprint(condicion.Valor), "==", "1", True)
		generador.Goto(False)
		generador.imprimirEtiqueta(True)
		v.Visit(ctx.Block(), new_ts, generador)
		generador.Goto(Salida)
		generador.imprimirEtiqueta(False)
		v.Visit(ctx.If_instr(), ts, generador)
		generador.imprimirEtiqueta(Salida)
	} else if condicion.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("Expresión no válida en el IF\n       > %v\n", condicion.Valor),
		})
	} else {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("La expresión '%v' en el IF no es booleana\n", condicion.Valor),
		})
	}
	return Valor{}
}

func (v *Visitor) VisitSwitch(ctx *parser.Switch_instrContext, ts Scope, generador *Generador) interface{} {
	expr := v.Visit(ctx.Expr(), ts, generador).(Valor)
	new_ts := NuevoScope(ts, "Switch")
	for i := 0; ctx.Case_(i) != nil; i++ {
		exprCase := v.Visit(ctx.Case_(i).Expr(), new_ts, generador).(Valor)
		if exprCase.Tipo != expr.Tipo {
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.Case_(i).Expr().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.Case_(i).Expr().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("No se pueden comparar valores de distinto tipo en un SWITCH \n"),
			})
			return Valor{}
		}
	}

	generador.Comentario("Instrucción SWITCH")
	Default := generador.nuevaEtiqueta()
	Salida := generador.nuevaEtiqueta()
	var listaEtiquetas []string

	for i := 0; ctx.Case_(i) != nil; i++ {
		generador.Comentario("CASE")
		etq := generador.nuevaEtiqueta()
		exprCase := v.Visit(ctx.Case_(i).Expr(), new_ts, generador).(Valor)
		generador.If(fmt.Sprint(expr.Valor), "==", fmt.Sprint(exprCase.Valor), etq)
		listaEtiquetas = append(listaEtiquetas, etq)
	}
	generador.Goto(Default)
	for i := 0; ctx.Case_(i) != nil; i++ {
		generador.imprimirEtiqueta(listaEtiquetas[i])
		generador.Comentario("Instrucciones CASE")
		v.Visit(ctx.Case_(i).Block(), new_ts, generador)
		generador.Goto(Salida)
	}

	// Cuando no hizo match con ningun case
	generador.imprimirEtiqueta(Default)
	if ctx.Default_(0) != nil {
		generador.Comentario("Instrucciones DEFAULT")
		v.Visit(ctx.Default_(0).Block(), new_ts, generador)
		generador.Goto(Salida)
	}
	generador.imprimirEtiqueta(Salida)
	return Valor{}
}

func (v *Visitor) VisitWhile_instr(ctx *parser.While_instrContext, ts Scope, generador *Generador) interface{} {
	new_ts := NuevoScope(ts, "While")

	Ciclo := generador.nuevaEtiqueta()
	True := generador.nuevaEtiqueta()
	False := generador.nuevaEtiqueta()
	new_ts.etqSalida = False
	new_ts.etqCiclo = Ciclo

	generador.Comentario("Instrucción WHILE")
	generador.imprimirEtiqueta(Ciclo)
	condicion := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if condicion.Tipo == Bool {
		generador.If(fmt.Sprint(condicion.Valor), "==", "1", True)
		generador.Goto(False)
		generador.imprimirEtiqueta(True)
		v.Visit(ctx.Block(), new_ts, generador)
		generador.Goto(Ciclo)
		generador.imprimirEtiqueta(False)
		/* 		if res.Break {
			break
		} */
		return Valor{}
	} else if condicion.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("Expresión no válida en el WHILE\n       > %v\n", condicion.Valor),
		})
		return Valor{}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
		Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
		Mensaje: fmt.Sprintf("\nError: La expresión '%v' en el WHILE no es booleana\n", condicion.Valor),
	})
	return Valor{}
}

func (v *Visitor) VisitFor(ctx *parser.For_instrContext, ts Scope, generador *Generador) interface{} {
	ts_for := NuevoScope(ts, "For")
	// Viene un rango
	if ctx.Rango() != nil {
		expr1 := v.Visit(ctx.Rango().Expr(0), ts_for, generador).(Valor)
		expr2 := v.Visit(ctx.Rango().Expr(1), ts_for, generador).(Valor)
		if expr1.Tipo == Int && expr2.Tipo == Int { // Ambas expresiones son INT
			fmt.Println(expr1.Valor, ", ", expr2.Valor)
			//if expr1.Valor.(int) <= expr2.Valor.(int) { // La expresion 1 es menor que la expresion 2
			generador.Comentario("Instrucción FOR")
			pos := ts_for.agregarVariable(Variable{ctx.ID().GetText(), Int, expr1.Valor.(int), false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1})
			if pos == -1 {
				generador.Comentario("Error: La variable ya existe")
				return Valor{}
			}
			generador.Comentario("Declaración de variable: " + ctx.ID().GetText())
			generador.setStack(fmt.Sprint(pos), fmt.Sprint(expr1.Valor))
			generador.agregarCodigo("\n")

			Ciclo := generador.nuevaEtiqueta()
			True := generador.nuevaEtiqueta()
			False := generador.nuevaEtiqueta()
			temp := generador.nuevoTemporal()
			ts_for.etqSalida = False
			ts_for.etqCiclo = Ciclo

			//generador.Expresion(temp, fmt.Sprint(expr1.Valor), "", "")
			generador.getStack(temp, fmt.Sprint(pos))
			generador.imprimirEtiqueta(Ciclo)
			generador.If(temp, "<=", fmt.Sprint(expr2.Valor), True)
			generador.Goto(False)
			generador.imprimirEtiqueta(True)
			v.Visit(ctx.Block(), ts_for, generador)
			generador.Comentario("Aumento de variable: " + ctx.ID().GetText())
			generador.Expresion(temp, temp, "+", "1")
			generador.setStack(fmt.Sprint(pos), temp)
			generador.Goto(Ciclo)
			generador.imprimirEtiqueta(False)
			fmt.Println("LLEGO AQUI")
			return Valor{}

			//}
			/* 			// La expresion 1 es mayor que la expresion 2 -> Error
			   			listaErrores = append(listaErrores, Error_{
			   				Tipo:    "SEMANTICO",
			   				Linea:   fmt.Sprint(ctx.Rango().Expr(0).GetStart().GetLine()),
			   				Columna: fmt.Sprint(ctx.Rango().Expr(0).GetStart().GetColumn()),
			   				Mensaje: fmt.Sprintf("\nError: La expresión izquierda '%v' en el rango del FOR es mayor que la derecha\n", expr1.Valor),
			   			})
			   			return Valor{} */
		}
		// No es INT -> Error
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Rango().Expr(0).GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Rango().Expr(0).GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("\nError: La expresión '%v' o '%v' en el rango del FOR no es de tipo 'Int'\n", expr1.Valor, expr2.Valor),
		})
		return Valor{}
	}
	return Valor{}
}

func (v *Visitor) VisitGuard(ctx *parser.GuardContext, ts Scope, generador *Generador) interface{} {
	condicion := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if condicion.Tipo == Bool {
		Salida := generador.nuevaEtiqueta()
		True := generador.nuevaEtiqueta()
		False := generador.nuevaEtiqueta()
		new_ts := NuevoScope(ts, "Guard")
		//new_ts.etqSalida = Salida

		generador.Comentario("Instrucción GUARD")
		generador.If(fmt.Sprint(condicion.Valor), "==", "1", False)
		generador.Goto(True)
		generador.imprimirEtiqueta(True)
		v.Visit(ctx.Block(), new_ts, generador)
		generador.Goto(Salida)
		generador.imprimirEtiqueta(False)
		generador.imprimirEtiqueta(Salida)
		return Valor{}
	} else if condicion.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("Expresión no válida en el GUARD\n       > %v\n", condicion.Valor),
		})
		return Valor{}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
		Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
		Mensaje: fmt.Sprintf("\nError: La expresión '%v' en el GUARD no es booleana\n", condicion.Valor),
	})
	return Valor{}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext, ts Scope, generador *Generador) interface{} {
	res := ts.encontrarVariable(ctx.GetText(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	if res != nil {

		tmp1 := generador.nuevoTemporal()
		tmp2 := generador.nuevoTemporal()

		generador.Comentario("Accediendo a variable: " + res.(Variable).ID)
		generador.Expresion(tmp1, "P", "+", fmt.Sprint(res.(Variable).Posicion))
		generador.getStack(tmp2, tmp1)
		generador.agregarCodigo("\n")
		return Valor{Valor: tmp2, Tipo: res.(Variable).Tipo}
	}
	return Valor{Valor: 999999999, Tipo: Error}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext, ts Scope, generador *Generador) interface{} {
	value := strings.TrimPrefix(strings.TrimSuffix(ctx.GetText(), "\""), "\"") // quita las primeras y ultimas comillas
	value = strings.Replace(value, "\\\"", "\"", -1)                           // reemplaza \" por "
	value = strings.Replace(value, "\\n", "\n", -1)                            // reemplaza \n por nueva linea
	value = strings.Replace(value, "\\r", "\r", -1)                            // reemplaza \n por nueva linea
	value = strings.Replace(value, "\\t", "\t", -1)                            // reemplaza \n por tabulacion
	value = strings.Replace(value, "\\\\", "\\", -1)                           // reemplaza \\ por \                                                       // agrega nueva linea al final
	temp := generador.nuevoTemporal()
	generador.Comentario("Primitivo String")
	generador.Expresion(temp, "H", "", "")

	for i := 0; i < len(value); i++ {
		generador.setHeap("H", fmt.Sprintf("%d", value[i]))
		generador.nextHeap()
	}

	generador.setHeap("H", "-1")
	generador.nextHeap()
	generador.agregarCodigo("\n")
	return Valor{Valor: temp, Tipo: String}
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext, ts Scope, generador *Generador) interface{} {
	i, _ := strconv.Atoi(ctx.GetText())
	return Valor{Valor: i, Tipo: Int}
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext, ts Scope, generador *Generador) interface{} {
	i, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Valor{Valor: i, Tipo: Float}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext, ts Scope, generador *Generador) interface{} {
	value, _ := strconv.ParseBool(ctx.GetText())
	valor := 0
	if value {
		valor = 1
	}
	generador.agregarCodigo("\n")
	return Valor{Valor: valor, Tipo: Bool}
}

func (v *Visitor) VisitCharExpr(ctx *parser.CharExprContext, ts Scope, generador *Generador) interface{} {
	value := strings.TrimPrefix(strings.TrimSuffix(ctx.GetText(), "'"), "'") // quita las primeras y ultimas comillas
	value = strings.Replace(value, "\\\"", "\"", -1)                         // reemplaza \" por "
	value = strings.Replace(value, "\\n", "\n", -1)                          // reemplaza \n por nueva linea
	value = strings.Replace(value, "\\r", "\r", -1)                          // reemplaza \n por nueva linea
	value = strings.Replace(value, "\\t", "\t", -1)                          // reemplaza \n por tabulacion
	value = strings.Replace(value, "\\\\", "\\", -1)                         // reemplaza \\ por \
	return Valor{Valor: fmt.Sprintf("%d", value[0]), Tipo: Character}
}

func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext, ts Scope, generador *Generador) interface{} {
	return Valor{Valor: 999999999, Tipo: String}
}

func (v *Visitor) VisitUMenosExpr(ctx *parser.UmenosExprContext, ts Scope, generador *Generador) interface{} {
	value := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if value.Tipo == Int || value.Tipo == Float {
		temp := generador.nuevoTemporal()
		generador.Expresion(temp, fmt.Sprint(value.Valor), "*", "-1")
		return Valor{Valor: temp, Tipo: value.Tipo}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
		Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
		Mensaje: "Negación Unaria solo puede realizarse con datos númericos\n",
	})
	generador.Comentario("Error: Negación Unaria solo puede realizarse con datos númericos")
	return Valor{Valor: 999999999, Tipo: Error}
}

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext, ts Scope, generador *Generador) interface{} {
	value := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if value.Tipo == Bool {
		True := generador.nuevaEtiqueta()
		False := generador.nuevaEtiqueta()
		Salida := generador.nuevaEtiqueta()
		temp := generador.nuevoTemporal()
		generador.If(fmt.Sprint(value.Valor), "==", "1", True)
		generador.Goto(False)
		generador.imprimirEtiqueta(True)
		generador.Expresion(temp, "0", "", "")
		generador.Goto(Salida)
		generador.imprimirEtiqueta(False)
		generador.Expresion(temp, "1", "", "")
		generador.imprimirEtiqueta(Salida)
		return Valor{Valor: temp, Tipo: Bool}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
		Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
		Mensaje: "Solo se puede realizar la operación NOT con tipos Bool",
	})
	generador.Comentario("Error: Solo se puede realizar la operación NOT con tipos Bool")
	return Valor{Valor: 999999999, Tipo: Error}
}

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext, ts Scope, generador *Generador) interface{} {
	return v.Visit(ctx.Expr(), ts, generador)
}

func (v *Visitor) VisitIntCastExpr(ctx *parser.IntCastExprContext, ts Scope, generador *Generador) interface{} {
	expr := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if expr.Tipo == Int {
		return Valor{Valor: expr.Valor, Tipo: Int}
	}
	if expr.Tipo == Float {
		return Valor{Valor: expr.Valor, Tipo: Int}
	}
	if expr.Tipo == String {
		generador.Comentario("Casteo de String a Int")
		temp1 := generador.nuevoTemporal() //pos en heap de inicio del numero
		temp2 := generador.nuevoTemporal() //temp generado para almacenar el numero casteado
		temp3 := generador.nuevoTemporal() //temp generado para almacenar el valor leido del heap
		ciclo := generador.nuevaEtiqueta()
		salida := generador.nuevaEtiqueta()
		generador.Expresion(temp1, fmt.Sprint(expr.Valor), "", "")
		generador.Expresion(temp2, "H", "", "") //temp2 = H
		generador.setHeap("H", "0")             //almacena en el heap el numero 0
		generador.imprimirEtiqueta(ciclo)
		generador.getHeap(temp3, temp1)              //temp3 = valor leido del heap
		generador.If(temp3, "==", "-1", salida)      //si es -1, termina el ciclo
		temp4 := generador.nuevoTemporal()           //temp generado para almacenar el valor numerico del caracter
		generador.Expresion(temp4, temp3, "-", "48") //se resta 48 para obtener su valor numerico, no ascii
		temp5 := generador.nuevoTemporal()           //temp generado para obtener el valor guardado
		generador.getHeap(temp5, "H")                //temp5 = valor guardado en el heap
		temp6 := generador.nuevoTemporal()
		generador.Expresion(temp6, temp5, "*", "10")  //multiplica el valor guardado * 10
		generador.Expresion(temp6, temp6, "+", temp4) //suma el valor guardado * 10 + el valor numerico del caracter
		generador.setHeap("H", temp6)                 //guarda el nuevo numero en el heap
		generador.Expresion(temp1, temp1, "+", "1")   //avanza al siguiente caracter
		generador.Goto(ciclo)
		generador.imprimirEtiqueta(salida)
		generador.nextHeap()
		generador.getHeap(temp2, temp2)
		generador.Comentario("Fin casteo de String a Int\n")
		return Valor{Valor: temp2, Tipo: Int}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
		Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
		Mensaje: fmt.Sprintf("No se puede convertir la expresion '%v' a Int", expr.Valor),
	})
	return Valor{}
}

func (v *Visitor) VisitFloatCastExpr(ctx *parser.FloatCastExprContext, ts Scope, generador *Generador) interface{} {
	expr := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if expr.Tipo == Float {
		return Valor{Valor: expr.Valor, Tipo: Float}
	}
	if expr.Tipo == Int {
		return Valor{Valor: expr.Valor, Tipo: Float}
	}
	if expr.Tipo == String {
		generador.Comentario("Casteo de String a Float")
		temp1 := generador.nuevoTemporal() //pos en heap de inicio del numero
		temp2 := generador.nuevoTemporal() //temp generado para almacenar el numero casteado
		temp3 := generador.nuevoTemporal() //temp generado para almacenar el valor leido del heap
		temp4 := generador.nuevoTemporal() //temp generado para almacenar el valor numerico del caracter
		temp5 := generador.nuevoTemporal() //temp generado para obtener el valor guardado
		temp6 := generador.nuevoTemporal() //temp generado para almacenar el nuevo numero
		temp7 := generador.nuevoTemporal() //variable aux para los decimales
		ciclo := generador.nuevaEtiqueta()
		aux := generador.nuevaEtiqueta()
		ciclo2 := generador.nuevaEtiqueta()
		salida := generador.nuevaEtiqueta()

		generador.Expresion(temp1, fmt.Sprint(expr.Valor), "", "")
		generador.Expresion(temp2, "H", "", "") //temp2 = H
		generador.setHeap("H", "0")             //almacena en el heap el numero 0
		generador.imprimirEtiqueta(ciclo)
		generador.getHeap(temp3, temp1)               //temp3 = valor leido del heap
		generador.If(temp3, "==", "-1", salida)       //si es -1, termina el ciclo
		generador.If(temp3, "==", "46", aux)          //si es ., va al aux
		generador.Expresion(temp4, temp3, "-", "48")  //se resta 48 para obtener su valor numerico, no ascii
		generador.getHeap(temp5, "H")                 //temp5 = valor guardado en el heap
		generador.Expresion(temp6, temp5, "*", "10")  //multiplica el valor guardado * 10
		generador.Expresion(temp6, temp6, "+", temp4) //suma el valor guardado * 10 + el valor numerico del caracter
		generador.setHeap("H", temp6)                 //guarda el nuevo numero en el heap
		generador.Expresion(temp1, temp1, "+", "1")   //avanza al siguiente caracter
		generador.Goto(ciclo)

		generador.imprimirEtiqueta(aux)
		generador.Expresion(temp1, temp1, "+", "1") //avanza al siguiente caracter
		generador.Expresion(temp7, "0.1", "", "")   //temp7 = 0.1

		generador.imprimirEtiqueta(ciclo2)
		generador.getHeap(temp3, temp1)               //temp3 = valor leido del heap
		generador.If(temp3, "==", "-1", salida)       //si es -1, termina el ciclo
		generador.Expresion(temp4, temp3, "-", "48")  //se resta 48 para obtener su valor numerico, no ascii
		generador.getHeap(temp5, "H")                 //temp5 = valor guardado en el heap
		generador.Expresion(temp6, temp4, "*", temp7) //multiplica el valor leido por aux -> leido*t7
		generador.Expresion(temp6, temp5, "+", temp6) //guardado+(leido*t7) -> es el nuevo numero
		generador.setHeap("H", temp6)                 //guarda el nuevo numero en el heap
		generador.Expresion(temp7, temp7, "*", "0.1") //temp7=temp7*0.1
		generador.Expresion(temp1, temp1, "+", "1")   //avanza al siguiente caracter
		generador.Goto(ciclo2)

		generador.imprimirEtiqueta(salida)
		generador.nextHeap()
		generador.getHeap(temp2, temp2)
		generador.Comentario("Fin casteo de String a Float\n")
		return Valor{Valor: temp2, Tipo: Float}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.Expr().GetStart().GetLine()),
		Columna: fmt.Sprint(ctx.Expr().GetStart().GetColumn()),
		Mensaje: fmt.Sprintf("No se puede convertir la expresion '%v' a Float", expr.Valor),
	})
	return Valor{}
}

func (v *Visitor) VisitStringCastExpr(ctx *parser.StringCastExprContext, ts Scope, generador *Generador) interface{} {
	expr := v.Visit(ctx.Expr(), ts, generador).(Valor)
	if expr.Tipo == Int {
		generador.Comentario("Casteo de Int a String")
		temp1 := generador.nuevoTemporal()
		temp2 := generador.nuevoTemporal()
		temp3 := generador.nuevoTemporal()
		temp4 := generador.nuevoTemporal()
		temp5 := generador.nuevoTemporal()
		temp6 := generador.nuevoTemporal() //almacena el inicio del string nuevo
		ciclo1 := generador.nuevaEtiqueta()
		ciclo2 := generador.nuevaEtiqueta()
		salida1 := generador.nuevaEtiqueta()
		salida2 := generador.nuevaEtiqueta()

		generador.setHeap("H", "-1")
		generador.nextHeap()
		generador.Expresion(temp1, fmt.Sprint(expr.Valor), "", "")
		generador.imprimirEtiqueta(ciclo1)
		generador.If(temp1, "==", "0", salida1)
		generador.Expresion(temp2, "(int)"+temp1, "%", "10")
		generador.Expresion(temp3, temp1, "-", temp2)
		generador.Expresion(temp1, temp3, "*", "0.1")
		generador.Expresion(temp2, temp2, "+", "48") //se suma 48 para guardar su valor ascii
		generador.setHeap("H", temp2)
		generador.nextHeap()
		generador.Goto(ciclo1)
		generador.imprimirEtiqueta(salida1)
		generador.Expresion(temp4, "H", "-", "1")
		generador.Expresion(temp6, "H", "", "")
		generador.imprimirEtiqueta(ciclo2)
		generador.If(temp4, "==", "-1", salida2)
		generador.getHeap(temp5, temp4)
		generador.setHeap("H", temp5)
		generador.Expresion(temp4, temp4, "-", "1")
		generador.nextHeap()
		generador.Goto(ciclo2)
		generador.imprimirEtiqueta(salida2)
		generador.nextHeap()
		generador.setHeap("H", "-1")
		generador.nextHeap()
		generador.Comentario("Fin Casteo de Int a String\n")
		return Valor{Valor: temp6, Tipo: String}
	} else if expr.Tipo == Float {
		return Valor{Valor: expr.Valor, Tipo: Float}
	} else if expr.Tipo == Bool {
		generador.Comentario("Casteo de Bool a String")
		temp := generador.nuevoTemporal()
		Salida := generador.nuevaEtiqueta()
		True := generador.nuevaEtiqueta()
		False := generador.nuevaEtiqueta()
		generador.Expresion(temp, "H", "", "")
		generador.If(fmt.Sprint(expr.Valor), "==", "1", True)
		generador.Goto(False)
		generador.imprimirEtiqueta(True)
		generador.setHeap("H", "116")
		generador.nextHeap()
		generador.setHeap("H", "114")
		generador.nextHeap()
		generador.setHeap("H", "117")
		generador.nextHeap()
		generador.setHeap("H", "101")
		generador.Goto(Salida)
		generador.imprimirEtiqueta(False)
		generador.setHeap("H", "102")
		generador.nextHeap()
		generador.setHeap("H", "97")
		generador.nextHeap()
		generador.setHeap("H", "108")
		generador.nextHeap()
		generador.setHeap("H", "115")
		generador.nextHeap()
		generador.setHeap("H", "101")
		generador.imprimirEtiqueta(Salida)
		generador.nextHeap()
		generador.setHeap("H", "-1")
		generador.nextHeap()
		generador.Comentario("Fin Casteo de Bool a String")
		return Valor{Valor: temp, Tipo: String}
	}
	return Valor{}
}

func (v *Visitor) VisitFuncion(ctx *parser.FuncionContext, ts Scope, generador *Generador) interface{} {
	var tipo int // Variable para almacenar el tipo de la variable
	if ctx.Tipo() != nil {
		switch ctx.Tipo().GetText() {
		case "String":
			tipo = String
		case "Bool":
			tipo = Bool
		case "Character":
			tipo = Character
		case "Int":
			tipo = Int
		case "Float":
			tipo = Float
		}
	} else {
		tipo = Void
	}
	ts.agregarFuncion(Funcion{ctx.ID().GetText(), ctx.Parametros(), ctx.Block(), tipo, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})

	Salida := generador.nuevaEtiqueta()
	ts_func := NuevoScope(ts, fmt.Sprintf("Función '%v'", ctx.ID().GetText()))
	ts_func.etqSalida = Salida
	generador.inicioFuncion(ctx.ID().GetText())

	//fmt.Println("parametros declarados: ", ctx.Parametros())
	if ctx.Parametros() != nil {
		// OBTENER EL PRIMER PARAMETRO
		// ID externo
		var id_externo string
		var id_interno string
		if ctx.Parametros().ID(1) != nil {
			id_externo = ctx.Parametros().ID(0).GetText()
			if id_externo == "_" {
				id_externo = ""
			}
			id_interno = ctx.Parametros().ID(1).GetText()
		} else {
			id_externo = ""
			id_interno = ctx.Parametros().ID(0).GetText()
		}

		TIPO := ctx.Parametros().Tipo()
		var tipo int // Variable para almacenar el tipo de la variable
		for TIPO != nil {
			switch TIPO.GetText() {
			case "String":
				tipo = String
			case "Bool":
				tipo = Bool
			case "Character":
				tipo = Character
			case "Int":
				tipo = Int
			case "Float":
				tipo = Float
			default:
				tipo = 10
			}
			if tipo != 10 {
				break
			}
			TIPO = TIPO.Tipo()
		}

		// Por referencia o por valor
		/* 		if ctx.Parametros().INOUT() != nil { // POR REFERENCIA
		   			// Agrega el parametro al scope de la funcion como constante
		   			//ts_func.agregarVariablePorReferencia(id_interno, Variable{id_interno, tipo, 0, false, 0, ctx.Parametros().ID(0).GetSymbol().GetLine(), ctx.Parametros().ID(0).GetSymbol().GetColumn(),false})
		   		} else {
		   			// Agrego el parametro al scope de la funcion como variable
		   			//ts_func.agregarVariable(Variable{id_interno, tipo, 0, false, 0, ctx.Parametros().ID(0).GetSymbol().GetLine(), ctx.Parametros().ID(0).GetSymbol().GetColumn(),false})
		   		} */
		ts_func.agregarParam(Variable{id_interno, tipo, 0, false, 0, ctx.Parametros().ID(0).GetSymbol().GetLine(), ctx.Parametros().ID(0).GetSymbol().GetColumn(), false, 1})

		params := ctx.Parametros()
		/// OBTENER LOS PARAMETROS ANIDADOS
		if params.Parametros() != nil {
			for true {
				// ID externo e interno
				var id_externo string
				var id_interno string
				if params.Parametros().ID(1) != nil {
					id_externo = params.Parametros().ID(0).GetText()
					if id_externo == "_" {
						id_externo = ""
					}
					id_interno = params.Parametros().ID(1).GetText()
				} else {
					id_externo = ""
					id_interno = params.Parametros().ID(0).GetText()
				}

				TIPO := params.Parametros().Tipo()
				var tipo int // Variable para almacenar el tipo de la variable
				for TIPO != nil {
					switch TIPO.GetText() {
					case "String":
						tipo = String
					case "Bool":
						tipo = Bool
					case "Character":
						tipo = Character
					case "Int":
						tipo = Int
					case "Float":
						tipo = Float
					default:
						tipo = 10
					}
					if tipo != 10 {
						break
					}
					fmt.Println("PRINT TIPO:", TIPO.GetText())
					TIPO = TIPO.Tipo()
				}
				/* 				// Por referencia o por valor
				   				if params.Parametros().INOUT() != nil { // POR REFERENCIA
				   					// Agrega el parametro al scope de la funcion como constante
				   					ts_func.agregarVariablePorReferencia(id_interno, Variable{id_interno, tipo, 0, false, 0, params.ID(0).GetSymbol().GetLine(), params.ID(0).GetSymbol().GetColumn(),false})
				   				} else {
				   					// Agrego el parametro al scope de la funcion como variable
				   					ts_func.agregarVariable(Variable{id_interno, tipo, 0, false, 0, params.ID(0).GetSymbol().GetLine(), params.ID(0).GetSymbol().GetColumn(),false})
				   				} */
				ts_func.agregarParam(Variable{id_interno, tipo, 0, false, 0, params.ID(0).GetSymbol().GetLine(), params.ID(0).GetSymbol().GetColumn(), false, 1})

				params = params.Parametros()
				if params.Parametros() == nil {
					break
				}
			}
		}
	}

	//v.Visit(ctx.Block(), ts, generador)
	ts_funciones[ctx.ID().GetText()] = ts_func
	v.Visit(ctx.Block(), ts_func, generador)
	generador.imprimirEtiqueta(ts_func.etqSalida)
	generador.finFuncion()
	return Valor{}
}

func (v *Visitor) VisitLlamada_func(ctx *parser.Llamada_funcContext, ts Scope, generador *Generador) interface{} {
	funcion := ts.encontrarFuncion(ctx.ID().GetText())
	if funcion.ID != "" {
		//ts_func := ts_funciones[ctx.ID().GetText()] // NuevoScope(ts, fmt.Sprintf("Función '%v'", ctx.ID().GetText()))
		//ts_func := NuevoScope(ts, fmt.Sprintf("Función '%v'", ctx.ID().GetText()))

		// ----------------------OBTENER LOS PARAMETROS DE LA LLAMADA DE LA FUNCION---------------------------------------
		// El atributo 'Constante' hara referencia a si es por referencia o por valor
		var parametros_llamada []Variable
		params_llamada := ctx.Parametros_llamada()

		if params_llamada != nil {
			// OBTENER EL PRIMER PARAMETRO
			// ID
			var id string
			if params_llamada.ID() != nil {
				id = params_llamada.ID().GetText()
			} else {
				id = params_llamada.Expr().GetText()
			}
			//  Valor y tipo
			expr := v.Visit(params_llamada.Expr(), ts, generador).(Valor)

			// Por referencia o por valor
			if params_llamada.REF() != nil {
				// POR REFERENCIA
				parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1})
			} else {
				// POR VALOR
				parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, true, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1})
			}

			/// OBTENER LOS PARAMETROS ANIDADOS
			if params_llamada.Parametros_llamada() != nil {
				for true {
					// ID
					var id string
					if params_llamada.Parametros_llamada().ID() != nil {
						id = params_llamada.Parametros_llamada().ID().GetText()
					} else {
						id = params_llamada.Parametros_llamada().Expr().GetText()
					}
					//  Valor y tipo
					expr := v.Visit(params_llamada.Parametros_llamada().Expr(), ts, generador).(Valor)

					// Por referencia o por valor
					if params_llamada.Parametros_llamada().REF() != nil { // POR REFERENCIA
						// Agrega el parametro al scope de la funcion como constante
						parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1})

					} else {
						// Agrego el parametro al scope de la funcion como variable
						parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, true, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), false, 1})
					}
					params_llamada = params_llamada.Parametros_llamada()
					if params_llamada.Parametros_llamada() == nil {
						break
					}
				}
			}
		}

		// ----------------------OBTENER LOS PARAMETROS DE LA DECLARACION DE LA FUNCION-----------------------------------

		var parametros_declarados []Parametro

		if funcion.Parametros != nil {
			// OBTENER EL PRIMER PARAMETRO
			// ID externo
			var id_externo string
			var id_interno string
			if funcion.Parametros.ID(1) != nil {
				id_externo = funcion.Parametros.ID(0).GetText()
				if id_externo == "_" {
					id_externo = ""
				}
				id_interno = funcion.Parametros.ID(1).GetText()
			} else {
				id_externo = ""
				id_interno = funcion.Parametros.ID(0).GetText()
			}

			TIPO := funcion.Parametros.Tipo()
			var tipo int // Variable para almacenar el tipo de la variable
			for TIPO != nil {
				switch TIPO.GetText() {
				case "String":
					tipo = String
				case "Bool":
					tipo = Bool
				case "Character":
					tipo = Character
				case "Int":
					tipo = Int
				case "Float":
					tipo = Float
				default:
					tipo = 10
				}
				if tipo != 10 {
					break
				}
				TIPO = TIPO.Tipo()
			}

			// Por referencia o por valor
			if funcion.Parametros.INOUT() != nil { // POR REFERENCIA
				// Agrega el parametro al scope de la funcion como constante
				parametros_declarados = append(parametros_declarados, Parametro{id_externo, id_interno, tipo, false})
			} else {
				// Agrego el parametro al scope de la funcion como variable
				parametros_declarados = append(parametros_declarados, Parametro{id_externo, id_interno, tipo, true})
			}

			/// OBTENER LOS PARAMETROS ANIDADOS
			if funcion.Parametros.Parametros() != nil {
				for true {
					// ID externo e interno
					var id_externo string
					var id_interno string
					if funcion.Parametros.Parametros().ID(1) != nil {
						id_externo = funcion.Parametros.Parametros().ID(0).GetText()
						if id_externo == "_" {
							id_externo = ""
						}
						id_interno = funcion.Parametros.Parametros().ID(1).GetText()
					} else {
						id_externo = ""
						id_interno = funcion.Parametros.Parametros().ID(0).GetText()
					}

					TIPO := funcion.Parametros.Parametros().Tipo()
					var tipo int // Variable para almacenar el tipo de la variable
					for TIPO != nil {
						switch TIPO.GetText() {
						case "String":
							tipo = String
						case "Bool":
							tipo = Bool
						case "Character":
							tipo = Character
						case "Int":
							tipo = Int
						case "Float":
							tipo = Float
						default:
							tipo = 10
						}
						if tipo != 10 {
							break
						}
						fmt.Println("PRINT TIPO:", TIPO.GetText())
						TIPO = TIPO.Tipo()
					}
					// Por referencia o por valor
					if funcion.Parametros.Parametros().INOUT() != nil { // POR REFERENCIA
						// Agrega el parametro al scope de la funcion como constante
						parametros_declarados = append(parametros_declarados, Parametro{id_externo, id_interno, tipo, false})
					} else {
						// Agrego el parametro al scope de la funcion como variable
						parametros_declarados = append(parametros_declarados, Parametro{id_externo, id_interno, tipo, true})
					}

					funcion.Parametros = funcion.Parametros.Parametros()
					if funcion.Parametros.Parametros() == nil {
						break
					}
				}
			}
		}

		// ----------------------COMPARAR LOS PARAMETROS DE LA LLAMADA CON LOS DE LA DECLARACION---------------------------
		referencia := false
		if parametros_declarados != nil && parametros_llamada != nil {
			// Verifica que la cantidad de parametros sea la misma
			if eq_size := len(parametros_declarados) == len(parametros_llamada); eq_size {
				for i := 0; i < len(parametros_declarados); i++ {
					// Verifica que los tipos coincidan en orden
					if parametros_declarados[i].Tipo != parametros_llamada[i].Tipo {
						listaErrores = append(listaErrores, Error_{
							Tipo:    "SEMANTICO",
							Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
							Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
							Mensaje: fmt.Sprintf("El tipo del parametro '%v' de la función '%v' no coincide con el de la llamada '%v'\n> %v\n", parametros_declarados[i].ID_Interno, funcion.ID, parametros_llamada[i].ID, parametros_llamada[i].Valor),
						})
						return Valor{}
					}
					// Verifica que los parametros de referecia declaradas si vengan por referencia en la llamada
					if condition := parametros_declarados[i].Referencia && !parametros_llamada[i].Constante; condition {
						listaErrores = append(listaErrores, Error_{
							Tipo:    "SEMANTICO",
							Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
							Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
							Mensaje: fmt.Sprintf("El parametro '%v' de la función '%v' debe ser por referencia\n", parametros_llamada[i].ID, funcion.ID),
						})
						return Valor{}
					} else if condition := !parametros_declarados[i].Referencia && parametros_llamada[i].Constante; condition {
						listaErrores = append(listaErrores, Error_{
							Tipo:    "SEMANTICO",
							Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
							Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
							Mensaje: fmt.Sprintf("El parametro '%v' de la función '%v' debe ser por valor\n", parametros_llamada[i].ID, funcion.ID),
						})
						return Valor{}
					}
					// Verifica que si viene un id externo en la declaracion, este sea igual al de la llamada
					if parametros_declarados[i].ID_Externo != "" {
						if parametros_declarados[i].ID_Externo != parametros_llamada[i].ID {
							listaErrores = append(listaErrores, Error_{
								Tipo:    "SEMANTICO",
								Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
								Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
								Mensaje: fmt.Sprintf("El id externo del parametro '%v' no coincide con el de la llamada '%v'\n", parametros_declarados[i].ID_Externo, parametros_llamada[i].ID),
							})
							return Valor{}
						}
					}
					// Verifica que el id interno no se repita en la declaracion
					seen := make(map[string]Parametro) // Crea un mapa para realizar un seguimiento de la cantidad de veces que se repite un id
					// Recorrer el slice
					for _, param := range parametros_declarados {
						// Verificar si el elemento ya se encuentra en el map
						if _, exists := seen[param.ID_Interno]; exists {
							listaErrores = append(listaErrores, Error_{
								Tipo:    "SEMANTICO",
								Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
								Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
								Mensaje: fmt.Sprintf("El parámetro %v se repite en la declaración de la función '%v'\n", param.ID_Interno, funcion.ID),
							})
							return Valor{}
						} else {
							// Si no se ha visto antes, registrar su aparición
							seen[param.ID_Interno] = param
						}
					}
				}
			} else {
				listaErrores = append(listaErrores, Error_{
					Tipo:    "SEMANTICO",
					Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
					Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
					Mensaje: fmt.Sprintf("La cantidad de parametros '%v' en la declaración de la función '%v' no coincide con los de la llamada '%v'\n", len(parametros_declarados), funcion.ID, len(parametros_llamada)),
				})
				return Valor{}
			}

			// ------------------------------AGREGAR LOS PARAMETROS AL AMBITO DE LA FUNCION--------------------------------------
			var count int = 0
			generador.Comentario("DECLARACION DE PARAMETROS DE LA FUNCION " + funcion.ID)
			for i := 0; i < len(parametros_declarados); i++ {

				if !parametros_llamada[i].Constante { //Si es por referencia
					referencia = true
				}
				generador.Comentario("Declaración de variable: " + parametros_declarados[i].ID_Interno)
				posicion := generador.nuevoTemporal()
				generador.Expresion(posicion, "P", "+", fmt.Sprint(ts.Size["Size"]+count))
				generador.setStack(posicion, fmt.Sprint(parametros_llamada[i].Valor))
				generador.agregarCodigo("\n")
				count++
				//ts_func.agregarVariable(Variable{parametros_llamada[i].ID, parametros_llamada[i].Tipo, parametros_llamada[i].Valor, false, 0, parametros_llamada[i].Linea, parametros_llamada[i].Columna})
			}
			generador.Comentario("FIN DECLARACION DE PARAMETROS")
		}

		// -----------------------------------------EJECUTAR EL BLOQUE DE LA FUNCION-----------------------------------------

		generador.Comentario("LLAMADA A LA FUNCION: " + funcion.ID)
		value := generador.nuevoTemporal()
		if !referencia {
			generador.nuevoAmbito(fmt.Sprint(ts.Size["Size"]))
		} else {
			generador.nuevoAmbito(fmt.Sprint(ts.Size["Size"] - 1))
		}
		generador.getFuncion(funcion.ID)
		generador.Comentario("Valor de retorno")
		generador.getStack(value, "P")
		if !referencia {
			generador.getAmbito(fmt.Sprint(ts.Size["Size"]))
		} else {
			generador.getAmbito(fmt.Sprint(ts.Size["Size"] - 1))
		}
		//generador.getStack(value, "P")
		generador.Comentario("FIN LLAMADA")
		return Valor{Valor: value, Tipo: funcion.TipoReturn}

	}
	return Valor{}
}

func (v *Visitor) VisitLlamadaFuncExpr(ctx *parser.LlamadaFuncExprContext, ts Scope, generador *Generador) interface{} {
	return v.Visit(ctx.Llamada_func(), ts, generador)
}

func (v *Visitor) VisitDeclaracion_vector(ctx *parser.Dec_vectorContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	var tipo int // Variable para almacenar el tipo de la variable
	switch ctx.Tipo().GetText() {
	case "String":
		tipo = String
	case "Bool":
		tipo = Bool
	case "Character":
		tipo = Character
	case "Int":
		tipo = Int
	case "Float":
		tipo = Float
	}
	//var valores []interface{}

	// Si vienen valores
	if ctx.Expr(0) != nil {
		// Verifica Tipos
		for i := 0; ctx.Expr(i) != nil; i++ {
			if v.Visit(ctx.Expr(i), ts, generador).(Valor).Tipo != tipo {
				listaErrores = append(listaErrores, Error_{
					Tipo:    "SEMANTICO",
					Linea:   fmt.Sprint(ctx.Expr(i).GetStart().GetLine()),
					Columna: fmt.Sprint(ctx.Expr(i).GetStart().GetColumn()),
					Mensaje: fmt.Sprintf("El tipo del elemento '%v' no coincide con el tipo del vector '%v'\n", v.Visit(ctx.Expr(i), ts, generador).(Valor).Valor, id),
				})
				return Valor{}
			}
		}
		generador.Comentario("Declaración de Vector")
		temp := generador.nuevoTemporal()
		cont := 0
		generador.Expresion(temp, "H", "", "")
		// Agrega los valores si todos son del mismo tipo
		for i := 0; ctx.Expr(i) != nil; i++ {
			value := v.Visit(ctx.Expr(i), ts, generador).(Valor).Valor
			generador.setHeap("H", fmt.Sprint(value))
			generador.nextHeap()
			cont++
		}
		//generador.setHeap("H", "-999999999") //valor que indica el final del vector
		//generador.nextHeap()
		generador.agregarCodigo("\n")

		// Agrega el vector al scope
		ts.agregarVariable(Variable{id, tipo, temp, false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), true, cont})
		return Valor{}
	}
	// Si no vienen valores
	ts.agregarVariable(Variable{id, tipo, 999999999, false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), true, 0})
	return Valor{}
}

func (v *Visitor) VisitAppend(ctx *parser.AppendContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	vector := ts.encontrarVariable(id, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()).(Variable)
	// Verifica que el ID sea de un vector
	if vector.enHeap {
		expr := v.Visit(ctx.Expr(), ts, generador).(Valor)
		generador.Comentario("Append")
		nuevoVector := generador.nuevoTemporal()
		salida := generador.nuevaEtiqueta()
		ciclo := generador.nuevaEtiqueta()
		contador := generador.nuevoTemporal()
		tmpH := generador.nuevoTemporal()
		tmp := generador.nuevoTemporal()

		generador.Expresion(nuevoVector, "H", "", "")               //inicio en el heap del nuevo Vector
		generador.Expresion(tmpH, fmt.Sprint(vector.Valor), "", "") //inicio en el heap del vector
		generador.imprimirEtiqueta(ciclo)
		generador.getHeap(tmp, tmpH)
		generador.If(contador, "==", fmt.Sprint(vector.Tam), salida)
		generador.Expresion(tmpH, tmpH, "+", "1")
		generador.Expresion(contador, contador, "+", "1")
		generador.setHeap("H", tmp)
		generador.nextHeap()
		generador.Goto(ciclo)
		generador.imprimirEtiqueta(salida)
		generador.setHeap("H", fmt.Sprint(expr.Valor))
		generador.nextHeap()
		generador.agregarCodigo("\n")
		ts.modificarVector(Variable{vector.ID, vector.Tipo, nuevoVector, vector.Constante, vector.Tam, vector.Linea, vector.Columna, vector.enHeap, vector.Tam + 1}, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
		return Valor{Valor: nuevoVector, Tipo: vector.Tipo}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
		Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	})
	return Valor{Valor: fmt.Sprintf("Error: La variable '%v' no es un Vector\n", id), Tipo: Error}
}

func (v *Visitor) VisitRemoveAt(ctx *parser.RemoveAtContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	vector := ts.encontrarVariable(id, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()).(Variable)
	indice := v.Visit(ctx.Expr(), ts, generador).(Valor)

	// Verifica que el ID sea de un vector
	if vector.enHeap {
		// Verifica que el indice sea de tipo int
		if indice.Tipo == Int {
			// Verifica si el índice está dentro de los límites del vector
			generador.Comentario("Remove At")
			tam := generador.nuevoTemporal()
			salida := generador.nuevaEtiqueta()
			ciclo := generador.nuevaEtiqueta()
			Err := generador.nuevaEtiqueta()
			NoErr := generador.nuevaEtiqueta()
			eliminar := generador.nuevaEtiqueta()
			contador := generador.nuevoTemporal()
			tmpH := generador.nuevoTemporal()
			tmp := generador.nuevoTemporal()
			nuevoVector := generador.nuevoTemporal()

			generador.Expresion(tam, fmt.Sprint(vector.Tam), "-", "1")
			generador.If(tam, "<", fmt.Sprint(indice.Valor), Err) // Si el indice es mayor al tamaño del vector
			generador.If(fmt.Sprint(indice.Valor), "<", "0", Err) // Si el indice es menor a 0
			generador.Goto(NoErr)
			generador.imprimirEtiqueta(Err)

			generador.Printf("c", "", "66")  //B
			generador.Printf("c", "", "111") //o
			generador.Printf("c", "", "117") //u
			generador.Printf("c", "", "110") //n
			generador.Printf("c", "", "100") //d
			generador.Printf("c", "", "115") //s
			generador.Printf("c", "", "69")  //E
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "111") //o
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "46")  //.
			generador.Printf("c", "", "10")  //salto

			generador.Expresion(nuevoVector, fmt.Sprint(vector.Valor), "", "")
			generador.Goto(salida)
			generador.imprimirEtiqueta(NoErr)
			// Elimina el elemento creando un nuevo vector excluyendo el elemento indicado

			generador.Expresion(nuevoVector, "H", "", "")               //inicio en el heap del nuevo Vector
			generador.Expresion(tmpH, fmt.Sprint(vector.Valor), "", "") //inicio en el heap del vector anterior
			generador.imprimirEtiqueta(ciclo)
			generador.getHeap(tmp, tmpH)
			generador.If(contador, "==", fmt.Sprint(indice.Valor), eliminar)
			generador.If(contador, "==", fmt.Sprint(vector.Tam), salida)
			generador.Expresion(tmpH, tmpH, "+", "1")
			generador.Expresion(contador, contador, "+", "1")
			generador.setHeap("H", tmp)
			generador.nextHeap()
			generador.Goto(ciclo)
			generador.imprimirEtiqueta(eliminar)
			generador.Expresion(tmpH, tmpH, "+", "1")
			generador.Expresion(contador, contador, "+", "1")
			generador.Goto(ciclo)
			generador.imprimirEtiqueta(salida)
			generador.agregarCodigo("\n")

			ts.modificarVector(Variable{vector.ID, vector.Tipo, nuevoVector, vector.Constante, vector.Tam, vector.Linea, vector.Columna, vector.enHeap, vector.Tam - 1}, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
			return Valor{Valor: nuevoVector, Tipo: vector.Tipo}
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
			Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
			Mensaje: fmt.Sprintf("El índice '%v' no es de tipo Int\n", indice.Valor),
		})
		return Valor{Valor: fmt.Sprintf("Error: El índice '%v' no es de tipo Int\n", indice.Valor), Tipo: Error}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
		Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	})
	return Valor{}
}

func (v *Visitor) VisitRemoveLast(ctx *parser.RemoveLastContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	vector := ts.encontrarVariable(id, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()).(Variable)

	// Verifica que el ID sea de un vector
	if vector.enHeap {
		if vector.Tam > 0 {
			generador.Comentario("Remove Last")
			nuevoVector := generador.nuevoTemporal()
			salida := generador.nuevaEtiqueta()
			ciclo := generador.nuevaEtiqueta()
			contador := generador.nuevoTemporal()
			tmpH := generador.nuevoTemporal()
			tmp := generador.nuevoTemporal()

			generador.Expresion(nuevoVector, "H", "", "")               //inicio en el heap del nuevo Vector
			generador.Expresion(tmpH, fmt.Sprint(vector.Valor), "", "") //inicio en el heap del vector
			generador.imprimirEtiqueta(ciclo)
			generador.getHeap(tmp, tmpH)
			generador.If(contador, "==", fmt.Sprint(vector.Tam-1), salida)
			generador.Expresion(tmpH, tmpH, "+", "1")
			generador.Expresion(contador, contador, "+", "1")
			generador.setHeap("H", tmp)
			generador.nextHeap()
			generador.Goto(ciclo)
			generador.imprimirEtiqueta(salida)
			generador.Comentario("FIN Remove Last")
			generador.agregarCodigo("\n")
			ts.modificarVector(Variable{vector.ID, vector.Tipo, nuevoVector, vector.Constante, vector.Tam, vector.Linea, vector.Columna, vector.enHeap, vector.Tam - 1}, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
			return Valor{Valor: nuevoVector, Tipo: vector.Tipo}
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
			Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
			Mensaje: fmt.Sprintf("El vector '%v' está vacío\n", id),
		})
		return Valor{}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
		Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	})
	return Valor{}
}

func (v *Visitor) VisitModificacion_vector(ctx *parser.Modificacion_vectorContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	/* 	vector := ts.encontrarVariable(id)
	   	indice := v.Visit(ctx.Expr(0), ts, generador)
	   	expr := v.Visit(ctx.Expr(1), ts, generador)
	   	// Verifica que el ID sea de un vector
	   	if reflect.TypeOf(vector.Valor).Kind() == reflect.Slice {
	   		// Si el indice es un numero
	   		if indice.Tipo == Int {
	   			// Verifica Tipos
	   			if expr.Tipo != vector.Tipo {
	   				listaErrores = append(listaErrores, Error_{
	   					Tipo:    "SEMANTICO",
	   					Linea:   "",
	   					Columna: "",
	   					Mensaje: fmt.Sprintf("El tipo del elemento '%v' no coincide con el tipo del vector '%v'\n", expr.Valor, id),
	   				})
	   				return Valor{Valor: fmt.Sprintf("Error: El tipo del elemento '%v' no coincide con el tipo del vector '%v'\n", expr.Valor, id), Tipo: Error}
	   			}
	   			// Verifica si el índice está dentro de los límites del slice
	   			if indice.Valor.(int) >= 0 && indice.Valor.(int) < len(vector.Valor.([]interface{})) {
	   				fmt.Printf("TABLA DE SIMBOLOS ANTES MODIFICAR VEC: %v\n", ts.Variables)
	   				// Agrega el vector al scope
	   				vector.Valor.([]interface{})[indice.Valor.(int)] = expr.Valor
	   				//ts.modificarVariable(id, vector)
	   				fmt.Printf("TABLA DE SIMBOLOS DESPUES MODIFICAR VEC: %v\n", ts.Variables)
	   				return Valor{Valor: 999999999}
	   			}
	   			listaErrores = append(listaErrores, Error_{
	   				Tipo:    "SEMANTICO",
	   				Linea:   "",
	   				Columna: "",
	   				Mensaje: fmt.Sprintf("El índice '%v' está fuera de los límites del vector '%v'\n", indice.Valor, id),
	   			})
	   			return Valor{Valor: fmt.Sprintf("Error: El índice '%v' está fuera de los límites del vector '%v'\n", indice.Valor, id), Tipo: Error}
	   		}
	   		// Si no es int
	   		listaErrores = append(listaErrores, Error_{
	   			Tipo:    "SEMANTICO",
	   			Linea:   "",
	   			Columna: "",
	   			Mensaje: fmt.Sprintf("El índice '%v' no es de tipo Int\n", indice.Valor),
	   		})
	   		return Valor{Valor: fmt.Sprintf("Error: El índice '%v' no es de tipo Int\n", indice.Valor), Tipo: Error}
	   	}
	   	listaErrores = append(listaErrores, Error_{
	   		Tipo:    "SEMANTICO",
	   		Linea:   "",
	   		Columna: "",
	   		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	   	}) */
	return Valor{Valor: fmt.Sprintf("Error: La variable '%v' no es un Vector\n", id), Tipo: Error}
}

func (v *Visitor) VisitAccesoVector(ctx *parser.AccesoVectorContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	vector := ts.encontrarVariable(id, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()).(Variable)
	indice := v.Visit(ctx.Expr(), ts, generador).(Valor)

	// Verifica que el ID sea de un vector
	if vector.enHeap {
		if indice.Tipo == Int {
			// Verifica si el índice está dentro de los límites del vector
			generador.Comentario("Acceso a Vector")
			tam := generador.nuevoTemporal()
			salida := generador.nuevaEtiqueta()
			ciclo := generador.nuevaEtiqueta()
			Err := generador.nuevaEtiqueta()
			NoErr := generador.nuevaEtiqueta()
			contador := generador.nuevoTemporal()
			tmpH := generador.nuevoTemporal()
			tmp := generador.nuevoTemporal()

			generador.Expresion(tam, fmt.Sprint(vector.Tam), "-", "1")
			generador.If(tam, "<", fmt.Sprint(indice.Valor), Err) // Si el indice es mayor al tamaño del vector
			generador.If(fmt.Sprint(indice.Valor), "<", "0", Err) // Si el indice es menor a 0
			generador.Goto(NoErr)
			generador.imprimirEtiqueta(Err)

			generador.Printf("c", "", "66")  //B
			generador.Printf("c", "", "111") //o
			generador.Printf("c", "", "117") //u
			generador.Printf("c", "", "110") //n
			generador.Printf("c", "", "100") //d
			generador.Printf("c", "", "115") //s
			generador.Printf("c", "", "69")  //E
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "111") //o
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "46")  //.
			generador.Printf("c", "", "32")  //espacio

			generador.Expresion(tmp, "999999999", "", "") //nil
			generador.Goto(salida)
			generador.imprimirEtiqueta(NoErr)

			// Retorna el valor del vector
			generador.Expresion(tmpH, fmt.Sprint(vector.Valor), "", "") //inicio en el heap del vector
			generador.imprimirEtiqueta(ciclo)
			generador.getHeap(tmp, tmpH)
			generador.If(contador, "==", fmt.Sprint(indice.Valor), salida)
			generador.Expresion(tmpH, tmpH, "+", "1")
			generador.Expresion(contador, contador, "+", "1")
			generador.Goto(ciclo)
			generador.imprimirEtiqueta(salida)
			generador.agregarCodigo("\n")

			return Valor{Valor: tmp, Tipo: vector.Tipo}
		}
		// Si no es int
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
			Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
			Mensaje: fmt.Sprintf("El índice '%v' no es de tipo Int\n", indice.Valor),
		})
		return Valor{}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
		Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	})
	return Valor{}
}

func (v *Visitor) VisitIsEmpty(ctx *parser.IsEmptyContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	vector := ts.encontrarVariable(ctx.ID().GetText(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()).(Variable)
	// Verifica que el ID sea de un vector
	if vector.enHeap {
		temp := generador.nuevoTemporal()
		if vector.Tam == 0 {
			generador.Expresion(temp, "1", "", "")
			return Valor{Valor: temp, Tipo: Bool}
		}
		generador.Expresion(temp, "0", "", "")
		return Valor{Valor: temp, Tipo: Bool}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
		Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	})
	return Valor{}
}

func (v *Visitor) VisitCount(ctx *parser.CountContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	vector := ts.encontrarVariable(ctx.ID().GetText(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()).(Variable)
	// Verifica que el ID sea de un vector
	if vector.enHeap {
		temp := generador.nuevoTemporal()
		generador.Expresion(temp, fmt.Sprint(vector.Tam), "", "")
		return Valor{Valor: temp, Tipo: Int}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
		Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	})
	return Valor{}
}

func (v *Visitor) VisitDeclaracion_matriz(ctx *parser.Dec_matrizContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	var tipo int // Variable para almacenar el tipo de la variable
	fmt.Println("Declaracion de matriz")
	var cont int
	aux := ctx.Tipo()
	for aux.Tipo() != nil {
		aux = aux.Tipo()
		cont++
	}
	fmt.Println("Dimensiones:", cont)
	fmt.Println("Tipo:", aux.GetText())
	switch aux.GetText() {
	case "String":
		tipo = String
	case "Bool":
		tipo = Bool
	case "Character":
		tipo = Character
	case "Int":
		tipo = Int
	case "Float":
		tipo = Float
	}
	fmt.Println("Tipo:", tipo)
	fmt.Println("Expresion:", ctx.Def_matriz().GetText())
	if cont == 1 {
		if ctx.Def_matriz().Expr(0) != nil {
			// Verifica Tipos
			for i := 0; ctx.Def_matriz().Expr(i) != nil; i++ {
				if v.Visit(ctx.Def_matriz().Expr(i), ts, generador).(Valor).Tipo != tipo {
					listaErrores = append(listaErrores, Error_{
						Tipo:    "SEMANTICO",
						Linea:   fmt.Sprint(ctx.Def_matriz().Expr(i).GetStart().GetLine()),
						Columna: fmt.Sprint(ctx.Def_matriz().Expr(i).GetStart().GetColumn()),
						Mensaje: fmt.Sprintf("El tipo del elemento '%v' no coincide con el tipo del vector '%v'\n", v.Visit(ctx.Def_matriz().Expr(i), ts, generador).(Valor).Valor, id),
					})
					return Valor{}
				}
			}
			generador.Comentario("Declaración de Matriz 1D")
			temp := generador.nuevoTemporal()
			cont := 0
			generador.Expresion(temp, "H", "", "")
			// Agrega los valores si todos son del mismo tipo
			for i := 0; ctx.Def_matriz().Expr(i) != nil; i++ {
				value := v.Visit(ctx.Def_matriz().Expr(i), ts, generador).(Valor).Valor
				generador.setHeap("H", fmt.Sprint(value))
				generador.nextHeap()
				cont++
			}
			generador.agregarCodigo("\n")
			// Agrega la matriz al scope
			ts.agregarVariable(Variable{id, tipo, temp, false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), true, cont})
			return Valor{}
		}
	} else if cont == 2 {
		tam := 0
		generador.Comentario("Declaración de Matriz " + fmt.Sprint(cont) + "D")
		temp := generador.nuevoTemporal()
		generador.Expresion(temp, "H", "", "")
		for j := 0; ctx.Def_matriz().Def_matriz(j) != nil; j++ {
			vect := ctx.Def_matriz().Def_matriz(j)
			for i := 0; vect.Expr(i) != nil; i++ {
				if v.Visit(vect.Expr(i), ts, generador).(Valor).Tipo != tipo {
					listaErrores = append(listaErrores, Error_{
						Tipo:    "SEMANTICO",
						Linea:   fmt.Sprint(vect.Expr(i).GetStart().GetLine()),
						Columna: fmt.Sprint(vect.Expr(i).GetStart().GetColumn()),
						Mensaje: fmt.Sprintf("El tipo del elemento '%v' no coincide con el tipo del vector '%v'\n", v.Visit(vect.Expr(i), ts, generador).(Valor).Valor, id),
					})
					return Valor{}
				}
			}
			// Agrega los valores si todos son del mismo tipo
			tam = 0
			for i := 0; vect.Expr(i) != nil; i++ {
				value := v.Visit(vect.Expr(i), ts, generador).(Valor).Valor
				generador.setHeap("H", fmt.Sprint(value))
				generador.nextHeap()
				tam++
			}
			generador.agregarCodigo("\n")
		}
		// Agrega la matriz al scope
		ts.agregarVariable(Variable{id, tipo, temp, false, 0, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), true, tam})
		return Valor{}
	}
	return Valor{}
}

func (v *Visitor) VisitAccesoMatriz(ctx *parser.AccesoMatrizContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	matriz := ts.encontrarVariable(id, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()).(Variable)

	/* 	for i := 0; ctx.Expr(i) != nil; i++ {
		indice = v.Visit(ctx.Expr(i), ts, generador).(Valor).Valor.(int) + indice
		fmt.Println("indice:", indice)
	} */
	in1 := v.Visit(ctx.Expr(0), ts, generador).(Valor)
	in2 := v.Visit(ctx.Expr(1), ts, generador).(Valor)

	// Verifica que el ID sea de la matriz
	if matriz.enHeap {
		if in1.Tipo == Int && in2.Tipo == Int {
			in1 := in1.Valor.(int)
			in2 := in2.Valor.(int)
			indice := in1*matriz.Tam + in2
			fmt.Println("indice:", indice)
			// Verifica si el índice está dentro de los límites de la matriz
			generador.Comentario("Acceso a Matriz")
			//tam := generador.nuevoTemporal()
			salida := generador.nuevaEtiqueta()
			ciclo := generador.nuevaEtiqueta()
			//Err := generador.nuevaEtiqueta()
			//NoErr := generador.nuevaEtiqueta()
			contador := generador.nuevoTemporal()
			tmpH := generador.nuevoTemporal()
			tmp := generador.nuevoTemporal()

			/* generador.Expresion(tam, fmt.Sprint(matriz.Tam), "-", "1")
			generador.If(tam, "<", fmt.Sprint(indice), Err) // Si el indice es mayor al tamaño del vector
			generador.If(fmt.Sprint(indice), "<", "0", Err) // Si el indice es menor a 0
			generador.Goto(NoErr)
			generador.imprimirEtiqueta(Err)

			generador.Printf("c", "", "66")  //B
			generador.Printf("c", "", "111") //o
			generador.Printf("c", "", "117") //u
			generador.Printf("c", "", "110") //n
			generador.Printf("c", "", "100") //d
			generador.Printf("c", "", "115") //s
			generador.Printf("c", "", "69")  //E
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "111") //o
			generador.Printf("c", "", "114") //r
			generador.Printf("c", "", "46")  //.
			generador.Printf("c", "", "32")  //espacio

			generador.Expresion(tmp, "999999999", "", "") //nil
			generador.Goto(salida)
			generador.imprimirEtiqueta(NoErr) */

			// Retorna el valor del vector
			generador.Expresion(tmpH, fmt.Sprint(matriz.Valor), "", "") //inicio en el heap del vector
			generador.imprimirEtiqueta(ciclo)
			generador.getHeap(tmp, tmpH)
			generador.If(contador, "==", fmt.Sprint(indice), salida)
			generador.Expresion(tmpH, tmpH, "+", "1")
			generador.Expresion(contador, contador, "+", "1")
			generador.Goto(ciclo)
			generador.imprimirEtiqueta(salida)
			generador.agregarCodigo("\n")

			return Valor{Valor: tmp, Tipo: matriz.Tipo}
		}
		// Si no es int
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
			Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
			Mensaje: fmt.Sprintf("El índice no es de tipo Int\n"),
		})
		return Valor{}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   fmt.Sprint(ctx.ID().GetSymbol().GetLine()),
		Columna: fmt.Sprint(ctx.ID().GetSymbol().GetColumn()),
		Mensaje: fmt.Sprintf("La variable '%v' no es una matriz\n", id),
	})
	return Valor{}
}

func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext, ts Scope, generador *Generador) interface{} {
	l := v.Visit(ctx.GetLeft(), ts, generador).(Valor)
	r := v.Visit(ctx.GetRight(), ts, generador).(Valor)
	op := ctx.GetOp().GetText()

	if l.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
			Mensaje: fmt.Sprintf("No se puede realizar la operación\n       > " + l.Valor.(string) + "\n"),
		})
		//return Valor{Valor: "Error: No se puede realizar la operación\n       > " + l.Valor.(string) + "\n", Tipo: Error}
		generador.Comentario("Error: No se puede realizar la operación > " + l.Valor.(string) + "\n")
		return Valor{Valor: 999999999, Tipo: Error}
	} else if r.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("No se puede realizar la operación %v", r.Valor),
		})
		//return Valor{Valor: "Error: No se puede realizar la operación\n       > " + r.Valor.(string) + "\n", Tipo: Error}
		generador.Comentario("Error: No se puede realizar la operación > " + r.Valor.(string) + "\n")
		return Valor{Valor: 999999999, Tipo: Error}
	} else if l.Valor == "nil" || r.Valor == "nil" {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
			Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
			Mensaje: "No se pueden hacer operaciones con nil\n",
		})
		//return Valor{Valor: "nil", Tipo: Error}
		generador.Comentario("Error: No se pueden hacer operaciones con nil\n")
		return Valor{Valor: 999999999, Tipo: Error}
	} else {
		temp := generador.nuevoTemporal()
		if op == "+" {
			//generador.Comentario("Operacion +")8
			if l.Tipo == String && r.Tipo == String {
				generador.Comentario("Concatenación de Strings")
				generador.Expresion(temp, "H", "", "") //Se almacena el inicio de la nueva cadena

				//Primer Palabra
				generador.Comentario("Primera Palabra")
				ciclo := generador.nuevaEtiqueta()
				salida := generador.nuevaEtiqueta()
				caracter := generador.nuevoTemporal()

				generador.imprimirEtiqueta(ciclo)
				generador.getHeap(caracter, fmt.Sprint(l.Valor))
				generador.If(caracter, "==", "-1", salida)
				generador.setHeap("H", caracter)
				generador.nextHeap()
				generador.Expresion(fmt.Sprint(l.Valor), fmt.Sprint(l.Valor), "+", "1")
				generador.Goto(ciclo)
				generador.imprimirEtiqueta(salida)

				//Segunda Palabra
				generador.Comentario("Segunda Palabra")
				ciclo = generador.nuevaEtiqueta()
				salida = generador.nuevaEtiqueta()
				caracter = generador.nuevoTemporal()

				generador.imprimirEtiqueta(ciclo)
				generador.getHeap(caracter, fmt.Sprint(r.Valor))
				generador.If(caracter, "==", "-1", salida)
				generador.setHeap("H", caracter)
				generador.nextHeap()
				generador.Expresion(fmt.Sprint(r.Valor), fmt.Sprint(r.Valor), "+", "1")
				generador.Goto(ciclo)
				generador.imprimirEtiqueta(salida)

				//Fin de la concatenacion
				generador.setHeap("H", "-1")
				generador.nextHeap()
				generador.agregarCodigo("\n")
				return Valor{Valor: temp, Tipo: String}
			} else if l.Tipo == Int && r.Tipo == Int {
				generador.Expresion(temp, "(int)"+fmt.Sprint(l.Valor), "+", "(int)"+fmt.Sprint(r.Valor))
				return Valor{Valor: temp, Tipo: Int}
			} else if (l.Tipo == Float && r.Tipo == Float) || (l.Tipo == Float && r.Tipo == Int) || (l.Tipo == Int && r.Tipo == Float) {
				generador.Expresion(temp, fmt.Sprint(l.Valor), "+", fmt.Sprint(r.Valor))
				return Valor{Valor: temp, Tipo: Float}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("No se puede realizar la suma entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor),
			})
			generador.Comentario(fmt.Sprintf("Error: No se puede realizar la suma entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor))
			return Valor{Valor: 999999999, Tipo: Error}
		} else if op == "-" {
			//generador.Comentario("Operacion -")
			if l.Tipo == Int && r.Tipo == Int {
				generador.Expresion(temp, "(int)"+fmt.Sprint(l.Valor), "-", "(int)"+fmt.Sprint(r.Valor))
				return Valor{Valor: temp, Tipo: Int}
			} else if (l.Tipo == Float && r.Tipo == Float) || (l.Tipo == Float && r.Tipo == Int) || (l.Tipo == Int && r.Tipo == Float) {
				generador.Expresion(temp, fmt.Sprint(l.Valor), "-", fmt.Sprint(r.Valor))
				return Valor{Valor: temp, Tipo: Float}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("No se puede realizar la resta entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor),
			})
			generador.Comentario(fmt.Sprintf("Error: No se puede realizar la resta entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor))
			return Valor{Valor: 999999999, Tipo: Error}
		} else if op == "*" {
			//generador.Comentario("Operacion *")
			if l.Tipo == Int && r.Tipo == Int {
				generador.Expresion(temp, "(int)"+fmt.Sprint(l.Valor), "*", "(int)"+fmt.Sprint(r.Valor))
				return Valor{Valor: temp, Tipo: Int}
			} else if (l.Tipo == Float && r.Tipo == Float) || (l.Tipo == Float && r.Tipo == Int) || (l.Tipo == Int && r.Tipo == Float) {
				generador.Expresion(temp, fmt.Sprint(l.Valor), "*", fmt.Sprint(r.Valor))
				return Valor{Valor: temp, Tipo: Float}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("No se puede realizar la multiplicación entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor),
			})
			generador.Comentario(fmt.Sprintf("Error: No se puede realizar la multiplicación entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor))
			return Valor{Valor: 999999999, Tipo: Error}
		} else if op == "/" {
			//generador.Comentario("Operacion /")
			if l.Tipo == Int && r.Tipo == Int {
				correcto := generador.nuevaEtiqueta()
				salida := generador.nuevaEtiqueta()
				generador.If(fmt.Sprint(r.Valor), "!=", "0", correcto)
				generador.Printf("c", "", "77")                //M
				generador.Printf("c", "", "97")                //a
				generador.Printf("c", "", "116")               //t
				generador.Printf("c", "", "104")               //h
				generador.Printf("c", "", "69")                //E
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "111")               //o
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "58")                //:
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "68")                //D
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "118")               //v
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "115")               //s
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "111")               //o
				generador.Printf("c", "", "110")               //n
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "98")                //b
				generador.Printf("c", "", "121")               //y
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "48")                //0
				generador.Printf("c", "", "46")                //.
				generador.Printf("c", "", "32")                //espacio
				generador.Expresion(temp, "999999999", "", "") //nil
				generador.Goto(salida)
				generador.imprimirEtiqueta(correcto)
				generador.Expresion(temp, "(int)"+fmt.Sprint(l.Valor), "/", "(int)"+fmt.Sprint(r.Valor))
				generador.imprimirEtiqueta(salida)
				return Valor{Valor: temp, Tipo: Int}
			} else if (l.Tipo == Float && r.Tipo == Float) || (l.Tipo == Float && r.Tipo == Int) || (l.Tipo == Int && r.Tipo == Float) {
				correcto := generador.nuevaEtiqueta()
				salida := generador.nuevaEtiqueta()
				generador.If(fmt.Sprint(r.Valor), "!=", "0", correcto)
				generador.Printf("c", "", "77")                //M
				generador.Printf("c", "", "97")                //a
				generador.Printf("c", "", "116")               //t
				generador.Printf("c", "", "104")               //h
				generador.Printf("c", "", "69")                //E
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "111")               //o
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "58")                //:
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "68")                //D
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "118")               //v
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "115")               //s
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "111")               //o
				generador.Printf("c", "", "110")               //n
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "98")                //b
				generador.Printf("c", "", "121")               //y
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "48")                //0
				generador.Printf("c", "", "46")                //.
				generador.Printf("c", "", "32")                //espacio
				generador.Expresion(temp, "999999999", "", "") //nil
				generador.Goto(salida)
				generador.imprimirEtiqueta(correcto)
				generador.Expresion(temp, fmt.Sprint(l.Valor), "/", fmt.Sprint(r.Valor))
				generador.imprimirEtiqueta(salida)
				return Valor{Valor: temp, Tipo: Float}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("No se puede realizar la división entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor),
			})
			generador.Comentario(fmt.Sprintf("Error: No se puede realizar la división entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor))
			return Valor{Valor: 999999999, Tipo: Error}
		} else if op == "%" {
			//generador.Comentario("Operacion %")
			if l.Tipo == Int && r.Tipo == Int {
				correcto := generador.nuevaEtiqueta()
				salida := generador.nuevaEtiqueta()
				generador.If(fmt.Sprint(r.Valor), "!=", "0", correcto)
				generador.Printf("c", "", "77")                //M
				generador.Printf("c", "", "97")                //a
				generador.Printf("c", "", "116")               //t
				generador.Printf("c", "", "104")               //h
				generador.Printf("c", "", "69")                //E
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "111")               //o
				generador.Printf("c", "", "114")               //r
				generador.Printf("c", "", "58")                //:
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "68")                //D
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "118")               //v
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "115")               //s
				generador.Printf("c", "", "105")               //i
				generador.Printf("c", "", "111")               //o
				generador.Printf("c", "", "110")               //n
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "98")                //b
				generador.Printf("c", "", "121")               //y
				generador.Printf("c", "", "32")                //espacio
				generador.Printf("c", "", "48")                //0
				generador.Printf("c", "", "46")                //.
				generador.Printf("c", "", "32")                //espacio
				generador.Expresion(temp, "999999999", "", "") //nil
				generador.Goto(salida)
				generador.imprimirEtiqueta(correcto)
				generador.Expresion(temp, "(int)"+fmt.Sprint(l.Valor), "%", "(int)"+fmt.Sprint(r.Valor))
				generador.imprimirEtiqueta(salida)
				return Valor{Valor: temp, Tipo: Int}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: "No se puede realizar el módulo entre tipos no válidos\n",
			})
			generador.Comentario(fmt.Sprintf("Error: No se puede realizar el módulo entre %v y %v, combinación de tipos no válida", l.Valor, r.Valor))
			return Valor{Valor: 999999999, Tipo: Error}
		} else if op == "==" || op == "!=" {
			if l.Tipo == r.Tipo {
				generador.Comentario(fmt.Sprintf("Operacion %v", op))
				True := generador.nuevaEtiqueta()
				False := generador.nuevaEtiqueta()
				Salida := generador.nuevaEtiqueta()
				generador.If(fmt.Sprint(l.Valor), op, fmt.Sprint(r.Valor), True)
				generador.Goto(False)
				generador.imprimirEtiqueta(True)
				generador.Expresion(temp, "1", "", "")
				generador.Goto(Salida)
				generador.imprimirEtiqueta(False)
				generador.Expresion(temp, "0", "", "")
				generador.imprimirEtiqueta(Salida)
				return Valor{Valor: temp, Tipo: Bool}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("Solo se puede realizar '%v' entre tipos iguales\n", op),
			})
			generador.Comentario(fmt.Sprintf("Error: Solo se puede realizar '%v' entre tipos iguales", op))
			return Valor{Valor: 999999999, Tipo: Error}

		} else if op == "<" || op == "<=" || op == ">" || op == ">=" {
			if l.Tipo == Bool || r.Tipo == Bool {
				listaErrores = append(listaErrores, Error_{
					Tipo:    "SEMANTICO",
					Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
					Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
					Mensaje: fmt.Sprintf("No se admite tipo 'Bool' en la operación '%v'\n", op),
				})
				generador.Comentario(fmt.Sprintf("No se admite tipo 'Bool' en la operación '%v'\n", op))
				return Valor{Valor: 999999999, Tipo: Error}
			} else if (l.Tipo == Int && r.Tipo == Int) || (l.Tipo == Float && r.Tipo == Float) || (l.Tipo == String && r.Tipo == String) || (l.Tipo == Character && r.Tipo == Character) {
				generador.Comentario(fmt.Sprintf("Operacion %v", op))
				True := generador.nuevaEtiqueta()
				False := generador.nuevaEtiqueta()
				Salida := generador.nuevaEtiqueta()
				generador.If(fmt.Sprint(l.Valor), op, fmt.Sprint(r.Valor), True)
				generador.Goto(False)
				generador.imprimirEtiqueta(True)
				generador.Expresion(temp, "1", "", "")
				generador.Goto(Salida)
				generador.imprimirEtiqueta(False)
				generador.Expresion(temp, "0", "", "")
				generador.imprimirEtiqueta(Salida)
				return Valor{Valor: temp, Tipo: Bool}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("Solo se admiten tipos iguales en la operacion '%v'\n", op),
			})
			generador.Comentario(fmt.Sprintf("Solo se admiten tipos iguales en la operacion '%v'\n", op))
			return Valor{Valor: 999999999, Tipo: Error}

		} else if op == "&&" {
			if l.Tipo == Bool && r.Tipo == Bool {
				True := generador.nuevaEtiqueta()
				False := generador.nuevaEtiqueta()
				TrueAux := generador.nuevaEtiqueta()
				Salida := generador.nuevaEtiqueta()
				generador.If(fmt.Sprint(l.Valor), "==", "1", TrueAux)
				generador.Goto(False)
				generador.imprimirEtiqueta(TrueAux)
				generador.If(fmt.Sprint(r.Valor), "==", "1", True)
				generador.Goto(False)
				generador.imprimirEtiqueta(True)
				generador.Expresion(temp, "1", "", "")
				generador.Goto(Salida)
				generador.imprimirEtiqueta(False)
				generador.Expresion(temp, "0", "", "")
				generador.imprimirEtiqueta(Salida)
				return Valor{Valor: temp, Tipo: Bool}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: fmt.Sprintf("Solo se admiten tipos 'Bool' en la operacion '&&'\n"),
			})
			generador.Comentario("Error: Solo se admiten tipos 'Bool' en la operacion '&&'")
			return Valor{Valor: 999999999, Tipo: Error}
		} else if op == "||" {
			if l.Tipo == Bool && r.Tipo == Bool {
				True := generador.nuevaEtiqueta()
				False := generador.nuevaEtiqueta()
				TrueAux := generador.nuevaEtiqueta()
				Salida := generador.nuevaEtiqueta()
				generador.If(fmt.Sprint(l.Valor), "==", "1", True)
				generador.Goto(TrueAux)
				generador.imprimirEtiqueta(TrueAux)
				generador.If(fmt.Sprint(r.Valor), "==", "1", True)
				generador.Goto(False)
				generador.imprimirEtiqueta(True)
				generador.Expresion(temp, "1", "", "")
				generador.Goto(Salida)
				generador.imprimirEtiqueta(False)
				generador.Expresion(temp, "0", "", "")
				generador.imprimirEtiqueta(Salida)
				return Valor{Valor: temp, Tipo: Bool}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   fmt.Sprint(ctx.GetLeft().GetStart().GetLine()),
				Columna: fmt.Sprint(ctx.GetLeft().GetStart().GetColumn()),
				Mensaje: "Solo se admiten tipos 'Bool' en la operacion '||'\n",
			})
			generador.Comentario("Error: Solo se admiten tipos 'Bool' en la operacion '||'")
			return Valor{Valor: 999999999, Tipo: Error}
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   fmt.Sprint(ctx.GetOp().GetLine()),
			Columna: fmt.Sprint(ctx.GetOp().GetColumn()),
			Mensaje: fmt.Sprintf("Operando no reconocido -> '%v'\n", op),
		})
		generador.Comentario(fmt.Sprintf("Error: Operando no reconocido -> '%v'\n", op))
		return Valor{Valor: 999999999, Tipo: Error}
	}
}

type Mensaje struct {
	Contenido string `json:"contenido"`
}

type Respuesta struct {
	Salida string `json:"Salida"`
}

type Response struct {
	Salida Scope `json:"Salida"`
}

type ResponseErrores struct {
	Salida []Error_ `json:"Salida"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ejecutar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		listaErrores = listaErrores[:0]
		var codigo Mensaje
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &codigo)
		input := antlr.NewInputStream(codigo.Contenido)
		lexer := parser.NewSwiftGrammarLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewSwiftGrammarParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
		tree := p.Prog()
		eval := Visitor{}
		ts := Scope{Variables: make(map[string]Variable), Nombre: "Global", Anterior: nil, Funciones: make(map[string]Funcion), Size: make(map[string]int)}
		generador := Generador{}
		ts_funciones = make(map[string]Scope)
		eval.Visit(tree, ts, &generador)
		res := generador.crearCodigoFinal()
		tablaSimbolos = ts
		w.WriteHeader(http.StatusOK)
		respuesta := Respuesta{
			Salida: fmt.Sprintf("%v", res),
		}
		jsonData, err := json.Marshal(respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})

	mux.HandleFunc("/ts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		respuesta := Response{
			Salida: tablaSimbolos,
		}
		jsonData, err := json.Marshal(respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(jsonData)
		w.Write(jsonData)

	})

	mux.HandleFunc("/errores", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		respuesta := ResponseErrores{
			Salida: listaErrores,
		}
		jsonData, err := json.Marshal(respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)

	})

	fmt.Println("Servidor corriendo en el puerto 5000")
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
