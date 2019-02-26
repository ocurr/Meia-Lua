// Code generated from Lua.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lua

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLuaListener is a complete listener for a parse tree produced by LuaParser.
type BaseLuaListener struct{}

var _ LuaListener = &BaseLuaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLuaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLuaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLuaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLuaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChunk is called when production chunk is entered.
func (s *BaseLuaListener) EnterChunk(ctx *ChunkContext) {}

// ExitChunk is called when production chunk is exited.
func (s *BaseLuaListener) ExitChunk(ctx *ChunkContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLuaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLuaListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseLuaListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseLuaListener) ExitStat(ctx *StatContext) {}

// EnterRetstat is called when production retstat is entered.
func (s *BaseLuaListener) EnterRetstat(ctx *RetstatContext) {}

// ExitRetstat is called when production retstat is exited.
func (s *BaseLuaListener) ExitRetstat(ctx *RetstatContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseLuaListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseLuaListener) ExitLabel(ctx *LabelContext) {}

// EnterFuncname is called when production funcname is entered.
func (s *BaseLuaListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BaseLuaListener) ExitFuncname(ctx *FuncnameContext) {}

// EnterVarlist is called when production varlist is entered.
func (s *BaseLuaListener) EnterVarlist(ctx *VarlistContext) {}

// ExitVarlist is called when production varlist is exited.
func (s *BaseLuaListener) ExitVarlist(ctx *VarlistContext) {}

// EnterTypedvarlist is called when production typedvarlist is entered.
func (s *BaseLuaListener) EnterTypedvarlist(ctx *TypedvarlistContext) {}

// ExitTypedvarlist is called when production typedvarlist is exited.
func (s *BaseLuaListener) ExitTypedvarlist(ctx *TypedvarlistContext) {}

// EnterNamelist is called when production namelist is entered.
func (s *BaseLuaListener) EnterNamelist(ctx *NamelistContext) {}

// ExitNamelist is called when production namelist is exited.
func (s *BaseLuaListener) ExitNamelist(ctx *NamelistContext) {}

// EnterExplist is called when production explist is entered.
func (s *BaseLuaListener) EnterExplist(ctx *ExplistContext) {}

// ExitExplist is called when production explist is exited.
func (s *BaseLuaListener) ExitExplist(ctx *ExplistContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseLuaListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseLuaListener) ExitExp(ctx *ExpContext) {}

// EnterTypeLiteral is called when production typeLiteral is entered.
func (s *BaseLuaListener) EnterTypeLiteral(ctx *TypeLiteralContext) {}

// ExitTypeLiteral is called when production typeLiteral is exited.
func (s *BaseLuaListener) ExitTypeLiteral(ctx *TypeLiteralContext) {}

// EnterPrefixexp is called when production prefixexp is entered.
func (s *BaseLuaListener) EnterPrefixexp(ctx *PrefixexpContext) {}

// ExitPrefixexp is called when production prefixexp is exited.
func (s *BaseLuaListener) ExitPrefixexp(ctx *PrefixexpContext) {}

// EnterFunctioncall is called when production functioncall is entered.
func (s *BaseLuaListener) EnterFunctioncall(ctx *FunctioncallContext) {}

// ExitFunctioncall is called when production functioncall is exited.
func (s *BaseLuaListener) ExitFunctioncall(ctx *FunctioncallContext) {}

// EnterVarOrExp is called when production varOrExp is entered.
func (s *BaseLuaListener) EnterVarOrExp(ctx *VarOrExpContext) {}

// ExitVarOrExp is called when production varOrExp is exited.
func (s *BaseLuaListener) ExitVarOrExp(ctx *VarOrExpContext) {}

// EnterVarId is called when production varId is entered.
func (s *BaseLuaListener) EnterVarId(ctx *VarIdContext) {}

// ExitVarId is called when production varId is exited.
func (s *BaseLuaListener) ExitVarId(ctx *VarIdContext) {}

// EnterTypedvar is called when production typedvar is entered.
func (s *BaseLuaListener) EnterTypedvar(ctx *TypedvarContext) {}

// ExitTypedvar is called when production typedvar is exited.
func (s *BaseLuaListener) ExitTypedvar(ctx *TypedvarContext) {}

// EnterVarSuffix is called when production varSuffix is entered.
func (s *BaseLuaListener) EnterVarSuffix(ctx *VarSuffixContext) {}

// ExitVarSuffix is called when production varSuffix is exited.
func (s *BaseLuaListener) ExitVarSuffix(ctx *VarSuffixContext) {}

// EnterNameAndArgs is called when production nameAndArgs is entered.
func (s *BaseLuaListener) EnterNameAndArgs(ctx *NameAndArgsContext) {}

// ExitNameAndArgs is called when production nameAndArgs is exited.
func (s *BaseLuaListener) ExitNameAndArgs(ctx *NameAndArgsContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseLuaListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseLuaListener) ExitArgs(ctx *ArgsContext) {}

// EnterFunctiondef is called when production functiondef is entered.
func (s *BaseLuaListener) EnterFunctiondef(ctx *FunctiondefContext) {}

// ExitFunctiondef is called when production functiondef is exited.
func (s *BaseLuaListener) ExitFunctiondef(ctx *FunctiondefContext) {}

// EnterFuncbody is called when production funcbody is entered.
func (s *BaseLuaListener) EnterFuncbody(ctx *FuncbodyContext) {}

// ExitFuncbody is called when production funcbody is exited.
func (s *BaseLuaListener) ExitFuncbody(ctx *FuncbodyContext) {}

// EnterParlist is called when production parlist is entered.
func (s *BaseLuaListener) EnterParlist(ctx *ParlistContext) {}

// ExitParlist is called when production parlist is exited.
func (s *BaseLuaListener) ExitParlist(ctx *ParlistContext) {}

// EnterTableconstructor is called when production tableconstructor is entered.
func (s *BaseLuaListener) EnterTableconstructor(ctx *TableconstructorContext) {}

// ExitTableconstructor is called when production tableconstructor is exited.
func (s *BaseLuaListener) ExitTableconstructor(ctx *TableconstructorContext) {}

// EnterFieldlist is called when production fieldlist is entered.
func (s *BaseLuaListener) EnterFieldlist(ctx *FieldlistContext) {}

// ExitFieldlist is called when production fieldlist is exited.
func (s *BaseLuaListener) ExitFieldlist(ctx *FieldlistContext) {}

// EnterField is called when production field is entered.
func (s *BaseLuaListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLuaListener) ExitField(ctx *FieldContext) {}

// EnterFieldsep is called when production fieldsep is entered.
func (s *BaseLuaListener) EnterFieldsep(ctx *FieldsepContext) {}

// ExitFieldsep is called when production fieldsep is exited.
func (s *BaseLuaListener) ExitFieldsep(ctx *FieldsepContext) {}

// EnterOperatorOr is called when production operatorOr is entered.
func (s *BaseLuaListener) EnterOperatorOr(ctx *OperatorOrContext) {}

// ExitOperatorOr is called when production operatorOr is exited.
func (s *BaseLuaListener) ExitOperatorOr(ctx *OperatorOrContext) {}

// EnterOperatorAnd is called when production operatorAnd is entered.
func (s *BaseLuaListener) EnterOperatorAnd(ctx *OperatorAndContext) {}

// ExitOperatorAnd is called when production operatorAnd is exited.
func (s *BaseLuaListener) ExitOperatorAnd(ctx *OperatorAndContext) {}

// EnterOperatorComparison is called when production operatorComparison is entered.
func (s *BaseLuaListener) EnterOperatorComparison(ctx *OperatorComparisonContext) {}

// ExitOperatorComparison is called when production operatorComparison is exited.
func (s *BaseLuaListener) ExitOperatorComparison(ctx *OperatorComparisonContext) {}

// EnterOperatorStrcat is called when production operatorStrcat is entered.
func (s *BaseLuaListener) EnterOperatorStrcat(ctx *OperatorStrcatContext) {}

// ExitOperatorStrcat is called when production operatorStrcat is exited.
func (s *BaseLuaListener) ExitOperatorStrcat(ctx *OperatorStrcatContext) {}

// EnterOperatorAddSub is called when production operatorAddSub is entered.
func (s *BaseLuaListener) EnterOperatorAddSub(ctx *OperatorAddSubContext) {}

// ExitOperatorAddSub is called when production operatorAddSub is exited.
func (s *BaseLuaListener) ExitOperatorAddSub(ctx *OperatorAddSubContext) {}

// EnterOperatorMulDivMod is called when production operatorMulDivMod is entered.
func (s *BaseLuaListener) EnterOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// ExitOperatorMulDivMod is called when production operatorMulDivMod is exited.
func (s *BaseLuaListener) ExitOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// EnterOperatorBitwise is called when production operatorBitwise is entered.
func (s *BaseLuaListener) EnterOperatorBitwise(ctx *OperatorBitwiseContext) {}

// ExitOperatorBitwise is called when production operatorBitwise is exited.
func (s *BaseLuaListener) ExitOperatorBitwise(ctx *OperatorBitwiseContext) {}

// EnterOperatorUnary is called when production operatorUnary is entered.
func (s *BaseLuaListener) EnterOperatorUnary(ctx *OperatorUnaryContext) {}

// ExitOperatorUnary is called when production operatorUnary is exited.
func (s *BaseLuaListener) ExitOperatorUnary(ctx *OperatorUnaryContext) {}

// EnterOperatorPower is called when production operatorPower is entered.
func (s *BaseLuaListener) EnterOperatorPower(ctx *OperatorPowerContext) {}

// ExitOperatorPower is called when production operatorPower is exited.
func (s *BaseLuaListener) ExitOperatorPower(ctx *OperatorPowerContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseLuaListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseLuaListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseLuaListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseLuaListener) ExitStringLiteral(ctx *StringLiteralContext) {}
