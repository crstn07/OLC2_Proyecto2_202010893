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
	case *parser.Copia_vectorContext:
		return v.VisitCopia_vector(val, ts, generador)
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
		//res := v.Visit(ctx.Instr(i), ts, generador)
		//out += fmt.Sprintf("%v", res.Valor)
		/* if res.Break {
			return Valor{Valor: out, Break: true}
		}
		if res.Continue {
			return Valor{Valor: out, Continue: true}
		}
		if res.Return {
			fmt.Println("Retorno en BLOQUE CON RETORNO TRUE: ", res.ReturnVal, res.ReturnTipo)
			return Valor{Valor: out, ReturnVal: res.ReturnVal, ReturnTipo: res.ReturnTipo, Return: true}
		} */
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
		//return Valor{Valor: 999999999, Continue: true}
	}
	if ctx.BREAK() != nil {
		//return Valor{Valor: 999999999, Break: true}
	}
	if ctx.RETURN() != nil {
		/* 		if ctx.Expr() != nil {
		   			retorno := v.Visit(ctx.Expr(), ts, generador)
		   			fmt.Println("Retorno ctx.RETURN(): ", retorno.Valor, retorno.Tipo)
		   			return Valor{ReturnVal: retorno.Valor, ReturnTipo: retorno.Tipo, Return: true}
		   		}
		   		return Valor{ReturnVal: nil, ReturnTipo: Void, Return: true} */
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
	heap := true
	if ctx.LET() == nil {
		constante = false
	}
	var tipo int // Variable para almacenar el tipo de la variable
	switch ctx.Tipo().GetText() {
	case "String":
		tipo = String
		heap = true
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
		//return ts.agregarVariable(Variable{ctx.ID().GetText(), tipo, valores, false, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
	} */
	var pos int
	if tipo == expr.Tipo {
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), tipo, expr.Valor, constante, 0, heap, "", "", ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) // Agrega la variable al scope
	} else if tipo == Float && expr.Tipo == Int {
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), Float, expr.Valor, constante, 0, false, "", "", ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) // Agrega la variable al scope de tipo float
	} else if tipo == Character && expr.Tipo == String && len(expr.Valor.(string)) == 1 {
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), Character, expr.Valor, constante, 0, false, "", "", ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) // Agrega la variable al scope de tipo Character
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
	heap := false
	if ctx.LET() == nil {
		var tipo int // Variable para almacenar el tipo de la variable
		switch ctx.Tipo().GetText() {
		case "String":
			tipo = String
			heap = true
		case "Bool":
			tipo = Bool
		case "Character":
			tipo = Character
		case "Int":
			tipo = Int
		case "Float":
			tipo = Float
		}
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), tipo, "nil", false, 0, heap, "", "", ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) // Agrega la variable al scope con valor nil
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
	heap := false
	if expr.Tipo != Error {
		constante := true
		if ctx.LET() == nil {
			constante = false
		}
		if expr.Tipo == String {
			heap = true
		}
		pos = ts.agregarVariable(Variable{ctx.ID().GetText(), expr.Tipo, expr.Valor, constante, 0, heap, "", "", ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
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
				generador.getStack(temp, fmt.Sprint(variable.(Variable).Posicion))
				generador.Expresion(temp2, temp, "+", fmt.Sprint(expr.Valor))
				res = Valor{Valor: temp2, Tipo: Int}
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
			generador.imprimirString()

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
		ts.etqSalida = Salida

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
		ts.etqSalida = Salida

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
		ts.etqSalida = Salida

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
	/* out := ""
	condicion := v.Visit(ctx.Expr(), ts, generador)
	if condicion.Tipo == Bool {
		for v.Visit(ctx.Expr(), ts, generador).Valor.(bool) { //WHILE
			new_ts := NuevoScope(ts, "While")
			//out += v.Visit(ctx.Block(), new_ts).Valor.(string)
			res := v.Visit(ctx.Block(), new_ts)
			out += res.Valor.(string)
			if res.Break {
				break
			}
		}
		return Valor{Valor: out}
	} else if condicion.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("Expresión no válida en el WHILE\n       > %v\n", condicion.Valor),
		})
		return Valor{Valor: fmt.Sprintf("\nError: Expresión no válida en el WHILE\n       > %v\n", condicion.Valor), Tipo: Error}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: fmt.Sprintf("\nError: La expresión '%v' en el WHILE no es booleana\n", condicion.Valor),
	})
	return Valor{Valor: fmt.Sprintf("\nError: La expresión '%v' en el WHILE no es booleana\n", condicion.Valor), Tipo: Error} */
	return nil
}

func (v *Visitor) VisitFor(ctx *parser.For_instrContext, ts Scope, generador *Generador) interface{} {
	/* ts_for := NuevoScope(ts, "For")
	out := ""
	id := ctx.ID().GetText()
	// Viene una Expresion. Solo puede ser String o Array
	if ctx.Expr() != nil {
		if expr := v.Visit(ctx.Expr(), ts_for); expr.Tipo == String { //Falta: Si expr.TIpo == Array
			//Ejecuta el for
			ts_for.agregarVariable(Variable{id, Character, " ", true, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) // Agrega la variable al ambito del FOR
			for i := 0; i < len(expr.Valor.(string)); i++ {
				caracter := expr.Valor.(string)[i]
				ts_for.modificarConstante(id, Valor{fmt.Sprintf("%c", caracter), Character, false, false, false, nil, 0}) // Cambia el valor del caracter actual
				// Ejecuta el bloque
				ts_block := NuevoScope(ts_for, "Bloque For") // Crea un nuevo scope para el bloque del FOR
				//out += v.Visit(ctx.Block(), ts_block).Valor.(string)

				res := v.Visit(ctx.Block(), ts_block)
				out += res.Valor.(string)
				if res.Break {
					break
				}
			}
			return Valor{Valor: out}
		}
		//Si expr.TIpo == Array :
		//Bloque de codigo
	}
	// Viene un rango
	if ctx.Rango() != nil {
		expr1 := v.Visit(ctx.Rango().Expr(0), ts_for)
		expr2 := v.Visit(ctx.Rango().Expr(1), ts_for)
		if expr1.Tipo == Int && expr2.Tipo == Int { // Ambas expresiones son INT
			if expr1.Valor.(int) <= expr2.Valor.(int) { // La expresion 1 es menor que la expresion 2
				//Ejecuta el for
				ts_for.agregarVariable(Variable{id, Int, expr1.Valor.(int), true, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) // Agrega la variable al ambito del FOR
				for ts_for.encontrarVariable(id).Valor.(int) <= expr2.Valor.(int) {
					// Ejecuta el bloque
					ts_block := NuevoScope(ts_for, "Bloque For") // Crea un nuevo scope para el bloque del FOR
					//out += v.Visit(ctx.Block(), ts_block).Valor.(string)
					res := v.Visit(ctx.Block(), ts_block)
					out += res.Valor.(string)
					if res.Break {
						break
					}
					ts_for.modificarConstante(id, Valor{ts_for.encontrarVariable(id).Valor.(int) + 1, Int, false, false, false, nil, 0}) // Aumenta en 1 la variable del FOR
				}
				return Valor{Valor: out}
			}
			// La expresion 1 es mayor que la expresion 2 -> Error
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   "",
				Columna: "",
				Mensaje: fmt.Sprintf("\nError: La expresión izquierda '%v' en el rango del FOR es mayor que la derecha\n", expr1.Valor),
			})
			return Valor{Valor: fmt.Sprintf("\nError: La expresión izquierda '%v' en el rango del FOR es mayor que la derecha\n", expr1.Valor), Tipo: Error}
		}
		// No es INT -> Error
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("\nError: La expresión '%v' o '%v' en el rango del FOR no es de tipo 'Int'\n", expr1.Valor, expr2.Valor),
		})
		return Valor{Valor: fmt.Sprintf("\nError: La expresion '%v' o '%v' en el rango del FOR no es de tipo 'Int'\n", expr1.Valor, expr2.Valor), Tipo: Error}
	}
	return Valor{Valor: 999999999} */
	return nil
}

func (v *Visitor) VisitGuard(ctx *parser.GuardContext, ts Scope, generador *Generador) interface{} {
	/* new_ts := NuevoScope(ts, "Guard")
	condicion := v.Visit(ctx.Expr(), ts, generador)
	if condicion.Tipo == Bool {
		if !condicion.Valor.(bool) { //GUARD
			return v.Visit(ctx.Block(), new_ts)
		}
		return Valor{Valor: 999999999}
	} else if condicion.Tipo == Error {
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("Expresión no válida en el GUARD\n       > %v\n", condicion.Valor),
		})
		return Valor{Valor: fmt.Sprintf("\nError: Expresión no válida en el GUARD\n       > %v\n", condicion.Valor), Tipo: Error}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: fmt.Sprintf("\nError: La expresión '%v' en el GUARD no es booleana\n", condicion.Valor),
	})
	return Valor{Valor: fmt.Sprintf("\nError: La expresión '%v' en el GUARD no es booleana\n", condicion.Valor), Tipo: Error} */
	return nil
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
		return Valor{Valor: temp, Tipo: Int}
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
	/* expr := v.Visit(ctx.Expr(), ts, generador)
	if expr.Tipo == Int {
		return Valor{Valor: expr.Valor, Tipo: Int}
	}
	if expr.Tipo == Float {
		return Valor{Valor: int(expr.Valor.(float64)), Tipo: Int}
	}
	if expr.Tipo == String {
		entero, err := strconv.Atoi(expr.Valor.(string))
		if err == nil {
			return Valor{Valor: entero, Tipo: Int}
		}
		decimal, err := strconv.ParseFloat(expr.Valor.(string), 64)
		if err == nil {
			entero = int(decimal)
			return Valor{Valor: entero, Tipo: Int}
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("No se puede convertir la expresion '%v' a Int", expr.Valor),
		})
		return Valor{Valor: "nil", Tipo: Int}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: fmt.Sprintf("No se puede convertir la expresion '%v' a Int", expr.Valor),
	})
	return Valor{Valor: "nil", Tipo: Int} */
	return nil
}

func (v *Visitor) VisitFloatCastExpr(ctx *parser.FloatCastExprContext, ts Scope, generador *Generador) interface{} {
	/* expr := v.Visit(ctx.Expr(), ts, generador)
	if expr.Tipo == Float {
		return Valor{Valor: expr.Valor, Tipo: Float}
	}
	if expr.Tipo == Int {
		return Valor{Valor: float64(expr.Valor.(int)), Tipo: Float}
	}
	if expr.Tipo == String {
		decimal, err := strconv.ParseFloat(expr.Valor.(string), 64)
		if err == nil {
			return Valor{Valor: decimal, Tipo: Float}
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("No se puede convertir la expresion '%v' a Float", expr.Valor),
		})
		return Valor{Valor: "nil", Tipo: Int}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: fmt.Sprintf("No se puede convertir la expresion '%v' a Float", expr.Valor),
	})
	return Valor{Valor: "nil", Tipo: Int} */
	return nil
}

func (v *Visitor) VisitStringCastExpr(ctx *parser.StringCastExprContext, ts Scope, generador *Generador) interface{} {
	//return Valor{Valor: fmt.Sprintf("%v", v.Visit(ctx.Expr(), ts, generador).Valor), Tipo: String}
	return nil
}

func (v *Visitor) VisitFuncion(ctx *parser.FuncionContext, ts Scope, generador *Generador) interface{} {
	/* 	if ctx.Tipo() != nil {
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
	   		return ts.agregarFuncion(Funcion{ctx.ID().GetText(), ctx.Parametros(), ctx.Block(), tipo, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
	   	}
	   	return ts.agregarFuncion(Funcion{ctx.ID().GetText(), ctx.Parametros(), ctx.Block(), Void, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
	*/
	return nil
}

func (v *Visitor) VisitLlamada_func(ctx *parser.Llamada_funcContext, ts Scope, generador *Generador) interface{} {
	/* 	funcion := ts.encontrarFuncion(ctx.ID().GetText())
	   	if funcion.ID != "" {
	   		ts_func := NuevoScope(ts, fmt.Sprintf("Función '%v'", ctx.ID().GetText()))

	   		// ----------------------OBTENER LOS PARAMETROS DE LA LLAMADA DE LA FUNCION---------------------------------------

	   		var parametros_llamada []Variable // En este caso el atributo 'Constante' hara referencia a si es por referencia o por valor
	   		params_llamada := ctx.Parametros_llamada()
	   		// OBTENER EL PRIMER PARAMETRO
	   		// ID
	   		var id string
	   		if params_llamada.ID() != nil {
	   			id = params_llamada.ID().GetText()
	   		} else {
	   			id = params_llamada.Expr().GetText()
	   		}
	   		//  Valor y tipo
	   		expr := v.Visit(params_llamada.Expr(), ts, generador)

	   		// Por referencia o por valor
	   		if params_llamada.REF() != nil { // POR REFERENCIA
	   			parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, false, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
	   		} else {
	   			parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, true, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
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
	   				expr := v.Visit(params_llamada.Parametros_llamada().Expr(), ts, generador)

	   				// Por referencia o por valor
	   				if params_llamada.Parametros_llamada().REF() != nil { // POR REFERENCIA
	   					// Agrega el parametro al scope de la funcion como constante
	   					parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, false, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})

	   				} else {
	   					// Agrego el parametro al scope de la funcion como variable
	   					parametros_llamada = append(parametros_llamada, Variable{id, expr.Tipo, expr.Valor, true, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
	   				}
	   				params_llamada = params_llamada.Parametros_llamada()
	   				if params_llamada.Parametros_llamada() == nil {
	   					break
	   				}
	   			}
	   		}

	   		// ----------------------OBTENER LOS PARAMETROS DE LA DECLARACION DE LA FUNCION-----------------------------------

	   		var parametros_declarados []Parametro

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

	   		// ----------------------COMPARAR LOS PARAMETROS DE LA LLAMADA CON LOS DE LA DECLARACION---------------------------

	   		// Verifica que la cantidad de parametros sea la misma
	   		if eq_size := len(parametros_declarados) == len(parametros_llamada); eq_size {
	   			for i := 0; i < len(parametros_declarados); i++ {
	   				// Verifica que los tipos que los tipos coincidan en orden
	   				if parametros_declarados[i].Tipo != parametros_llamada[i].Tipo {
	   					listaErrores = append(listaErrores, Error_{
	   						Tipo:    "SEMANTICO",
	   						Linea:   "",
	   						Columna: "",
	   						Mensaje: fmt.Sprintf("El tipo del parametro '%v' de la función '%v' no coincide con el de la llamada '%v'\n       > %v\n", parametros_declarados[i].ID_Interno, funcion.ID, parametros_llamada[i].ID, parametros_llamada[i].Valor),
	   					})
	   					return Valor{Valor: fmt.Sprintf("Error: El tipo del parametro '%v' de la función '%v' no coincide con el de la llamada '%v'\n       > %v\n", parametros_declarados[i].ID_Interno, funcion.ID, parametros_llamada[i].ID, parametros_llamada[i].Valor), Tipo: Error}
	   				}
	   				// Verifica que los parametros de referecia declaradas si vengan por referencia en la llamada
	   				if condition := parametros_declarados[i].Referencia && !parametros_llamada[i].Constante; condition {
	   					listaErrores = append(listaErrores, Error_{
	   						Tipo:    "SEMANTICO",
	   						Linea:   "",
	   						Columna: "",
	   						Mensaje: fmt.Sprintf("El parametro '%v' de la función '%v' debe ser por referencia\n", parametros_llamada[i].ID, funcion.ID),
	   					})
	   					return Valor{Valor: fmt.Sprintf("Error: El parametro '%v' de la función '%v' debe ser por referencia\n", parametros_llamada[i].ID, funcion.ID), Tipo: Error}
	   				} else if condition := !parametros_declarados[i].Referencia && parametros_llamada[i].Constante; condition {
	   					listaErrores = append(listaErrores, Error_{
	   						Tipo:    "SEMANTICO",
	   						Linea:   "",
	   						Columna: "",
	   						Mensaje: fmt.Sprintf("El parametro '%v' de la función '%v' debe ser por valor\n", parametros_llamada[i].ID, funcion.ID),
	   					})
	   					return Valor{Valor: fmt.Sprintf("Error: El parametro '%v' de la función '%v' debe ser por valor\n", parametros_llamada[i].ID, funcion.ID), Tipo: Error}
	   				}
	   				// Verifica que si viene un id externo en la declaracion, este sea igual al de la llamada
	   				if parametros_declarados[i].ID_Externo != "" {
	   					if parametros_declarados[i].ID_Externo != parametros_llamada[i].ID {
	   						listaErrores = append(listaErrores, Error_{
	   							Tipo:    "SEMANTICO",
	   							Linea:   "",
	   							Columna: "",
	   							Mensaje: fmt.Sprintf("El id externo del parametro '%v' no coincide con el de la llamada '%v'\n", parametros_declarados[i].ID_Externo, parametros_llamada[i].ID),
	   						})
	   						return Valor{Valor: fmt.Sprintf("Error: El id externo del parametro '%v' no coincide con el de la llamada '%v'\n", parametros_declarados[i].ID_Externo, parametros_llamada[i].ID), Tipo: Error}
	   					}
	   				}
	   				// Verificaerifica que el id interno no se repita en la declaracion
	   				seen := make(map[string]Parametro) // Crea un mapa para realizar un seguimiento de la cantidad de veces que se repite un id
	   				// Recorrer el slice
	   				for _, param := range parametros_declarados {
	   					// Verificar si el elemento ya se encuentra en el map
	   					if _, exists := seen[param.ID_Interno]; exists {
	   						listaErrores = append(listaErrores, Error_{
	   							Tipo:    "SEMANTICO",
	   							Linea:   "",
	   							Columna: "",
	   							Mensaje: fmt.Sprintf("El parámetro %v se repite en la declaración de la función '%v'\n", param.ID_Interno, funcion.ID),
	   						})
	   						return Valor{Valor: fmt.Sprintf("Error: El parámetro %v se repite en la declaración de la función '%v'\n", param.ID_Interno, funcion.ID), Tipo: Error}
	   					} else {
	   						// Si no se ha visto antes, registrar su aparición
	   						seen[param.ID_Interno] = param
	   					}
	   				}
	   			}
	   		} else {
	   			listaErrores = append(listaErrores, Error_{
	   				Tipo:    "SEMANTICO",
	   				Linea:   "",
	   				Columna: "",
	   				Mensaje: fmt.Sprintf("La cantidad de parametros '%v' en la declaración de la función '%v' no coincide con los de la llamada '%v'\n", len(parametros_declarados), funcion.ID, len(parametros_llamada)),
	   			})
	   			return Valor{Valor: fmt.Sprintf("Error: La cantidad de parametros '%v' en la declaración de la función '%v' no coincide con los de la llamada '%v'\n", len(parametros_declarados), funcion.ID, len(parametros_llamada)), Tipo: Error}
	   		}

	   		// ------------------------------AGREGAR LOS PARAMETROS AL AMBITO DE LA FUNCION--------------------------------------

	   		for i := 0; i < len(parametros_declarados); i++ {
	   			if !parametros_llamada[i].Constante { //Si es por referencia
	   				ts_func.agregarVariablePorReferencia(parametros_declarados[i].ID_Interno, Variable{parametros_llamada[i].ID, parametros_declarados[i].Tipo, parametros_llamada[i].Valor, parametros_llamada[i].Constante, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
	   			} else {
	   				ts_func.agregarVariable(Variable{parametros_declarados[i].ID_Interno, parametros_declarados[i].Tipo, parametros_llamada[i].Valor, parametros_llamada[i].Constante, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()})
	   			}
	   		}

	   		// -----------------------------------------EJECUTAR EL BLOQUE DE LA FUNCION-----------------------------------------

	   		if funcion.TipoReturn == Void {
	   			retorno := v.Visit(funcion.Bloque, ts_func)
	   			//fmt.Println("Retorno EN FUNCION VOID: ", retorno.ReturnVal)
	   			fmt.Println("TABLA DE SIMBOLOS FUNCION:", ts_func.Nombre, ts_func.Variables)
	   			fmt.Println("TABLA DE SIMBOLOS TS:", ts.Nombre, ts.Variables)
	   			if retorno.ReturnVal == nil {
	   				// Si los parametros son por referencia, actualizar los valores en el scope padre
	   				for _, valor := range ts_func.Variables {
	   					if !valor.Constante {
	   						ts.cambiarValPorReferencia(valor.ID, valor.Valor)
	   						fmt.Println("-----------SE MODIFICO LA VARIABLE------------------------->", valor.ID)
	   						fmt.Println("TABLA DE SIMBOLOS MODIFICADA TS:", ts.Nombre, ts.Variables)
	   					}
	   				}
	   				return retorno
	   			}
	   			return Valor{Valor: "Error: La función no debe tener un valor de retorno\n", Tipo: Error}

	   		} else {
	   			retorno := v.Visit(funcion.Bloque, ts_func)
	   			//fmt.Println("Retorno EN FUNCION: ", retorno.ReturnVal, retorno.ReturnTipo)
	   			if retorno.ReturnTipo == Void {
	   				listaErrores = append(listaErrores, Error_{
	   					Tipo:    "SEMANTICO",
	   					Linea:   "",
	   					Columna: "",
	   					Mensaje: fmt.Sprintf("La función '%v' debe tener un valor de retorno\n", funcion.ID),
	   				})
	   				return Valor{Valor: fmt.Sprintf("Error: La función '%v' debe tener un valor de retorno\n", funcion.ID), Tipo: Error}
	   			} else if retorno.ReturnTipo == funcion.TipoReturn {
	   				return Valor{Valor: retorno.ReturnVal, Tipo: retorno.ReturnTipo, ReturnVal: retorno.ReturnVal, ReturnTipo: retorno.ReturnTipo, Return: true}
	   			} else {
	   				listaErrores = append(listaErrores, Error_{
	   					Tipo:    "SEMANTICO",
	   					Linea:   "",
	   					Columna: "",
	   					Mensaje: fmt.Sprintf("El tipo de la expresión de retorno no coincide con el tipo de la función:  '%v'\n", funcion.ID),
	   				})
	   				return Valor{Valor: fmt.Sprintf("Error: El tipo de la expresión de retorno no coincide con el tipo de la función:  '%v'\n", funcion.ID), Tipo: Error}
	   			}
	   		}
	   	}

	   	listaErrores = append(listaErrores, Error_{
	   		Tipo:    "SEMANTICO",
	   		Linea:   "",
	   		Columna: "",
	   		Mensaje: fmt.Sprintf("No existe la función %s", ctx.ID().GetText()),
	   	}) */
	return Valor{Valor: fmt.Sprintf("Error: No existe la función %s", ctx.ID().GetText()), Tipo: Error}

}

func (v *Visitor) VisitLlamadaFuncExpr(ctx *parser.LlamadaFuncExprContext, ts Scope, generador *Generador) interface{} {
	return v.Visit(ctx.Llamada_func(), ts, generador)
}

func (v *Visitor) VisitDeclaracion_vector(ctx *parser.Dec_vectorContext, ts Scope, generador *Generador) interface{} {
	/* id := ctx.ID().GetText()
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
	var valores []interface{}

	// Si vienen valores
	if ctx.Expr(0) != nil {
		// Verifica Tipos
		for i := 0; ctx.Expr(i) != nil; i++ {
			if v.Visit(ctx.Expr(i), ts, generador).Tipo != tipo {
				listaErrores = append(listaErrores, Error_{
					Tipo:    "SEMANTICO",
					Linea:   "",
					Columna: "",
					Mensaje: fmt.Sprintf("El tipo del elemento '%v' no coincide con el tipo del vector '%v'\n", v.Visit(ctx.Expr(i), ts, generador).Valor, id),
				})
				return Valor{Valor: fmt.Sprintf("Error: El tipo del elemento '%v' no coincide con el tipo del vector '%v'\n", v.Visit(ctx.Expr(i), ts, generador).Valor, id), Tipo: Error}
			}
		}
		// Agrega los valores si todos son del mismo tipo
		for i := 0; ctx.Expr(i) != nil; i++ {
			valores = append(valores, v.Visit(ctx.Expr(i), ts, generador).Valor)
		}
		// Agrega el vector al scope
		return ts.agregarVariable(Variable{id, tipo, valores, false, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) // quitar el return ?
	}
	// Si no vienen valores
	return ts.agregarVariable(Variable{id, tipo, valores, false, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()}) */

	return nil
}

func (v *Visitor) VisitCopia_vector(ctx *parser.Copia_vectorContext, ts Scope, generador *Generador) interface{} {
	/* id := ctx.ID().GetText()
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
	expr := v.Visit(ctx.Expr(), ts, generador)
	// Verifica tipos
	if expr.Tipo == tipo {
		// Si es un array Copia los valores
		if reflect.TypeOf(expr.Valor).Kind() == reflect.Slice {
			valores := make([]interface{}, len(expr.Valor.([]interface{})))
			copy(valores, expr.Valor.([]interface{}))
			ts.agregarVariable(Variable{id, tipo, valores, false})
			return Valor{Valor: 999999999}
		}
		return Valor{Valor: fmt.Sprintf("Error: La expresión '%v' no es un vector\n", expr.Valor), Tipo: Error}
	}
	return Valor{Valor: fmt.Sprintf("Error: El tipo de la expresión '%v' no coincide con el tipo del vector '%v'\n", expr, id), Tipo: Error} */
	return Valor{Valor: 999999999}
}

func (v *Visitor) VisitAppend(ctx *parser.AppendContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	/* 	vector := ts.encontrarVariable(id)
	   	// Verifica que el ID sea de un vector
	   	if reflect.TypeOf(vector.Valor).Kind() == reflect.Slice {
	   		ts.append(id, v.Visit(ctx.Expr(), ts, generador).Valor)
	   		return Valor{Valor: 999999999}
	   	}
	   	listaErrores = append(listaErrores, Error_{
	   		Tipo:    "SEMANTICO",
	   		Linea:   "",
	   		Columna: "",
	   		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	   	}) */
	return Valor{Valor: fmt.Sprintf("Error: La variable '%v' no es un Vector\n", id), Tipo: Error}
}

func (v *Visitor) VisitRemoveAt(ctx *parser.RemoveAtContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	/* 	vector := ts.encontrarVariable(id)
	   	indice := v.Visit(ctx.Expr(), ts, generador)
	   	fmt.Printf("TABLA DE SIMBOLOS ANTES: %v\n", ts.Variables)

	   	// Verifica que el ID sea de un vector
	   	if reflect.TypeOf(vector.Valor).Kind() == reflect.Slice {
	   		// Verifica que el indice sea de tipo int
	   		if indice.Tipo == Int {
	   			// Verifica si el índice está dentro de los límites del slice
	   			if indice.Valor.(int) >= 0 && indice.Valor.(int) < len(vector.Valor.([]interface{})) {
	   				// Elimina el elemento creando un nuevo slice excluyendo el elemento indicado
	   				vector.Valor = append(vector.Valor.([]interface{})[:indice.Valor.(int)], vector.Valor.([]interface{})[indice.Valor.(int)+1:]...)
	   				ts.modificarVariable(id, vector)
	   				fmt.Printf("TABLA DE SIMBOLOS DESPUES: %v\n", ts.Variables)
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

func (v *Visitor) VisitRemoveLast(ctx *parser.RemoveLastContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	/* vector := ts.encontrarVariable(id)

	// Verifica que el ID sea de un vector
	if reflect.TypeOf(vector.Valor).Kind() == reflect.Slice {
		// Elimina el elemento creando un nuevo slice excluyendo el elemento indicado
		if len(vector.Valor.([]interface{})) > 0 {
			vector.Valor = vector.Valor.([]interface{})[:len(vector.Valor.([]interface{}))-1]
			ts.modificarVariable(id, vector)
			fmt.Printf("TABLA DE SIMBOLOS DESPUES: %v\n", ts.Variables)
			return Valor{Valor: 999999999}
		}
		listaErrores = append(listaErrores, Error_{
			Tipo:    "SEMANTICO",
			Linea:   "",
			Columna: "",
			Mensaje: fmt.Sprintf("El vector '%v' está vacío\n", id),
		})
		return Valor{Valor: fmt.Sprintf("Error: El vector '%v' está vacío\n", id), Tipo: Error}
	}
	listaErrores = append(listaErrores, Error_{
		Tipo:    "SEMANTICO",
		Linea:   "",
		Columna: "",
		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	}) */
	return Valor{Valor: fmt.Sprintf("Error: La variable '%v' no es un Vector\n", id), Tipo: Error}
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
	/* vector := ts.encontrarVariable(id)
	indice := v.Visit(ctx.Expr(), ts, generador)

	// Verifica que el ID sea de un vector
	if reflect.TypeOf(vector.Valor).Kind() == reflect.Slice {
		if indice.Tipo == Int {
			// Verifica si el índice está dentro de los límites del slice
			if indice.Valor.(int) >= 0 && indice.Valor.(int) < len(vector.Valor.([]interface{})) {
				// Retorna el valor del vector
				return Valor{Valor: vector.Valor.([]interface{})[indice.Valor.(int)], Tipo: vector.Tipo}
			}
			listaErrores = append(listaErrores, Error_{
				Tipo:    "SEMANTICO",
				Linea:   "",
				Columna: "",
				Mensaje: fmt.Sprintf("El índice '%v' está fuera de los límites del vector '%v'\n", indice.Valor, id),
			})
			return Valor{Valor: "nil", Tipo: vector.Tipo}
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

func (v *Visitor) VisitIsEmpty(ctx *parser.IsEmptyContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	/* 	vector := ts.encontrarVariable(ctx.ID().GetText())
	   	// Verifica que el ID sea de un vector
	   	if reflect.TypeOf(vector.Valor).Kind() == reflect.Slice {
	   		if len(vector.Valor.([]interface{})) == 0 {
	   			return Valor{Valor: true, Tipo: Bool}
	   		}
	   		return Valor{Valor: false, Tipo: Bool}
	   	}
	   	listaErrores = append(listaErrores, Error_{
	   		Tipo:    "SEMANTICO",
	   		Linea:   "",
	   		Columna: "",
	   		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	   	}) */
	return Valor{Valor: fmt.Sprintf("Error: La variable '%v' no es un Vector\n", id), Tipo: Error}
}

func (v *Visitor) VisitCount(ctx *parser.CountContext, ts Scope, generador *Generador) interface{} {
	id := ctx.ID().GetText()
	/* 	vector := ts.encontrarVariable(ctx.ID().GetText())
	   	// Verifica que el ID sea de un vector
	   	if reflect.TypeOf(vector.Valor).Kind() == reflect.Slice {
	   		return Valor{Valor: len(vector.Valor.([]interface{})), Tipo: Int} // Retorna la cantidad de elementos del vector
	   	}
	   	listaErrores = append(listaErrores, Error_{
	   		Tipo:    "SEMANTICO",
	   		Linea:   "",
	   		Columna: "",
	   		Mensaje: fmt.Sprintf("La variable '%v' no es un Vector\n", id),
	   	}) */
	return Valor{Valor: fmt.Sprintf("Error: La variable '%v' no es un Vector\n", id), Tipo: Error}
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
			Mensaje: fmt.Sprintf("No se puede realizar la operación\n       > " + r.Valor.(string) + "\n"),
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
				generador.Expresion(temp, fmt.Sprint(l.Valor), "+", fmt.Sprint(r.Valor))
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
				generador.Expresion(temp, fmt.Sprint(l.Valor), "-", fmt.Sprint(r.Valor))
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
				generador.Expresion(temp, fmt.Sprint(l.Valor), "*", fmt.Sprint(r.Valor))
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
				generador.Expresion(temp, fmt.Sprint(l.Valor), "/", fmt.Sprint(r.Valor))
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
