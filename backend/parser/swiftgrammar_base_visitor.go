// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseSwiftGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSwiftGrammarVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitInstr(ctx *InstrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitDeclaracionTipoValor(ctx *DeclaracionTipoValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitDeclaracionTipo(ctx *DeclaracionTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitDeclaracionValor(ctx *DeclaracionValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitAsignacion(ctx *AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitPrint_instr(ctx *Print_instrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitIf(ctx *IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitIfElse(ctx *IfElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitElseIf(ctx *ElseIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitSwitch_instr(ctx *Switch_instrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitCase(ctx *CaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitDefault(ctx *DefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitWhile_instr(ctx *While_instrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitFor_instr(ctx *For_instrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitRango(ctx *RangoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitGuard(ctx *GuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitFuncion(ctx *FuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitLlamada_func(ctx *Llamada_funcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitParametros_llamada(ctx *Parametros_llamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitDec_vector(ctx *Dec_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitCopia_vector(ctx *Copia_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitModificacion_vector(ctx *Modificacion_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitAppend(ctx *AppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitRemoveLast(ctx *RemoveLastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitRemoveAt(ctx *RemoveAtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitIsEmpty(ctx *IsEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitFloatCastExpr(ctx *FloatCastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitCount(ctx *CountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitOpExpr(ctx *OpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitCharExpr(ctx *CharExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitAccesoVector(ctx *AccesoVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitUmenosExpr(ctx *UmenosExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitLlamadaFuncExpr(ctx *LlamadaFuncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitStringCastExpr(ctx *StringCastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitIntCastExpr(ctx *IntCastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitStrExpr(ctx *StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}
