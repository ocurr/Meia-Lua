// Code generated from Lua.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lua

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLuaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLuaVisitor) VisitChunk(ctx *ChunkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitRetstat(ctx *RetstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFuncname(ctx *FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitIfstat(ctx *IfstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitElseifstat(ctx *ElseifstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitElsestat(ctx *ElsestatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitWhilestat(ctx *WhilestatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVarlist(ctx *VarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitTypedvarlist(ctx *TypedvarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitNamelist(ctx *NamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExplist(ctx *ExplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExp(ctx *ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitTypeLiteral(ctx *TypeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitPrefixexp(ctx *PrefixexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFunctioncall(ctx *FunctioncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVarOrExp(ctx *VarOrExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVarId(ctx *VarIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitTypedvar(ctx *TypedvarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVarSuffix(ctx *VarSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitNameAndArgs(ctx *NameAndArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFunctiondef(ctx *FunctiondefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFuncbody(ctx *FuncbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitParlist(ctx *ParlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitTableconstructor(ctx *TableconstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFieldlist(ctx *FieldlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitFieldsep(ctx *FieldsepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorOr(ctx *OperatorOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorAnd(ctx *OperatorAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorComparison(ctx *OperatorComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorStrcat(ctx *OperatorStrcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorAddSub(ctx *OperatorAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorMulDivMod(ctx *OperatorMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorBitwise(ctx *OperatorBitwiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorUnary(ctx *OperatorUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitOperatorPower(ctx *OperatorPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
