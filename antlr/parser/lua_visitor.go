// Code generated from Lua.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lua

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LuaParser.
type LuaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LuaParser#chunk.
	VisitChunk(ctx *ChunkContext) interface{}

	// Visit a parse tree produced by LuaParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LuaParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by LuaParser#retstat.
	VisitRetstat(ctx *RetstatContext) interface{}

	// Visit a parse tree produced by LuaParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by LuaParser#funcname.
	VisitFuncname(ctx *FuncnameContext) interface{}

	// Visit a parse tree produced by LuaParser#ifstat.
	VisitIfstat(ctx *IfstatContext) interface{}

	// Visit a parse tree produced by LuaParser#elseifstat.
	VisitElseifstat(ctx *ElseifstatContext) interface{}

	// Visit a parse tree produced by LuaParser#elsestat.
	VisitElsestat(ctx *ElsestatContext) interface{}

	// Visit a parse tree produced by LuaParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by LuaParser#varlist.
	VisitVarlist(ctx *VarlistContext) interface{}

	// Visit a parse tree produced by LuaParser#typedvarlist.
	VisitTypedvarlist(ctx *TypedvarlistContext) interface{}

	// Visit a parse tree produced by LuaParser#namelist.
	VisitNamelist(ctx *NamelistContext) interface{}

	// Visit a parse tree produced by LuaParser#explist.
	VisitExplist(ctx *ExplistContext) interface{}

	// Visit a parse tree produced by LuaParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by LuaParser#boolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by LuaParser#typeLiteral.
	VisitTypeLiteral(ctx *TypeLiteralContext) interface{}

	// Visit a parse tree produced by LuaParser#prefixexp.
	VisitPrefixexp(ctx *PrefixexpContext) interface{}

	// Visit a parse tree produced by LuaParser#functioncall.
	VisitFunctioncall(ctx *FunctioncallContext) interface{}

	// Visit a parse tree produced by LuaParser#varOrExp.
	VisitVarOrExp(ctx *VarOrExpContext) interface{}

	// Visit a parse tree produced by LuaParser#varId.
	VisitVarId(ctx *VarIdContext) interface{}

	// Visit a parse tree produced by LuaParser#typedvar.
	VisitTypedvar(ctx *TypedvarContext) interface{}

	// Visit a parse tree produced by LuaParser#varSuffix.
	VisitVarSuffix(ctx *VarSuffixContext) interface{}

	// Visit a parse tree produced by LuaParser#nameAndArgs.
	VisitNameAndArgs(ctx *NameAndArgsContext) interface{}

	// Visit a parse tree produced by LuaParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by LuaParser#functiondef.
	VisitFunctiondef(ctx *FunctiondefContext) interface{}

	// Visit a parse tree produced by LuaParser#funcbody.
	VisitFuncbody(ctx *FuncbodyContext) interface{}

	// Visit a parse tree produced by LuaParser#parlist.
	VisitParlist(ctx *ParlistContext) interface{}

	// Visit a parse tree produced by LuaParser#tableconstructor.
	VisitTableconstructor(ctx *TableconstructorContext) interface{}

	// Visit a parse tree produced by LuaParser#fieldlist.
	VisitFieldlist(ctx *FieldlistContext) interface{}

	// Visit a parse tree produced by LuaParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by LuaParser#fieldsep.
	VisitFieldsep(ctx *FieldsepContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorOr.
	VisitOperatorOr(ctx *OperatorOrContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorAnd.
	VisitOperatorAnd(ctx *OperatorAndContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorComparison.
	VisitOperatorComparison(ctx *OperatorComparisonContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorStrcat.
	VisitOperatorStrcat(ctx *OperatorStrcatContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorAddSub.
	VisitOperatorAddSub(ctx *OperatorAddSubContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorMulDivMod.
	VisitOperatorMulDivMod(ctx *OperatorMulDivModContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorBitwise.
	VisitOperatorBitwise(ctx *OperatorBitwiseContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorUnary.
	VisitOperatorUnary(ctx *OperatorUnaryContext) interface{}

	// Visit a parse tree produced by LuaParser#operatorPower.
	VisitOperatorPower(ctx *OperatorPowerContext) interface{}

	// Visit a parse tree produced by LuaParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by LuaParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}
}
