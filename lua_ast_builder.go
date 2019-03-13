package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/senior-project/antlr/parser"
)

type LuaASTBuilder struct {
	*antlr.BaseParseTreeVisitor
}

func NewLuaASTBuilder() *LuaASTBuilder {
	return new(LuaASTBuilder)
}

func (v *LuaASTBuilder) VisitChunk(ctx *parser.ChunkContext) interface{} {
	return ChunkC{Block: v.Visit(ctx.Block()).(BlockC)}
	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	stats := ctx.AllStat()
	statLst := make([]Stat, len(stats))
	for i, _ := range stats {
		statLst[i] = v.Visit(stats[i]).(Stat)
	}
	return BlockC{StatLst: statLst}
	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitStat(ctx *parser.StatContext) interface{} {

	if t := ctx.Typedvarlist(); t != nil {
		e := ctx.Explist()

		allTVar := v.Visit(t).(IdLst)
		allExp := v.Visit(e).(ExpLst)

		if len(allTVar.List) != len(allExp.List) {
			fmt.Println("ERROR: AST: the var list is not the same length as the expression list")
			return DefLst{}
		}

		lst := make([]DefC, len(allTVar.List))

		for i := 0; i < len(allTVar.List); i++ {
			lst[i] = DefC{Id: allTVar.List[i], Exp: allExp.List[i]}
		}

		return DefLst{List: lst}

	}

	return DefC{}
	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitRetstat(ctx *parser.RetstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitLabel(ctx *parser.LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFuncname(ctx *parser.FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarlist(ctx *parser.VarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTypedvarlist(ctx *parser.TypedvarlistContext) interface{} {

	allVars := ctx.AllTypedvar()

	allIdC := make([]IdC, len(allVars))

	for i := 0; i < len(allVars); i++ {
		allIdC[i] = v.Visit(allVars[i]).(IdC)
	}

	return IdLst{List: allIdC}

	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitExplist(ctx *parser.ExplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitExp(ctx *parser.ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTypeLiteral(ctx *parser.TypeLiteralContext) interface{} {
	switch ctx.GetText() {
	case "float":
		return FloatT{}
	case "int":
		return IntT{}
	default:
		fmt.Printf("ERROR type %s is not supported\n", ctx.GetText())
		return ErrorT{}
	}
}

func (v *LuaASTBuilder) VisitPrefixexp(ctx *parser.PrefixexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFunctioncall(ctx *parser.FunctioncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarOrExp(ctx *parser.VarOrExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarId(ctx *parser.VarIdContext) interface{} {
	return ctx.GetText()
	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTypedvar(ctx *parser.TypedvarContext) interface{} {

	return IdC{Id: v.Visit(ctx.VarId()).(string), TypeId: v.Visit(ctx.TypeLiteral()).(TypeT)}

	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarSuffix(ctx *parser.VarSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNameAndArgs(ctx *parser.NameAndArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitArgs(ctx *parser.ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFunctiondef(ctx *parser.FunctiondefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFuncbody(ctx *parser.FuncbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitParlist(ctx *parser.ParlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTableconstructor(ctx *parser.TableconstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFieldlist(ctx *parser.FieldlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFieldsep(ctx *parser.FieldsepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorOr(ctx *parser.OperatorOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorAnd(ctx *parser.OperatorAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorComparison(ctx *parser.OperatorComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorStrcat(ctx *parser.OperatorStrcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorAddSub(ctx *parser.OperatorAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorMulDivMod(ctx *parser.OperatorMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorBitwise(ctx *parser.OperatorBitwiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorUnary(ctx *parser.OperatorUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorPower(ctx *parser.OperatorPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNumberLiteral(ctx *parser.NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
