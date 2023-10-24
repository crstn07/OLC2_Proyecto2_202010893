// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftGrammarParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#instr.
	VisitInstr(ctx *InstrContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#DeclaracionTipoValor.
	VisitDeclaracionTipoValor(ctx *DeclaracionTipoValorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#DeclaracionTipo.
	VisitDeclaracionTipo(ctx *DeclaracionTipoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#DeclaracionValor.
	VisitDeclaracionValor(ctx *DeclaracionValorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#print_instr.
	VisitPrint_instr(ctx *Print_instrContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#If.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IfElse.
	VisitIfElse(ctx *IfElseContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ElseIf.
	VisitElseIf(ctx *ElseIfContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#switch_instr.
	VisitSwitch_instr(ctx *Switch_instrContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#while_instr.
	VisitWhile_instr(ctx *While_instrContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#for_instr.
	VisitFor_instr(ctx *For_instrContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#rango.
	VisitRango(ctx *RangoContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#guard.
	VisitGuard(ctx *GuardContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#funcion.
	VisitFuncion(ctx *FuncionContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#llamada_func.
	VisitLlamada_func(ctx *Llamada_funcContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#parametros_llamada.
	VisitParametros_llamada(ctx *Parametros_llamadaContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#dec_vector.
	VisitDec_vector(ctx *Dec_vectorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#copia_vector.
	VisitCopia_vector(ctx *Copia_vectorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#modificacion_vector.
	VisitModificacion_vector(ctx *Modificacion_vectorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#append.
	VisitAppend(ctx *AppendContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#removeLast.
	VisitRemoveLast(ctx *RemoveLastContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#removeAt.
	VisitRemoveAt(ctx *RemoveAtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IsEmpty.
	VisitIsEmpty(ctx *IsEmptyContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#FloatCastExpr.
	VisitFloatCastExpr(ctx *FloatCastExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#NilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#Count.
	VisitCount(ctx *CountContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#CharExpr.
	VisitCharExpr(ctx *CharExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#AccesoVector.
	VisitAccesoVector(ctx *AccesoVectorContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#UmenosExpr.
	VisitUmenosExpr(ctx *UmenosExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#LlamadaFuncExpr.
	VisitLlamadaFuncExpr(ctx *LlamadaFuncExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#StringCastExpr.
	VisitStringCastExpr(ctx *StringCastExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IntCastExpr.
	VisitIntCastExpr(ctx *IntCastExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}
}
