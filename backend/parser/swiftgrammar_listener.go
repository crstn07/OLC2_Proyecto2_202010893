// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstr is called when entering the instr production.
	EnterInstr(c *InstrContext)

	// EnterDeclaracionTipoValor is called when entering the DeclaracionTipoValor production.
	EnterDeclaracionTipoValor(c *DeclaracionTipoValorContext)

	// EnterDeclaracionTipo is called when entering the DeclaracionTipo production.
	EnterDeclaracionTipo(c *DeclaracionTipoContext)

	// EnterDeclaracionValor is called when entering the DeclaracionValor production.
	EnterDeclaracionValor(c *DeclaracionValorContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterPrint_instr is called when entering the print_instr production.
	EnterPrint_instr(c *Print_instrContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterIfElse is called when entering the IfElse production.
	EnterIfElse(c *IfElseContext)

	// EnterElseIf is called when entering the ElseIf production.
	EnterElseIf(c *ElseIfContext)

	// EnterSwitch_instr is called when entering the switch_instr production.
	EnterSwitch_instr(c *Switch_instrContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the default production.
	EnterDefault(c *DefaultContext)

	// EnterWhile_instr is called when entering the while_instr production.
	EnterWhile_instr(c *While_instrContext)

	// EnterFor_instr is called when entering the for_instr production.
	EnterFor_instr(c *For_instrContext)

	// EnterRango is called when entering the rango production.
	EnterRango(c *RangoContext)

	// EnterGuard is called when entering the guard production.
	EnterGuard(c *GuardContext)

	// EnterFuncion is called when entering the funcion production.
	EnterFuncion(c *FuncionContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterLlamada_func is called when entering the llamada_func production.
	EnterLlamada_func(c *Llamada_funcContext)

	// EnterParametros_llamada is called when entering the parametros_llamada production.
	EnterParametros_llamada(c *Parametros_llamadaContext)

	// EnterDec_vector is called when entering the dec_vector production.
	EnterDec_vector(c *Dec_vectorContext)

	// EnterCopia_vector is called when entering the copia_vector production.
	EnterCopia_vector(c *Copia_vectorContext)

	// EnterModificacion_vector is called when entering the modificacion_vector production.
	EnterModificacion_vector(c *Modificacion_vectorContext)

	// EnterAppend is called when entering the append production.
	EnterAppend(c *AppendContext)

	// EnterRemoveLast is called when entering the removeLast production.
	EnterRemoveLast(c *RemoveLastContext)

	// EnterRemoveAt is called when entering the removeAt production.
	EnterRemoveAt(c *RemoveAtContext)

	// EnterIsEmpty is called when entering the IsEmpty production.
	EnterIsEmpty(c *IsEmptyContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterFloatCastExpr is called when entering the FloatCastExpr production.
	EnterFloatCastExpr(c *FloatCastExprContext)

	// EnterFloatExpr is called when entering the FloatExpr production.
	EnterFloatExpr(c *FloatExprContext)

	// EnterNilExpr is called when entering the NilExpr production.
	EnterNilExpr(c *NilExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterCount is called when entering the Count production.
	EnterCount(c *CountContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterCharExpr is called when entering the CharExpr production.
	EnterCharExpr(c *CharExprContext)

	// EnterAccesoVector is called when entering the AccesoVector production.
	EnterAccesoVector(c *AccesoVectorContext)

	// EnterUmenosExpr is called when entering the UmenosExpr production.
	EnterUmenosExpr(c *UmenosExprContext)

	// EnterLlamadaFuncExpr is called when entering the LlamadaFuncExpr production.
	EnterLlamadaFuncExpr(c *LlamadaFuncExprContext)

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterStringCastExpr is called when entering the StringCastExpr production.
	EnterStringCastExpr(c *StringCastExprContext)

	// EnterIntCastExpr is called when entering the IntCastExpr production.
	EnterIntCastExpr(c *IntCastExprContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstr is called when exiting the instr production.
	ExitInstr(c *InstrContext)

	// ExitDeclaracionTipoValor is called when exiting the DeclaracionTipoValor production.
	ExitDeclaracionTipoValor(c *DeclaracionTipoValorContext)

	// ExitDeclaracionTipo is called when exiting the DeclaracionTipo production.
	ExitDeclaracionTipo(c *DeclaracionTipoContext)

	// ExitDeclaracionValor is called when exiting the DeclaracionValor production.
	ExitDeclaracionValor(c *DeclaracionValorContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitPrint_instr is called when exiting the print_instr production.
	ExitPrint_instr(c *Print_instrContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitIfElse is called when exiting the IfElse production.
	ExitIfElse(c *IfElseContext)

	// ExitElseIf is called when exiting the ElseIf production.
	ExitElseIf(c *ElseIfContext)

	// ExitSwitch_instr is called when exiting the switch_instr production.
	ExitSwitch_instr(c *Switch_instrContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the default production.
	ExitDefault(c *DefaultContext)

	// ExitWhile_instr is called when exiting the while_instr production.
	ExitWhile_instr(c *While_instrContext)

	// ExitFor_instr is called when exiting the for_instr production.
	ExitFor_instr(c *For_instrContext)

	// ExitRango is called when exiting the rango production.
	ExitRango(c *RangoContext)

	// ExitGuard is called when exiting the guard production.
	ExitGuard(c *GuardContext)

	// ExitFuncion is called when exiting the funcion production.
	ExitFuncion(c *FuncionContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitLlamada_func is called when exiting the llamada_func production.
	ExitLlamada_func(c *Llamada_funcContext)

	// ExitParametros_llamada is called when exiting the parametros_llamada production.
	ExitParametros_llamada(c *Parametros_llamadaContext)

	// ExitDec_vector is called when exiting the dec_vector production.
	ExitDec_vector(c *Dec_vectorContext)

	// ExitCopia_vector is called when exiting the copia_vector production.
	ExitCopia_vector(c *Copia_vectorContext)

	// ExitModificacion_vector is called when exiting the modificacion_vector production.
	ExitModificacion_vector(c *Modificacion_vectorContext)

	// ExitAppend is called when exiting the append production.
	ExitAppend(c *AppendContext)

	// ExitRemoveLast is called when exiting the removeLast production.
	ExitRemoveLast(c *RemoveLastContext)

	// ExitRemoveAt is called when exiting the removeAt production.
	ExitRemoveAt(c *RemoveAtContext)

	// ExitIsEmpty is called when exiting the IsEmpty production.
	ExitIsEmpty(c *IsEmptyContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitFloatCastExpr is called when exiting the FloatCastExpr production.
	ExitFloatCastExpr(c *FloatCastExprContext)

	// ExitFloatExpr is called when exiting the FloatExpr production.
	ExitFloatExpr(c *FloatExprContext)

	// ExitNilExpr is called when exiting the NilExpr production.
	ExitNilExpr(c *NilExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitCount is called when exiting the Count production.
	ExitCount(c *CountContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitCharExpr is called when exiting the CharExpr production.
	ExitCharExpr(c *CharExprContext)

	// ExitAccesoVector is called when exiting the AccesoVector production.
	ExitAccesoVector(c *AccesoVectorContext)

	// ExitUmenosExpr is called when exiting the UmenosExpr production.
	ExitUmenosExpr(c *UmenosExprContext)

	// ExitLlamadaFuncExpr is called when exiting the LlamadaFuncExpr production.
	ExitLlamadaFuncExpr(c *LlamadaFuncExprContext)

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitStringCastExpr is called when exiting the StringCastExpr production.
	ExitStringCastExpr(c *StringCastExprContext)

	// ExitIntCastExpr is called when exiting the IntCastExpr production.
	ExitIntCastExpr(c *IntCastExprContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)
}
