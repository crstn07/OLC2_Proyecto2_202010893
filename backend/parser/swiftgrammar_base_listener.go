// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseSwiftGrammarListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseSwiftGrammarListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstr is called when production instr is entered.
func (s *BaseSwiftGrammarListener) EnterInstr(ctx *InstrContext) {}

// ExitInstr is called when production instr is exited.
func (s *BaseSwiftGrammarListener) ExitInstr(ctx *InstrContext) {}

// EnterDeclaracionTipoValor is called when production DeclaracionTipoValor is entered.
func (s *BaseSwiftGrammarListener) EnterDeclaracionTipoValor(ctx *DeclaracionTipoValorContext) {}

// ExitDeclaracionTipoValor is called when production DeclaracionTipoValor is exited.
func (s *BaseSwiftGrammarListener) ExitDeclaracionTipoValor(ctx *DeclaracionTipoValorContext) {}

// EnterDeclaracionTipo is called when production DeclaracionTipo is entered.
func (s *BaseSwiftGrammarListener) EnterDeclaracionTipo(ctx *DeclaracionTipoContext) {}

// ExitDeclaracionTipo is called when production DeclaracionTipo is exited.
func (s *BaseSwiftGrammarListener) ExitDeclaracionTipo(ctx *DeclaracionTipoContext) {}

// EnterDeclaracionValor is called when production DeclaracionValor is entered.
func (s *BaseSwiftGrammarListener) EnterDeclaracionValor(ctx *DeclaracionValorContext) {}

// ExitDeclaracionValor is called when production DeclaracionValor is exited.
func (s *BaseSwiftGrammarListener) ExitDeclaracionValor(ctx *DeclaracionValorContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseSwiftGrammarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseSwiftGrammarListener) ExitTipo(ctx *TipoContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseSwiftGrammarListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseSwiftGrammarListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterPrint_instr is called when production print_instr is entered.
func (s *BaseSwiftGrammarListener) EnterPrint_instr(ctx *Print_instrContext) {}

// ExitPrint_instr is called when production print_instr is exited.
func (s *BaseSwiftGrammarListener) ExitPrint_instr(ctx *Print_instrContext) {}

// EnterIf is called when production If is entered.
func (s *BaseSwiftGrammarListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production If is exited.
func (s *BaseSwiftGrammarListener) ExitIf(ctx *IfContext) {}

// EnterIfElse is called when production IfElse is entered.
func (s *BaseSwiftGrammarListener) EnterIfElse(ctx *IfElseContext) {}

// ExitIfElse is called when production IfElse is exited.
func (s *BaseSwiftGrammarListener) ExitIfElse(ctx *IfElseContext) {}

// EnterElseIf is called when production ElseIf is entered.
func (s *BaseSwiftGrammarListener) EnterElseIf(ctx *ElseIfContext) {}

// ExitElseIf is called when production ElseIf is exited.
func (s *BaseSwiftGrammarListener) ExitElseIf(ctx *ElseIfContext) {}

// EnterSwitch_instr is called when production switch_instr is entered.
func (s *BaseSwiftGrammarListener) EnterSwitch_instr(ctx *Switch_instrContext) {}

// ExitSwitch_instr is called when production switch_instr is exited.
func (s *BaseSwiftGrammarListener) ExitSwitch_instr(ctx *Switch_instrContext) {}

// EnterCase is called when production case is entered.
func (s *BaseSwiftGrammarListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseSwiftGrammarListener) ExitCase(ctx *CaseContext) {}

// EnterDefault is called when production default is entered.
func (s *BaseSwiftGrammarListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production default is exited.
func (s *BaseSwiftGrammarListener) ExitDefault(ctx *DefaultContext) {}

// EnterWhile_instr is called when production while_instr is entered.
func (s *BaseSwiftGrammarListener) EnterWhile_instr(ctx *While_instrContext) {}

// ExitWhile_instr is called when production while_instr is exited.
func (s *BaseSwiftGrammarListener) ExitWhile_instr(ctx *While_instrContext) {}

// EnterFor_instr is called when production for_instr is entered.
func (s *BaseSwiftGrammarListener) EnterFor_instr(ctx *For_instrContext) {}

// ExitFor_instr is called when production for_instr is exited.
func (s *BaseSwiftGrammarListener) ExitFor_instr(ctx *For_instrContext) {}

// EnterRango is called when production rango is entered.
func (s *BaseSwiftGrammarListener) EnterRango(ctx *RangoContext) {}

// ExitRango is called when production rango is exited.
func (s *BaseSwiftGrammarListener) ExitRango(ctx *RangoContext) {}

// EnterGuard is called when production guard is entered.
func (s *BaseSwiftGrammarListener) EnterGuard(ctx *GuardContext) {}

// ExitGuard is called when production guard is exited.
func (s *BaseSwiftGrammarListener) ExitGuard(ctx *GuardContext) {}

// EnterFuncion is called when production funcion is entered.
func (s *BaseSwiftGrammarListener) EnterFuncion(ctx *FuncionContext) {}

// ExitFuncion is called when production funcion is exited.
func (s *BaseSwiftGrammarListener) ExitFuncion(ctx *FuncionContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseSwiftGrammarListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseSwiftGrammarListener) ExitParametros(ctx *ParametrosContext) {}

// EnterLlamada_func is called when production llamada_func is entered.
func (s *BaseSwiftGrammarListener) EnterLlamada_func(ctx *Llamada_funcContext) {}

// ExitLlamada_func is called when production llamada_func is exited.
func (s *BaseSwiftGrammarListener) ExitLlamada_func(ctx *Llamada_funcContext) {}

// EnterParametros_llamada is called when production parametros_llamada is entered.
func (s *BaseSwiftGrammarListener) EnterParametros_llamada(ctx *Parametros_llamadaContext) {}

// ExitParametros_llamada is called when production parametros_llamada is exited.
func (s *BaseSwiftGrammarListener) ExitParametros_llamada(ctx *Parametros_llamadaContext) {}

// EnterDec_vector is called when production dec_vector is entered.
func (s *BaseSwiftGrammarListener) EnterDec_vector(ctx *Dec_vectorContext) {}

// ExitDec_vector is called when production dec_vector is exited.
func (s *BaseSwiftGrammarListener) ExitDec_vector(ctx *Dec_vectorContext) {}

// EnterCopia_vector is called when production copia_vector is entered.
func (s *BaseSwiftGrammarListener) EnterCopia_vector(ctx *Copia_vectorContext) {}

// ExitCopia_vector is called when production copia_vector is exited.
func (s *BaseSwiftGrammarListener) ExitCopia_vector(ctx *Copia_vectorContext) {}

// EnterModificacion_vector is called when production modificacion_vector is entered.
func (s *BaseSwiftGrammarListener) EnterModificacion_vector(ctx *Modificacion_vectorContext) {}

// ExitModificacion_vector is called when production modificacion_vector is exited.
func (s *BaseSwiftGrammarListener) ExitModificacion_vector(ctx *Modificacion_vectorContext) {}

// EnterAppend is called when production append is entered.
func (s *BaseSwiftGrammarListener) EnterAppend(ctx *AppendContext) {}

// ExitAppend is called when production append is exited.
func (s *BaseSwiftGrammarListener) ExitAppend(ctx *AppendContext) {}

// EnterRemoveLast is called when production removeLast is entered.
func (s *BaseSwiftGrammarListener) EnterRemoveLast(ctx *RemoveLastContext) {}

// ExitRemoveLast is called when production removeLast is exited.
func (s *BaseSwiftGrammarListener) ExitRemoveLast(ctx *RemoveLastContext) {}

// EnterRemoveAt is called when production removeAt is entered.
func (s *BaseSwiftGrammarListener) EnterRemoveAt(ctx *RemoveAtContext) {}

// ExitRemoveAt is called when production removeAt is exited.
func (s *BaseSwiftGrammarListener) ExitRemoveAt(ctx *RemoveAtContext) {}

// EnterDec_matriz is called when production dec_matriz is entered.
func (s *BaseSwiftGrammarListener) EnterDec_matriz(ctx *Dec_matrizContext) {}

// ExitDec_matriz is called when production dec_matriz is exited.
func (s *BaseSwiftGrammarListener) ExitDec_matriz(ctx *Dec_matrizContext) {}

// EnterDef_matriz is called when production def_matriz is entered.
func (s *BaseSwiftGrammarListener) EnterDef_matriz(ctx *Def_matrizContext) {}

// ExitDef_matriz is called when production def_matriz is exited.
func (s *BaseSwiftGrammarListener) ExitDef_matriz(ctx *Def_matrizContext) {}

// EnterIsEmpty is called when production IsEmpty is entered.
func (s *BaseSwiftGrammarListener) EnterIsEmpty(ctx *IsEmptyContext) {}

// ExitIsEmpty is called when production IsEmpty is exited.
func (s *BaseSwiftGrammarListener) ExitIsEmpty(ctx *IsEmptyContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseSwiftGrammarListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseSwiftGrammarListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterFloatCastExpr is called when production FloatCastExpr is entered.
func (s *BaseSwiftGrammarListener) EnterFloatCastExpr(ctx *FloatCastExprContext) {}

// ExitFloatCastExpr is called when production FloatCastExpr is exited.
func (s *BaseSwiftGrammarListener) ExitFloatCastExpr(ctx *FloatCastExprContext) {}

// EnterFloatExpr is called when production FloatExpr is entered.
func (s *BaseSwiftGrammarListener) EnterFloatExpr(ctx *FloatExprContext) {}

// ExitFloatExpr is called when production FloatExpr is exited.
func (s *BaseSwiftGrammarListener) ExitFloatExpr(ctx *FloatExprContext) {}

// EnterNilExpr is called when production NilExpr is entered.
func (s *BaseSwiftGrammarListener) EnterNilExpr(ctx *NilExprContext) {}

// ExitNilExpr is called when production NilExpr is exited.
func (s *BaseSwiftGrammarListener) ExitNilExpr(ctx *NilExprContext) {}

// EnterAccesoMatriz is called when production AccesoMatriz is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoMatriz(ctx *AccesoMatrizContext) {}

// ExitAccesoMatriz is called when production AccesoMatriz is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoMatriz(ctx *AccesoMatrizContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseSwiftGrammarListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseSwiftGrammarListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterCount is called when production Count is entered.
func (s *BaseSwiftGrammarListener) EnterCount(ctx *CountContext) {}

// ExitCount is called when production Count is exited.
func (s *BaseSwiftGrammarListener) ExitCount(ctx *CountContext) {}

// EnterOpExpr is called when production OpExpr is entered.
func (s *BaseSwiftGrammarListener) EnterOpExpr(ctx *OpExprContext) {}

// ExitOpExpr is called when production OpExpr is exited.
func (s *BaseSwiftGrammarListener) ExitOpExpr(ctx *OpExprContext) {}

// EnterCharExpr is called when production CharExpr is entered.
func (s *BaseSwiftGrammarListener) EnterCharExpr(ctx *CharExprContext) {}

// ExitCharExpr is called when production CharExpr is exited.
func (s *BaseSwiftGrammarListener) ExitCharExpr(ctx *CharExprContext) {}

// EnterAccesoVector is called when production AccesoVector is entered.
func (s *BaseSwiftGrammarListener) EnterAccesoVector(ctx *AccesoVectorContext) {}

// ExitAccesoVector is called when production AccesoVector is exited.
func (s *BaseSwiftGrammarListener) ExitAccesoVector(ctx *AccesoVectorContext) {}

// EnterUmenosExpr is called when production UmenosExpr is entered.
func (s *BaseSwiftGrammarListener) EnterUmenosExpr(ctx *UmenosExprContext) {}

// ExitUmenosExpr is called when production UmenosExpr is exited.
func (s *BaseSwiftGrammarListener) ExitUmenosExpr(ctx *UmenosExprContext) {}

// EnterLlamadaFuncExpr is called when production LlamadaFuncExpr is entered.
func (s *BaseSwiftGrammarListener) EnterLlamadaFuncExpr(ctx *LlamadaFuncExprContext) {}

// ExitLlamadaFuncExpr is called when production LlamadaFuncExpr is exited.
func (s *BaseSwiftGrammarListener) ExitLlamadaFuncExpr(ctx *LlamadaFuncExprContext) {}

// EnterParExpr is called when production ParExpr is entered.
func (s *BaseSwiftGrammarListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production ParExpr is exited.
func (s *BaseSwiftGrammarListener) ExitParExpr(ctx *ParExprContext) {}

// EnterStringCastExpr is called when production StringCastExpr is entered.
func (s *BaseSwiftGrammarListener) EnterStringCastExpr(ctx *StringCastExprContext) {}

// ExitStringCastExpr is called when production StringCastExpr is exited.
func (s *BaseSwiftGrammarListener) ExitStringCastExpr(ctx *StringCastExprContext) {}

// EnterIntCastExpr is called when production IntCastExpr is entered.
func (s *BaseSwiftGrammarListener) EnterIntCastExpr(ctx *IntCastExprContext) {}

// ExitIntCastExpr is called when production IntCastExpr is exited.
func (s *BaseSwiftGrammarListener) ExitIntCastExpr(ctx *IntCastExprContext) {}

// EnterStrExpr is called when production StrExpr is entered.
func (s *BaseSwiftGrammarListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production StrExpr is exited.
func (s *BaseSwiftGrammarListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseSwiftGrammarListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseSwiftGrammarListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseSwiftGrammarListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseSwiftGrammarListener) ExitIntExpr(ctx *IntExprContext) {}
